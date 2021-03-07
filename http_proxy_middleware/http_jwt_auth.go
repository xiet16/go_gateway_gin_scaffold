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

func HttpJwtAuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}

		serviceDetail := serviceInterface.(*dao.ServiceDetail)

		//decode jwt token
		//app_id 与app_list
		//appinfo 放到gin.Context

		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer", "")
		fmt.Println("token:", token)
		appMathched := false
		if token != "" {
			claims, err := public.JwtDecode(token)
			fmt.Println("claims:", claims)
			if err != nil {
				middleware.ResponseError(c, 2001, err)
				c.Abort()
				return
			}

			appList := dao.AppManagerHandler.GetAppList()
			for _, appInfo := range appList {
				if appInfo.AppID == claims.Issuer {
					appMathched = true
					c.Set("appDetail", appInfo)
					break
				}
			}
		}

		if serviceDetail.AccessControl.OpenAuth == 1 && !appMathched {
			middleware.ResponseError(c, 2003, errors.New("not match valid app"))
			c.Abort()
			return
		}
		c.Next()
	}
}
