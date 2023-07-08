// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id            uint64      `json:"id"              description:""`
	UserName      string      `json:"user_name"       description:"用户名"`
	Mobile        string      `json:"mobile"          description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string      `json:"user_nickname"   description:"用户昵称"`
	Birthday      int         `json:"birthday"        description:"生日"`
	UserPassword  string      `json:"user_password"   description:"登录密码;cmf_password加密"`
	UserSalt      string      `json:"user_salt"       description:"加密盐"`
	UserStatus    uint        `json:"user_status"     description:"用户状态;0:禁用,1:正常,2:未验证"`
	UserEmail     string      `json:"user_email"      description:"用户登录邮箱"`
	Sex           int         `json:"sex"             description:"性别;0:保密,1:男,2:女"`
	Avatar        string      `json:"avatar"          description:"用户头像"`
	DeptId        uint64      `json:"dept_id"         description:"部门id"`
	Remark        string      `json:"remark"          description:"备注"`
	IsAdmin       int         `json:"is_admin"        description:"是否后台管理员 1 是  0   否"`
	Address       string      `json:"address"         description:"联系地址"`
	Describe      string      `json:"describe"        description:"描述信息"`
	LastLoginIp   string      `json:"last_login_ip"   description:"最后登录ip"`
	LastLoginTime *gtime.Time `json:"last_login_time" description:"最后登录时间"`
	CreatedAt     *gtime.Time `json:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updated_at"      description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deleted_at"      description:"删除时间"`
}
