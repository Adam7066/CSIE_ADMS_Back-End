// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SemesterRank is the golang structure of table semester_rank for DAO operations like Where/Data.
type SemesterRank struct {
	g.Meta                `orm:"table:semester_rank, do:true"`
	Id                    interface{} //
	StuId                 interface{} //
	Year                  interface{} //
	Semester              interface{} //
	TotalPoints           interface{} //
	TotalCredits          interface{} //
	EarnedCredits         interface{} //
	FailedCredits         interface{} //
	AverageScore          interface{} //
	DepartmentRank        interface{} //
	DepartmentRankPercent interface{} //
	CurrentStudentStatus  interface{} //
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
	DeletedAt             *gtime.Time //
}
