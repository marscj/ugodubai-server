package model

import (
	"ugodubai-server/internal/app/system/model/do"
	"ugodubai-server/internal/app/system/model/entity"
)

type SysProduct struct {
	*entity.SysProduct
	Meta      []*SysProductMeta   `orm:"with:product_id=product_id" json:"meta"`
	Terms     []*SysProductLookup `orm:"with:product_id=product_id" json:"terms"`
	Variation []*SysVariation     `orm:"with:product_id=product_id" json:"variation"`
}

type SysProductMeta struct {
	*entity.SysProductMeta
}

type SysVariation struct {
	*entity.SysVariation
	Price []*SysVariationPrice `orm:"with:variation_id=variation_id" json:"price"`
}

type SysVariationPrice struct {
	*entity.SysVariationPrice
	Agent     *entity.SysAgent         `orm:"with:agent_id=agent_id" json:"agent"`
	Attribute []*SysVariationAttribute `orm:"with:attribute_id=attribute_id" json:"attribute"`
}

type SysVariationAttribute struct {
	*entity.SysVariationAttribute
}

type SysProductTerms struct {
	*do.SysProductTerms
}

type SysProductLookup struct {
	*do.SysProductLookup
	Terms []*entity.SysProductTerms `orm:"with:term_id=term_id" json:"terms"`
}
