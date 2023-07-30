package system

import (
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ProductListReq struct {
	g.Meta   `path:"/product/list" tags:"产品管理" method:"get" summary:"产品列表"`
	TermsIDs []int64 `json:"terms_ids"`
	Keyword  *string `json:"keyword"`
	Status   *int    `json:"status"`
	commonApi.PageReq
}

type ProductListRes struct {
	g.Meta   `mime:"application/json"`
	Products []model.SysProductList `json:"product"`
	commonApi.ListRes
}

type ProductGetReq struct {
	g.Meta `path:"/product/get" tags:"产品管理" method:"get" summary:"获取产品"`
	Id     uint64 `p:"id" v:"required#产品id不能为空"`
}

type ProductGetRes struct {
	g.Meta  `mime:"application/json"`
	Product *model.SysProduct `json:"product"`
}
