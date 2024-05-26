package admin

import "github.com/gogf/gf/v2/frame/g"

type AddOrUpdateSettingReq struct {
	g.Meta `path:"/setting" tags:"后台/设置" summary:"新增或更新键值对"`

	Key   string `json:"key" v:"required#键不能为空"`
	Value string `json:"value" v:"required#值不能为空"`
}

type GetSettingReq struct {
	g.Meta `path:"/setting" tags:"后台/设置" summary:"查询设置"`

	Key string `json:"key" v:"required#键不能为空"`
}

type SettingRes struct {
	Value string `json:"value"`
}
