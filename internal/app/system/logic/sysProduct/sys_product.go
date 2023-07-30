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
func (s *sSysProduct) List(ctx context.Context, req *system.ProductListReq) (total interface{}, productList []*model.SysProductList, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysProduct.Ctx(ctx).Where("status = ?", 1)
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "产品列表获取失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}

		err = m.Page(req.PageNum, req.PageSize).OrderAsc("order").WithAll().Scan(&productList)
		liberr.ErrIsNil(ctx, err, "产品列表获取失败")
	})
	return
}

// // List 产品列表
// func (s *sSysProduct) List(ctx context.Context, req *system.ProductListReq) (total interface{}, productList []*model.SysProductList, err error) {

// 	err = g.Try(ctx, func(ctx context.Context) {
// 		m := dao.SysProduct.Ctx(ctx)
// 		total, err = m.Count()
// 		liberr.ErrIsNil(ctx, err, "产品列表获取失败")

// 		if req.PageNum == 0 {
// 			req.PageNum = 1
// 		}
// 		if req.PageSize == 0 {
// 			req.PageSize = consts.PageSize
// 		}

// 		user := service.Context().GetLoginUser(ctx)

// 		if user.IsAdmin == 0 {
// 			err = m.Page(req.PageNum, req.PageSize).OrderAsc("product_id").WithAll().Scan(&productList)
// 			liberr.ErrIsNil(ctx, err, "产品列表获取失败")
// 		} else {
// 			err = m.Page(req.PageNum, req.PageSize).OrderAsc("product_id").WithAll().Scan(&productList)
// 			liberr.ErrIsNil(ctx, err, "产品列表获取失败")
// 		}

// 		variables := g.Map{
// 			"agentId": 1, // 指定 agent_id
// 			// "agentId": nil, // 不指定 agent_id
// 		}

// 		query := g.Model("sys_product", "p").
// 			Distinct().
// 			Fields("p.product_id, p.name_en, p.name_cn, p.description_en, p.description_cn, p.content_en, p.content_cn, p.status, p.image").
// 			InnerJoin("sys_product_lookup", "vp", "p.product_id = vp.product_id")

// 		if agentId := variables["agentId"]; agentId != nil {
// 			query = query.Where("vp.agent_id = ?", agentId)
// 		}

// 		var products []*model.SysProduct

// 		err = query.Scan(&products)

// 		if err != nil {
// 			panic(err)
// 		}

// 		// 打印查询结果
// 		for _, product := range products {
// 			fmt.Printf("Product ID: %d\n", product.ProductId)
// 			fmt.Printf("Name (EN): %s\n", product.NameEn)
// 			fmt.Printf("Name (CN): %s\n", product.NameCn)
// 			fmt.Printf("Description (EN): %s\n", product.DescriptionEn)
// 			fmt.Printf("Description (CN): %s\n", product.DescriptionCn)
// 			fmt.Printf("Content (EN): %s\n", product.ContentEn)
// 			fmt.Printf("Content (CN): %s\n", product.ContentCn)
// 			fmt.Printf("Status: %d\n", product.Status)
// 			fmt.Printf("Image: %s\n", product.Image)
// 			fmt.Println("-----------------------")
// 		}
// 	})
// 	return
// }

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
