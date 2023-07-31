package model

import (
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SysProductList struct {
	g.Meta        `orm:"table:sys_product, do:true"`
	ProductId     uint64                   `json:"productId"     description:""`
	NameEn        string                   `json:"nameEn"        description:"名称"`
	NameCn        string                   `json:"nameCn"        description:"名称"`
	DescriptionEn string                   `json:"descriptionEn" description:"产品简介"`
	DescriptionCn string                   `json:"descriptionCn" description:"产品简介"`
	Image         string                   `json:"image"         description:"缩略图"`
	Price         *SysProductPriceLookup   `json:"price"`
	Terms         []*SysProductTermsLookup `orm:"with:product_id=product_id" json:"term"`
}

type SysProduct struct {
	*entity.SysProduct
	Price  *SysProductPriceLookup   `orm:"with:product_id=product_id" json:"price"`
	Meta   []*SysProductMeta        `orm:"with:product_id=product_id" json:"meta"`
	Lookup []*SysProductLookup      `orm:"with:product_id=product_id" json:"lookup"`
	Term   []*SysProductTermsLookup `orm:"with:product_id=product_id" json:"term"`
}

type SysProductMeta struct {
	*entity.SysProductMeta
}

type SysProductPriceLookup struct {
	*entity.SysProductPriceLookup
}

type SysProductLookup struct {
	*entity.SysProductLookup
	Variation *SysVariation          `orm:"with:variation_id=variation_id" json:"variation"`
	Price     *SysVariationPrice     `orm:"with:variation_price_id=variation_price_id" json:"price"`
	Attribute *SysVariationAttribute `orm:"with:attribute_id=attribute_id" json:"attribute"`
	Agent     *SysAgentList          `orm:"with:agent_id=agent_id json:agent" `
}

type SysVariation struct {
	*entity.SysVariation
}

type SysVariationPrice struct {
	*entity.SysVariationPrice
}

type SysVariationAttribute struct {
	*entity.SysVariationAttribute
}

type SysProductTermsLookup struct {
	*entity.SysProductTermsLookup
	Terms []*entity.SysProductTerms `orm:"with:term_id=term_id" json:"terms"`
}
