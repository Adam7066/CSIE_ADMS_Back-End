package user

import (
	"CSIE_ADMS_Back-End/internal/dao"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/user/v1"
)

func (c *ControllerV1) GetUsers(ctx context.Context, req *v1.GetUsersReq) (res *v1.GetUsersRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.GetUsersRes{
			Error: err.Error(),
		})
	}

	id := r.GetCtxVar("id").Int()
	if !utility.IsAdmin(id) {
		r.Response.WriteJsonExit(v1.GetUsersRes{
			Error: "permission denied",
		})
	}

	resData := make([]v1.GetUserResData, 0)
	m := dao.Users.Ctx(ctx)
	err = m.Scan(&resData)
	if err != nil {
		r.Response.WriteJsonExit(v1.GetUsersRes{
			Error: err.Error(),
		})
	}

	r.Response.WriteJson(v1.GetUsersRes{
		Data: resData,
	})
	return
}
