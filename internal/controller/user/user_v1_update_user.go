package user

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/user/v1"
)

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.UpdateUserRes{
			Error: err.Error(),
		})
	}

	tokenID := r.GetCtxVar("id").Int()
	id := tokenID
	if req.Id > 0 {
		if req.Id != id {
			if utility.IsAdmin(id) {
				r.Response.WriteJsonExit(v1.UpdateUserRes{
					Error: "permission denied",
				})
			}
			id = req.Id
		}
	}

	data := g.Map{}
	m := dao.Users.Ctx(ctx)
	// need admin permission
	if utility.IsAdmin(tokenID) {
		if req.Role != "" {
			data["role"] = req.Role
		}
	}
	if req.Email != "" {
		data["email"] = req.Email
	}
	if req.Password != "" {
		data["password"] = req.Password
	}
	if req.Username != "" {
		data["username"] = req.Username
	}
	_, err = m.Update(data, "id", id)
	if err != nil {
		r.Response.WriteJsonExit(v1.UpdateUserRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.UpdateUserRes{
		Data: "Success",
	})
	return
}
