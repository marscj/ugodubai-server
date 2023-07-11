package model

import "ugodubai-server/internal/app/system/model/entity"

type SysAgent struct {
	*entity.SysAgent
	User []*SysUserSimple `orm:"with:agent_id=id" json:"user"`
}
