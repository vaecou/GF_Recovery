// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReProtocolDao is the data access object for table re_protocol.
type ReProtocolDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ReProtocolColumns // columns contains all the column names of Table for convenient usage.
}

// ReProtocolColumns defines and stores column names for table re_protocol.
type ReProtocolColumns struct {
	Id        string //
	Content   string //
	CreatedAt string //
	UpdatedAt string //
}

// reProtocolColumns holds the columns for table re_protocol.
var reProtocolColumns = ReProtocolColumns{
	Id:        "id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewReProtocolDao creates and returns a new DAO object for table data access.
func NewReProtocolDao() *ReProtocolDao {
	return &ReProtocolDao{
		group:   "default",
		table:   "re_protocol",
		columns: reProtocolColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReProtocolDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReProtocolDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReProtocolDao) Columns() ReProtocolColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReProtocolDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReProtocolDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReProtocolDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
