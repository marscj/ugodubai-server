// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysTranslation is the golang structure for table sys_translation.
type SysTranslation struct {
	Id           uint64 `json:"id"           description:"唯一标识符"`
	RecordType   string `json:"recordType"   description:""`
	RecordId     uint64 `json:"recordId"     description:""`
	LanguageCode string `json:"languageCode" description:""`
	MetaKey      string `json:"metaKey"      description:""`
	MetaValue    string `json:"metaValue"    description:""`
}
