// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SemesterRankDao is the data access object for table semester_rank.
type SemesterRankDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SemesterRankColumns // columns contains all the column names of Table for convenient usage.
}

// SemesterRankColumns defines and stores column names for table semester_rank.
type SemesterRankColumns struct {
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

// semesterRankColumns holds the columns for table semester_rank.
var semesterRankColumns = SemesterRankColumns{
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

// NewSemesterRankDao creates and returns a new DAO object for table data access.
func NewSemesterRankDao() *SemesterRankDao {
	return &SemesterRankDao{
		group:   "default",
		table:   "semester_rank",
		columns: semesterRankColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SemesterRankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SemesterRankDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SemesterRankDao) Columns() SemesterRankColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SemesterRankDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SemesterRankDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SemesterRankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
