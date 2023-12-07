// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// YearsRankDao is the data access object for table years_rank.
type YearsRankDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns YearsRankColumns // columns contains all the column names of Table for convenient usage.
}

// YearsRankColumns defines and stores column names for table years_rank.
type YearsRankColumns struct {
	Id                    string //
	StuId                 string //
	Year                  string //
	Semester              string //
	TotalPoints           string //
	TotalCredits          string //
	EarnedCredits         string //
	FailedCredits         string //
	AverageScore          string //
	DepartmentRank        string //
	DepartmentRankPercent string //
	CurrentStudentStatus  string //
	CreatedAt             string //
	UpdatedAt             string //
	DeletedAt             string //
}

// yearsRankColumns holds the columns for table years_rank.
var yearsRankColumns = YearsRankColumns{
	Id:                    "id",
	StuId:                 "stu_id",
	Year:                  "year",
	Semester:              "semester",
	TotalPoints:           "total_points",
	TotalCredits:          "total_credits",
	EarnedCredits:         "earned_credits",
	FailedCredits:         "failed_credits",
	AverageScore:          "average_score",
	DepartmentRank:        "department_rank",
	DepartmentRankPercent: "department_rank_percent",
	CurrentStudentStatus:  "current_student_status",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
	DeletedAt:             "deleted_at",
}

// NewYearsRankDao creates and returns a new DAO object for table data access.
func NewYearsRankDao() *YearsRankDao {
	return &YearsRankDao{
		group:   "default",
		table:   "years_rank",
		columns: yearsRankColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *YearsRankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *YearsRankDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *YearsRankDao) Columns() YearsRankColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *YearsRankDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *YearsRankDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *YearsRankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
