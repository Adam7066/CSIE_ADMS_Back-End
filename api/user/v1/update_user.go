package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateUserReq struct {
	g.Meta   `path:"/update_user" tags:"UpdateUser" method:"post"`
	Id       int    `p:"id"`
	Email    string `p:"email" v:"email"`
	Password string `p:"password" v:"password|length:10,18"`
	Username string `p:"username"`
	Role     string `p:"role"`
}

type UpdateUserRes struct {
	Error string `json:"error"`
	Data  any    `json:"data"`
}
