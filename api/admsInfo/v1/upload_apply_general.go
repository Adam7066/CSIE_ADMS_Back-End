package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadApplyGeneralReq struct {
	g.Meta `path:"/upload_apply_general" tags:"UploadApplyGeneral" method:"post" mime:"multipart/form-data"`
	Year   int               `p:"year" v:"required"`
	File   *ghttp.UploadFile `p:"file" type:"file"`
}

type UploadApplyGeneralRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}
