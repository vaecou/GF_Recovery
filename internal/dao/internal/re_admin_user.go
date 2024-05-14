// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReAdminUserDao is the data access object for table re_admin_user.
type ReAdminUserDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ReAdminUserColumns // columns contains all the column names of Table for convenient usage.
}

// ReAdminUserColumns defines and stores column names for table re_admin_user.
type ReAdminUserColumns struct {
	Id        string // ID
	Name      string // 昵称
	Account   string // 账户
	Password  string // 密码
	Salt      string // 盐值
	Status    string // 状态值
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// reAdminUserColumns holds the columns for table re_admin_user.
var reAdminUserColumns = ReAdminUserColumns{
	Id:        "id",
	Name:      "name",
	Account:   "account",
	Password:  "password",
	Salt:      "salt",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewReAdminUserDao creates and returns a new DAO object for table data access.
func NewReAdminUserDao() *ReAdminUserDao {
	return &ReAdminUserDao{
		group:   "default",
		table:   "re_admin_user",
		columns: reAdminUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReAdminUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReAdminUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReAdminUserDao) Columns() ReAdminUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReAdminUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReAdminUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReAdminUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
