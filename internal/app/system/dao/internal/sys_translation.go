// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTranslationDao is the data access object for table sys_translation.
type SysTranslationDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysTranslationColumns // columns contains all the column names of Table for convenient usage.
}

// SysTranslationColumns defines and stores column names for table sys_translation.
type SysTranslationColumns struct {
	Id           string // 唯一标识符
	RecordType   string //
	RecordId     string //
	LanguageCode string //
	MetaKey      string //
	MetaValue    string //
}

// sysTranslationColumns holds the columns for table sys_translation.
var sysTranslationColumns = SysTranslationColumns{
	Id:           "id",
	RecordType:   "record_type",
	RecordId:     "record_id",
	LanguageCode: "language_code",
	MetaKey:      "meta_key",
	MetaValue:    "meta_value",
}

// NewSysTranslationDao creates and returns a new DAO object for table data access.
func NewSysTranslationDao() *SysTranslationDao {
	return &SysTranslationDao{
		group:   "default",
		table:   "sys_translation",
		columns: sysTranslationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysTranslationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysTranslationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysTranslationDao) Columns() SysTranslationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysTranslationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysTranslationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysTranslationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
