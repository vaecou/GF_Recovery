package admin

import "github.com/gogf/gf/v2/frame/g"

type GetMiniUserListReq struct {
	g.Meta `path:"/mini_user/list" tags:"用户" summary:"查询用户列表"`

	Page  int `json:"page" v:"required#页不能为空"`
	Limit int `json:"limit" v:"required#数量不能为空"`
}

type MiniUserRes struct {
	Total int                `json:"total"`
	List  []*MiniUserListRes `json:"list"`
}

type MiniUserListRes struct {
	ID        int    `json:"id"`
	Balance   int    `json:"balance"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	CreatedAt string `json:"created"`
}
