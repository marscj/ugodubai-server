package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
)

type ISysOrder interface {
	List(ctx context.Context, req *system.OrderListReq) (total interface{}, orderList []*model.SysOrder, err error)
	Get(ctx context.Context, id uint64) (order *model.SysOrder, err error)
}

var localSysOrder ISysOrder

func SysOrder() ISysOrder {
	if localSysOrder == nil {
		panic("implement not found for interface ISysOrder, forgot register?")
	}
	return localSysOrder
}

func RegisterSysOrder(i ISysOrder) {
	localSysOrder = i
}
