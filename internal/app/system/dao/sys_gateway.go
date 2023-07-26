// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// internalSysGatewayDao is internal type for wrapping internal DAO implements.
type internalSysGatewayDao = *internal.SysGatewayDao

// sysGatewayDao is the data access object for table sys_gateway.
// You can define custom methods on it to extend its functionality as you wish.
type sysGatewayDao struct {
	internalSysGatewayDao
}

var (
	// SysGateway is globally public accessible object for table sys_gateway operations.
	SysGateway = sysGatewayDao{
		internal.NewSysGatewayDao(),
	}
)

// Fill with you ideas below.