// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReAbout is the golang structure of table re_about for DAO operations like Where/Data.
type ReAbout struct {
	g.Meta    `orm:"table:re_about, do:true"`
	Id        interface{} //
	Content   interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
