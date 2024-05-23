package admin

import "github.com/gogf/gf/v2/frame/g"

type UpdateMiniUserReq struct {
	g.Meta `path:"/mini_user" tags:"用户" summary:"更新用户"`

	ID     int   `json:"id" v:"required#ID不能为空"`
	Status *bool `json:"status" v:"required#Status不能为空"`
}

type GetMiniUserListReq struct {
	g.Meta `path:"/mini_user/list" tags:"用户" summary:"查询用户列表"`

	Title  string `json:"title"`
	Select string `json:"select"`
	Page   int    `json:"page" v:"required#页不能为空"`
	Limit  int    `json:"limit" v:"required#数量不能为空"`
}

type MiniUserRes struct {
	Total int                `json:"total"`
	List  []*MiniUserListRes `json:"list"`
}

type MiniUserListRes struct {
	ID        int     `json:"id"`
	Balance   float64 `json:"balance"`
	Phone     string  `json:"phone"`
	Status    bool    `json:"status"`
	CreatedAt string  `json:"created"`
}
