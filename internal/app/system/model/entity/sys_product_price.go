// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductPrice is the golang structure for table sys_product_price.
type SysProductPrice struct {
	Id           uint64      `json:"id"           description:""`
	ProductId    uint64      `json:"productId"    description:""`
	VariationId  uint64      `json:"variationId"  description:""`
	AgentId      uint64      `json:"agentId"      description:""`
	StartDate    *gtime.Time `json:"startDate"    description:""`
	EndDate      *gtime.Time `json:"endDate"      description:""`
	CostPrice    float64     `json:"costPrice"    description:"成本价"`
	SellingPrice float64     `json:"sellingPrice" description:"销售价"`
	SpecialPrice float64     `json:"specialPrice" description:"预付价格"`
	Currency     string      `json:"currency"     description:"货币"`
	Status       int         `json:"status"       description:"状态 0.下线 1.上线"`
	Limit        int         `json:"limit"        description:"状态 0.使用库存 1.无限"`
	Stock        int         `json:"stock"        description:""`
}
