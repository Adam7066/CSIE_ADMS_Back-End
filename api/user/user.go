// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"context"
	
	"CSIE_ADMS_Back-End/api/user/v1"
)

type IUserV1 interface {
	CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error)
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
}


