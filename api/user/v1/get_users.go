package v1

import "github.com/gogf/gf/v2/frame/g"

type GetUsersReq struct {
	g.Meta `path:"/users" tags:"GetUsers" method:"get"`
}

type GetUsersRes struct {
	Error string           `json:"error"`
	Data  []GetUserResData `json:"data"`
}
