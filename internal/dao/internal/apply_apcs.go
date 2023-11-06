// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyApcsDao is the data access object for table apply_apcs.
type ApplyApcsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ApplyApcsColumns // columns contains all the column names of Table for convenient usage.
}

// ApplyApcsColumns defines and stores column names for table apply_apcs.
type ApplyApcsColumns struct {
	Id                  string //
	StuId               string //
	GsatScoreId         string //
	SchoolDeptCode      string //
	GsatCalScore        string //
	ProjectScore1       string //
	ProjectScore2       string //
	ProjectTestScore    string //
	SelectionTotalScore string //
	AdmsStatus          string //
	AdmsRank            string //
	AdmsTotalRank       string //
}

// applyApcsColumns holds the columns for table apply_apcs.
var applyApcsColumns = ApplyApcsColumns{
	Id:                  "id",
	StuId:               "stu_id",
	GsatScoreId:         "gsat_score_id",
	SchoolDeptCode:      "school_dept_code",
	GsatCalScore:        "gsat_cal_score",
	ProjectScore1:       "project_score_1",
	ProjectScore2:       "project_score_2",
	ProjectTestScore:    "project_test_score",
	SelectionTotalScore: "selection_total_score",
	AdmsStatus:          "adms_status",
	AdmsRank:            "adms_rank",
	AdmsTotalRank:       "adms_total_rank",
}

// NewApplyApcsDao creates and returns a new DAO object for table data access.
func NewApplyApcsDao() *ApplyApcsDao {
	return &ApplyApcsDao{
		group:   "default",
		table:   "apply_apcs",
		columns: applyApcsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplyApcsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplyApcsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplyApcsDao) Columns() ApplyApcsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplyApcsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplyApcsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplyApcsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
