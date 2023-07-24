// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysBooking is the golang structure of table sys_booking for DAO operations like Where/Data.
type SysBooking struct {
	g.Meta    `orm:"table:sys_booking, do:true"`
	Id        interface{} // 代理商ID
	ParentId  interface{} // 父
	RelatedId interface{} // 关联订单ID
	AgentId   interface{} // 代理商ID
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
