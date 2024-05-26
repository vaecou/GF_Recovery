// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReAddressList is the golang structure of table re_address_list for DAO operations like Where/Data.
type ReAddressList struct {
	g.Meta    `orm:"table:re_address_list, do:true"`
	UserId    interface{} // 用户ID
	Id        interface{} // ID
	Name      interface{} // 联系人
	Phone     interface{} // 手机号
	Provinces interface{} // 省
	Citys     interface{} // 市
	Areas     interface{} // 区
	Detail    interface{} // 详细地址
	Default   interface{} // 是否默认
}
