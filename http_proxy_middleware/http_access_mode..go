package http_proxy_middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

func HttpAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := dao.ServiceManagerHandler.HttpAccessMode(c)
		if err != nil {
			middleware.ResponseError(c, 1001, err)
			c.Abort()
			return
		}
		fmt.Println("matched service: ", public.ObjToJson(service))
		c.Set("service", service)
		c.Next()
	}
}
