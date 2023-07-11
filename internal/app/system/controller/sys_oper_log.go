package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/service"
)

var OperLog = new(operateLogController)

type operateLogController struct {
	BaseController
}

// List 列表
func (c *operateLogController) List(ctx context.Context, req *system.SysOperLogListReq) (res *system.SysOperLogListRes, err error) {
	res, err = service.OperateLog().List(ctx, req)
	return
}

// Get 获取操作日志
func (c *operateLogController) Get(ctx context.Context, req *system.SysOperLogGetReq) (res *system.SysOperLogGetRes, err error) {
	res = new(system.SysOperLogGetRes)
	res.SysOperLogInfo, err = service.OperateLog().GetByOperId(ctx, req.OperId)
	return
}

func (c *operateLogController) Delete(ctx context.Context, req *system.SysOperLogDeleteReq) (res *system.SysOperLogDeleteRes, err error) {
	err = service.OperateLog().DeleteByIds(ctx, req.OperIds)
	return
}

func (c *operateLogController) Clear(ctx context.Context, req *system.SysOperLogClearReq) (res *system.SysOperLogClearRes, err error) {
	err = service.OperateLog().ClearLog(ctx)
	return
}
