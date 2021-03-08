package tcp_proxy_router

import (
	"context"
	"fmt"
	"log"

	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/reverse_proxy"
	"github.com/xiet16/go_gateway_gin_scaffold/tcp_proxy_middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/tcp_server"
)

var tcpServerList = []*tcp_server.TcpServer{}

func TcpServerRun() {
	serviceList := dao.ServiceManagerHandler.GetTcpServiceList()
	for _, serviceItem := range serviceList {
		tempItem := serviceItem
		log.Printf("tcp proxy server run %v \n", tempItem.TcpRule.Port)

		go func(serviceDetail *dao.ServiceDetail) {
			addr := fmt.Sprintf(":%d", serviceDetail.TcpRule.Port)

			//构建负载均衡器
			rb, err := dao.LoadBalanceHandler.GetLoadBalancer(serviceDetail)
			if err != nil {
				log.Fatalf("tcp proxy getloadbalancer run %v error %v \n", addr, err)
				return
			}

			//构建路由并且设置中间件
			router := tcp_proxy_middleware.NewTcpSliceRouter()
			router.Group("/").Use(
				tcp_proxy_middleware.TCPFlowCountMiddleware(),
				tcp_proxy_middleware.TCPFlowLimitMiddleware(),
				tcp_proxy_middleware.TCPWhiteListMiddleware(),
				tcp_proxy_middleware.TCPFlowCountMiddleware(),
			)

			//构建回调Handler
			routerHandler := tcp_proxy_middleware.NewTcpSliceRouterHandler(func(c *tcp_proxy_middleware.TcpSliceRouterContext) tcp_server.TCPHandler {
				return reverse_proxy.NewTcpLoadBalanceReverseProxy(c, rb)
			}, router)
			baseCtx := context.WithValue(context.Background(), "service", serviceDetail)
			tcpserver := &tcp_server.TcpServer{
				Addr:    addr,
				Handler: routerHandler,
				BaseCtx: baseCtx,
			}
			tcpServerList = append(tcpServerList, tcpserver)
			if err := tcpserver.ListenAndServe(); err != nil && err != tcp_server.ErrServerClosed {
				log.Fatalf("tcp proxy server run %v error %v \n", tempItem.TcpRule.Port, err)
			}
		}(tempItem)
	}
}

func TcpServerStop() {
	for _, tcpServer := range tcpServerList {
		tcpServer.Close()
		log.Printf("tcp proxy server stop: %v \n", tcpServer.Addr)
	}
}
