// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure of table sys_post for DAO operations like Where/Data.
type SysPost struct {
	g.Meta    `orm:"table:sys_post, do:true"`
	PostId    interface{} // 岗位ID
	PostCode  interface{} // 岗位编码
	PostName  interface{} // 岗位名称
	PostSort  interface{} // 显示顺序
	Status    interface{} // 状态（0正常 1停用）
	Remark    interface{} // 备注
	CreatedBy interface{} // 创建人
	UpdatedBy interface{} // 修改人
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
