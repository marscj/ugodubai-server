// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysProductPriceLookup is the golang structure for table sys_product_price_lookup.
type SysProductPriceLookup struct {
	Id        uint64  `json:"id"        description:""`
	ProductId uint64  `json:"productId" description:""`
	MinPrice  float64 `json:"minPrice"  description:"最低价格"`
	MaxPrice  float64 `json:"maxPrice"  description:"最高价格"`
}
