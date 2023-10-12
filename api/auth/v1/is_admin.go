package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IsAdminReq struct {
	g.Meta `path:"/is_admin" tags:"IsAdmin" method:"GET"`
}

type IsAdminRes struct {
	Error string `json:"error"`
	Data  bool   `json:"data"`
}
