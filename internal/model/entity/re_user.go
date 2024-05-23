// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReUser is the golang structure for table re_user.
type ReUser struct {
	Id        int         `json:"id"         orm:"id"         ` // ID
	Balance   float64     `json:"balance"    orm:"balance"    ` // 余额
	Phone     string      `json:"phone"      orm:"phone"      ` // 手机号
	Openid    string      `json:"openid"     orm:"openid"     ` // 小程序用户OpenID
	Unionid   string      `json:"unionid"    orm:"unionid"    ` // 小程序用户UnionID
	Status    bool        `json:"status"     orm:"status"     ` // 状态值
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" ` // 删除时间
}
