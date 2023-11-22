package v1

import "github.com/gogf/gf/v2/frame/g"

type GetYearsApplyReq struct {
	g.Meta `path:"/adms_info/years/apply" tags:"GetYearsApply" method:"get"`
}

type GetYearsApplyRes struct {
	Error string `json:"error"`
	Data  []int  `json:"data"`
}
