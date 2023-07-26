package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
)

var Product = productController{}

type productController struct {
	BaseController
}

// List 代理商列u表
func (c *productController) List(ctx context.Context, req *system.ProductListReq) (res *system.ProductListRes, err error) {
	var (
		total       interface{}
		productList []*model.SysProduct
	)

	res = new(system.ProductListRes)

	total, productList, err = service.SysProduct().List(ctx, req)
	res.Total = total
	res.Products = productList

	return
}

// Get 获取代理商信息
func (c *productController) Get(ctx context.Context, req *system.ProductGetReq) (res *system.ProductGetRes, err error) {
	res = new(system.ProductGetRes)
	res.Product, err = service.SysProduct().Get(ctx, req.Id)
	return
}
