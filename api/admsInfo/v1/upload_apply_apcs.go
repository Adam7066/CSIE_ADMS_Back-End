package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadApplyApcsReq struct {
	g.Meta `path:"/upload_apply_apcs" tags:"UploadApplyApcs" method:"post" mime:"multipart/form-data"`
	Year   int               `p:"year" v:"required"`
	File   *ghttp.UploadFile `p:"file" type:"file"`
}

type UploadApplyApcsRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
