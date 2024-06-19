package admin

import "github.com/gogf/gf/v2/frame/g"

type AddAboutReq struct {
	g.Meta `path:"/about" tags:"后台/关于我们" summary:"增加关于我们"`

	Content string `json:"content" v:"required#内容不能为空"`
}

type UpdateAboutReq struct {
	g.Meta `path:"/about" tags:"后台/关于我们" summary:"更新关于我们"`

	ID      int    `json:"id" v:"required#ID不能为空"`
	Content string `json:"content" v:"required#内容不能为空"`
}

type GetAboutReq struct {
	g.Meta `path:"/about" tags:"后台/关于我们" summary:"查询关于我们"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type GetAboutListReq struct {
	g.Meta `path:"/about/list" tags:"后台/关于我们" summary:"查询关于我们列表"`

	Page  int `json:"page" v:"required#页不能为空"`
	Limit int `json:"limit" v:"required#数量不能为空"`
}

type DeleteAboutReq struct {
	g.Meta `path:"/about" tags:"后台/关于我们" summary:"删除关于我们"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type AboutRes struct {
	Total int             `json:"total"`
	List  []*AboutListRes `json:"list"`
}

type AboutListRes struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created"`
}
