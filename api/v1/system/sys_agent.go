package system

import (
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AgentSearchReq struct {
	g.Meta `path:"/agent/list" tags:"代理商管理" method:"get" summary:"代理商列表"`
	commonApi.PageReq
}

type AgentSearchRes struct {
	g.Meta    `mime:"application/json"`
	AgentList []*model.SysAgentRes
	commonApi.ListRes
}
