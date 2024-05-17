// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReRegions is the golang structure of table re_regions for DAO operations like Where/Data.
type ReRegions struct {
	g.Meta   `orm:"table:re_regions, do:true"`
	Id       interface{} // ID
	ParentId interface{} // 父ID
	Sort     interface{} // 排序
	Regions  interface{} // 值
}
