package mini

import "github.com/gogf/gf/v2/frame/g"

type GetAboutInfoReq struct {
	g.Meta `path:"/about" tags:"小程序/关于我们" summary:"获取关于我们"`
}

type GetProtocolInfoReq struct {
	g.Meta `path:"/protocol" tags:"小程序/回收协议" summary:"获取回收协议"`
}

type MoreInfoRes struct {
	Value string `json:"value"`
}
