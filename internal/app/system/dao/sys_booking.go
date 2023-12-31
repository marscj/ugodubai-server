// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// internalSysBookingDao is internal type for wrapping internal DAO implements.
type internalSysBookingDao = *internal.SysBookingDao

// sysBookingDao is the data access object for table sys_booking.
// You can define custom methods on it to extend its functionality as you wish.
type sysBookingDao struct {
	internalSysBookingDao
}

var (
	// SysBooking is globally public accessible object for table sys_booking operations.
	SysBooking = sysBookingDao{
		internal.NewSysBookingDao(),
	}
)

// Fill with you ideas below.
