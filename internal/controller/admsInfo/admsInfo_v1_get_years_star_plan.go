package admsInfo

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/admsInfo/v1"
)

func (c *ControllerV1) GetYearsStarPlan(ctx context.Context, req *v1.GetYearsStarPlanReq) (res *v1.GetYearsStarPlanRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetYearsStarPlanRes{
			Error: err.Error(),
		})
	}

	m := dao.Students.Ctx(ctx)
	admissionMethodId, _ := g.Model("admission_methods").
		Fields("id").Where("name", "繁星推薦").Value()
	admissionYears, err := m.Fields("admission_year").
		Where("admission_method_id", admissionMethodId).
		Distinct().OrderAsc("admission_year").All()
	if err != nil {
		r.Response.WriteJsonExit(v1.GetYearsStarPlanRes{
			Error: err.Error(),
		})
	}

	resData := make([]int, 0)
	for _, admissionYear := range admissionYears {
		resData = append(resData, admissionYear["admission_year"].Int())
	}

	r.Response.WriteJson(v1.GetYearsStarPlanRes{
		Data: resData,
	})
	return
}
