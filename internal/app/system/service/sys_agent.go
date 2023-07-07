package service

import (
	"context"

	"ugodubai-server/api/v1/system"
)

type ISysAgent interface {
	List(ctx context.Context, req *system.AgentSearchReq) (res *system.AgentSearchRes, err error)
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
