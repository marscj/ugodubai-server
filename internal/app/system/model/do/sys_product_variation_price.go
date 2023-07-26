// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductVariationPrice is the golang structure of table sys_product_variation_price for DAO operations like Where/Data.
type SysProductVariationPrice struct {
	g.Meta       `orm:"table:sys_product_variation_price, do:true"`
	Id           interface{} //
	AttributeId  interface{} //
	VariationId  interface{} //
	AgentId      interface{} //
	StartDate    *gtime.Time //
	EndDate      *gtime.Time //
	CostPrice    interface{} // 成本价
	SpecialPrice interface{} // 预付价格
	SellingPrice interface{} // 销售价
	Currency     interface{} // 货币
	Status       interface{} // 状态 0.下线 1.上线
	Stock        interface{} //
}
