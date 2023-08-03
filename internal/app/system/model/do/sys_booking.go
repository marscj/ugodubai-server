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
	g.Meta           `orm:"table:sys_booking, do:true"`
	BookingId        interface{} // 代理商ID
	ParentId         interface{} // 父
	RelatedId        interface{} // 关联订单ID
	AgentId          interface{} // 代理商ID
	SupplierId       interface{} // 代理商ID
	ActionDate       *gtime.Time // 执行日期
	ProductId        interface{} // 产品ID
	VariationId      interface{} // 变体产品
	VariationPriceId interface{} // 变体产品价格
	FitNumber        interface{} // FIT
	Sku              interface{} // SKU
	ContactName      interface{} // 联系人姓名
	ContactPhone     interface{} // 联系方式
	UnitPrice        interface{} // 单价
	Quantity         interface{} // 数量
	TotalPrice       interface{} // 总价
	Tax              interface{} // 税
	Currency         interface{} // 货币
	BookingStatus    interface{} // 订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消
	PaymentStatus    interface{} // 支付状态 0.等待支付 1.未支付 2.已支付 3.已退款
	PaymentMethod    interface{} // 支付方式 0.余额 1.信用 2.支付宝 3.微信 4.公司转账
	SupplierStatus   interface{} // 提供商支付状态 0.未支付 1.已支付 2.已退款
	Remark           interface{} // 备注
	CreatedBy        interface{} // 创建者
	UpdatedBy        interface{} // 更新者
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time // 更新时间
}
