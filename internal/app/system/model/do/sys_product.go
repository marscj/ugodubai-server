// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysProduct is the golang structure of table sys_product for DAO operations like Where/Data.
type SysProduct struct {
	g.Meta           `orm:"table:sys_product, do:true"`
	Id               interface{} //
	Name             interface{} // 名称
	Description      interface{} //
	ShortDescription interface{} //
}
