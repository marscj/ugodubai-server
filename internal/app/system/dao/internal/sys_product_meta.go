// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductMetaDao is the data access object for table sys_product_meta.
type SysProductMetaDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysProductMetaColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductMetaColumns defines and stores column names for table sys_product_meta.
type SysProductMetaColumns struct {
	MetaId    string //
	ProductId string //
	MetaKey   string //
	MetaValue string //
}

// sysProductMetaColumns holds the columns for table sys_product_meta.
var sysProductMetaColumns = SysProductMetaColumns{
	MetaId:    "meta_id",
	ProductId: "product_id",
	MetaKey:   "meta_key",
	MetaValue: "meta_value",
}

// NewSysProductMetaDao creates and returns a new DAO object for table data access.
func NewSysProductMetaDao() *SysProductMetaDao {
	return &SysProductMetaDao{
		group:   "default",
		table:   "sys_product_meta",
		columns: sysProductMetaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductMetaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductMetaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductMetaDao) Columns() SysProductMetaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductMetaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductMetaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductMetaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
