package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model/entity"
	"ugodubai-server/internal/app/system/service"
)

var Agent = agentController{}

type agentController struct {
	BaseController
}

// List 代理商列u表
func (c *agentController) List(ctx context.Context, req *system.AgentListReq) (res *system.AgentListRes, err error) {
	var (
		total     interface{}
		agentList []*entity.SysAgent
	)

	res = new(system.AgentListRes)

	total, agentList, err = service.SysAgent().List(ctx, req)
	res.Total = total
	res.Agents = agentList

	return
}

// Get 获取代理商信息
func (c *agentController) Get(ctx context.Context, req *system.AgentGetReq) (res *system.AgentGetRes, err error) {
	res = new(system.AgentGetRes)
	res.Agent, err = service.SysAgent().Get(ctx, req.Id)

	if err != nil {
		return
	}

	if req.User {
		res.User, err = service.SysUser().GetAgentUsers(ctx, req.Id)
	}
	return
}
