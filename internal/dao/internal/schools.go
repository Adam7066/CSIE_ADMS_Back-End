// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SchoolsDao is the data access object for table schools.
type SchoolsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SchoolsColumns // columns contains all the column names of Table for convenient usage.
}

// SchoolsColumns defines and stores column names for table schools.
type SchoolsColumns struct {
	Id         string //
	DegreeId   string //
	Name       string //
	SchoolCode string //
	CityId     string //
	AreaId     string //
}

// schoolsColumns holds the columns for table schools.
var schoolsColumns = SchoolsColumns{
	Id:         "id",
	DegreeId:   "degree_id",
	Name:       "name",
	SchoolCode: "school_code",
	CityId:     "city_id",
	AreaId:     "area_id",
}

// NewSchoolsDao creates and returns a new DAO object for table data access.
func NewSchoolsDao() *SchoolsDao {
	return &SchoolsDao{
		group:   "default",
		table:   "schools",
		columns: schoolsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SchoolsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SchoolsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SchoolsDao) Columns() SchoolsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SchoolsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SchoolsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SchoolsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
