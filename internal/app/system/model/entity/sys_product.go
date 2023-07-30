// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysProduct is the golang structure for table sys_product.
type SysProduct struct {
	ProductId     uint64 `json:"productId"     description:""`
	NameEn        string `json:"nameEn"        description:"名称"`
	NameCn        string `json:"nameCn"        description:"名称"`
	DescriptionEn string `json:"descriptionEn" description:"产品简介"`
	DescriptionCn string `json:"descriptionCn" description:"产品简介"`
	ContentEn     string `json:"contentEn"     description:"产品内容"`
	ContentCn     string `json:"contentCn"     description:"产品内容"`
	Status        int    `json:"status"        description:"状态 0.下线 1.上线"`
	Image         string `json:"image"         description:"缩略图"`
	Order         int    `json:"order"         description:"显示顺序"`
}
