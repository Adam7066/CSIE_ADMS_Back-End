// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CitiesDao is the data access object for table cities.
type CitiesDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns CitiesColumns // columns contains all the column names of Table for convenient usage.
}

// CitiesColumns defines and stores column names for table cities.
type CitiesColumns struct {
	Id   string //
	Name string //
}

// citiesColumns holds the columns for table cities.
var citiesColumns = CitiesColumns{
	Id:   "id",
	Name: "name",
}

// NewCitiesDao creates and returns a new DAO object for table data access.
func NewCitiesDao() *CitiesDao {
	return &CitiesDao{
		group:   "default",
		table:   "cities",
		columns: citiesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CitiesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CitiesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CitiesDao) Columns() CitiesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CitiesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CitiesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CitiesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
