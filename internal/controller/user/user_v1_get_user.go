package user

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/user/v1"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetUserRes{
			Error: err.Error(),
		})
	}

	id := r.GetCtxVar("id").Int()
	if req.Id > 0 {
		if req.Id != id {
			if !utility.IsAdmin(id) {
				r.Response.WriteJsonExit(v1.GetUserRes{
					Error: "permission denied",
				})
			}
			id = req.Id
		}
	}

	resData := v1.GetUserResData{}
	m := dao.Users.Ctx(ctx)
	err = m.Scan(&resData, "id", id)
	if err != nil {
		r.Response.WriteJsonExit(v1.GetUserRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.GetUserRes{
		Data: &resData,
	})
	return
}
