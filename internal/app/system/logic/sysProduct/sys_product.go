package sysProduct

import (
	"context"
	"strconv"
	"strings"
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

// 请编写一个例子

// List 产品列表
func (s *sSysProduct) List(ctx context.Context, req *system.ProductListReq) (total interface{}, productList []*model.SysProductList, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		// 获取登录用户
		user := service.Context().GetLoginUser(ctx)
		agentID := user.AgentId

		var whereStr string
		var args []interface{}

		// 关键字查询
		// if req.Keyword != nil {
		// 	// Search for products with the keyword
		// 	whereStr += "`product_id` IN (SELECT `product_id` FROM `sys_product_keywords` WHERE `keyword` LIKE ?) "
		// 	args = append(args, "%"+*req.Keyword+"%")
		// }
		if req.Keyword != nil {
			// Search for products with the keyword
			whereStr += "p.product_id IN (SELECT product_id FROM sys_product_keywords WHERE keyword LIKE ?) "
			args = append(args, "%"+*req.Keyword+"%")
		}

		// terms 查询
		// if len(req.TermsIDs) > 0 {
		// 	if whereStr != "" {
		// 		whereStr += " AND "
		// 	}
		// 	placeholders := strings.Trim(strings.Repeat("?,", len(req.TermsIDs)), ",")
		// 	whereStr += "`product_id` IN (SELECT `product_id` FROM `sys_product_lookup_terms` WHERE `term_id` IN(" + placeholders + ")) "
		// 	for _, id := range req.TermsIDs {
		// 		args = append(args, id)
		// 	}
		// }
		if len(req.TermsIDs) > 0 {
			if whereStr != "" {
				whereStr += " AND "
			}
			placeholders := strings.Trim(strings.Repeat("?,", len(req.TermsIDs)), ",")
			whereStr += "p.product_id IN (SELECT product_id FROM sys_product_lookup_terms WHERE term_id IN(" + placeholders + ")) "
			for _, id := range req.TermsIDs {
				args = append(args, id)
			}
		}

		//状态查询
		// if req.Status != nil {
		// 	if whereStr != "" {
		// 		whereStr += " AND "
		// 	}
		// 	whereStr += "`status` = ? "
		// 	args = append(args, *req.Status)
		// }
		if req.Status != nil {
			if whereStr != "" {
				whereStr += " AND "
			}
			whereStr += "p.status = ? "
			args = append(args, *req.Status)
		}

		//查询价格
		// var joinStr string
		// if agentID != 0 {
		// 	joinStr = "LEFT JOIN `sys_product_price_lookup` ON `sys_product`.`product_id`=`sys_product_price_lookup`.`product_id` AND (`sys_product_price_lookup`.`agent_id`=" + strconv.FormatInt(agentID, 10) + " OR `sys_product_price_lookup`.`agent_id` IS NULL)"
		// } else {
		// 	joinStr = "LEFT JOIN `sys_product_price_lookup` ON `sys_product`.`product_id`=`sys_product_price_lookup`.`product_id` AND `sys_product_price_lookup`.`agent_id` IS NULL"
		// }
		// m := dao.SysProduct.Ctx(ctx).Alias("sp").LeftJoin("sys_product_price_lookup spl", joinStr).Fields("sp.*, spl.min_price, spl.max_price").Where(whereStr, args...).Where("is_deleted", 0)
		// m := dao.SysProduct.Ctx(ctx).Fields("sys_product.*, sys_product_price_lookup.min_price, sys_product_price_lookup.max_price").Where(whereStr, args...).Where("is_deleted", 0).LeftJoin(joinStr)
		// joinStr := fmt.Sprintf("LEFT JOIN sys_product_price_lookup spp ON p.product_id = spp.product_id AND (spp.agent_id = %s OR spp.agent_id IS NULL)", agentIDString)
		// m := g.Model("sys_product", "p").Fields("p.*,spp.min_price, spp.max_price").LeftJoin(joinStr).Where(whereStr, args...)
		// if agentID != 0 {
		// 	g.Model("sys_product", "p").Fields("p.*").InnerJoin("sys_product_price_lookup", "pp", "p.product_id = pp.product_id") ...
		// } else {
		// }
		m := g.Model("sys_product", "p").Fields("p.*, spp.min_price, spp.max_price")
		if agentID != 0 {
			m = m.LeftJoin("sys_product_price_lookup", "spp", "p.product_id = spp.product_id AND (spp.agent_id = ? OR spp.agent_id = 0)", strconv.FormatInt(agentID, 10)).
				Where(whereStr, args...)
		} else {
			m = m.
				LeftJoin("sys_product_price_lookup", "spp", "p.product_id = spp.product_id AND spp.agent_id = 0")
			// Where(whereStr, args...)
		}

		// m := g.Model("sys_product", "p").Where(whereStr, args...).Fields("p.*, spp.min_price, spp.max_price")
		// m = m.LeftJoin("sys_product_price_lookup", "spp", "p.product_id = spp.product_id AND spp.agent_id = 0")

		// total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "产品列表获取失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}

		query, _ := m.All()

		print(query)

		err = m.Page(req.PageNum, req.PageSize).OrderAsc("order").With(model.SysProductPriceLookup{}).Scan(&productList)
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
