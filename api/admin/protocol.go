package admin

import "github.com/gogf/gf/v2/frame/g"

type AddProtocolReq struct {
	g.Meta `path:"/protocol" tags:"后台/回收协议" summary:"增加回收协议"`

	Content string `json:"content" v:"required#内容不能为空"`
}

type UpdateProtocolReq struct {
	g.Meta `path:"/protocol" tags:"后台/回收协议" summary:"更新回收协议"`

	ID      int    `json:"id" v:"required#ID不能为空"`
	Content string `json:"content" v:"required#内容不能为空"`
}

type GetProtocolReq struct {
	g.Meta `path:"/protocol" tags:"后台/回收协议" summary:"查询回收协议"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type GetProtocolListReq struct {
	g.Meta `path:"/protocol/list" tags:"后台/回收协议" summary:"查询回收协议列表"`

	Page  int `json:"page" v:"required#页不能为空"`
	Limit int `json:"limit" v:"required#数量不能为空"`
}

type DeleteProtocolReq struct {
	g.Meta `path:"/protocol" tags:"后台/回收协议" summary:"删除回收协议"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type ProtocolRes struct {
	Total int                `json:"total"`
	List  []*ProtocolListRes `json:"list"`
}

type ProtocolListRes struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created"`
}
