// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductLookup is the golang structure of table sys_product_lookup for DAO operations like Where/Data.
type SysProductLookup struct {
	g.Meta            `orm:"table:sys_product_lookup, do:true"`
	VariationLookupId interface{} //
	ProductId         interface{} //
	VariationId       interface{} //
	AttributeId       interface{} //
	VariationPriceId  interface{} //
	AgentId           interface{} //
}
