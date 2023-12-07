// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// YearsRank is the golang structure for table years_rank.
type YearsRank struct {
	Id                    int         `json:"id"                    ` //
	StuId                 int         `json:"stuId"                 ` //
	Year                  int         `json:"year"                  ` //
	Semester              int         `json:"semester"              ` //
	TotalPoints           float64     `json:"totalPoints"           ` //
	TotalCredits          int         `json:"totalCredits"          ` //
	EarnedCredits         int         `json:"earnedCredits"         ` //
	FailedCredits         int         `json:"failedCredits"         ` //
	AverageScore          float64     `json:"averageScore"          ` //
	DepartmentRank        int         `json:"departmentRank"        ` //
	DepartmentRankPercent float64     `json:"departmentRankPercent" ` //
	CurrentStudentStatus  string      `json:"currentStudentStatus"  ` //
	CreatedAt             *gtime.Time `json:"createdAt"             ` //
	UpdatedAt             *gtime.Time `json:"updatedAt"             ` //
	DeletedAt             *gtime.Time `json:"deletedAt"             ` //
}
