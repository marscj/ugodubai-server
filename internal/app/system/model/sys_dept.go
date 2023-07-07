package model

import "ugodubai-server/internal/app/system/model/entity"

type SysDeptTreeRes struct {
	*entity.SysDept
	Children []*SysDeptTreeRes `json:"children"`
}
