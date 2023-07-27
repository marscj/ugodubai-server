// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysVariationLookupDao is the data access object for table sys_variation_lookup.
type SysVariationLookupDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns SysVariationLookupColumns // columns contains all the column names of Table for convenient usage.
}

// SysVariationLookupColumns defines and stores column names for table sys_variation_lookup.
type SysVariationLookupColumns struct {
	VariationLookupId string //
	VariationId       string //
	AttributeId       string //
	VariationPriceId  string //
	AgentId           string //
}

// sysVariationLookupColumns holds the columns for table sys_variation_lookup.
var sysVariationLookupColumns = SysVariationLookupColumns{
	VariationLookupId: "variation_lookup_id",
	VariationId:       "variation_id",
	AttributeId:       "attribute_id",
	VariationPriceId:  "variation_price_id",
	AgentId:           "agent_id",
}

// NewSysVariationLookupDao creates and returns a new DAO object for table data access.
func NewSysVariationLookupDao() *SysVariationLookupDao {
	return &SysVariationLookupDao{
		group:   "default",
		table:   "sys_variation_lookup",
		columns: sysVariationLookupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysVariationLookupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysVariationLookupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysVariationLookupDao) Columns() SysVariationLookupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysVariationLookupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysVariationLookupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysVariationLookupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
