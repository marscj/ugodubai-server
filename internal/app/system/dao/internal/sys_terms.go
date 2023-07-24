// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTermsDao is the data access object for table sys_terms.
type SysTermsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SysTermsColumns // columns contains all the column names of Table for convenient usage.
}

// SysTermsColumns defines and stores column names for table sys_terms.
type SysTermsColumns struct {
	TermId    string //
	Taxonomy  string //
	Parent    string //
	Name      string //
	Slug      string //
	TermOrder string //
}

// sysTermsColumns holds the columns for table sys_terms.
var sysTermsColumns = SysTermsColumns{
	TermId:    "term_id",
	Taxonomy:  "taxonomy",
	Parent:    "parent",
	Name:      "name",
	Slug:      "slug",
	TermOrder: "term_order",
}

// NewSysTermsDao creates and returns a new DAO object for table data access.
func NewSysTermsDao() *SysTermsDao {
	return &SysTermsDao{
		group:   "default",
		table:   "sys_terms",
		columns: sysTermsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysTermsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysTermsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysTermsDao) Columns() SysTermsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysTermsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysTermsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysTermsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
