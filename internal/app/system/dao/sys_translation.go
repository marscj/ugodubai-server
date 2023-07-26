// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// internalSysTranslationDao is internal type for wrapping internal DAO implements.
type internalSysTranslationDao = *internal.SysTranslationDao

// sysTranslationDao is the data access object for table sys_translation.
// You can define custom methods on it to extend its functionality as you wish.
type sysTranslationDao struct {
	internalSysTranslationDao
}

var (
	// SysTranslation is globally public accessible object for table sys_translation operations.
	SysTranslation = sysTranslationDao{
		internal.NewSysTranslationDao(),
	}
)

// Fill with you ideas below.