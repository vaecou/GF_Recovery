// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReSetting is the golang structure of table re_setting for DAO operations like Where/Data.
type ReSetting struct {
	g.Meta `orm:"table:re_setting, do:true"`
	Key    interface{} // Key
	Value  interface{} // Value
}
