package sysProduct

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/consts"
	"ugodubai-server/internal/app/system/dao"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
	"ugodubai-server/library/liberr"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysProduct(New())
}

func New() *sSysProduct {
	return &sSysProduct{}
}

type sSysProduct struct {
}

// List 产品列表
func (s *sSysProduct) List(ctx context.Context, req *system.ProductListReq) (total interface{}, productList []*model.SysProduct, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysProduct.Ctx(ctx)
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "产品列表获取失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		user := service.Context().GetLoginUser(ctx)
		print(user.IsAdmin)

		if user.IsAdmin == 1 {
			err = m.Page(req.PageNum, req.PageSize).Order("id asc").WithAll().Scan(&productList)
			liberr.ErrIsNil(ctx, err, "产品列表获取失败")
		} else {
			err = m.Page(req.PageNum, req.PageSize).Order("id asc").Scan(&productList)
			liberr.ErrIsNil(ctx, err, "产品列表获取失败")
		}
	})
	return
}

//  通过Id获取产品信息
func (s *sSysProduct) Get(ctx context.Context, id uint64) (product *model.SysProduct, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysProduct.Ctx(ctx).WherePri(id).WithAll().Scan(&product)

		if err != nil {
			liberr.ErrIsNil(ctx, err, "产品信息获取失败")
		}
	})
	return
}