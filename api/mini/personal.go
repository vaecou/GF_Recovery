package mini

import "github.com/gogf/gf/v2/frame/g"

type GetInfoReq struct {
	g.Meta `path:"/personal" tags:"小程序/个人页面" summary:"获取个人页面信息"`
}

type InfoRes struct {
	KG      float64 `json:"kg"`
	Balance int     `json:"balance"`
	One     float64 `json:"one"`
	Two     float64 `json:"two"`
	Three   float64 `json:"three"`
}
