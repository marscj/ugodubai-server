package model

import "ugodubai-server/internal/app/system/model/entity"

type SysDeptTree struct {
	*entity.SysDept
	Children []*SysDeptTree `json:"children"`
}

type SysDept struct {
	*entity.SysDept
}
