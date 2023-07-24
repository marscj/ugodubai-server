// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysProduct is the golang structure of table sys_product for DAO operations like Where/Data.
type SysProduct struct {
	g.Meta      `orm:"table:sys_product, do:true"`
	Id          interface{} //
	Sku         interface{} // SKU
	Name        interface{} // 名称
	Description interface{} // 产品简介
	Content     interface{} // 产品内容
	Status      interface{} // 状态 0.下线 1.上线
}
