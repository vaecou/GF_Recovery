package admin

import "github.com/gogf/gf/v2/frame/g"

type AddRegionsReq struct {
	g.Meta `path:"/regions" tags:"地区" summary:"增加地区"`

	ParentID int    `json:"parent_id"`
	Regions  string `json:"regions" v:"required#地区不能为空"`
}

type UpdateRegionsReq struct {
	g.Meta `path:"/regions" tags:"地区" summary:"修改地区"`

	ID       int    `json:"id" v:"required#ID不能为空"`
	ParentID int    `json:"parent_id" v:"required#ParentID不能为空"`
	Sort     int    `json:"sort" v:"required#Sort不能为空"`
	Regions  string `json:"regions" v:"required#地区不能为空"`
}

type DeleteRegionsReq struct {
	g.Meta `path:"/regions" tags:"地区" summary:"删除地区"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type GetRegionsReq struct {
	g.Meta `path:"/regions" tags:"地区" summary:"获取地区"`

	ID int `json:"id" v:"required#ID不能为空"`
}

type GetRegionsListReq struct {
	g.Meta `path:"/regions/list" tags:"地区" summary:"获取地区列表"`
}

type RegionsRes struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Sort     int    `json:"sort"`
	Regions  string `json:"regions"`

	Children []*RegionsRes `json:"children"`
}
