// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysTerms is the golang structure of table sys_terms for DAO operations like Where/Data.
type SysTerms struct {
	g.Meta    `orm:"table:sys_terms, do:true"`
	TermId    interface{} //
	Taxonomy  interface{} //
	Parent    interface{} //
	Name      interface{} //
	Slug      interface{} //
	TermOrder interface{} //
}
