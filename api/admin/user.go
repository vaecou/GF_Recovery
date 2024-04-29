package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddAdminUserReq struct {
	g.Meta `path:"/user" tags:"管理员" summary:"新增管理员"`

	Name     string `json:"name" v:"required#名字不能为空"`
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"required|password#密码不能为空|密码长度在6-18位之间"`
	Salt     string `xorm:"salt"`
}

type UpdateAdminUserReq struct {
	g.Meta `path:"/user" tags:"管理员" summary:"更新管理员"`

	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password" v:"password#密码长度在6-18位之间"`
	Salt     string `xorm:"salt"`
}

type GetAdminUserReq struct {
	g.Meta `path:"/user" tags:"管理员" summary:"查询管理员"`

	Id int `json:"id" v:"required#ID不能为空"`
}

type AdminUserRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Account string `json:"account"`
}
