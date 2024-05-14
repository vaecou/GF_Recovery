// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReQuestionDao is the data access object for table re_question.
type ReQuestionDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ReQuestionColumns // columns contains all the column names of Table for convenient usage.
}

// ReQuestionColumns defines and stores column names for table re_question.
type ReQuestionColumns struct {
	Id        string //
	Title     string // 标题
	Content   string // 内容
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// reQuestionColumns holds the columns for table re_question.
var reQuestionColumns = ReQuestionColumns{
	Id:        "id",
	Title:     "title",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewReQuestionDao creates and returns a new DAO object for table data access.
func NewReQuestionDao() *ReQuestionDao {
	return &ReQuestionDao{
		group:   "default",
		table:   "re_question",
		columns: reQuestionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReQuestionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReQuestionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReQuestionDao) Columns() ReQuestionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReQuestionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReQuestionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReQuestionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
