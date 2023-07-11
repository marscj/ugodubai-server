package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/service"
)

var UserOnline = new(SysUserOnlineController)

type SysUserOnlineController struct{}

func (c *SysUserOnlineController) List(ctx context.Context, req *system.SysUserOnlineListReqstRes) (res *system.SysUserOnlineListRes, err error) {
	res, err = service.SysUserOnline().GetOnlineListPage(ctx, req)
	return
}

func (c *SysUserOnlineController) ForceLogout(ctx context.Context, req *system.SysUserOnlineForceLogoutReq) (res *system.SysUserOnlineForceLogoutRes, err error) {
	err = service.SysUserOnline().ForceLogout(ctx, req.Ids)
	return
}
