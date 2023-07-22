// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAgentDao is the data access object for table sys_agent.
type SysAgentDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SysAgentColumns // columns contains all the column names of Table for convenient usage.
}

// SysAgentColumns defines and stores column names for table sys_agent.
type SysAgentColumns struct {
	Id                 string // 代理商ID
	Name               string // 名称
	Email              string // 邮箱
	ContactName        string // 联系人姓名
	ContactPhone       string // 联系人电话
	Address            string // 地址
	Nationality        string // 国籍
	AgentCode          string // 代理商代码
	AvailableLimit     string // 可用额度
	CreditLimit        string // 信用额度
	OutstandingBalance string // 未结算额度
	AccountBlance      string // 账户余额
	Status             string // 状态 0.不可用 1.可用
	AdminId            string // 管理员ID
	LicenseUrl         string // 许可证URL
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
}

// sysAgentColumns holds the columns for table sys_agent.
var sysAgentColumns = SysAgentColumns{
	Id:                 "id",
	Name:               "name",
	Email:              "email",
	ContactName:        "contact_name",
	ContactPhone:       "contact_phone",
	Address:            "address",
	Nationality:        "nationality",
	AgentCode:          "agent_code",
	AvailableLimit:     "available_limit",
	CreditLimit:        "credit_limit",
	OutstandingBalance: "outstanding_balance",
	AccountBlance:      "account_blance",
	Status:             "status",
	AdminId:            "admin_id",
	LicenseUrl:         "license_url",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewSysAgentDao creates and returns a new DAO object for table data access.
func NewSysAgentDao() *SysAgentDao {
	return &SysAgentDao{
		group:   "default",
		table:   "sys_agent",
		columns: sysAgentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAgentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAgentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAgentDao) Columns() SysAgentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAgentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAgentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAgentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
