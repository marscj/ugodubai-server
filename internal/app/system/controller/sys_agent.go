package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/service"
)

var Agent = agentController{}

type agentController struct {
	BaseController
}

// List 代理商列u表
func (c *agentController) List(ctx context.Context, req *system.AgentSearchReq) (res *system.AgentSearchRes, err error) {
	res, err = service.SysAgent().List(ctx, req)
	return
}

// Get 获取角色信息
func (c *agentController) Get(ctx context.Context, req *system.AgentGetReq) (res *system.AgentGetRes, err error) {
	res = new(system.AgentGetRes)
	res.Agent, err = service.SysAgent().Get(ctx, req.Id)
	return
}
