package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/public"
)

type AdminLoginInput struct {
	UserName string `json:"userName" form:"userName" comment:"用户名" example:"用户名" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" example:"密码" validate:"required"`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
