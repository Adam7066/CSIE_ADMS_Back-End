// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GsatScoreDao is the data access object for table gsat_score.
type GsatScoreDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns GsatScoreColumns // columns contains all the column names of Table for convenient usage.
}

// GsatScoreColumns defines and stores column names for table gsat_score.
type GsatScoreColumns struct {
	Id             string //
	GsatTestNumber string //
	Chinese        string //
	English        string //
	Math           string //
	Nature         string //
	Society        string //
	Listening      string //
	SumScore       string //
}

// gsatScoreColumns holds the columns for table gsat_score.
var gsatScoreColumns = GsatScoreColumns{
	Id:             "id",
	GsatTestNumber: "gsat_test_number",
	Chinese:        "chinese",
	English:        "english",
	Math:           "math",
	Nature:         "nature",
	Society:        "society",
	Listening:      "listening",
	SumScore:       "sum_score",
}

// NewGsatScoreDao creates and returns a new DAO object for table data access.
func NewGsatScoreDao() *GsatScoreDao {
	return &GsatScoreDao{
		group:   "default",
		table:   "gsat_score",
		columns: gsatScoreColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GsatScoreDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GsatScoreDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GsatScoreDao) Columns() GsatScoreColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GsatScoreDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GsatScoreDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GsatScoreDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
