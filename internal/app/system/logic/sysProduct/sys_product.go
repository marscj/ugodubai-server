package sysProduct

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/consts"
	"ugodubai-server/internal/app/system/dao"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
	"ugodubai-server/library/liberr"

	"github.com/gogf/gf/util/gconv"
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

// 请编写一个例子

// List 产品列表
func (s *sSysProduct) List(ctx context.Context, req *system.ProductListReq) (total interface{}, productList []*model.SysProductList, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysProduct.Ctx(ctx)

		var whereStr string
		var args []interface{}

		// Keyword search
		if req.Keyword != "" {
			// Search for products with the keyword
			whereStr += "`product_id` IN (SELECT `product_id` FROM `sys_product_keywords` WHERE `keyword` LIKE ?) "
			args = append(args, "%"+req.Keyword+"%")
		}

		//获取User
		user := service.Context().GetLoginUser(ctx)

		//前端客户只能获取已上线产品
		if user.IsAdmin == 0 {
			m = m.Where("status = ?", 1)
		}

		//查询terms
		var productIDs []int64
		if len(req.TermsIDs) > 0 {

			ids, err := g.Model("sys_product_lookup_terms").Fields("product_id").Where("term_id IN(?)", req.TermsIDs).Distinct().Array("product_id")
			if err != nil {
				liberr.ErrIsNil(ctx, err, "产品列表获取失败")
				return
			}
			productIDs = append(productIDs, gconv.Int64s(ids)...)
		}

		if len(productIDs) > 0 {
			m = m.Where("product_id IN(?)", productIDs)
			if err != nil {
				liberr.ErrIsNil(ctx, err, "产品列表获取失败")
				return
			}
		}

		//查询关键字
		if req.Keyword != "" {
			m = m.Where("name_cn LIKE ? OR name_en LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
		}

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
