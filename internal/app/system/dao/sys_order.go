// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// internalSysOrderDao is internal type for wrapping internal DAO implements.
type internalSysOrderDao = *internal.SysOrderDao

// sysOrderDao is the data access object for table sys_order.
// You can define custom methods on it to extend its functionality as you wish.
type sysOrderDao struct {
	internalSysOrderDao
}

var (
	// SysOrder is globally public accessible object for table sys_order operations.
	SysOrder = sysOrderDao{
		internal.NewSysOrderDao(),
	}
)

// Fill with you ideas below.
