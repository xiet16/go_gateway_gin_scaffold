package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/public"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"关键词" validate:""`
	PageNo   string `json:"pageNo" form:"pageNo" comment:"页数" example:"页数" validate:""`
	PageSize string `json:"pageSize" form:"pageSize" comment:"每页条数" example:"每页条数" validate:""`
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id" comment:"关键词" example:"关键词" validate:""`
	ServiceName string `json:"serviceName" form:"serviceName" comment:"页数" example:"页数" validate:""`
	ServiceDesc string `json:"pageSize" form:"pageSize" comment:"每页条数" example:"每页条数" validate:""`
	LoadType    string `json:"loadType" form:"loadType" comment:"页数" example:"页数" validate:""`
	ServiceAddr string `json:"serviceAddr" form:"serviceAddr" comment:"每页条数" example:"每页条数" validate:""`
	Qps         int64  `json:"qps" form:"qps" comment:"页数" example:"页数" validate:""`
	Qpd         int64  `json:"qpd" form:"qpd" comment:"每页条数" example:"每页条数" validate:""`
	TotalNode   int64  `json:"totalNode" form:"totalNode" comment:"每页条数" example:"每页条数" validate:""`
}

type ServiceListOutput struct {
	Total       string `json:"total" form:"total" comment:"总条数" example:"总条数" validate:""`
	ServiceName string `json:"serviceName" form:"serviceName" comment:"服务名" example:"服务名" validate:""`
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
