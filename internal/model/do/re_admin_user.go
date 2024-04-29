// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReAdminUser is the golang structure of table re_admin_user for DAO operations like Where/Data.
type ReAdminUser struct {
	g.Meta    `orm:"table:re_admin_user, do:true"`
	Id        interface{} // ID
	Name      interface{} // 名字
	Account   interface{} // 账户
	Password  interface{} // 密码
	Salt      interface{} // 盐值
	Status    interface{} // 状态值
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
