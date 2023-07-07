// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAgent is the golang structure of table sys_agent for DAO operations like Where/Data.
type SysAgent struct {
	g.Meta         `orm:"table:sys_agent, do:true"`
	Id             interface{} // 代理商ID
	Name           interface{} // 名称
	Email          interface{} // 邮箱
	ContactName    interface{} // 联系人姓名
	ContactPhone   interface{} // 联系人电话
	Address        interface{} // 地址
	Nationality    interface{} // 国籍
	AgentCode      interface{} // 代理商代码
	AvailableLimit interface{} // 可用额度
	CreditLimit    interface{} // 信用额度
	UsedLimit      interface{} // 已使用额度
	Status         interface{} // 状态
	AdminId        interface{} // 管理员ID
	LicenseUrl     interface{} // 许可证URL
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
