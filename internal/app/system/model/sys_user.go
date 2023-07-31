package model

import (
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

// LoginUser 登录返回
type LoginUser struct {
	*entity.SysUser
	Id           uint64 `orm:"id,primary"       json:"id"`           //
	UserName     string `orm:"user_name,unique" json:"userName"`     // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	UserPassword string `orm:"user_password"    json:"userPassword"` // 登录密码;cmf_password加密
	UserSalt     string `orm:"user_salt"        json:"userSalt"`     // 加密盐
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`   // 用户状态;0:禁用,1:正常,2:未验证
	IsAdmin      int    `orm:"is_admin"         json:"isAdmin"`      // 是否后台管理员 1 是  0   否
	Avatar       string `orm:"avatar" json:"avatar"`                 //头像
	DeptId       uint64 `orm:"dept_id"       json:"deptId"`          //部门id
	AgentId      int64  `orm:"agent_id" json:"agentId"       `       //代理商ID
}

type SysUser struct {
	// gmeta.Meta `orm:"table:sys_user"`
	*entity.SysUser
	Agent *SysAgent  `orm:"with:id=agent_id" json:"agent"`
	Dept  *SysDept   `orm:"with:dept_id=dept_id" json:"dept"`
	Role  []*SysRole `orm:"with:id=id" json:"role"`
}

type SysUserSimple struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           uint64 `orm:"id"       json:"id"`                   //
	Avatar       string `orm:"avatar" json:"avatar"`                 // 头像
	Sex          int    `orm:"sex" json:"sex"`                       // 性别
	UserName     string `orm:"user_name" json:"userName"`            // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	AgentId      uint64 `orm:"agent_id" json:"agentId"`
}
