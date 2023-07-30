// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductPriceLookupDao is the data access object for table sys_product_price_lookup.
type SysProductPriceLookupDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns SysProductPriceLookupColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductPriceLookupColumns defines and stores column names for table sys_product_price_lookup.
type SysProductPriceLookupColumns struct {
	Id        string //
	ProductId string //
	AgentId   string //
	MinPrice  string // 最低价格
	MaxPrice  string // 最高价格
}

// sysProductPriceLookupColumns holds the columns for table sys_product_price_lookup.
var sysProductPriceLookupColumns = SysProductPriceLookupColumns{
	Id:        "id",
	ProductId: "product_id",
	AgentId:   "agent_id",
	MinPrice:  "min_price",
	MaxPrice:  "max_price",
}

// NewSysProductPriceLookupDao creates and returns a new DAO object for table data access.
func NewSysProductPriceLookupDao() *SysProductPriceLookupDao {
	return &SysProductPriceLookupDao{
		group:   "default",
		table:   "sys_product_price_lookup",
		columns: sysProductPriceLookupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductPriceLookupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductPriceLookupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductPriceLookupDao) Columns() SysProductPriceLookupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductPriceLookupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductPriceLookupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductPriceLookupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
