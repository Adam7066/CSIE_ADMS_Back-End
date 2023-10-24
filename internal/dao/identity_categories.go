// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"CSIE_ADMS_Back-End/internal/dao/internal"
)

// internalIdentityCategoriesDao is internal type for wrapping internal DAO implements.
type internalIdentityCategoriesDao = *internal.IdentityCategoriesDao

// identityCategoriesDao is the data access object for table identity_categories.
// You can define custom methods on it to extend its functionality as you wish.
type identityCategoriesDao struct {
	internalIdentityCategoriesDao
}

var (
	// IdentityCategories is globally public accessible object for table identity_categories operations.
	IdentityCategories = identityCategoriesDao{
		internal.NewIdentityCategoriesDao(),
	}
)

// Fill with you ideas below.