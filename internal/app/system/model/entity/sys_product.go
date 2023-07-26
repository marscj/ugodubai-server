// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysProduct is the golang structure for table sys_product.
type SysProduct struct {
	Id            uint64 `json:"id"            description:""`
	Sku           string `json:"sku"           description:"SKU"`
	NameEn        string `json:"nameEn"        description:"名称"`
	NameCn        string `json:"nameCn"        description:"名称"`
	DescriptionEn string `json:"descriptionEn" description:"产品简介"`
	DescriptionCn string `json:"descriptionCn" description:"产品简介"`
	ContentEn     string `json:"contentEn"     description:"产品内容"`
	ContentCn     string `json:"contentCn"     description:"产品内容"`
	Status        int    `json:"status"        description:"状态 0.下线 1.上线"`
}