package student

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/student/v1"
)

func (c *ControllerV1) GetStudentByYear(ctx context.Context, req *v1.GetStudentByYearReq) (res *v1.GetStudentByYearRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetStudentYearsRes{
			Error: err.Error(),
		})
	}

	m := dao.Students.Ctx(ctx)
	students, err := m.Where("admission_year", req.Year).OrderAsc("student_code").All()
	if err != nil {
		r.Response.WriteJsonExit(v1.GetStudentByYearRes{
			Error: err.Error(),
		})
	}

	resData := make([]v1.GetStudentByYearResData, 0)
	for _, student := range students {
		tempData := v1.GetStudentByYearResData{}
		tempData.Id = student["id"].Int()
		tempData.Name = student["name"].String()
		tempData.StudentCode = student["student_code"].String()
		tempData.AdmissionYear = student["admission_year"].Int()

		admissionMethod, _ := g.Model("admission_methods").
			Fields("name").Where("id", student["admission_method_id"]).Value()
		tempData.AdmissionMethod = admissionMethod.String()

		if student["admission_group_id"].Int() != 0 {
			admissionGroup, _ := g.Model("admission_groups").
				Fields("name").Where("id", student["admission_group_id"]).Value()
			tempData.AdmissionGroup = admissionGroup.String()
		}
		if student["identity_category_id"].Int() != 0 {
			identityCategory, _ := g.Model("identity_categories").
				Fields("name").Where("id", student["identity_category_id"]).Value()
			tempData.IdentityCategory = identityCategory.String()
		}
		if student["graduated_school_id"].Int() != 0 {
			graduatedSchool, _ := g.Model("schools").
				Fields("name").Where("id", student["graduated_school_id"]).Value()
			tempData.GraduatedSchool = graduatedSchool.String()
		}
		resData = append(resData, tempData)
	}

	r.Response.WriteJson(v1.GetStudentByYearRes{
		Data: resData,
	})
	return
}
