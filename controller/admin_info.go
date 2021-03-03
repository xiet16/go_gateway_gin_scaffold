package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xiet16/gin_scaffold/dto"
	"github.com/xiet16/gin_scaffold/middleware"
	"github.com/xiet16/gin_scaffold/public"
)

type AdminInfoController struct{}

func AdminInfoRegister(group *gin.RouterGroup) {
	adminInfo := &AdminInfoController{}
	group.GET("/admin_info", adminInfo.AdminInfo)
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
func (adminlogin *AdminInfoController) AdminInfo(c *gin.Context) {
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
