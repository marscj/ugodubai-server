// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysTerms is the golang structure for table sys_terms.
type SysTerms struct {
	TermId    uint64 `json:"termId"    description:""`
	Taxonomy  string `json:"taxonomy"  description:""`
	Parent    uint64 `json:"parent"    description:""`
	Name      string `json:"name"      description:""`
	Slug      string `json:"slug"      description:""`
	TermOrder int    `json:"termOrder" description:""`
}
