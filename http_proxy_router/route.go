package http_proxy_router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/http_proxy_middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use(http_proxy_middleware.HttpAccessModeMiddleware(),
		http_proxy_middleware.HttpHeaderTransferMiddleware(),
		http_proxy_middleware.HTTPReverseProxyMiddleware(),
		http_proxy_middleware.HttpStripUriMiddleware())

	return router
}
