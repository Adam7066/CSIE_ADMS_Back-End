package v1

import "github.com/gogf/gf/v2/frame/g"

type GetStudentByYearResData struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	StudentCode      string `json:"student_code"`
	AdmissionYear    int    `json:"admission_year"`
	AdmissionMethod  string `json:"admission_method"`
	AdmissionGroup   string `json:"admission_group"`
	IdentityCategory string `json:"identity_category"`
	GraduatedSchool  string `json:"graduated_school"`
}

type GetStudentByYearReq struct {
	g.Meta `path:"/students" tags:"GetStudentByYear" method:"get"`
	Year   int `json:"year" v:"required"`
}

type GetStudentByYearRes struct {
	Error string                    `json:"error"`
	Data  []GetStudentByYearResData `json:"data"`
}
