// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysGateway is the golang structure of table sys_gateway for DAO operations like Where/Data.
type SysGateway struct {
	g.Meta    `orm:"table:sys_gateway, do:true"`
	GatewayId interface{} //
	Name      interface{} // 名称
}
