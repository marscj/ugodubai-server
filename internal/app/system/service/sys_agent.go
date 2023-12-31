package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
)

type ISysAgent interface {
	List(ctx context.Context, req *system.AgentListReq) (total interface{}, agentList []*model.SysAgent, err error)
	Get(ctx context.Context, id uint64) (agent *model.SysAgent, err error)
}

var localSysAgent ISysAgent

func SysAgent() ISysAgent {
	if localSysAgent == nil {
		panic("implement not found for interface ISysAgent, forgot register?")
	}
	return localSysAgent
}

func RegisterSysAgent(i ISysAgent) {
	localSysAgent = i
}
