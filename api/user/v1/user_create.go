package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateUserReq struct {
	g.Meta   `path:"/create_user" tags:"CreateUser" method:"post"`
	Email    string `p:"email" v:"required|email"`
	Password string `p:"password" v:"required|password|length:10,18"`
	Username string `p:"username" v:"required"`
	Role     string `p:"role"`
}

type CreateUserRes struct {
	Error string `json:"error"`
	Data  any    `json:"data"`
}
