// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReOrderList is the golang structure of table re_order_list for DAO operations like Where/Data.
type ReOrderList struct {
	g.Meta     `orm:"table:re_order_list, do:true"`
	OrderId    interface{} // 订单ID
	UserId     interface{} // 用户ID
	RecyclerId interface{} // 回收员ID
	Kg         interface{} // 公斤
	Balance    interface{} // 红包
	Name       interface{} // 名字
	Phone      interface{} // 手机号
	Provinces  interface{} // 省
	Citys      interface{} // 市
	Areas      interface{} // 区
	Detail     interface{} // 详细地址
	Day        interface{} // 月日
	Week       interface{} // 星期
	Time       interface{} // 时间
	Estimate   interface{} // 预估重量
	Type       interface{} // 状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
