package model

import (
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SysAgent struct {
	*entity.SysAgent
	User []*SysUserSimple `orm:"with:agent_id=agent_id" json:"user"`
}

type SysAgentList struct {
	g.Meta  `orm:"table:sys_agent, do:true"`
	AgentId uint64 `json:"agentId"            description:"代理商ID"`
	Name    string `json:"name"               description:"名称"`
}
