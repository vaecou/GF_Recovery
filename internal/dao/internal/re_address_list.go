// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReAddressListDao is the data access object for table re_address_list.
type ReAddressListDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ReAddressListColumns // columns contains all the column names of Table for convenient usage.
}

// ReAddressListColumns defines and stores column names for table re_address_list.
type ReAddressListColumns struct {
	UserId    string // 用户ID
	Id        string // ID
	Name      string // 联系人
	Phone     string // 手机号
	Provinces string // 省
	Citys     string // 市
	Areas     string // 区
	Detail    string // 详细地址
	Default   string // 是否默认
}

// reAddressListColumns holds the columns for table re_address_list.
var reAddressListColumns = ReAddressListColumns{
	UserId:    "user_id",
	Id:        "id",
	Name:      "name",
	Phone:     "phone",
	Provinces: "provinces",
	Citys:     "citys",
	Areas:     "areas",
	Detail:    "detail",
	Default:   "default",
}

// NewReAddressListDao creates and returns a new DAO object for table data access.
func NewReAddressListDao() *ReAddressListDao {
	return &ReAddressListDao{
		group:   "default",
		table:   "re_address_list",
		columns: reAddressListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReAddressListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReAddressListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReAddressListDao) Columns() ReAddressListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReAddressListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReAddressListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReAddressListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
