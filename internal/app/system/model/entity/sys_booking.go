// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysBooking is the golang structure for table sys_booking.
type SysBooking struct {
	BookingId        uint64      `json:"bookingId"        description:"代理商ID"`
	ParentId         uint64      `json:"parentId"         description:"父"`
	RelatedId        string      `json:"relatedId"        description:"关联订单ID"`
	AgentId          uint64      `json:"agentId"          description:"代理商ID"`
	SupplierId       uint64      `json:"supplierId"       description:"代理商ID"`
	ActionDate       *gtime.Time `json:"actionDate"       description:"执行日期"`
	ProductId        uint64      `json:"productId"        description:"产品ID"`
	VariationId      uint64      `json:"variationId"      description:"变体产品"`
	VariationPriceId uint64      `json:"variationPriceId" description:"变体产品价格"`
	FitNumber        string      `json:"fitNumber"        description:"FIT"`
	Sku              string      `json:"sku"              description:"SKU"`
	ContactName      string      `json:"contactName"      description:"联系人姓名"`
	ContactPhone     string      `json:"contactPhone"     description:"联系方式"`
	UnitPrice        float64     `json:"unitPrice"        description:"单价"`
	Quantity         int         `json:"quantity"         description:"数量"`
	TotalPrice       float64     `json:"totalPrice"       description:"总价"`
	Tax              float64     `json:"tax"              description:"税"`
	Currency         string      `json:"currency"         description:"货币"`
	BookingStatus    int         `json:"bookingStatus"    description:"订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消"`
	PaymentStatus    int         `json:"paymentStatus"    description:"支付状态 0.等待支付 1.未支付 2.已支付 3.已退款"`
	PaymentMethod    int         `json:"paymentMethod"    description:"支付方式 0.余额 1.信用 2.支付宝 3.微信 4.公司转账"`
	SupplierStatus   int         `json:"supplierStatus"   description:"提供商支付状态 1.未支付 2.已支付 3.已退款"`
	Remark           string      `json:"remark"           description:"备注"`
	CreatedBy        uint64      `json:"createdBy"        description:"创建者"`
	UpdatedBy        uint64      `json:"updatedBy"        description:"更新者"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:"更新时间"`
}
