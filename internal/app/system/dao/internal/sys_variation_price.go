// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysVariationPriceDao is the data access object for table sys_variation_price.
type SysVariationPriceDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns SysVariationPriceColumns // columns contains all the column names of Table for convenient usage.
}

// SysVariationPriceColumns defines and stores column names for table sys_variation_price.
type SysVariationPriceColumns struct {
	VariationPriceId string //
	VariationId      string //
	AttributeId      string //
	TimeId           string //
	AgentId          string //
	StartDate        string //
	EndDate          string //
	CostPrice        string // 成本价
	SpecialPrice     string // 预付价格
	SellingPrice     string // 销售价
	Currency         string // 货币
	Stock            string //
	UpdatedAt        string // 更新时间
}

// sysVariationPriceColumns holds the columns for table sys_variation_price.
var sysVariationPriceColumns = SysVariationPriceColumns{
	VariationPriceId: "variation_price_id",
	VariationId:      "variation_id",
	AttributeId:      "attribute_id",
	TimeId:           "time_id",
	AgentId:          "agent_id",
	StartDate:        "start_date",
	EndDate:          "end_date",
	CostPrice:        "cost_price",
	SpecialPrice:     "special_price",
	SellingPrice:     "selling_price",
	Currency:         "currency",
	Stock:            "stock",
	UpdatedAt:        "updated_at",
}

// NewSysVariationPriceDao creates and returns a new DAO object for table data access.
func NewSysVariationPriceDao() *SysVariationPriceDao {
	return &SysVariationPriceDao{
		group:   "default",
		table:   "sys_variation_price",
		columns: sysVariationPriceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysVariationPriceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysVariationPriceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysVariationPriceDao) Columns() SysVariationPriceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysVariationPriceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysVariationPriceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysVariationPriceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
