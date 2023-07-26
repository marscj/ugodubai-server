package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
)

type ISysProduct interface {
	List(ctx context.Context, req *system.ProductListReq) (total interface{}, productList []*model.SysProduct, err error)
	Get(ctx context.Context, id uint64) (product *model.SysProduct, err error)
}

var localSysProduct ISysProduct

func SysProduct() ISysProduct {
	if localSysProduct == nil {
		panic("implement not found for interface ISysProduct, forgot register?")
	}
	return localSysProduct
}

func RegisterSysProduct(i ISysProduct) {
	localSysProduct = i
}
