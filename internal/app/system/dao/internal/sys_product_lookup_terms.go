// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductLookupTermsDao is the data access object for table sys_product_lookup_terms.
type SysProductLookupTermsDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns SysProductLookupTermsColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductLookupTermsColumns defines and stores column names for table sys_product_lookup_terms.
type SysProductLookupTermsColumns struct {
	Id        string //
	ProductId string //
	TermId    string //
}

// sysProductLookupTermsColumns holds the columns for table sys_product_lookup_terms.
var sysProductLookupTermsColumns = SysProductLookupTermsColumns{
	Id:        "id",
	ProductId: "product_id",
	TermId:    "term_id",
}

// NewSysProductLookupTermsDao creates and returns a new DAO object for table data access.
func NewSysProductLookupTermsDao() *SysProductLookupTermsDao {
	return &SysProductLookupTermsDao{
		group:   "default",
		table:   "sys_product_lookup_terms",
		columns: sysProductLookupTermsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductLookupTermsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductLookupTermsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductLookupTermsDao) Columns() SysProductLookupTermsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductLookupTermsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductLookupTermsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductLookupTermsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
