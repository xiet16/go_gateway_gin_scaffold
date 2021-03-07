package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

type OAuthInput struct {
	UserName string `json:"userName" form:"userName" comment:"用户名" example:"用户名" validate:"required,valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"密码" validate:"required"`
}

func (param *OAuthInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type TokensInput struct {
	GrantType string `json:"grant_type" form:"grant_type" comment:"授权类型" example:"client_credentials" validate:"required"`
	Scope     string `json:"scope" form:"scope" comment:"范围权限" example:"read_write" validate:"required"`
}

func (param *TokensInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type TokensOutput struct {
	Accesstoken string `json:"access_token" form:"access_token" `
	ExpiresIn   int    `json:"expires_in" form:"expires_in" `
	TokenType   string `json:"token_type" form:"token_type" `
	Scope       string `json:"scope" form:"scope" `
}
