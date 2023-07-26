// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductAttributeDao is the data access object for table sys_product_attribute.
type SysProductAttributeDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns SysProductAttributeColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductAttributeColumns defines and stores column names for table sys_product_attribute.
type SysProductAttributeColumns struct {
	AttributeId      string //
	AttributeKey     string //
	AttributeValueEn string //
	AttributeValueCn string //
}

// sysProductAttributeColumns holds the columns for table sys_product_attribute.
var sysProductAttributeColumns = SysProductAttributeColumns{
	AttributeId:      "attribute_id",
	AttributeKey:     "attribute_key",
	AttributeValueEn: "attribute_value_en",
	AttributeValueCn: "attribute_value_cn",
}

// NewSysProductAttributeDao creates and returns a new DAO object for table data access.
func NewSysProductAttributeDao() *SysProductAttributeDao {
	return &SysProductAttributeDao{
		group:   "default",
		table:   "sys_product_attribute",
		columns: sysProductAttributeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductAttributeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductAttributeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductAttributeDao) Columns() SysProductAttributeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductAttributeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductAttributeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductAttributeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
