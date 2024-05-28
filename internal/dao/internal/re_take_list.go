// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReTakeListDao is the data access object for table re_take_list.
type ReTakeListDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ReTakeListColumns // columns contains all the column names of Table for convenient usage.
}

// ReTakeListColumns defines and stores column names for table re_take_list.
type ReTakeListColumns struct {
	UserId    string //
	Balance   string //
	Type      string //
	CreatedAt string //
}

// reTakeListColumns holds the columns for table re_take_list.
var reTakeListColumns = ReTakeListColumns{
	UserId:    "user_id",
	Balance:   "balance",
	Type:      "type",
	CreatedAt: "created_at",
}

// NewReTakeListDao creates and returns a new DAO object for table data access.
func NewReTakeListDao() *ReTakeListDao {
	return &ReTakeListDao{
		group:   "default",
		table:   "re_take_list",
		columns: reTakeListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReTakeListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReTakeListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReTakeListDao) Columns() ReTakeListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReTakeListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReTakeListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReTakeListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
