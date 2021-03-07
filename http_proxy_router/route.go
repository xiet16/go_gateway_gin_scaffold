package http_proxy_router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/controller"
	"github.com/xiet16/go_gateway_gin_scaffold/http_proxy_middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	oauth := router.Group("/oauth")
	oauth.Use(middleware.TranslationMiddleware())
	{
		controller.OAuthRegister(oauth)
	}

	//root := router.Group("/")
	router.Use(http_proxy_middleware.HttpAccessModeMiddleware(),

		http_proxy_middleware.HttpFlowCountMiddleware(),
		http_proxy_middleware.HttpFlowLimitMiddleware(),

		http_proxy_middleware.HttpJwtAuthTokenMiddleware(),
		http_proxy_middleware.HttpWhiteListMiddleware(),
		http_proxy_middleware.HttpBlackListMiddleware(),
		http_proxy_middleware.HttpJwtFlowCountMiddleware(),

		http_proxy_middleware.HttpHeaderTransferMiddleware(),
		http_proxy_middleware.HttpStripUriMiddleware(),
		http_proxy_middleware.HTTPUrlRewriteMiddleware(),

		http_proxy_middleware.HTTPReverseProxyMiddleware())

	return router
}
