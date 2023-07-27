// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysVariationLookup is the golang structure of table sys_variation_lookup for DAO operations like Where/Data.
type SysVariationLookup struct {
	g.Meta            `orm:"table:sys_variation_lookup, do:true"`
	VariationLookupId interface{} //
	VariationId       interface{} //
	AttributeId       interface{} //
	VariationPriceId  interface{} //
	AgentId           interface{} //
}
