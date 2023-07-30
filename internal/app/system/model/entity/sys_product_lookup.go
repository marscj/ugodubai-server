// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysProductLookup is the golang structure for table sys_product_lookup.
type SysProductLookup struct {
	VariationLookupId uint64 `json:"variationLookupId" description:""`
	ProductId         uint64 `json:"productId"         description:""`
	VariationId       uint64 `json:"variationId"       description:""`
	AttributeId       uint64 `json:"attributeId"       description:""`
	VariationPriceId  uint64 `json:"variationPriceId"  description:""`
	AgentId           uint64 `json:"agentId"           description:""`
}
