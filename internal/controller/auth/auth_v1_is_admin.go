package auth

import (
	"CSIE_ADMS_Back-End/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"CSIE_ADMS_Back-End/api/auth/v1"
)

func (c *ControllerV1) IsAdmin(ctx context.Context, req *v1.IsAdminReq) (res *v1.IsAdminRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.IsAdminRes{
			Error: err.Error(),
		})
	}
	r.Response.WriteJson(v1.IsAdminRes{
		Data: utility.IsAdmin(r.GetCtxVar("id").Int()),
	})
	return
}
