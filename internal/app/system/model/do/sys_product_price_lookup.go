// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductPriceLookup is the golang structure of table sys_product_price_lookup for DAO operations like Where/Data.
type SysProductPriceLookup struct {
	g.Meta    `orm:"table:sys_product_price_lookup, do:true"`
	Id        interface{} //
	ProductId interface{} //
	AgentId   interface{} //
	MinPrice  interface{} // 最低价格
	MaxPrice  interface{} // 最高价格
}
