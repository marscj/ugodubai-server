// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrder is the golang structure for table sys_order.
type SysOrder struct {
	Id            uint64      `json:"id"            description:"唯一标识符"`
	Uuid          string      `json:"uuid"          description:"UUID"`
	RelatedId     string      `json:"relatedId"     description:"关联订单ID"`
	FitNumber     string      `json:"fitNumber"     description:"订单号"`
	Sku           string      `json:"sku"           description:"SKU"`
	ActionDate    *gtime.Time `json:"actionDate"    description:"执行日期"`
	ActionTime    *gtime.Time `json:"actionTime"    description:"执行时间"`
	GuestName     string      `json:"guestName"     description:"客人姓名"`
	GuestContact  string      `json:"guestContact"  description:"客人联系方式"`
	AgentId       uint64      `json:"agentId"       description:"代理商ID"`
	AgentCode     string      `json:"agentCode"     description:"代理商代码"`
	ProductId     uint64      `json:"productId"     description:"产品ID"`
	ProductName   string      `json:"productName"   description:"产品名称"`
	UnitPrice     float64     `json:"unitPrice"     description:"单价"`
	Quantity      int         `json:"quantity"      description:"数量"`
	TotalPrice    float64     `json:"totalPrice"    description:"总价"`
	OrderStatus   int         `json:"orderStatus"   description:"订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消"`
	PaymentStatus int         `json:"paymentStatus" description:"支付状态 0.等待支付 1.未支付 2.已支付 3.已退款"`
	PaymentMethod int         `json:"paymentMethod" description:"支付方式 0.余额 1.信用 2.支付宝 3.微信 4.公司转账"`
	Remark        string      `json:"remark"        description:"备注"`
	Currency      string      `json:"currency"      description:"货币"`
	CreatedBy     uint64      `json:"createdBy"     description:"创建者"`
	UpdatedBy     uint64      `json:"updatedBy"     description:"更新者"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
}
