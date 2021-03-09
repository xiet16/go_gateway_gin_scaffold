package reverse_proxy

import (
	"context"
	"log"

	"github.com/e421083458/grpc-proxy/proxy"
	"github.com/xiet16/go_gateway_gin_scaffold/reverse_proxy/load_balance"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewGrpcLoadBalanceHandler(lb load_balance.LoadBalance) grpc.StreamHandler {
	return func() grpc.StreamHandler {
		nextAddr, err := lb.Get("")
		if err != nil {
			log.Fatal("get next addr fail")
		}
		director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
			c, err := grpc.DialContext(ctx, nextAddr, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
			md, _ := metadata.FromIncomingContext(ctx)
			outCtx, _ := context.WithCancel(ctx)
			outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())
			return outCtx, c, err
		}
		return proxy.TransparentHandler(director)
	}()
}
