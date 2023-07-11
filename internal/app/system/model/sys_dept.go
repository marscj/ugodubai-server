package model

import "ugodubai-server/internal/app/system/model/entity"

type SysDeptTreeModel struct {
	*entity.SysDept
	Children []*SysDeptTreeModel `json:"children"`
}
