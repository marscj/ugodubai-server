package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
)

var Order = orderController{}

type orderController struct {
	BaseController
}

// List 代理商列u表
func (c *orderController) List(ctx context.Context, req *system.OrderListReq) (res *system.OrderListRes, err error) {
	var (
		total     interface{}
		orderList []*model.SysOrder
	)

	res = new(system.OrderListRes)

	total, orderList, err = service.SysOrder().List(ctx, req)
	res.Total = total
	res.Orders = orderList

	return
}

// Get 获取代理商信息
func (c *orderController) Get(ctx context.Context, req *system.OrderGetReq) (res *system.OrderGetRes, err error) {
	res = new(system.OrderGetRes)
	res.Order, err = service.SysOrder().Get(ctx, req.Id)
	return
}
