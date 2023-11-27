package v1

import "github.com/gogf/gf/v2/frame/g"

type GetYearsStarPlanReq struct {
	g.Meta `path:"/adms_info/years/star_plan" tags:"GetYearsStarPlan" method:"get"`
}

type GetYearsStarPlanRes struct {
	Error string `json:"error"`
	Data  []int  `json:"data"`
}
