// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"CSIE_ADMS_Back-End/internal/dao/internal"
)

// internalApplyApcsDao is internal type for wrapping internal DAO implements.
type internalApplyApcsDao = *internal.ApplyApcsDao

// applyApcsDao is the data access object for table apply_apcs.
// You can define custom methods on it to extend its functionality as you wish.
type applyApcsDao struct {
	internalApplyApcsDao
}

var (
	// ApplyApcs is globally public accessible object for table apply_apcs operations.
	ApplyApcs = applyApcsDao{
		internal.NewApplyApcsDao(),
	}
)

// Fill with you ideas below.