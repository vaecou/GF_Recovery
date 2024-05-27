// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReOrderListDao is the data access object for table re_order_list.
type ReOrderListDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ReOrderListColumns // columns contains all the column names of Table for convenient usage.
}

// ReOrderListColumns defines and stores column names for table re_order_list.
type ReOrderListColumns struct {
	OrderId    string // 订单ID
	UserId     string // 用户ID
	RecyclerId string // 回收员ID
	Kg         string // 公斤
	Balance    string // 红包
	Name       string // 名字
	Phone      string // 手机号
	Provinces  string // 省
	Citys      string // 市
	Areas      string // 区
	Detail     string // 详细地址
	Day        string // 月日
	Week       string // 星期
	Time       string // 时间
	Estimate   string // 预估重量
	Type       string // 状态
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// reOrderListColumns holds the columns for table re_order_list.
var reOrderListColumns = ReOrderListColumns{
	OrderId:    "order_id",
	UserId:     "user_id",
	RecyclerId: "recycler_id",
	Kg:         "kg",
	Balance:    "balance",
	Name:       "name",
	Phone:      "phone",
	Provinces:  "provinces",
	Citys:      "citys",
	Areas:      "areas",
	Detail:     "detail",
	Day:        "day",
	Week:       "week",
	Time:       "time",
	Estimate:   "estimate",
	Type:       "type",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewReOrderListDao creates and returns a new DAO object for table data access.
func NewReOrderListDao() *ReOrderListDao {
	return &ReOrderListDao{
		group:   "default",
		table:   "re_order_list",
		columns: reOrderListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReOrderListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReOrderListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReOrderListDao) Columns() ReOrderListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReOrderListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReOrderListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReOrderListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
