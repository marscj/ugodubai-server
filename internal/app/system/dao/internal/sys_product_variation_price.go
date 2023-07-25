// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductVariationPriceDao is the data access object for table sys_product_variation_price.
type SysProductVariationPriceDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns SysProductVariationPriceColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductVariationPriceColumns defines and stores column names for table sys_product_variation_price.
type SysProductVariationPriceColumns struct {
	Id              string //
	VariationMetaId string //
	AgentId         string //
	StartDate       string //
	EndDate         string //
	CostPrice       string // 成本价
	SpecialPrice    string // 预付价格
	SellingPrice    string // 销售价
	Currency        string // 货币
}

// sysProductVariationPriceColumns holds the columns for table sys_product_variation_price.
var sysProductVariationPriceColumns = SysProductVariationPriceColumns{
	Id:              "id",
	VariationMetaId: "variation_meta_id",
	AgentId:         "agent_id",
	StartDate:       "start_date",
	EndDate:         "end_date",
	CostPrice:       "cost_price",
	SpecialPrice:    "special_price",
	SellingPrice:    "selling_price",
	Currency:        "currency",
}

// NewSysProductVariationPriceDao creates and returns a new DAO object for table data access.
func NewSysProductVariationPriceDao() *SysProductVariationPriceDao {
	return &SysProductVariationPriceDao{
		group:   "default",
		table:   "sys_product_variation_price",
		columns: sysProductVariationPriceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductVariationPriceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductVariationPriceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductVariationPriceDao) Columns() SysProductVariationPriceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductVariationPriceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductVariationPriceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductVariationPriceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
