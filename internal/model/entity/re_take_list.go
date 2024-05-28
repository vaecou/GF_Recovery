// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReTakeList is the golang structure for table re_take_list.
type ReTakeList struct {
	UserId    int         `json:"user_id"    orm:"user_id"    ` //
	Balance   int         `json:"balance"    orm:"balance"    ` //
	Type      int         `json:"type"       orm:"type"       ` //
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` //
}
