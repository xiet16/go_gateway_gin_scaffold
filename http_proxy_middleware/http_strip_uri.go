package http_proxy_middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

func HttpStripUriMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}

		serviceDetail := serviceInterface.(*dao.ServiceDetail)
		if serviceDetail.HttpRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HttpRule.NeedStripUri == 1 {
			fmt.Println("replace before:", c.Request.URL.Path)
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, serviceDetail.HttpRule.Rule, "", 1)
			fmt.Println("replace after:", c.Request.URL.Path)
		}
		//http://127.0.0.1:8080/test.xxxx/abc
		//http://127.0.0.1:8080/abc
		c.Next()
	}
}
