package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
)

var Booking = bookingController{}

type bookingController struct {
	BaseController
}

// List 代理商列u表
func (c *bookingController) List(ctx context.Context, req *system.BookingListReq) (res *system.BookingListRes, err error) {
	var (
		total       interface{}
		bookingList []*model.SysBooking
	)

	res = new(system.BookingListRes)

	total, bookingList, err = service.SysBooking().List(ctx, req)
	res.Total = total
	res.Bookings = bookingList

	return
}

// Get 获取代理商信息
func (c *bookingController) Get(ctx context.Context, req *system.BookingGetReq) (res *system.BookingGetRes, err error) {
	res = new(system.BookingGetRes)
	res.Booking, err = service.SysBooking().Get(ctx, req.Id)
	return
}

func (c *bookingController) Checkout(ctx context.Context, req *system.CheckOutReq) (res *system.CheckOutRes, err error) {
	res = new(system.CheckOutRes)
	res, err = service.SysBooking().Checkout(ctx, req)
	return
}
