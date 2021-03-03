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

// AdminLogin godoc
// @Summary 管理员登录接口
// @Description 管理员登录接口
// @Tags 管理员登录接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
	}
	out := &dto.AdminLoginOutput{Token: params.UserName}
	middleware.ResponseSuccess(c, out)
}
