// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admsInfo

import (
	"context"
	
	"CSIE_ADMS_Back-End/api/admsInfo/v1"
)

type IAdmsInfoV1 interface {
	UploadApplyApcs(ctx context.Context, req *v1.UploadApplyApcsReq) (res *v1.UploadApplyApcsRes, err error)
	UploadStarPlan(ctx context.Context, req *v1.UploadStarPlanReq) (res *v1.UploadStarPlanRes, err error)
}

