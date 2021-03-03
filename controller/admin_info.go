package controller

import (
	"encoding/json"
	"fmt"

	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/dao"
	"github.com/xiet16/gin_scaffold/dto"
	"github.com/xiet16/gin_scaffold/middleware"
	"github.com/xiet16/gin_scaffold/public"
)

type AdminInfoController struct{}

func AdminInfoRegister(group *gin.RouterGroup) {
	adminInfo := &AdminInfoController{}
	group.GET("/admin_info", adminInfo.AdminInfo)
	group.POST("/changepwd", adminInfo.ChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息接口
// @Description 管理员信息接口
// @Tags 管理员信息接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (adminInfo *AdminInfoController) AdminInfo(c *gin.Context) {
	//通过session从缓存中查询登录信息，转换成对应的结构体
	//从结构体中取数据，做封装返回
	session := sessions.Default(c)
	//sessionInfo := session.Get(public.AdminSessionInfoKey).(string)
	sessionInfo := session.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessionInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "",
		Introduction: "i am admin",
		Roles:        []string{"管理员"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 更改管理员密码
// @Description 更改管理员密码
// @Tags 更改管理员密码
// @ID /admin/changepwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/changepwd [post]
func (ac *AdminInfoController) ChangePwd(c *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1. session 读取到用户信息
	//2  读取数据库信息
	//3 验证密码

	session := sessions.Default(c)
	sessionInfo := session.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessionInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, &dao.Admin{UserName: adminSessionInfo.UserName})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	newPassword := public.GenSecretPassword(adminInfo.Salt, params.Password)
	adminInfo.Password = newPassword
	err = adminInfo.Save(c, tx)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	middleware.ResponseSuccess(c, "")
}
