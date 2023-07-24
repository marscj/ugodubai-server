// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysProductVariation is the golang structure for table sys_product_variation.
type SysProductVariation struct {
	VariationId uint64 `json:"variationId" description:""`
	ProductId   uint64 `json:"productId"   description:""`
	Name        string `json:"name"        description:"名称"`
	Sku         string `json:"sku"         description:"SKU"`
	Status      int    `json:"status"      description:"状态 0.下线 1.上线"`
	Limit       int    `json:"limit"       description:"状态 0.无限 1.使用库存"`
	Stock       int    `json:"stock"       description:""`
}
