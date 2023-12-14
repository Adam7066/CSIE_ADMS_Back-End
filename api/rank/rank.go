// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package rank

import (
	"context"

	"CSIE_ADMS_Back-End/api/rank/v1"
)

type IRankV1 interface {
	GetRankByStuId(ctx context.Context, req *v1.GetRankByStuIdReq) (res *v1.GetRankByStuIdRes, err error)
	UploadSemesterRank(ctx context.Context, req *v1.UploadSemesterRankReq) (res *v1.UploadSemesterRankRes, err error)
	UploadYearsRank(ctx context.Context, req *v1.UploadYearsRankReq) (res *v1.UploadYearsRankRes, err error)
}
