// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// internalSysVariationAttributeDao is internal type for wrapping internal DAO implements.
type internalSysVariationAttributeDao = *internal.SysVariationAttributeDao

// sysVariationAttributeDao is the data access object for table sys_variation_attribute.
// You can define custom methods on it to extend its functionality as you wish.
type sysVariationAttributeDao struct {
	internalSysVariationAttributeDao
}

var (
	// SysVariationAttribute is globally public accessible object for table sys_variation_attribute operations.
	SysVariationAttribute = sysVariationAttributeDao{
		internal.NewSysVariationAttributeDao(),
	}
)

// Fill with you ideas below.