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
	ActionDate    *gtime.Time // 执行日期
	ActionTime    *gtime.Time // 执行时间
	AgentId       interface{} // 代理商ID
	AgentCode     interface{} // 代理商代码
	ProductName   interface{} // 产品名称
	GuestName     interface{} // 客人姓名
	GuestContact  interface{} // 客人联系方式
	UnitPrice     interface{} // 单价
	Quantity      interface{} // 数量
	TotalPrice    interface{} // 总价
	OrderStatus   interface{} // 订单状态
	PaymentStatus interface{} // 支付状态
	Remark        interface{} // 备注
	Currency      interface{} // 货币
	CreatedBy     interface{} // 创建者
	UpdatedBy     interface{} // 更新者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
