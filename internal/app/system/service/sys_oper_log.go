package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IOperateLog interface {
		OperationLog(r *ghttp.Request)
		Invoke(ctx context.Context, data *model.SysOperLogAdd)
		List(ctx context.Context, req *system.SysOperLogListReq) (listRes *system.SysOperLogListRes, err error)
		GetByOperId(ctx context.Context, operId uint64) (res *model.SysOperLogInfoRes, err error)
		DeleteByIds(ctx context.Context, ids []uint64) (err error)
		ClearLog(ctx context.Context) (err error)
	}
)

var (
	localOperateLog IOperateLog
)

func OperateLog() IOperateLog {
	if localOperateLog == nil {
		panic("implement not found for interface IOperateLog, forgot register?")
	}
	return localOperateLog
}

func RegisterOperateLog(i IOperateLog) {
	localOperateLog = i
}
