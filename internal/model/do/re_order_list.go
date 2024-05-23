// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReOrderList is the golang structure of table re_order_list for DAO operations like Where/Data.
type ReOrderList struct {
	g.Meta `orm:"table:re_order_list, do:true"`
	UserId interface{} // 订单ID
}
