// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReUserDao is the data access object for table re_user.
type ReUserDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ReUserColumns // columns contains all the column names of Table for convenient usage.
}

// ReUserColumns defines and stores column names for table re_user.
type ReUserColumns struct {
	Id         string // ID
	Balance    string // 余额
	Phone      string // 手机号
	Openid     string // 小程序用户OpenID
	Unionid    string // 小程序用户UnionID
	Status     string // 状态值
	LastActive string // 最近活跃时间
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// reUserColumns holds the columns for table re_user.
var reUserColumns = ReUserColumns{
	Id:         "id",
	Balance:    "balance",
	Phone:      "phone",
	Openid:     "openid",
	Unionid:    "unionid",
	Status:     "status",
	LastActive: "last_active",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewReUserDao creates and returns a new DAO object for table data access.
func NewReUserDao() *ReUserDao {
	return &ReUserDao{
		group:   "default",
		table:   "re_user",
		columns: reUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReUserDao) Columns() ReUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
