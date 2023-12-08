// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package student

import (
	"context"

	"CSIE_ADMS_Back-End/api/student/v1"
)

type IStudentV1 interface {
	GetStudentByYear(ctx context.Context, req *v1.GetStudentByYearReq) (res *v1.GetStudentByYearRes, err error)
	GetStudentYears(ctx context.Context, req *v1.GetStudentYearsReq) (res *v1.GetStudentYearsRes, err error)
}
