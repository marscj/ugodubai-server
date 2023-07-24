// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysBooking is the golang structure for table sys_booking.
type SysBooking struct {
	Id            uint64      `json:"id"            description:"代理商ID"`
	ParentId      uint64      `json:"parentId"      description:"父"`
	RelatedId     string      `json:"relatedId"     description:"关联订单ID"`
	AgentId       uint64      `json:"agentId"       description:"代理商ID"`
	ActionDate    *gtime.Time `json:"actionDate"    description:"执行日期"`
	ProductId     uint64      `json:"productId"     description:"产品ID"`
	VariationId   uint64      `json:"variationId"   description:"变体产品"`
	FitNumber     string      `json:"fitNumber"     description:"FIT"`
	Sku           string      `json:"sku"           description:"SKU"`
	GuestName     string      `json:"guestName"     description:"客人姓名"`
	GuestContact  string      `json:"guestContact"  description:"客人联系方式"`
	ProductName   string      `json:"productName"   description:"产品名称"`
	UnitPrice     float64     `json:"unitPrice"     description:"单价"`
	Quantity      int         `json:"quantity"      description:"数量"`
	TotalPrice    float64     `json:"totalPrice"    description:"总价"`
	Tax           float64     `json:"tax"           description:"税"`
	Currency      string      `json:"currency"      description:"货币"`
	BookingStatus int         `json:"bookingStatus" description:"订单状态 0.待核单 1.已核单出票中 2.已出票 3.取消待确认 4.已取消"`
	Remark        string      `json:"remark"        description:"备注"`
	CreatedBy     uint64      `json:"createdBy"     description:"创建者"`
	UpdatedBy     uint64      `json:"updatedBy"     description:"更新者"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
}
