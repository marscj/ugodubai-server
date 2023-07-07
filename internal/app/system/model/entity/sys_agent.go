// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAgent is the golang structure for table sys_agent.
type SysAgent struct {
	Id             uint        `json:"id"             description:"代理商ID"`
	Name           string      `json:"name"           description:"名称"`
	Email          string      `json:"email"          description:"邮箱"`
	ContactName    string      `json:"contactName"    description:"联系人姓名"`
	ContactPhone   string      `json:"contactPhone"   description:"联系人电话"`
	Address        string      `json:"address"        description:"地址"`
	Nationality    string      `json:"nationality"    description:"国籍"`
	AgentCode      string      `json:"agentCode"      description:"代理商代码"`
	AvailableLimit float64     `json:"availableLimit" description:"可用额度"`
	CreditLimit    float64     `json:"creditLimit"    description:"信用额度"`
	UsedLimit      float64     `json:"usedLimit"      description:"已使用额度"`
	AdminId        int         `json:"adminId"        description:"管理员ID"`
	LicenseUrl     string      `json:"licenseUrl"     description:"许可证URL"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
}
