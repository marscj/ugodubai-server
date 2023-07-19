package system

import (
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta `path:"/order/list" tags:"代理商管理" method:"get" summary:"代理商列表"`
	commonApi.PageReq
}

type OrderListRes struct {
	g.Meta `mime:"application/json"`
	Orders []*model.SysOrder `json:"order"`
	commonApi.ListRes
}

type OrderGetReq struct {
	g.Meta `path:"/order/get" tags:"代理商管理" method:"get" summary:"获取代理商信息"`
	Id     uint64 `p:"id" v:"required#代理商id不能为空""`
}

type OrderGetRes struct {
	g.Meta `mime:"application/json"`
	Order  *model.SysOrder `json:"order"`
}
