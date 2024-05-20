package mini

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddMiniUserReq struct {
	g.Meta `path:"/user" tags:"用户" summary:"新增用户"`

	Code string `json:"code" v:"required#Code不能为空"`
}

type MiniUserRes struct {
}
