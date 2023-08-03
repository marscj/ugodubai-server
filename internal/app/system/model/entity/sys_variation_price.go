// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysVariationPrice is the golang structure for table sys_variation_price.
type SysVariationPrice struct {
	VariationPriceId uint64      `json:"variationPriceId" description:""`
	VariationId      uint64      `json:"variationId"      description:""`
	AttributeId      uint64      `json:"attributeId"      description:""`
	AgentId          uint64      `json:"agentId"          description:""`
	StartDate        *gtime.Time `json:"startDate"        description:""`
	EndDate          *gtime.Time `json:"endDate"          description:""`
	CostPrice        float64     `json:"costPrice"        description:"成本价"`
	SpecialPrice     float64     `json:"specialPrice"     description:"预付价格"`
	SellingPrice     float64     `json:"sellingPrice"     description:"销售价"`
	Currency         string      `json:"currency"         description:"货币"`
	Stock            int         `json:"stock"            description:""`
	Tiems            string      `json:"tiems"            description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:"更新时间"`
}
