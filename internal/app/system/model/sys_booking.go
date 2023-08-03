package model

import (
	"time"
	"ugodubai-server/internal/app/system/model/entity"
)

type SysBooking struct {
	*entity.SysBooking
}

type PreCheckOutItem struct {
	Quantity         uint   `v:"required#请输入产品数量"`
	VariationPriceId uint64 `v:"required#请选择产品项"`
	ActionDate       string `v:"required|date-format:Y-m-d H:i#请输入执行日期|请输入正确的日期，如2020-01-01 23:00"`
	Price            string
	Tax              string
	Timestamp        time.Time `json:"timestamp"`
	Signature        string    `dc:"签名"`
}

type CheckOutItem struct {
	Quantity         uint   `v:"required#请输入产品数量"`
	VariationPriceId uint64 `v:"required#请选择产品项"`
	ContactName      string `v:"required|length:2,100#请输入客户姓名|客户姓名长度为2-100位"`
	ContactPhone     string `v:"max-length:32#最大32位"`
	Currency         string `v:"max-length:3#最大3位"`
	Remark           string `v:"max-length:255#最大255位"`
	RelatedId        string `v:"max-length:64#最大64位"`
	ActionDate       string `v:"required|date-format:Y-m-d H:i#请输入执行日期|请输入正确的日期，如2020-01-01 23:00"`
}
