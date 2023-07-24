// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPaymentTokenmetaDao is the data access object for table sys_payment_tokenmeta.
type SysPaymentTokenmetaDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns SysPaymentTokenmetaColumns // columns contains all the column names of Table for convenient usage.
}

// SysPaymentTokenmetaColumns defines and stores column names for table sys_payment_tokenmeta.
type SysPaymentTokenmetaColumns struct {
	MetaId         string //
	PaymentTokenId string //
	MetaKey        string //
	MetaValue      string //
}

// sysPaymentTokenmetaColumns holds the columns for table sys_payment_tokenmeta.
var sysPaymentTokenmetaColumns = SysPaymentTokenmetaColumns{
	MetaId:         "meta_id",
	PaymentTokenId: "payment_token_id",
	MetaKey:        "meta_key",
	MetaValue:      "meta_value",
}

// NewSysPaymentTokenmetaDao creates and returns a new DAO object for table data access.
func NewSysPaymentTokenmetaDao() *SysPaymentTokenmetaDao {
	return &SysPaymentTokenmetaDao{
		group:   "default",
		table:   "sys_payment_tokenmeta",
		columns: sysPaymentTokenmetaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPaymentTokenmetaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPaymentTokenmetaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysPaymentTokenmetaDao) Columns() SysPaymentTokenmetaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPaymentTokenmetaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPaymentTokenmetaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysPaymentTokenmetaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
