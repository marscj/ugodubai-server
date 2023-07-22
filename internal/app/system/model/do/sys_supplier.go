// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysSupplier is the golang structure of table sys_supplier for DAO operations like Where/Data.
type SysSupplier struct {
	g.Meta   `orm:"table:sys_supplier, do:true"`
	Id       interface{} // 唯一标识符
	Name     interface{} // 名称
	Banlance interface{} // 余额
}
