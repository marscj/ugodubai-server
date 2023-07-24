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
	SpecialPrice float64     `json:"specialPrice" description:"预付价格"`
	SellingPrice float64     `json:"sellingPrice" description:"销售价"`
	Currency     string      `json:"currency"     description:"货币"`
}
