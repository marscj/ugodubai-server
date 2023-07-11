package model

import "ugodubai-server/internal/app/system/model/entity"

type SysAgent struct {
	*entity.SysAgent
}

type SysAgentList struct {
	Id   int    `json:"id"                 description:"代理商ID"`
	Name string `json:"name"               description:"名称"`
}
