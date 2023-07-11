package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/model/entity"
)

type ISysDept interface {
	GetList(ctx context.Context, req *system.DeptListReq) (list []*entity.SysDept, err error)
	GetFromCache(ctx context.Context) (list []*entity.SysDept, err error)
	Add(ctx context.Context, req *system.DeptAddReq) (err error)
	Edit(ctx context.Context, req *system.DeptEditReq) (err error)
	Delete(ctx context.Context, id uint64) (err error)
	FindSonByParentId(deptList []*entity.SysDept, deptId uint64) []*entity.SysDept
	GetListTree(pid uint64, list []*entity.SysDept) (deptTree []*model.SysDeptTreeRes)
	GetByDeptId(ctx context.Context, deptId uint64) (dept *entity.SysDept, err error)
}

var localSysDept ISysDept

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}
