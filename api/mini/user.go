package mini

import "github.com/gogf/gf/v2/frame/g"

type GetUserPhoneReq struct {
	g.Meta `path:"/user/phone" tags:"小程序/用户" summary:"用户手机号是否存在"`
}

type UserPhoneRes struct {
	Result bool `json:"result"`
}

type SaveUserPhoneReq struct {
	g.Meta `path:"/user/phone" tags:"小程序/用户" summary:"添加手机号"`

	Code string `json:"code" v:"required#您已取消授权"`
}
