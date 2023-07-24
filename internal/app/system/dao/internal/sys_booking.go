// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysBookingDao is the data access object for table sys_booking.
type SysBookingDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SysBookingColumns // columns contains all the column names of Table for convenient usage.
}

// SysBookingColumns defines and stores column names for table sys_booking.
type SysBookingColumns struct {
	Id            string // 代理商ID
	ParentId      string // 父
	RelatedId     string // 关联订单ID
	AgentId       string // 代理商ID
	ActionDate    string // 执行日期
	ProductId     string // 产品ID
	VariationId   string // 变体产品
	FitNumber     string // FIT
	Sku           string // SKU
	GuestName     string // 客人姓名
	GuestContact  string // 客人联系方式
	ProductName   string // 产品名称
	UnitPrice     string // 单价
	Quantity      string // 数量
	TotalPrice    string // 总价
	Tax           string // 税
	Currency      string // 货币
	BookingStatus string // 订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消
	Remark        string // 备注
	CreatedBy     string // 创建者
	UpdatedBy     string // 更新者
	CreatedAt     string //
	UpdatedAt     string //
}

// sysBookingColumns holds the columns for table sys_booking.
var sysBookingColumns = SysBookingColumns{
	Id:            "id",
	ParentId:      "parent_id",
	RelatedId:     "related_id",
	AgentId:       "agent_id",
	ActionDate:    "action_date",
	ProductId:     "product_id",
	VariationId:   "variation_id",
	FitNumber:     "fit_number",
	Sku:           "sku",
	GuestName:     "guest_name",
	GuestContact:  "guest_contact",
	ProductName:   "product_name",
	UnitPrice:     "unit_price",
	Quantity:      "quantity",
	TotalPrice:    "total_price",
	Tax:           "tax",
	Currency:      "currency",
	BookingStatus: "booking_status",
	Remark:        "remark",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSysBookingDao creates and returns a new DAO object for table data access.
func NewSysBookingDao() *SysBookingDao {
	return &SysBookingDao{
		group:   "default",
		table:   "sys_booking",
		columns: sysBookingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysBookingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysBookingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysBookingDao) Columns() SysBookingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysBookingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysBookingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysBookingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
