// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysVariationAttributeDao is the data access object for table sys_variation_attribute.
type SysVariationAttributeDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns SysVariationAttributeColumns // columns contains all the column names of Table for convenient usage.
}

// SysVariationAttributeColumns defines and stores column names for table sys_variation_attribute.
type SysVariationAttributeColumns struct {
	AttributeId      string //
	AttributeKey     string //
	AttributeValueEn string //
	AttributeValueCn string //
}

// sysVariationAttributeColumns holds the columns for table sys_variation_attribute.
var sysVariationAttributeColumns = SysVariationAttributeColumns{
	AttributeId:      "attribute_id",
	AttributeKey:     "attribute_key",
	AttributeValueEn: "attribute_value_en",
	AttributeValueCn: "attribute_value_cn",
}

// NewSysVariationAttributeDao creates and returns a new DAO object for table data access.
func NewSysVariationAttributeDao() *SysVariationAttributeDao {
	return &SysVariationAttributeDao{
		group:   "default",
		table:   "sys_variation_attribute",
		columns: sysVariationAttributeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysVariationAttributeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysVariationAttributeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysVariationAttributeDao) Columns() SysVariationAttributeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysVariationAttributeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysVariationAttributeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysVariationAttributeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
