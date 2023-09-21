package user

import (
	"CSIE_ADMS_Back-End/api/user/v1"
	"CSIE_ADMS_Back-End/internal/dao"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.CreateUserRes{
			Error: err.Error(),
		})
	}

	// Check if admin
	if !utility.IsAdmin(r.GetCtxVar("id").Int()) {
		r.Response.WriteJsonExit(v1.CreateUserRes{
			Error: "Permission denied",
		})
	}

	// Check if email exists
	m := dao.Users.Ctx(ctx)
	findEmail, err := m.One("email", req.Email)
	if err != nil {
		r.Response.WriteJsonExit(v1.CreateUserRes{
			Error: err.Error(),
		})
	}
	if findEmail != nil {
		r.Response.WriteJsonExit(v1.CreateUserRes{
			Error: "Email already exists",
		})
	}

	//Create user
	if req.Role == "" {
		req.Role = "user"
	}
	encodedHash, err := utility.HashPWD(req.Password)
	if err != nil {
		r.Response.WriteJsonExit(v1.CreateUserRes{
			Error: err.Error(),
		})
	}
	_, err = m.Data(g.Map{
		"email":    req.Email,
		"password": encodedHash,
		"username": req.Username,
		"role":     req.Role,
	}).Insert()
	if err != nil {
		r.Response.WriteJsonExit(v1.CreateUserRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.CreateUserRes{
		Data: "Success",
	})
	return
}
