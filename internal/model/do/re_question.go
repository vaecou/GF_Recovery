// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReQuestion is the golang structure of table re_question for DAO operations like Where/Data.
type ReQuestion struct {
	g.Meta    `orm:"table:re_question, do:true"`
	Id        interface{} //
	Title     interface{} // 标题
	Content   interface{} // 内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
