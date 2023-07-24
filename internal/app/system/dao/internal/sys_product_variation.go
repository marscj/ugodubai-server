// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductVariationDao is the data access object for table sys_product_variation.
type SysProductVariationDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns SysProductVariationColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductVariationColumns defines and stores column names for table sys_product_variation.
type SysProductVariationColumns struct {
	VariationId string //
	ProductId   string //
	Name        string // 名称
	Sku         string // SKU
	Status      string // 状态 0.下线 1.上线
	Limit       string // 状态 0.无限 1.使用库存
	Stock       string //
}

// sysProductVariationColumns holds the columns for table sys_product_variation.
var sysProductVariationColumns = SysProductVariationColumns{
	VariationId: "variation_id",
	ProductId:   "product_id",
	Name:        "name",
	Sku:         "sku",
	Status:      "status",
	Limit:       "limit",
	Stock:       "stock",
}

// NewSysProductVariationDao creates and returns a new DAO object for table data access.
func NewSysProductVariationDao() *SysProductVariationDao {
	return &SysProductVariationDao{
		group:   "default",
		table:   "sys_product_variation",
		columns: sysProductVariationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductVariationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductVariationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductVariationDao) Columns() SysProductVariationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductVariationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductVariationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductVariationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
