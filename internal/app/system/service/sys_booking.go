package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
)

type ISysBooking interface {
	List(ctx context.Context, req *system.BookingListReq) (total interface{}, bookingList []*model.SysBooking, err error)
	Get(ctx context.Context, id uint64) (booking *model.SysBooking, err error)
	PreCheckout(ctx context.Context, req *system.PreCheckOutReq) (subTotal string, subTax string, Total string, items []*model.PreCheckOutItem, err error)
	Checkout(ctx context.Context, req *system.CheckOutReq) (err error)
}

var localSysBooking ISysBooking

func SysBooking() ISysBooking {
	if localSysBooking == nil {
		panic("implement not found for interface ISysBooking, forgot register?")
	}
	return localSysBooking
}

func RegisterSysBooking(i ISysBooking) {
	localSysBooking = i
}
