package model

import (
	"time"
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SysBooking struct {
	*entity.SysBooking
}

type SysVariationPriceByCheckout struct {
	g.Meta           `orm:"table:sys_variation_price, do:true"`
	VariationPriceId uint64      `json:"variationPriceId" description:""`
	VariationId      uint64      `json:"variationId"      description:""`
	AttributeId      uint64      `json:"attributeId"      description:""`
	StartDate        *gtime.Time `json:"startDate"        description:""`
	EndDate          *gtime.Time `json:"endDate"          description:""`
	SpecialPrice     float64     `json:"specialPrice"     description:"预付价格"`
	SellingPrice     float64     `json:"sellingPrice"     description:"销售价"`
}

type SysVariationByCheckout struct {
	g.Meta      `orm:"table:sys_variation, do:true"`
	VariationId uint64 `json:"variationId" description:""`
	ProductId   uint64 `json:"productId"   description:""`
	NameEn      string `json:"nameEn"      description:"默认名称"`
	NameCn      string `json:"nameCn"      description:"默认名称"`
	Sku         string `json:"sku"         description:"SKU"`
	Status      int    `json:"status"      description:"状态 0.下线 1.上线"`
}

type SysVariationAttributeByCheckout struct {
	g.Meta           `orm:"table:sys_variation_attribute, do:true"`
	AttributeId      uint64 `json:"attributeId"      description:""`
	AttributeKey     string `json:"attributeKey"     description:""`
	AttributeValueEn string `json:"attributeValueEn" description:""`
	AttributeValueCn string `json:"attributeValueCn" description:""`
}

type SysProductByCheckout struct {
	g.Meta    `orm:"table:sys_product, do:true"`
	ProductId uint64 `json:"productId"     description:""`
	NameEn    string `json:"nameEn"        description:"名称"`
	NameCn    string `json:"nameCn"        description:"名称"`
}

type SysPreCheckOutItemReq struct {
	Quantity         uint   `json:"quantity" v:"required#请输入产品数量"`
	VariationPriceId uint64 `json:"variationPriceId" v:"required#请选择产品项"`
	ActionDate       string `json:"actionDate" v:"required|date-format:Y-m-d H:i#请输入执行日期|请输入正确的日期，如2020-01-01 23:00"`
}

type SysPreCheckOutItemRes struct {
	Quantity         uint      `json:"quantity" v:"required#请输入产品数量"`
	VariationPriceId uint64    `json:"variationPriceId" v:"required#请选择产品项"`
	ActionDate       string    `json:"actionDate" v:"required|date-format:Y-m-d H:i#请输入执行日期|请输入正确的日期，如2020-01-01 23:00"`
	Timestamp        time.Time `json:"timestamp"`
	Signature        string    `json:"signature"`
}

type SysCheckOutItem struct {
	Quantity         uint   `json:"quantity" v:"required#请输入产品数量"`
	VariationPriceId uint64 `json:"variationPriceId" v:"required#请选择产品项"`
	ContactName      string `json:"contactName" v:"required|length:2,100#请输入客户姓名|客户姓名长度为2-100位"`
	ContactPhone     string `json:"contactPhone" v:"max-length:32#最大32位"`
	Currency         string `json:"currency" v:"max-length:3#最大3位"`
	Remark           string `json:"remark" v:"max-length:255#最大255位"`
	RelatedId        string `json:"relatedId" v:"max-length:64#最大64位"`
	ActionDate       string `json:"actionDate" v:"required|date-format:Y-m-d H:i#请输入执行日期|请输入正确的日期，如2020-01-01 23:00"`
}
