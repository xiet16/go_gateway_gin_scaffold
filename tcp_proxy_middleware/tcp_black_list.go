package tcp_proxy_middleware

import (
	"fmt"
	"strings"

	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

//匹配接入方式 基于请求信息
func TcpBlackListMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		splits := strings.Split(c.conn.RemoteAddr().String(), ":")
		clientIP := ""
		if len(splits) == 2 {
			clientIP = splits[0]
		}

		blackIplist := []string{}
		if serviceDetail.AccessControl.BlackList != "" {
			blackIplist = strings.Split(serviceDetail.AccessControl.BlackList, ",")
		}

		if serviceDetail.AccessControl.OpenAuth == 1 && len(blackIplist) > 0 {
			if public.InStringSlice(blackIplist, clientIP) {
				c.conn.Write([]byte(fmt.Sprintf("%s  in black ip list", clientIP)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
