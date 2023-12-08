package v1

import "github.com/gogf/gf/v2/frame/g"

type GetStudentYearsReq struct {
	g.Meta `path:"/student/years" tags:"GetStudentYears" method:"get"`
}

type GetStudentYearsRes struct {
	Error string `json:"error"`
	Data  []int  `json:"data"`
}
