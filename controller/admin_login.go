package controller

import (
	"encoding/json"
	"time"

	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/dto"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminlogin := &AdminLoginController{}
	group.POST("/login", adminlogin.AdminLogin)
	group.GET("/logout", adminlogin.AdminLogout)
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

	//设置session
	sessionInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessionInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	session := sessions.Default(c)
	session.Set(public.AdminSessionInfoKey, string(sessBts))
	session.Save()
	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// AdminLogout godoc
// @Summary 管理员退出接口
// @Description 管理员退出接口
// @Tags 管理员退出接口
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/logout [get]
func (adminlogin *AdminLoginController) AdminLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(public.AdminSessionInfoKey)
	session.Save()
	middleware.ResponseSuccess(c, "")
}
