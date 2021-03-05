package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"loginTime"`
	Avatar       string    `json:"avatar"`       //头像
	Introduction string    `json:"Introduction"` //介绍
	Roles        []string  `json:"roles"`        //角色
}

type ChangePwdInput struct {
	UserName string `json:"userName" form:"userName" comment:"用户名" example:"用户名" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" example:"密码" validate:"required"`
}

func (param *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
