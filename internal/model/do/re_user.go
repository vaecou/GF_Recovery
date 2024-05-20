// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReUser is the golang structure of table re_user for DAO operations like Where/Data.
type ReUser struct {
	g.Meta     `orm:"table:re_user, do:true"`
	Id         interface{} // ID
	Balance    interface{} // 余额
	Phone      interface{} // 手机号
	Openid     interface{} // 小程序用户OpenID
	Unionid    interface{} // 小程序用户UnionID
	Status     interface{} // 状态值
	LastActive *gtime.Time // 最近活跃时间
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
