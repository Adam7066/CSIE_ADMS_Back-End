package v1

import "github.com/gogf/gf/v2/frame/g"

type GetRankByStuIdResData struct {
	Semester map[string]int `json:"semester"`
	Years    map[string]int `json:"years"`
}

type GetRankByStuIdReq struct {
	g.Meta `path:"/ranks" tags:"GetRankByStuId" method:"get"`
	StuIds []int `json:"stu_ids" v:"required"`
}

type GetRankByStuIdRes struct {
	Error string                  `json:"error"`
	Data  []GetRankByStuIdResData `json:"data"`
}
