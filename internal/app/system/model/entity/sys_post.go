// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost struct {
	PostId    uint64      `json:"post_id"    description:"岗位ID"`
	PostCode  string      `json:"post_code"  description:"岗位编码"`
	PostName  string      `json:"post_name"  description:"岗位名称"`
	PostSort  int         `json:"post_sort"  description:"显示顺序"`
	Status    uint        `json:"status"     description:"状态（0正常 1停用）"`
	Remark    string      `json:"remark"     description:"备注"`
	CreatedBy uint64      `json:"created_by" description:"创建人"`
	UpdatedBy uint64      `json:"updated_by" description:"修改人"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deleted_at" description:"删除时间"`
}
