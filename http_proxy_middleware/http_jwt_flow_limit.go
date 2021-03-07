package http_proxy_middleware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

func HttpJwtFlowLimitMiddleware() gin.HandlerFunc {
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
		if appInfo.Qps > 0 {
			clientLimiter, err := public.FlowLimiterHandler.GetLimiter(
				public.FlowAppPrefix+appInfo.AppID+"_"+c.ClientIP(),
				float64(appInfo.Qps))
			if err != nil {
				middleware.ResponseError(c, 5001, err)
				c.Abort()
				return
			}
			if !clientLimiter.Allow() {
				middleware.ResponseError(c, 5002, errors.New(fmt.Sprintf("%v flow limit %v", c.ClientIP(), appInfo.Qps)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
