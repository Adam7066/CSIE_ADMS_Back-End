package admsInfo

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"CSIE_ADMS_Back-End/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/admsInfo/v1"
)

func (c *ControllerV1) GetStudentsStarPlan(ctx context.Context, req *v1.GetStudentsStarPlanReq) (res *v1.GetStudentsStarPlanRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetStudentsStarPlanRes{
			Error: err.Error(),
		})
	}

	admissionMethodId, _ := g.Model("admission_methods").
		Fields("id").Where("name", "繁星推薦").Value()

	resData := make([]v1.GetStudentsStarPlanResData, 0)

	for _, year := range req.Years {
		students := make([]entity.Students, 0)
		m := dao.Students.Ctx(ctx)
		err := m.Where("admission_year=? AND admission_method_id=?", year, admissionMethodId).Scan(&students)
		if err != nil {
			r.Response.WriteJsonExit(v1.GetStudentsStarPlanRes{
				Error: err.Error(),
			})
		}
		for _, student := range students {
			tmpResData := v1.GetStudentsStarPlanResData{}
			tmpResData.StudentId = student.Id
			tmpResData.AdmissionYear = student.AdmissionYear
			tmpResData.StudentCode = student.StudentCode
			tmpResData.Name = student.Name

			graduatedSchool, _ := g.Model("schools").
				Fields("name").Where("id", student.GraduatedSchoolId).Value()
			tmpResData.GraduatedSchool = graduatedSchool.String()

			identityCategory, _ := g.Model("identity_categories").
				Fields("name").Where("id", student.IdentityCategoryId).Value()
			tmpResData.IdentityCategory = identityCategory.String()

			gsatScoreId, _ := g.Model("star_plan").
				Fields("gsat_score_id").Where("stu_id", student.Id).Value()

			gsatScore := entity.GsatScore{}
			err = g.Model("gsat_score").Where("id", gsatScoreId).Scan(&gsatScore)
			if err != nil {
				r.Response.WriteJsonExit(v1.GetStudentsStarPlanRes{
					Error: err.Error(),
				})
			}

			tmpResData.Chinese = gsatScore.Chinese
			tmpResData.English = gsatScore.English
			tmpResData.Math = gsatScore.Math
			tmpResData.Nature = gsatScore.Nature
			tmpResData.Society = gsatScore.Society
			tmpResData.Listening = gsatScore.Listening
			tmpResData.SumScore = gsatScore.SumScore

			resData = append(resData, tmpResData)
		}
	}

	r.Response.WriteJson(v1.GetStudentsStarPlanRes{
		Data: resData,
	})
	return
}
