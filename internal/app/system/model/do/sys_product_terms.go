// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductTerms is the golang structure of table sys_product_terms for DAO operations like Where/Data.
type SysProductTerms struct {
	g.Meta   `orm:"table:sys_product_terms, do:true"`
	TermId   interface{} //
	Taxonomy interface{} //
	Parent   interface{} //
	NameEn   interface{} //
	NameCn   interface{} //
}
