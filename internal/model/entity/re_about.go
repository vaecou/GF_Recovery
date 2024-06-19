// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReAbout is the golang structure for table re_about.
type ReAbout struct {
	Id        int         `json:"id"         orm:"id"         ` //
	Content   string      `json:"content"    orm:"content"    ` //
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` //
}
