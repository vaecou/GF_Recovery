// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReAdminUser is the golang structure for table re_admin_user.
type ReAdminUser struct {
	Id        int         `json:"id"         orm:"id"         ` // ID
	Name      string      `json:"name"       orm:"name"       ` // 名字
	Account   string      `json:"account"    orm:"account"    ` // 账户
	Password  string      `json:"password"   orm:"password"   ` // 密码
	Salt      string      `json:"salt"       orm:"salt"       ` // 盐值
	Status    bool        `json:"status"     orm:"status"     ` // 状态值
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" ` // 删除时间
}
