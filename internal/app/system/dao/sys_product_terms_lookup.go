// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// internalSysProductTermsLookupDao is internal type for wrapping internal DAO implements.
type internalSysProductTermsLookupDao = *internal.SysProductTermsLookupDao

// sysProductTermsLookupDao is the data access object for table sys_product_terms_lookup.
// You can define custom methods on it to extend its functionality as you wish.
type sysProductTermsLookupDao struct {
	internalSysProductTermsLookupDao
}

var (
	// SysProductTermsLookup is globally public accessible object for table sys_product_terms_lookup operations.
	SysProductTermsLookup = sysProductTermsLookupDao{
		internal.NewSysProductTermsLookupDao(),
	}
)

// Fill with you ideas below.