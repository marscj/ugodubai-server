// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductVariationMetaDao is the data access object for table sys_product_variation_meta.
type SysProductVariationMetaDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of current DAO.
	columns SysProductVariationMetaColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductVariationMetaColumns defines and stores column names for table sys_product_variation_meta.
type SysProductVariationMetaColumns struct {
	MetaId      string //
	VariationId string //
	MetaKey     string //
	MetaValue   string //
}

// sysProductVariationMetaColumns holds the columns for table sys_product_variation_meta.
var sysProductVariationMetaColumns = SysProductVariationMetaColumns{
	MetaId:      "meta_id",
	VariationId: "variation_id",
	MetaKey:     "meta_key",
	MetaValue:   "meta_value",
}

// NewSysProductVariationMetaDao creates and returns a new DAO object for table data access.
func NewSysProductVariationMetaDao() *SysProductVariationMetaDao {
	return &SysProductVariationMetaDao{
		group:   "default",
		table:   "sys_product_variation_meta",
		columns: sysProductVariationMetaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductVariationMetaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductVariationMetaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductVariationMetaDao) Columns() SysProductVariationMetaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductVariationMetaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductVariationMetaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductVariationMetaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
