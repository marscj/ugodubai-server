// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysVariationPrice is the golang structure of table sys_variation_price for DAO operations like Where/Data.
type SysVariationPrice struct {
	g.Meta           `orm:"table:sys_variation_price, do:true"`
	VariationPriceId interface{} //
	VariationId      interface{} //
	AttributeId      interface{} //
	StartDate        *gtime.Time //
	EndDate          *gtime.Time //
	CostPrice        interface{} // 成本价
	SpecialPrice     interface{} // 预付价格
	SellingPrice     interface{} // 销售价
	Currency         interface{} // 货币
	Stock            interface{} //
}
