// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductVariationPrice is the golang structure for table sys_product_variation_price.
type SysProductVariationPrice struct {
	Id              uint64      `json:"id"              description:""`
	VariationMetaId uint64      `json:"variationMetaId" description:""`
	AgentId         uint64      `json:"agentId"         description:""`
	StartDate       *gtime.Time `json:"startDate"       description:""`
	EndDate         *gtime.Time `json:"endDate"         description:""`
	CostPrice       float64     `json:"costPrice"       description:"成本价"`
	SpecialPrice    float64     `json:"specialPrice"    description:"预付价格"`
	SellingPrice    float64     `json:"sellingPrice"    description:"销售价"`
	Currency        string      `json:"currency"        description:"货币"`
}
