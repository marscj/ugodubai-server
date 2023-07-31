package sysProduct

import (
	"context"
	"strings"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/consts"
	"ugodubai-server/internal/app/system/dao"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
	"ugodubai-server/library/liberr"

	commonService "ugodubai-server/internal/app/common/service"

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
func (s *sSysProduct) List(ctx context.Context, req *system.ProductListReq) (total interface{}, productList []*model.SysProductList, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 获取登录用户
		user := service.Context().GetLoginUser(ctx)
		agentID := user.AgentId

		var whereStr string
		var args []interface{}

		var products []*model.SysProductList
		var priceLookups []*model.SysProductPriceLookup

		// 获取所有的价格信息并存储在缓存中
		priceLookups, err = s.GetAllProductPrices(ctx)

		// 关键字查询
		if req.Keyword != nil {
			// Search for products with the keyword
			whereStr += "`p.product_id` IN (SELECT `product_id` FROM `sys_product_keywords` WHERE `keyword` LIKE ?) "
			args = append(args, "%"+*req.Keyword+"%")
		}

		// terms 查询
		if len(req.TermsIDs) > 0 {
			if whereStr != "" {
				whereStr += " AND "
			}
			placeholders := strings.Trim(strings.Repeat("?,", len(req.TermsIDs)), ",")
			whereStr += "`p.product_id` IN (SELECT `product_id` FROM `sys_product_lookup_terms` WHERE `term_id` IN(" + placeholders + ")) "
			for _, id := range req.TermsIDs {
				args = append(args, id)
			}
		}

		//状态查询
		if req.Status != nil {
			if whereStr != "" {
				whereStr += " AND "
			}
			whereStr += "`p.status` = ? "
			args = append(args, *req.Status)
		}

		m := dao.SysProduct.Ctx(ctx).Where(whereStr, args...)

		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "产品列表获取失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}

		err = m.Page(req.PageNum, req.PageSize).OrderAsc("order").WithAll().Scan(&products)
		liberr.ErrIsNil(ctx, err, "产品列表获取失败")

		// 遍历产品并从缓存中获取其价格信息
		for _, product := range products {
			for _, priceLookup := range priceLookups {
				if priceLookup.ProductId == product.ProductId && priceLookup.AgentId == agentID {
					product.Price = priceLookup
					break
				} else if priceLookup.ProductId == product.ProductId && priceLookup.AgentId == 0 {
					product.Price = priceLookup
				}
			}
		}
		productList = products
	})
	return total, productList, nil
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

// 获取所有产品关联的价格
func (s *sSysProduct) GetAllProductPrices(ctx context.Context) (list []*model.SysProductPriceLookup, err error) {
	cache := commonService.Cache()
	//从缓存获取
	data := cache.Get(ctx, "sys_product_price_lookup_key")
	if !data.IsNil() {
		err = data.Structs(&list)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysProductPriceLookup.Ctx(ctx).OrderAsc("Id").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取产品价格数据出错")
		//缓存
		cache.Set(ctx, "sys_product_price_lookup_key", list, 0, "sys_product_price_lookup_tag")
	})
	return
}
