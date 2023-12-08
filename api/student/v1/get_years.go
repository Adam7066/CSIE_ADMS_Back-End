package v1

import "github.com/gogf/gf/v2/frame/g"

type StudentGetYearsReq struct {
	g.Meta `path:"/student/years" tags:"StudentGetYears" method:"get"`
}

type StudentGetYearsRes struct {
	Error string `json:"error"`
	Data  []int  `json:"data"`
}
