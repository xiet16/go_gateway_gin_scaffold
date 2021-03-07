package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dao"
	"github.com/xiet16/go_gateway_gin_scaffold/dto"
	"github.com/xiet16/go_gateway_gin_scaffold/golang_common/lib"
	"github.com/xiet16/go_gateway_gin_scaffold/middleware"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

type OAuthContoller struct{}

func OAuthRegister(group *gin.RouterGroup) {
	oauth := &OAuthContoller{}
	group.POST("/tokens", oauth.Tokens)
}

// Tokens godoc
// @Summary 管理员信息接口
// @Description 管理员信息接口
// @Tags 管理员信息接口
// @ID /oauth/tokens
// @Accept  json
// @Produce  json
// @Param body body dto.TokensInput true "body"
// @Success 200 {object} middleware.Response{data=dto.TokensOutput} "success"
// @Router /oauth/tokens [get]
func (oauthController *OAuthContoller) Tokens(c *gin.Context) {
	params := &dto.TokensInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	splits := strings.Split(c.GetHeader("Authorization"), " ")
	if len(splits) != 2 {
		middleware.ResponseError(c, 2001, errors.New("用户名或密码格式错误"))
	}

	appSecret, err := base64.StdEncoding.DecodeString(splits[2])
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	fmt.Println("appSecret:", string(appSecret))

	//取出 app_id secrect
	//生成app_list
	//匹配app_id
	//jwt 生成token
	//生成 output

	parts := strings.Split(string(splits[1]), ":")
	if len(parts) != 2 {
		middleware.ResponseError(c, 2002, err)
		return
	}
	appList := dao.AppManagerHandler.GetAppList()
	for _, appInfo := range appList {
		if appInfo.AppID == parts[0] && appInfo.Secret == parts[1] {
			claims := jwt.StandardClaims{
				Issuer:    appInfo.AppID,
				ExpiresAt: time.Now().Add(public.JwtExpires * time.Second).In(lib.TimeLocation).Unix(),
			}
			token, err := public.JwtEncode(claims)
			if err != nil {
				middleware.ResponseError(c, 2004, err)
				return
			}
			output := &dto.TokensOutput{
				Accesstoken: token,
				ExpiresIn:   public.JwtExpires,
				TokenType:   "Bearer",
				Scope:       "read_write",
			}

			middleware.ResponseSuccess(c, output)
			return
		}
	}

	middleware.ResponseError(c, 2005, errors.New("未匹配正确的app 信息"))
}
