// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReTakeList is the golang structure of table re_take_list for DAO operations like Where/Data.
type ReTakeList struct {
	g.Meta    `orm:"table:re_take_list, do:true"`
	UserId    interface{} //
	Balance   interface{} //
	Type      interface{} //
	CreatedAt *gtime.Time //
}
