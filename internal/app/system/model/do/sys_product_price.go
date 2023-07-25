// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProductPrice is the golang structure of table sys_product_price for DAO operations like Where/Data.
type SysProductPrice struct {
	g.Meta            `orm:"table:sys_product_price, do:true"`
	Id                interface{} //
	ProductId         interface{} //
	VariationId       interface{} //
	AgentId           interface{} //
	StartDate         *gtime.Time //
	EndDate           *gtime.Time //
	CostPrice         interface{} // 成本价
	SpecialPrice      interface{} // 预付价格
	SellingPrice      interface{} // 销售价
	ChildCostPrice    interface{} // 儿童成本价
	ChildSpecialPrice interface{} // 儿童预付价格
	ChildSellingPrice interface{} // 儿童销售价
	Currency          interface{} // 货币
}
