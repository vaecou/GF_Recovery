package mini

import "github.com/gogf/gf/v2/frame/g"

type GetNumReq struct {
	g.Meta `path:"/num" tags:"小程序/首页" summary:"获取使用次数数量"`
}
type GetNumRes struct {
	Num int `json:"num"`
}
