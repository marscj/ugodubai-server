// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductPriceDao is the data access object for table sys_product_price.
type SysProductPriceDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SysProductPriceColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductPriceColumns defines and stores column names for table sys_product_price.
type SysProductPriceColumns struct {
	Id           string //
	ProductId    string //
	VariationId  string //
	AgentId      string //
	StartDate    string //
	EndDate      string //
	CostPrice    string // 成本价
	SellingPrice string // 销售价
	SpecialPrice string // 预付价格
	Currency     string // 货币
	Status       string // 状态 0.下线 1.上线
	Limit        string // 状态 0.使用库存 1.无限
	Stock        string //
}

// sysProductPriceColumns holds the columns for table sys_product_price.
var sysProductPriceColumns = SysProductPriceColumns{
	Id:           "id",
	ProductId:    "product_id",
	VariationId:  "variation_id",
	AgentId:      "agent_id",
	StartDate:    "start_date",
	EndDate:      "end_date",
	CostPrice:    "cost_price",
	SellingPrice: "selling_price",
	SpecialPrice: "special_price",
	Currency:     "currency",
	Status:       "status",
	Limit:        "limit",
	Stock:        "stock",
}

// NewSysProductPriceDao creates and returns a new DAO object for table data access.
func NewSysProductPriceDao() *SysProductPriceDao {
	return &SysProductPriceDao{
		group:   "default",
		table:   "sys_product_price",
		columns: sysProductPriceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductPriceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductPriceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductPriceDao) Columns() SysProductPriceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductPriceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductPriceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductPriceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
