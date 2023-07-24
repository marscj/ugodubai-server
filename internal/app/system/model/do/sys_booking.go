// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysBooking is the golang structure of table sys_booking for DAO operations like Where/Data.
type SysBooking struct {
	g.Meta        `orm:"table:sys_booking, do:true"`
	Id            interface{} // 代理商ID
	ParentId      interface{} // 父
	RelatedId     interface{} // 关联订单ID
	AgentId       interface{} // 代理商ID
	ActionDate    *gtime.Time // 执行日期
	ProductId     interface{} // 产品ID
	VariationId   interface{} // 变体产品
	FitNumber     interface{} // FIT
	Sku           interface{} // SKU
	GuestName     interface{} // 客人姓名
	GuestContact  interface{} // 客人联系方式
	ProductName   interface{} // 产品名称
	UnitPrice     interface{} // 单价
	Quantity      interface{} // 数量
	TotalPrice    interface{} // 总价
	Tax           interface{} // 税
	Currency      interface{} // 货币
	BookingStatus interface{} // 订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消
	Remark        interface{} // 备注
	CreatedBy     interface{} // 创建者
	UpdatedBy     interface{} // 更新者
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
