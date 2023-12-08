// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package student

import (
	"context"

	"CSIE_ADMS_Back-End/api/student/v1"
)

type IStudentV1 interface {
	StudentGetYears(ctx context.Context, req *v1.StudentGetYearsReq) (res *v1.StudentGetYearsRes, err error)
}
