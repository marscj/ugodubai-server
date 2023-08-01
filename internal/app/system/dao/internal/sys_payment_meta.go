// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPaymentMetaDao is the data access object for table sys_payment_meta.
type SysPaymentMetaDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysPaymentMetaColumns // columns contains all the column names of Table for convenient usage.
}

// SysPaymentMetaColumns defines and stores column names for table sys_payment_meta.
type SysPaymentMetaColumns struct {
	MetaId    string //
	PaymentId string //
	MetaKey   string //
	MetaValue string //
}

// sysPaymentMetaColumns holds the columns for table sys_payment_meta.
var sysPaymentMetaColumns = SysPaymentMetaColumns{
	MetaId:    "meta_id",
	PaymentId: "payment_id",
	MetaKey:   "meta_key",
	MetaValue: "meta_value",
}

// NewSysPaymentMetaDao creates and returns a new DAO object for table data access.
func NewSysPaymentMetaDao() *SysPaymentMetaDao {
	return &SysPaymentMetaDao{
		group:   "default",
		table:   "sys_payment_meta",
		columns: sysPaymentMetaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPaymentMetaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPaymentMetaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysPaymentMetaDao) Columns() SysPaymentMetaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPaymentMetaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPaymentMetaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysPaymentMetaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
