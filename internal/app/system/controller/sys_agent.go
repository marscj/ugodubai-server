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

//List 代理商列u表
func (c *agentController) List(ctx context.Context, req *system.AgentSearchReq) (res *system.AgentSearchRes, err error) {
	res, err = service.SysAgent().List(ctx, req)
	return
}
