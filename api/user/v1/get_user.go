package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type GetUserResData struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserReq struct {
	g.Meta `path:"/get_user" tags:"GetUser" method:"get"`
	Id     int `p:"id"`
}

type GetUserRes struct {
	Error string          `json:"error"`
	Data  *GetUserResData `json:"data"`
}
