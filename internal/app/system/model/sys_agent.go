package model

import (
	"ugodubai-server/internal/app/system/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

type SysAgent struct {
	*do.SysAgent
	User []*SysUserSimple `orm:"with:agent_id=agent_id" json:"user"`
}

type SysAgentList struct {
	g.Meta  `orm:"table:sys_agent, do:true"`
	AgentId interface{} // 代理商ID
	Name    interface{} // 名称
}
