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
	ProductId     uint64      `json:"productId"     description:"代理商ID"`
	ProductName   string      `json:"productName"   description:"产品名称"`
	UnitPrice     float64     `json:"unitPrice"     description:"单价"`
	Quantity      int         `json:"quantity"      description:"数量"`
	TotalPrice    float64     `json:"totalPrice"    description:"总价"`
	OrderStatus   int         `json:"orderStatus"   description:"订单状态"`
	PaymentStatus int         `json:"paymentStatus" description:"支付状态"`
	Remark        string      `json:"remark"        description:"备注"`
	Currency      string      `json:"currency"      description:"货币"`
	CreatedBy     uint64      `json:"createdBy"     description:"创建者"`
	UpdatedBy     uint64      `json:"updatedBy"     description:"更新者"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
}
