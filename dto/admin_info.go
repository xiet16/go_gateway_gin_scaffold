package dto

import "time"

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"loginTime"`
	Avatar       string    `json:"avatar"`       //头像
	Introduction string    `json:"Introduction"` //介绍
	Roles        []string  `json:"roles"`        //角色
}
