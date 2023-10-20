// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AreasDao is the data access object for table areas.
type AreasDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns AreasColumns // columns contains all the column names of Table for convenient usage.
}

// AreasColumns defines and stores column names for table areas.
type AreasColumns struct {
	Id   string //
	Name string //
}

// areasColumns holds the columns for table areas.
var areasColumns = AreasColumns{
	Id:   "id",
	Name: "name",
}

// NewAreasDao creates and returns a new DAO object for table data access.
func NewAreasDao() *AreasDao {
	return &AreasDao{
		group:   "default",
		table:   "areas",
		columns: areasColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AreasDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AreasDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AreasDao) Columns() AreasColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AreasDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AreasDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AreasDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
