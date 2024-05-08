package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddAdminUserReq struct {
	g.Meta `path:"/user" tags:"管理员" summary:"新增管理员"`

	Name     string `json:"name" v:"required#名字不能为空"`
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"required|password#密码不能为空|密码长度在6-18位之间"`
	Status   *bool  `json:"status"`
	Salt     string
}

type UpdateAdminUserReq struct {
	g.Meta `path:"/user" tags:"管理员" summary:"更新管理员"`

	ID       int    `json:"id"`
	Name     string `json:"name" v:"required#名字不能为空"`
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"password#密码长度在6-18位之间"`
	Status   *bool  `json:"status"`
	Salt     string
}

type GetAdminUserReq struct {
	g.Meta `path:"/user" tags:"管理员" summary:"查询管理员"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type GetAdminUserListReq struct {
	g.Meta `path:"/user/list" tags:"管理员" summary:"查询管理员列表"`

	Page  int `json:"page" v:"required#页不能为空"`
	Limit int `json:"limit" v:"required#数量不能为空"`
}

type AdminUserRes struct {
	Total int                 `json:"total"`
	List  []*AdminUserListRes `json:"list"`
}

type AdminUserListRes struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Account   string `json:"account"`
	Status    string `json:"status"`
	CreatedAt string `json:"created"`
}
