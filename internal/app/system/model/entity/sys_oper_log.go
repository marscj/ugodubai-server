// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperLog is the golang structure for table sys_oper_log.
type SysOperLog struct {
	OperId        uint64      `json:"oper_id"        description:"日志主键"`
	Title         string      `json:"title"          description:"模块标题"`
	BusinessType  int         `json:"business_type"  description:"业务类型（0其它 1新增 2修改 3删除）"`
	Method        string      `json:"method"         description:"方法名称"`
	RequestMethod string      `json:"request_method" description:"请求方式"`
	OperatorType  int         `json:"operator_type"  description:"操作类别（0其它 1后台用户 2手机端用户）"`
	OperName      string      `json:"oper_name"      description:"操作人员"`
	DeptName      string      `json:"dept_name"      description:"部门名称"`
	OperUrl       string      `json:"oper_url"       description:"请求URL"`
	OperIp        string      `json:"oper_ip"        description:"主机地址"`
	OperLocation  string      `json:"oper_location"  description:"操作地点"`
	OperParam     string      `json:"oper_param"     description:"请求参数"`
	ErrorMsg      string      `json:"error_msg"      description:"错误消息"`
	OperTime      *gtime.Time `json:"oper_time"      description:"操作时间"`
}
