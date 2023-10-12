// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package auth

import (
	"context"
	
	"CSIE_ADMS_Back-End/api/auth/v1"
)

type IAuthV1 interface {
	IsAdmin(ctx context.Context, req *v1.IsAdminReq) (res *v1.IsAdminRes, err error)
}


