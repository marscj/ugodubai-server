// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductDao is the data access object for table sys_product.
type SysProductDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SysProductColumns // columns contains all the column names of Table for convenient usage.
}

// SysProductColumns defines and stores column names for table sys_product.
type SysProductColumns struct {
	Id            string //
	Sku           string // SKU
	NameEn        string // 名称
	NameCn        string // 名称
	DescriptionEn string // 产品简介
	DescriptionCn string // 产品简介
	ContentEn     string // 产品内容
	ContentCn     string // 产品内容
	Status        string // 状态 0.下线 1.上线
}

// sysProductColumns holds the columns for table sys_product.
var sysProductColumns = SysProductColumns{
	Id:            "id",
	Sku:           "sku",
	NameEn:        "name_en",
	NameCn:        "name_cn",
	DescriptionEn: "description_en",
	DescriptionCn: "description_cn",
	ContentEn:     "content_en",
	ContentCn:     "content_cn",
	Status:        "status",
}

// NewSysProductDao creates and returns a new DAO object for table data access.
func NewSysProductDao() *SysProductDao {
	return &SysProductDao{
		group:   "default",
		table:   "sys_product",
		columns: sysProductColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProductDao) Columns() SysProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
