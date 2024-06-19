// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReOrderList is the golang structure for table re_order_list.
type ReOrderList struct {
	OrderId    string      `json:"order_id"    orm:"order_id"    ` // 订单ID
	UserId     int         `json:"user_id"     orm:"user_id"     ` // 用户ID
	RecyclerId int         `json:"recycler_id" orm:"recycler_id" ` // 回收员ID
	Kg         float64     `json:"kg"          orm:"kg"          ` // 公斤
	Balance    int         `json:"balance"     orm:"balance"     ` // 红包
	Name       string      `json:"name"        orm:"name"        ` // 名字
	Phone      string      `json:"phone"       orm:"phone"       ` // 手机号
	Provinces  string      `json:"provinces"   orm:"provinces"   ` // 省
	Citys      string      `json:"citys"       orm:"citys"       ` // 市
	Areas      string      `json:"areas"       orm:"areas"       ` // 区
	Detail     string      `json:"detail"      orm:"detail"      ` // 详细地址
	Day        string      `json:"day"         orm:"day"         ` // 月日
	Week       string      `json:"week"        orm:"week"        ` // 星期
	Time       string      `json:"time"        orm:"time"        ` // 时间
	Estimate   string      `json:"estimate"    orm:"estimate"    ` // 预估重量
	Type       string      `json:"type"        orm:"type"        ` // 状态
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  ` // 更新时间
	DeletedAt  *gtime.Time `json:"deleted_at"  orm:"deleted_at"  ` // 删除时间
}
