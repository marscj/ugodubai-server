// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysVariationDao is the data access object for table sys_variation.
type SysVariationDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysVariationColumns // columns contains all the column names of Table for convenient usage.
}

// SysVariationColumns defines and stores column names for table sys_variation.
type SysVariationColumns struct {
	VariationId string //
	NameEn      string // 默认名称
	NameCn      string // 默认名称
	Sku         string // SKU
	Status      string // 状态 0.下线 1.上线
}

// sysVariationColumns holds the columns for table sys_variation.
var sysVariationColumns = SysVariationColumns{
	VariationId: "variation_id",
	NameEn:      "name_en",
	NameCn:      "name_cn",
	Sku:         "sku",
	Status:      "status",
}

// NewSysVariationDao creates and returns a new DAO object for table data access.
func NewSysVariationDao() *SysVariationDao {
	return &SysVariationDao{
		group:   "default",
		table:   "sys_variation",
		columns: sysVariationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysVariationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysVariationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysVariationDao) Columns() SysVariationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysVariationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysVariationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysVariationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
