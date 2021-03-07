package http_proxy_middleware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

func HttpJwtFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// serviceInterface, ok := c.Get("service")
		// if !ok {
		// 	middleware.ResponseError(c, 2001, errors.New("service not found"))
		// 	c.Abort()
		// 	return
		// }

		// serviceDetail := serviceInterface.(*dao.ServiceDetail)

		appInterface, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		appInfo := appInterface.(*dao.App)
		appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + appInfo.AppID)
		if err != nil {
			middleware.ResponseError(c, 4001, err)
			c.Abort()
			return
		}

		appCounter.Increase()
		if appInfo.Qpd > 0 && appCounter.TotalCount > appInfo.Qpd {
			middleware.ResponseError(c, 2003, errors.New(fmt.Sprintf("租户日请求量限流 limit:%v,current:%v", appInfo.Qpd, appCounter.TotalCount)))
			c.Abort()
			return
		}
		fmt.Sprintf("appCounter Qpd:%v,current:%v", appCounter.QPS, appCounter.TotalCount)
		c.Next()
	}
}
