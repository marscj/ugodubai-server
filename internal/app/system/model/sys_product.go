package model

import (
	"ugodubai-server/internal/app/system/model/do"
	"ugodubai-server/internal/app/system/model/entity"
)

type SysProduct struct {
	*do.SysProduct
	Meta       []*SysProductMeta     `orm:"with:product_id=product_id" json:"meta"`
	Lookup     []*SysProductLookup   `orm:"with:product_id=product_id" json:"lookup"`
	Variations []*SysVariationLookup `orm:"with:product_id=product_id" json:"variations"`
}

type SysProductMeta struct {
	*do.SysProductMeta
}

type SysVariationLookup struct {
	*do.SysVariationLookup
	Variation *SysVariation          `orm:"with:variation_id=variation_id" json:"variation"`
	Price     *SysVariationPrice     `orm:"with:variation_price_id=variation_price_id" json:"price"`
	Attribute *SysVariationAttribute `orm:"with:attribute_id=attribute_id" json:"attribute"`
	Agent     *SysAgentList          `orm:"with:agent_id=agent_id json:agent" `
}

type SysVariation struct {
	*do.SysVariation
}

type SysVariationPrice struct {
	*do.SysVariationPrice
}

type SysVariationAttribute struct {
	*do.SysVariationAttribute
}

type SysProductLookup struct {
	*do.SysProductLookup
	Terms []*entity.SysProductTerms `orm:"with:term_id=term_id" json:"category"`
}
