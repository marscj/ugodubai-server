// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAgent is the golang structure for table sys_agent.
type SysAgent struct {
	Id             int         `json:"id"              description:"代理商ID"`
	Name           string      `json:"name"            description:"名称"`
	Email          string      `json:"email"           description:"邮箱"`
	ContactName    string      `json:"contact_name"    description:"联系人姓名"`
	ContactPhone   string      `json:"contact_phone"   description:"联系人电话"`
	Address        string      `json:"address"         description:"地址"`
	Nationality    string      `json:"nationality"     description:"国籍"`
	AgentCode      string      `json:"agent_code"      description:"代理商代码"`
	AvailableLimit float64     `json:"available_limit" description:"可用额度"`
	CreditLimit    float64     `json:"credit_limit"    description:"信用额度"`
	UsedLimit      float64     `json:"used_limit"      description:"已使用额度"`
	Status         int         `json:"status"          description:"状态"`
	AdminId        int         `json:"admin_id"        description:"管理员ID"`
	LicenseUrl     string      `json:"license_url"     description:"许可证URL"`
	CreatedAt      *gtime.Time `json:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updated_at"      description:"更新时间"`
}
