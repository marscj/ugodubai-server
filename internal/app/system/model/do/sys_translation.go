// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysTranslation is the golang structure of table sys_translation for DAO operations like Where/Data.
type SysTranslation struct {
	g.Meta       `orm:"table:sys_translation, do:true"`
	Id           interface{} // 唯一标识符
	RecordType   interface{} //
	RecordId     interface{} //
	LanguageCode interface{} //
	MetaKey      interface{} //
	MetaValue    interface{} //
}
