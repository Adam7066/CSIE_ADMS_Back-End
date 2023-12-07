package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadYearsRankReq struct {
	g.Meta   `path:"/rank/upload_years" tags:"UploadYearsRank" method:"post" mime:"multipart/form-data"`
	Year     int               `p:"year" v:"required"`
	Semester int               `p:"semester" v:"required"`
	File     *ghttp.UploadFile `p:"file" type:"file"`
}

type UploadYearsRankRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
