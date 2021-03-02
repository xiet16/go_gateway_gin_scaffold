package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/dto"
	"github.com/xiet16/gin_scaffold/middleware"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminlogin := &AdminLoginController{}
	group.POST("/login", adminlogin.AdminLogin)
}

func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
	}
	middleware.ResponseSuccess(c, "")
}
