// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// internalSysProductTermsDao is internal type for wrapping internal DAO implements.
type internalSysProductTermsDao = *internal.SysProductTermsDao

// sysProductTermsDao is the data access object for table sys_product_terms.
// You can define custom methods on it to extend its functionality as you wish.
type sysProductTermsDao struct {
	internalSysProductTermsDao
}

var (
	// SysProductTerms is globally public accessible object for table sys_product_terms operations.
	SysProductTerms = sysProductTermsDao{
		internal.NewSysProductTermsDao(),
	}
)

// Fill with you ideas below.