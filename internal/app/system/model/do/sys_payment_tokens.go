// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysPaymentTokens is the golang structure of table sys_payment_tokens for DAO operations like Where/Data.
type SysPaymentTokens struct {
	g.Meta    `orm:"table:sys_payment_tokens, do:true"`
	TokenId   interface{} //
	GatewayId interface{} //
	Token     interface{} //
	UserId    interface{} //
	AgentId   interface{} //
	Type      interface{} //
	IsDefault interface{} //
}
