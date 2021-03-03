package controller

import (
	"fmt"

	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/dao"
	"github.com/xiet16/gin_scaffold/dto"
	"github.com/xiet16/gin_scaffold/middleware"
	"github.com/xiet16/gin_scaffold/public"
)

type ServiceController struct {
}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务列表
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param pageSize query int true "页大小"
// @Param pageNo query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}

	list, total, err := serviceInfo.PageList(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
	}

	outList := []dto.ServiceListItemOutput{}

	for _, item := range list {
		serviceAddr := "unknown"
		// 1 http+后缀接入  clusterip+cluster_port+path
		// 2 http域名接入  domain
		// 3 grpc,tcp  ip+port
		serviceDetail, err := item.ServiceDetail(c, tx, &item)
		if err != nil {
			middleware.ResponseError(c, 2003, err)
			return
		}
		clusterIP := lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort := lib.GetStringConf("base.cluster.cluster_port")
		clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")

		if serviceDetail.Info.Loadtype == public.LoadTypeHTTP &&
			serviceDetail.HttpRule.RuleType == public.HTTPRulrTypePrefixURL &&
			serviceDetail.HttpRule.NeedHttps == 1 {
			serviceAddr = fmt.Sprint("%s:%s%s", clusterIP, clusterSSLPort, serviceDetail.HttpRule.Rule)
		}

		if serviceDetail.Info.Loadtype == public.LoadTypeHTTP &&
			serviceDetail.HttpRule.RuleType == public.HTTPRulrTypePrefixURL &&
			serviceDetail.HttpRule.NeedHttps == 0 {
			serviceAddr = fmt.Sprint("%s:%s%s", clusterIP, clusterPort, serviceDetail.HttpRule.Rule)
		}

		if serviceDetail.Info.Loadtype == public.LoadTypeHTTP &&
			serviceDetail.HttpRule.RuleType == public.HTTPRuleTypeDomain {
			serviceAddr = serviceDetail.HttpRule.Rule
		}

		if serviceDetail.Info.Loadtype == public.LoadTypeTCP {
			serviceAddr = fmt.Sprint("%s:%d", clusterIP, clusterPort)
		}

		ipList := serviceDetail.LoadBalance.GetIPListByModel(c, tx)
		outItem := dto.ServiceListItemOutput{
			ID:          item.ID,
			ServiceName: item.ServiceName,
			ServiceDesc: item.ServiceDesc,
			ServiceAddr: serviceAddr,
			Qps:         0,
			Qpd:         0,
			TotalNode:   len(ipList),
		}

		outList = append(outList, outItem)
	}

	out := &dto.ServiceListOutput{
		Total: total,
		List:  outList,
	}

	middleware.ResponseSuccess(c, out)
}
