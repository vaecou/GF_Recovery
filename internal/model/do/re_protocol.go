// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReProtocol is the golang structure of table re_protocol for DAO operations like Where/Data.
type ReProtocol struct {
	g.Meta    `orm:"table:re_protocol, do:true"`
	Id        interface{} //
	Content   interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
