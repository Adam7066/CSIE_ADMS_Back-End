// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Students is the golang structure of table students for DAO operations like Where/Data.
type Students struct {
	g.Meta             `orm:"table:students, do:true"`
	Id                 interface{} //
	Name               interface{} //
	GenderId           interface{} //
	StudentCode        interface{} //
	DegreeId           interface{} //
	AdmissionYear      interface{} //
	AdmissionMethodId  interface{} //
	AdmissionGroupId   interface{} //
	IdentityCategoryId interface{} //
	GraduatedSchoolId  interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
	DeletedAt          *gtime.Time //
}
