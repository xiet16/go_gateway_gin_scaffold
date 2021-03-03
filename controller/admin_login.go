package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/dao"
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
		return
	}
	admin := &dao.Admin{}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}
