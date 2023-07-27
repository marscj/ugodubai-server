// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductLookupDao is the data access object for table sys_product_lookup.
type SysProductLookupDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns SysProductLookupColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductLookupColumns defines and stores column names for table sys_product_lookup.
type SysProductLookupColumns struct {
	Id        string //
	ProductId string //
	TermId    string //
}

// sysProductLookupColumns holds the columns for table sys_product_lookup.
var sysProductLookupColumns = SysProductLookupColumns{
	Id:        "id",
	ProductId: "product_id",
	TermId:    "term_id",
}

// NewSysProductLookupDao creates and returns a new DAO object for table data access.
func NewSysProductLookupDao() *SysProductLookupDao {
	return &SysProductLookupDao{
		group:   "default",
		table:   "sys_product_lookup",
		columns: sysProductLookupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductLookupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductLookupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductLookupDao) Columns() SysProductLookupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductLookupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductLookupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductLookupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
