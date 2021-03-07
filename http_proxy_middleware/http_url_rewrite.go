package http_proxy_middleware

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
)

func HTTPUrlRewriteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}

		serviceDetail := serverInterface.(*dao.ServiceDetail)
		for _, item := range strings.Split(serviceDetail.HttpRule.UrlRewrite, ",") {
			items := strings.Split(item, " ")
			if len(items) != 2 {
				continue
			}
			regexp, err := regexp.Compile(items[0])
			if err != nil {
				fmt.Println("regexp.Complie err", err)
				continue
			}

			replacePath := regexp.ReplaceAll([]byte(c.Request.URL.Path), []byte(items[1]))
			c.Request.URL.Path = string(replacePath)
		}
		c.Next()
	}
}
