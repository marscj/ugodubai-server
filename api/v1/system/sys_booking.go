package system

import (
	"time"
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type BookingListReq struct {
	g.Meta `path:"/booking/list" tags:"订单管理" method:"get" summary:"订单列表"`
	commonApi.PageReq
}

type BookingListRes struct {
	g.Meta   `mime:"application/json"`
	Bookings []*model.SysBooking `json:"order"`
	commonApi.ListRes
}

type BookingGetReq struct {
	g.Meta `path:"/booking/get" tags:"订单管理" method:"get" summary:"获取订单商信息"`
	Id     uint64 `p:"id" v:"required#订单id不能为空"`
}

type BookingGetRes struct {
	g.Meta  `mime:"application/json"`
	Booking *model.SysBooking `json:"order"`
}

// 检查
type PreCheckOutReq struct {
	g.Meta `path:"/precheckout" tags:"订单管理" method:"post" summary:"结算"`
	Item   []*PreCheckOutItem `v:"required#请选择购买的产品"`
}

type PreCheckOutRes struct {
}

// 结算
type CheckOutReq struct {
	g.Meta    `path:"/checkout" tags:"订单管理" method:"post" summary:"结算"`
	GatewayId uint64          `v:"required#请选择支付方式"`
	Item      []*CheckOutItem `v:"required#请选择购买的产品"`
}

type CheckOutRes struct {
	g.Meta `mime:"application/json"`
}

type PreCheckOutItem struct {
	Quantity         uint   `v:"required#请输入产品数量"`
	VariationPriceId uint64 `v:"required#请选择产品项"`
	ActionDate       string `v:"required|date-format:Y-m-d H:i#请输入执行日期|请输入正确的日期，如2020-01-01 23:00"`
	Price            string
	Tax              string
	Timestamp        time.Time `json:"timestamp"`
	Signature        string    `dc:"签名"`
}

type CheckOutItem struct {
	Quantity         uint   `v:"required#请输入产品数量"`
	VariationPriceId uint64 `v:"required#请选择产品项"`
	ContactName      string `v:"required|length:2,100#请输入客户姓名|客户姓名长度为2-100位"`
	ContactPhone     string `v:"max-length:32#最大32位"`
	Currency         string `v:"max-length:3#最大3位"`
	Remark           string `v:"max-length:255#最大255位"`
	RelatedId        string `v:"max-length:64#最大64位"`
	ActionDate       string `v:"required|date-format:Y-m-d H:i#请输入执行日期|请输入正确的日期，如2020-01-01 23:00"`
}
