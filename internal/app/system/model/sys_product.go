package model

import "ugodubai-server/internal/app/system/model/entity"

type SysProduct struct {
	*entity.SysProduct
	Meta      []*SysProductMeta      `orm:"with:product_id=id" json:"meta"`
	Variation []*SysProductVariation `orm:"with:product_id=id" json:"variation"`
}

type SysProductMeta struct {
	*entity.SysProductMeta
}

type SysProductVariation struct {
	*entity.SysProductVariation
	Price []*SysProductVariationPrice `orm:"with:variation_id=variation_id" json:"price"`
}

type SysProductVariationPrice struct {
	*entity.SysProductVariationPrice
	Agent     *entity.SysAgent       `orm:"with:id=agent_id" json:"agent"`
	Attribute []*SysProductAttribute `orm:"with:attribute_id=attribute_id" json:"attribute"`
}

type SysProductAttribute struct {
	*entity.SysProductAttribute
}

type SysProductTerms struct {
	*entity.SysProductTerms
}
