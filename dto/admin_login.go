package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/public"
)

type AdminLoginInput struct {
	UserName string `json:"userName" form:"userName" comment:"用户名" example:"用户名" validate:"required,valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"密码" validate:"required"`
}

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"userName"`
	LoginTime time.Time `json:"loginTime"`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
}
