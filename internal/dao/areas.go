// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"CSIE_ADMS_Back-End/internal/dao/internal"
)

// internalAreasDao is internal type for wrapping internal DAO implements.
type internalAreasDao = *internal.AreasDao

// areasDao is the data access object for table areas.
// You can define custom methods on it to extend its functionality as you wish.
type areasDao struct {
	internalAreasDao
}

var (
	// Areas is globally public accessible object for table areas operations.
	Areas = areasDao{
		internal.NewAreasDao(),
	}
)

// Fill with you ideas below.