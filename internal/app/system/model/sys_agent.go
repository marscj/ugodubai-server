package model

import "ugodubai-server/internal/app/system/model/entity"

type SysAgentModel struct {
	*entity.SysAgent
}

type SysAgentListModel struct {
	Id   int    `json:"id"                 description:"代理商ID"`
	Name string `json:"name"               description:"名称"`
}
