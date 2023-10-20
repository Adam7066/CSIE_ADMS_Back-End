// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"CSIE_ADMS_Back-End/internal/dao/internal"
)

// internalSchoolsDao is internal type for wrapping internal DAO implements.
type internalSchoolsDao = *internal.SchoolsDao

// schoolsDao is the data access object for table schools.
// You can define custom methods on it to extend its functionality as you wish.
type schoolsDao struct {
	internalSchoolsDao
}

var (
	// Schools is globally public accessible object for table schools operations.
	Schools = schoolsDao{
		internal.NewSchoolsDao(),
	}
)

// Fill with you ideas below.
