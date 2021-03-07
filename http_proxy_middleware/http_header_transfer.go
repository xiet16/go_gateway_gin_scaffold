package http_proxy_middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
)

func HttpHeaderTransferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}

		serviceDetail := serviceInterface.(*dao.ServiceDetail)
		for _, item := range strings.Split(serviceDetail.HttpRule.HeaderTransfor, ",") {
			items := strings.Split(item, " ")
			if len(items) != 3 {
				continue
			}

			if items[0] == "add" || items[1] == "edit" {
				c.Request.Header.Set(items[1], items[2])
			}
			if items[0] == "del" {
				c.Request.Header.Del(items[1])
			}
		}
		c.Next()
	}
}
