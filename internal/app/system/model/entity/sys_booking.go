// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysBooking is the golang structure for table sys_booking.
type SysBooking struct {
	Id        uint64      `json:"id"        description:"代理商ID"`
	ParentId  uint64      `json:"parentId"  description:"父"`
	RelatedId string      `json:"relatedId" description:"关联订单ID"`
	AgentId   uint64      `json:"agentId"   description:"代理商ID"`
	CreatedBy uint64      `json:"createdBy" description:"创建者"`
	UpdatedBy uint64      `json:"updatedBy" description:"更新者"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
