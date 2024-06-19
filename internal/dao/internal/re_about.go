// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReAboutDao is the data access object for table re_about.
type ReAboutDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ReAboutColumns // columns contains all the column names of Table for convenient usage.
}

// ReAboutColumns defines and stores column names for table re_about.
type ReAboutColumns struct {
	Id        string //
	Content   string //
	CreatedAt string //
	UpdatedAt string //
}

// reAboutColumns holds the columns for table re_about.
var reAboutColumns = ReAboutColumns{
	Id:        "id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewReAboutDao creates and returns a new DAO object for table data access.
func NewReAboutDao() *ReAboutDao {
	return &ReAboutDao{
		group:   "default",
		table:   "re_about",
		columns: reAboutColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReAboutDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReAboutDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReAboutDao) Columns() ReAboutColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReAboutDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReAboutDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReAboutDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
