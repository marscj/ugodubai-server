package model

import (
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SysProductList struct {
	g.Meta        `orm:"table:sys_product, do:true"`
	ProductId     uint64                 `json:"productId"     description:""`
	NameEn        string                 `json:"nameEn"        description:"名称"`
	NameCn        string                 `json:"nameCn"        description:"名称"`
	DescriptionEn string                 `json:"descriptionEn" description:"产品简介"`
	DescriptionCn string                 `json:"descriptionCn" description:"产品简介"`
	Image         string                 `json:"image"         description:"缩略图"`
	Price         *SysProductPriceLookup `orm:"with:product_id=product_id" json:"price"`
}

type SysProduct struct {
	*entity.SysProduct
	MetaData  []*SysProductMeta        `orm:"with:product_id=product_id" json:"meta"`
	Term      []*SysProductTermsLookup `orm:"with:product_id=product_id" json:"term"`
	Variation []*SysVariation          `orm:"with:product_id=product_id" json:"variation"`
}

type SysProductMeta struct {
	*entity.SysProductMeta
}

type SysProductPriceLookup struct {
	*entity.SysProductPriceLookup
}

type SysVariation struct {
	*entity.SysVariation
	Price []*SysVariationPrice `orm:"with:variation_id=variation_id" json:"price"`
}

type SysVariationPrice struct {
	*entity.SysVariationPrice
	Attribute *SysVariationAttribute `orm:"with:attribute_id=attribute_id" json:"attribute"`
	Agent     *SysAgentList          `orm:"with:agent_id=agent_id json:agent" `
	Time      *SysVariationTime      `orm:"with:time_id=time_id json:agent" `
}

type SysVariationAttribute struct {
	*entity.SysVariationAttribute
}

type SysProductTermsLookup struct {
	*entity.SysProductTermsLookup
	Terms []*entity.SysProductTerms `orm:"with:term_id=term_id" json:"terms"`
}

type SysVariationTime struct {
	*entity.SysVariationTime
}
