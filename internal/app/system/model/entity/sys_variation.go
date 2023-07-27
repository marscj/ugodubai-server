// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysVariation is the golang structure for table sys_variation.
type SysVariation struct {
	VariationId uint64 `json:"variationId" description:""`
	ProductId   uint64 `json:"productId"   description:""`
	NameEn      string `json:"nameEn"      description:"默认名称"`
	NameCn      string `json:"nameCn"      description:"默认名称"`
	Sku         string `json:"sku"         description:"SKU"`
	Status      int    `json:"status"      description:"状态 0.下线 1.上线"`
}
