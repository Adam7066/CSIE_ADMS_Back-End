// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"CSIE_ADMS_Back-End/internal/dao/internal"
)

// internalApplyGeneralDao is internal type for wrapping internal DAO implements.
type internalApplyGeneralDao = *internal.ApplyGeneralDao

// applyGeneralDao is the data access object for table apply_general.
// You can define custom methods on it to extend its functionality as you wish.
type applyGeneralDao struct {
	internalApplyGeneralDao
}

var (
	// ApplyGeneral is globally public accessible object for table apply_general operations.
	ApplyGeneral = applyGeneralDao{
		internal.NewApplyGeneralDao(),
	}
)

// Fill with you ideas below.
