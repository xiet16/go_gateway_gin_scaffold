package http_proxy_middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

func HttpFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}

		serviceDetail := serviceInterface.(*dao.ServiceDetail)

		//全站 服务 租户

		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			middleware.ResponseError(c, 4001, errors.New("service not found"))
			c.Abort()
			return
		}
		totalCounter.Increase()
		dayCount, _ := totalCounter.GetDayData(time.Now())
		fmt.Printf("totalcounter qus:%v, daycount:%v", totalCounter.QPS, dayCount)

		serverCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			middleware.ResponseError(c, 4001, errors.New("service not found"))
			c.Abort()
			return
		}
		serverCounter.Increase()
		dayServiceCount, _ := serverCounter.GetDayData(time.Now())
		fmt.Printf("totalcounter qus:%v, daycount:%v", serverCounter.QPS, dayServiceCount)

		c.Next()
	}
}
