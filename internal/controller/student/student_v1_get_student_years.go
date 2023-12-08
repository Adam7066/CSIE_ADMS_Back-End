package student

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/student/v1"
)

func (c *ControllerV1) GetStudentYears(ctx context.Context, req *v1.GetStudentYearsReq) (res *v1.GetStudentYearsRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetStudentYearsRes{
			Error: err.Error(),
		})
	}

	m := dao.Students.Ctx(ctx)
	admissionYears, err := m.Fields("admission_year").
		Distinct().OrderAsc("admission_year").All()
	if err != nil {
		r.Response.WriteJsonExit(v1.GetStudentYearsRes{
			Error: err.Error(),
		})
	}

	resData := make([]int, 0)
	for _, admissionYear := range admissionYears {
		resData = append(resData, admissionYear["admission_year"].Int())
	}

	r.Response.WriteJson(v1.GetStudentYearsRes{
		Data: resData,
	})
	return
}
