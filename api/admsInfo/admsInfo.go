// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admsInfo

import (
	"context"
	
	"CSIE_ADMS_Back-End/api/admsInfo/v1"
)

type IAdmsInfoV1 interface {
	GetStudentsApply(ctx context.Context, req *v1.GetStudentsApplyReq) (res *v1.GetStudentsApplyRes, err error)
	GetStudentsStarPlan(ctx context.Context, req *v1.GetStudentsStarPlanReq) (res *v1.GetStudentsStarPlanRes, err error)
	GetYearsApply(ctx context.Context, req *v1.GetYearsApplyReq) (res *v1.GetYearsApplyRes, err error)
	GetYearsStarPlan(ctx context.Context, req *v1.GetYearsStarPlanReq) (res *v1.GetYearsStarPlanRes, err error)
	UploadApplyApcs(ctx context.Context, req *v1.UploadApplyApcsReq) (res *v1.UploadApplyApcsRes, err error)
	UploadApplyGeneral(ctx context.Context, req *v1.UploadApplyGeneralReq) (res *v1.UploadApplyGeneralRes, err error)
	UploadStarPlan(ctx context.Context, req *v1.UploadStarPlanReq) (res *v1.UploadStarPlanRes, err error)
}


