// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysProduct is the golang structure for table sys_product.
type SysProduct struct {
	Id          uint64 `json:"id"          description:""`
	Sku         string `json:"sku"         description:"SKU"`
	Name        string `json:"name"        description:"名称"`
	Description string `json:"description" description:"产品简介"`
	Content     string `json:"content"     description:"产品内容"`
	Status      int    `json:"status"      description:"状态 0.下线 1.上线"`
}
