package system

import (
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

type CheckOutReq struct {
	g.Meta    `path:"/checkout" tags:"订单管理" method:"post" summary:"结算"`
	GatewayId string `p:"gatewayId" v:"required#请选择支付方式"`
}

type CheckOutRes struct {
	g.Meta `mime:"application/json"`
}
