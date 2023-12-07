package rank

import (
	"CSIE_ADMS_Back-End/api/rank/v1"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"os"
)

func (c *ControllerV1) UploadSemesterRank(ctx context.Context, req *v1.UploadSemesterRankReq) (res *v1.UploadSemesterRankRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.UploadSemesterRankRes{
			Error: err.Error(),
		})
	}

	// Check if admin
	if !utility.IsAdmin(r.GetCtxVar("id").Int()) {
		r.Response.WriteJsonExit(v1.UploadSemesterRankRes{
			Error: "Permission denied",
		})
	}

	// Get upload file
	if req.File == nil {
		r.Response.WriteJsonExit(v1.UploadSemesterRankRes{
			Error: "No file uploaded",
		})
	}
	file := r.GetUploadFile("file")

	// Check file type
	fileType := file.Header.Get("content-type")
	if fileType != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		r.Response.WriteJsonExit(v1.UploadSemesterRankRes{
			Error: "Only .xlsx file is allowed",
		})
	}

	// Save file
	filename, err := file.Save("./upload/", true)
	if err != nil {
		r.Response.WriteJsonExit(v1.UploadSemesterRankRes{
			Error: err.Error(),
		})
	}

	// Handle file
	err = handleUploadRankFile("semester_rank", filename, req.Year, req.Semester)
	if err != nil {
		r.Response.WriteJsonExit(v1.UploadSemesterRankRes{
			Error: err.Error(),
		})
	}

	// Remove file if handled successfully
	os.Remove(fmt.Sprintf("./upload/%s", filename))

	r.Response.WriteJson(v1.UploadSemesterRankRes{
		Data: "Success",
	})
	return
}
