// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StarPlanDao is the data access object for table star_plan.
type StarPlanDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns StarPlanColumns // columns contains all the column names of Table for convenient usage.
}

// StarPlanColumns defines and stores column names for table star_plan.
type StarPlanColumns struct {
	Id                            string //
	StuId                         string //
	GsatScoreId                   string //
	SchoolDeptCode                string //
	RecommendedOrder              string //
	AdmsOrder                     string //
	SchoolRankingPercentage       string //
	ChineseRankingPercentage      string //
	EnglishRankingPercentage      string //
	MathRankingPercentage         string //
	PhysicsRankingPercentage      string //
	ChemistryRankingPercentage    string //
	BiologyRankingPercentage      string //
	EarthScienceRankingPercentage string //
	HistoryRankingPercentage      string //
	GeographyRankingPercentage    string //
	CitizenshipRankingPercentage  string //
	MusicRankingPercentage        string //
	ArtRankingPercentage          string //
	DanceRankingPercentage        string //
	PhysicalRankingPercentage     string //
	StudySubjectCourseClass       string //
}

// starPlanColumns holds the columns for table star_plan.
var starPlanColumns = StarPlanColumns{
	Id:                            "id",
	StuId:                         "stu_id",
	GsatScoreId:                   "gsat_score_id",
	SchoolDeptCode:                "school_dept_code",
	RecommendedOrder:              "recommended_order",
	AdmsOrder:                     "adms_order",
	SchoolRankingPercentage:       "school_ranking_percentage",
	ChineseRankingPercentage:      "chinese_ranking_percentage",
	EnglishRankingPercentage:      "english_ranking_percentage",
	MathRankingPercentage:         "math_ranking_percentage",
	PhysicsRankingPercentage:      "physics_ranking_percentage",
	ChemistryRankingPercentage:    "chemistry_ranking_percentage",
	BiologyRankingPercentage:      "biology_ranking_percentage",
	EarthScienceRankingPercentage: "earth_science_ranking_percentage",
	HistoryRankingPercentage:      "history_ranking_percentage",
	GeographyRankingPercentage:    "geography_ranking_percentage",
	CitizenshipRankingPercentage:  "citizenship_ranking_percentage",
	MusicRankingPercentage:        "music_ranking_percentage",
	ArtRankingPercentage:          "art_ranking_percentage",
	DanceRankingPercentage:        "dance_ranking_percentage",
	PhysicalRankingPercentage:     "physical_ranking_percentage",
	StudySubjectCourseClass:       "study_subject_course_class",
}

// NewStarPlanDao creates and returns a new DAO object for table data access.
func NewStarPlanDao() *StarPlanDao {
	return &StarPlanDao{
		group:   "default",
		table:   "star_plan",
		columns: starPlanColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StarPlanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StarPlanDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StarPlanDao) Columns() StarPlanColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StarPlanDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StarPlanDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StarPlanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
