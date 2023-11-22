package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetStudentsApplyResData struct {
	StudentId        int    `json:"student_id"`        // student_id
	AdmissionYear    int    `json:"admission_year"`    // 入學年度
	StudentCode      string `json:"student_code"`      // 學號
	Name             string `json:"name"`              // 姓名
	GraduatedSchool  string `json:"graduated_school"`  // 畢業學校
	AdmissionGroup   string `json:"admission_group"`   // 入學組別
	IdentityCategory string `json:"identity_category"` // 身分類別
	Chinese          int    `json:"chinese"`           // 國文級分
	English          int    `json:"english"`           // 英文級分
	Math             int    `json:"math"`              // 數學級分
	Nature           int    `json:"nature"`            // 自然級分
	Society          int    `json:"society"`           // 社會級分
	Listening        string `json:"listening"`         // 聽力測驗
	SumScore         int    `json:"sum_score"`         // 總級分
}

type GetStudentsApplyReq struct {
	g.Meta `path:"/adms_info/students/apply" tags:"GetStudentsApply" method:"post"`
	Years  []int `p:"years" v:"required"`
}

type GetStudentsApplyRes struct {
	Error string                    `json:"error"`
	Data  []GetStudentsApplyResData `json:"data"`
}
