package dao

import (
	"ugodubai-server/internal/app/system/dao/internal"
)

// sysJobDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysJobDao struct {
	*internal.SysJobDao
}

var (
	// SysJob is globally public accessible object for table tools_gen_table operations.
	SysJob = sysJobDao{
		internal.NewSysJobDao(),
	}
)

// Fill with you ideas below.
