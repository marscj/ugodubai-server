// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	DeptId    uint64      `json:"dept_id"    description:"部门id"`
	ParentId  uint64      `json:"parent_id"  description:"父部门id"`
	Ancestors string      `json:"ancestors"  description:"祖级列表"`
	DeptName  string      `json:"dept_name"  description:"部门名称"`
	OrderNum  int         `json:"order_num"  description:"显示顺序"`
	Leader    string      `json:"leader"     description:"负责人"`
	Phone     string      `json:"phone"      description:"联系电话"`
	Email     string      `json:"email"      description:"邮箱"`
	Status    uint        `json:"status"     description:"部门状态（0正常 1停用）"`
	CreatedBy uint64      `json:"created_by" description:"创建人"`
	UpdatedBy int64       `json:"updated_by" description:"修改人"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deleted_at" description:"删除时间"`
}
