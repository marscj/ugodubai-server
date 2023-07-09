package system

import (
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AgentSearchReq struct {
	g.Meta `path:"/agent/list" tags:"代理商管理" method:"get" summary:"代理商列表"`
	commonApi.PageReq
}

type AgentSearchRes struct {
	g.Meta `mime:"application/json"`
	Agents []*model.SysAgentRes `json:"agent"`
	commonApi.ListRes
}

type AgentGetReq struct {
	g.Meta `path:"/agent/get" tags:"代理商管理" method:"get" summary:"获取代理商信息"`
	Id     uint `p:"id" v:"required#角色id不能为空""`
}

type AgentGetRes struct {
	g.Meta `mime:"application/json"`
	Agent  *entity.SysAgent `json:"agent"`
}
