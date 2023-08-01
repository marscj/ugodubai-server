// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProduct is the golang structure of table sys_product for DAO operations like Where/Data.
type SysProduct struct {
	g.Meta        `orm:"table:sys_product, do:true"`
	ProductId     interface{} //
	NameEn        interface{} // 名称
	NameCn        interface{} // 名称
	DescriptionEn interface{} // 产品简介
	DescriptionCn interface{} // 产品简介
	ContentEn     interface{} // 产品内容
	ContentCn     interface{} // 产品内容
	Status        interface{} // 状态 0.下线 1.上线
	Image         interface{} // 缩略图
	Order         interface{} // 显示顺序
	IsDeleted     interface{} // 状态 0.未删除 1.已删除
	UpdatedAt     *gtime.Time // 更新时间
}
