package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/model/entity"
)

type (
	ISysUserOnline interface {
		Invoke(ctx context.Context, params *model.SysUserOnlineParams)
		SaveOnline(ctx context.Context, params *model.SysUserOnlineParams)
		CheckUserOnline(ctx context.Context)
		GetOnlineListPage(ctx context.Context, req *system.SysUserOnlineListReqstRes, hasToken ...bool) (res *system.SysUserOnlineListRes, err error)
		UserIsOnline(ctx context.Context, token string) bool
		DeleteOnlineByToken(ctx context.Context, token string) (err error)
		ForceLogout(ctx context.Context, ids []int) (err error)
		GetInfosByIds(ctx context.Context, ids []int) (onlineList []*entity.SysUserOnline, err error)
	}
)

var (
	localSysUserOnline ISysUserOnline
)

func SysUserOnline() ISysUserOnline {
	if localSysUserOnline == nil {
		panic("implement not found for interface ISysUserOnline, forgot register?")
	}
	return localSysUserOnline
}

func RegisterSysUserOnline(i ISysUserOnline) {
	localSysUserOnline = i
}
