// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrder is the golang structure of table sys_order for DAO operations like Where/Data.
type SysOrder struct {
	g.Meta        `orm:"table:sys_order, do:true"`
	Id            interface{} // 唯一标识符
	Uuid          interface{} // UUID
	RelatedId     interface{} // 关联订单ID
	FitNumber     interface{} // 订单号
	Sku           interface{} // SKU
	ActionDate    *gtime.Time // 执行日期
	ActionTime    *gtime.Time // 执行时间
	GuestName     interface{} // 客人姓名
	GuestContact  interface{} // 客人联系方式
	AgentId       interface{} // 代理商ID
	AgentCode     interface{} // 代理商代码
	ProductId     interface{} // 产品ID
	ProductName   interface{} // 产品名称
	UnitPrice     interface{} // 单价
	Quantity      interface{} // 数量
	TotalPrice    interface{} // 总价
	OrderStatus   interface{} // 订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消
	PaymentStatus interface{} // 支付状态 0.等待支付 1.未支付 2.已支付 3.已退款
	PaymentMethod interface{} // 支付方式 0.余额 1.信用 2.支付宝 3.微信 4.公司转账
	Remark        interface{} // 备注
	Currency      interface{} // 货币
	CreatedBy     interface{} // 创建者
	UpdatedBy     interface{} // 更新者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
