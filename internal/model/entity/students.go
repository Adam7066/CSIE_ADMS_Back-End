// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Students is the golang structure for table students.
type Students struct {
	Id                 int         `json:"id"                 ` //
	Name               string      `json:"name"               ` //
	GenderId           int         `json:"genderId"           ` //
	StudentCode        string      `json:"studentCode"        ` //
	DegreeId           int         `json:"degreeId"           ` //
	AdmissionMethodId  int         `json:"admissionMethodId"  ` //
	AdmissionGroupId   int         `json:"admissionGroupId"   ` //
	IdentityCategoryId int         `json:"identityCategoryId" ` //
	GraduatedSchoolId  int         `json:"graduatedSchoolId"  ` //
	CreatedAt          *gtime.Time `json:"createdAt"          ` //
	UpdatedAt          *gtime.Time `json:"updatedAt"          ` //
	DeletedAt          *gtime.Time `json:"deletedAt"          ` //
}
