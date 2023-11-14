// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyGeneralDao is the data access object for table apply_general.
type ApplyGeneralDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ApplyGeneralColumns // columns contains all the column names of Table for convenient usage.
}

// ApplyGeneralColumns defines and stores column names for table apply_general.
type ApplyGeneralColumns struct {
	Id                  string //
	StuId               string //
	GsatScoreId         string //
	SchoolDeptCode      string //
	GsatCalScore        string //
	ProjectScore1       string //
	ProjectTestScore    string //
	SelectionTotalScore string //
	AdmsStatus          string //
	AdmsRank            string //
	AdmsTotalRank       string //
}

// applyGeneralColumns holds the columns for table apply_general.
var applyGeneralColumns = ApplyGeneralColumns{
	Id:                  "id",
	StuId:               "stu_id",
	GsatScoreId:         "gsat_score_id",
	SchoolDeptCode:      "school_dept_code",
	GsatCalScore:        "gsat_cal_score",
	ProjectScore1:       "project_score_1",
	ProjectTestScore:    "project_test_score",
	SelectionTotalScore: "selection_total_score",
	AdmsStatus:          "adms_status",
	AdmsRank:            "adms_rank",
	AdmsTotalRank:       "adms_total_rank",
}

// NewApplyGeneralDao creates and returns a new DAO object for table data access.
func NewApplyGeneralDao() *ApplyGeneralDao {
	return &ApplyGeneralDao{
		group:   "default",
		table:   "apply_general",
		columns: applyGeneralColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplyGeneralDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplyGeneralDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplyGeneralDao) Columns() ApplyGeneralColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplyGeneralDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplyGeneralDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplyGeneralDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
