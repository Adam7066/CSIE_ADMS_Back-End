package rank

import (
	v1 "CSIE_ADMS_Back-End/api/rank/v1"
	"CSIE_ADMS_Back-End/internal/model/entity"
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetRankByStuId(ctx context.Context, req *v1.GetRankByStuIdReq) (res *v1.GetRankByStuIdRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetRankByStuIdRes{
			Error: err.Error(),
		})
	}

	resData := make([]v1.GetRankByStuIdResData, 0)
	for _, stuId := range req.StuIds {
		var semesterRanks []entity.SemesterRank
		err := g.Model("semester_rank").Where("stu_id", stuId).Scan(&semesterRanks)
		if err != nil {
			r.Response.WriteJsonExit(v1.GetRankByStuIdRes{
				Error: err.Error(),
			})
		}
		var yearsRanks []entity.YearsRank
		err = g.Model("years_rank").Where("stu_id", stuId).Scan(&yearsRanks)
		if err != nil {
			r.Response.WriteJsonExit(v1.GetRankByStuIdRes{
				Error: err.Error(),
			})
		}

		tmpResSemesterRank := make(map[string]int)
		for _, semesterRank := range semesterRanks {
			key := strconv.Itoa(semesterRank.Year) + "-" + strconv.Itoa(semesterRank.Semester)
			tmpResSemesterRank[key] = semesterRank.DepartmentRank
		}
		tmpResYearsRank := make(map[string]int)
		for _, yearsRank := range yearsRanks {
			key := strconv.Itoa(yearsRank.Year) + "-" + strconv.Itoa(yearsRank.Semester)
			tmpResYearsRank[key] = yearsRank.DepartmentRank
		}
		resData = append(resData, v1.GetRankByStuIdResData{
			Semester: tmpResSemesterRank,
			Years:    tmpResYearsRank,
		})
	}
	r.Response.WriteJson(v1.GetRankByStuIdRes{
		Data: resData,
	})
	return
}
