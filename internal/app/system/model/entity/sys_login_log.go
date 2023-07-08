// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginLog is the golang structure for table sys_login_log.
type SysLoginLog struct {
	InfoId        int64       `json:"info_id"        description:"访问ID"`
	LoginName     string      `json:"login_name"     description:"登录账号"`
	Ipaddr        string      `json:"ipaddr"         description:"登录IP地址"`
	LoginLocation string      `json:"login_location" description:"登录地点"`
	Browser       string      `json:"browser"        description:"浏览器类型"`
	Os            string      `json:"os"             description:"操作系统"`
	Status        int         `json:"status"         description:"登录状态（0成功 1失败）"`
	Msg           string      `json:"msg"            description:"提示消息"`
	LoginTime     *gtime.Time `json:"login_time"     description:"登录时间"`
	Module        string      `json:"module"         description:"登录模块"`
}
