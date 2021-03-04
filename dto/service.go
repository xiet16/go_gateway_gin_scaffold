package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/public"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"关键词" validate:""`
	PageNo   int64  `json:"pageNo" form:"pageNo" comment:"页数" example:"页数" validate:""`
	PageSize int64  `json:"pageSize" form:"pageSize" comment:"每页条数" example:"每页条数" validate:""`
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id" comment:"关键词" example:"" validate:""`
	ServiceName string `json:"serviceName" form:"serviceName" comment:"页数" example:"" validate:""`
	ServiceDesc string `json:"pageSize" form:"pageSize" comment:"每页条数" example:"" validate:""`
	LoadType    int    `json:"loadType" form:"loadType" comment:"页数" example:"" validate:""`
	ServiceAddr string `json:"serviceAddr" form:"serviceAddr" comment:"每页条数" example:"" validate:""`
	Qps         int64  `json:"qps" form:"qps" comment:"页数" example:"" validate:""`
	Qpd         int64  `json:"qpd" form:"qpd" comment:"每页条数" example:"" validate:""`
	TotalNode   int    `json:"totalNode" form:"totalNode" comment:"每页条数" example:"" validate:""`
}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总条数" example:"" validate:""`
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"服务列表" example:"" validate:""`
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务id" 服务id:"服务id" validate:"required"`
}

func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
