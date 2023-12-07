package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadSemesterRankReq struct {
	g.Meta   `path:"/rank/upload_semester" tags:"UploadSemesterRank" method:"post" mime:"multipart/form-data"`
	Year     int               `p:"year" v:"required"`
	Semester int               `p:"semester" v:"required"`
	File     *ghttp.UploadFile `p:"file" type:"file"`
}

type UploadSemesterRankRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
