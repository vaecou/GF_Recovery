// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReQuestion is the golang structure for table re_question.
type ReQuestion struct {
	Id        int         `json:"id"         orm:"id"         ` //
	Title     string      `json:"title"      orm:"title"      ` // 标题
	Content   string      `json:"content"    orm:"content"    ` // 内容
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` // 更新时间
}
