// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrderDao is the data access object for table sys_order.
type SysOrderDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SysOrderColumns // columns contains all the column names of Table for convenient usage.
}

// SysOrderColumns defines and stores column names for table sys_order.
type SysOrderColumns struct {
	Id            string // 唯一标识符
	Uuid          string // UUID
	RelatedId     string // 关联订单ID
	FitNumber     string // 订单号
	Sku           string // SKU
	ActionDate    string // 执行日期
	ActionTime    string // 执行时间
	GuestName     string // 客人姓名
	GuestContact  string // 客人联系方式
	AgentId       string // 代理商ID
	AgentCode     string // 代理商代码
	ProductId     string // 代理商ID
	ProductName   string // 产品名称
	UnitPrice     string // 单价
	Quantity      string // 数量
	TotalPrice    string // 总价
	OrderStatus   string // 订单状态
	PaymentStatus string // 支付状态
	Remark        string // 备注
	Currency      string // 货币
	CreatedBy     string // 创建者
	UpdatedBy     string // 更新者
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// sysOrderColumns holds the columns for table sys_order.
var sysOrderColumns = SysOrderColumns{
	Id:            "id",
	Uuid:          "uuid",
	RelatedId:     "related_id",
	FitNumber:     "fit_number",
	Sku:           "sku",
	ActionDate:    "action_date",
	ActionTime:    "action_time",
	GuestName:     "guest_name",
	GuestContact:  "guest_contact",
	AgentId:       "agent_id",
	AgentCode:     "agent_code",
	ProductId:     "product_id",
	ProductName:   "product_name",
	UnitPrice:     "unit_price",
	Quantity:      "quantity",
	TotalPrice:    "total_price",
	OrderStatus:   "order_status",
	PaymentStatus: "payment_status",
	Remark:        "remark",
	Currency:      "currency",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSysOrderDao creates and returns a new DAO object for table data access.
func NewSysOrderDao() *SysOrderDao {
	return &SysOrderDao{
		group:   "default",
		table:   "sys_order",
		columns: sysOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysOrderDao) Columns() SysOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
