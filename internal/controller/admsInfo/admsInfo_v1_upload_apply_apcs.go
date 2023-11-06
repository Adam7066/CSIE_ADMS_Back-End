package admsInfo

import (
	"CSIE_ADMS_Back-End/utility"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"

	"CSIE_ADMS_Back-End/api/admsInfo/v1"
)

func (c *ControllerV1) UploadApplyApcs(ctx context.Context, req *v1.UploadApplyApcsReq) (res *v1.UploadApplyApcsRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.UploadApplyApcsRes{
			Error: err.Error(),
		})
	}

	// Check if admin
	if !utility.IsAdmin(r.GetCtxVar("id").Int()) {
		r.Response.WriteJsonExit(v1.UploadApplyApcsRes{
			Error: "Permission denied",
		})
	}

	// Get upload file
	if req.File == nil {
		r.Response.WriteJsonExit(v1.UploadApplyApcsRes{
			Error: "No file uploaded",
		})
	}
	file := r.GetUploadFile("file")

	// Check file type
	fileType := file.Header.Get("content-type")
	if fileType != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		r.Response.WriteJsonExit(v1.UploadApplyApcsRes{
			Error: "Only .xlsx file is allowed",
		})
	}

	// Save file
	filename, err := file.Save("./upload/", true)
	if err != nil {
		r.Response.WriteJsonExit(v1.UploadApplyApcsRes{
			Error: err.Error(),
		})
	}

	// Handle file
	err = handleUploadApplyApcsFile(req.Year, filename)
	if err != nil {
		r.Response.WriteJsonExit(v1.UploadApplyApcsRes{
			Error: err.Error(),
		})
	}

	// Remove file if handled successfully
	os.Remove(fmt.Sprintf("./upload/%s", filename))

	r.Response.WriteJson(v1.UploadApplyApcsRes{
		Data: "Success",
	})
	return
}

func handleUploadApplyApcsFile(year int, filename string) error {
	path := fmt.Sprintf("./upload/%s", filename)
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		degreeId, _ := g.Model("degrees").Fields("id").Where("name", "大學").Value()
		admissionMethodId, _ := g.Model("admission_methods").Fields("id").Where("name", "申請入學").Value()
		admissionGroupId, _ := g.Model("admission_groups").Fields("id").Where("name", "APCS組").Value()
		student := g.Map{
			"degree_id":           degreeId,
			"admission_year":      year,
			"admission_method_id": admissionMethodId,
			"admission_group_id":  admissionGroupId,
		}
		gsatScore := g.Map{}
		applyApcs := g.Map{}
		for j, col := range row {
			col = strings.TrimSpace(col)
			if col == "" || col == "無成績" {
				col = "-1"
			}
			switch rows[0][j] {
			case "姓名":
				student["name"] = col
			case "學號":
				student["student_code"] = col
			case "高中代碼":
				student["graduated_school_id"], _ = g.Model("schools").Fields("id").Where("school_code", col).Value()

			case "學測應試號碼":
				gsatScore["gsat_test_number"] = col
			case "國文級分":
				gsatScore["chinese"], _ = strconv.Atoi(col)
			case "英文級分":
				gsatScore["english"], _ = strconv.Atoi(col)
			case "數學級分":
				gsatScore["math"], _ = strconv.Atoi(col)
			case "自然級分":
				gsatScore["nature"], _ = strconv.Atoi(col)
			case "社會級分":
				gsatScore["society"], _ = strconv.Atoi(col)
			case "英聽":
				gsatScore["listening"] = col

			case "校系代碼":
				applyApcs["school_dept_code"] = col
			case "學科能力測驗成績":
				applyApcs["gsat_cal_score"], _ = strconv.ParseFloat(col, 64)
			case "指定項目一成績":
				applyApcs["project_score_1"], _ = strconv.ParseFloat(col, 64)
			case "指定項目二成績":
				applyApcs["project_score_2"], _ = strconv.ParseFloat(col, 64)
			case "指定項目甄試成績":
				applyApcs["project_test_score"], _ = strconv.ParseFloat(col, 64)
			case "甄選總成績":
				applyApcs["selection_total_score"], _ = strconv.ParseFloat(col, 64)
			case "招生名額錄取狀態":
				applyApcs["adms_status"] = col
			case "招生名額名次":
				applyApcs["adms_rank"], _ = strconv.Atoi(col)
			case "招生名額總名次":
				applyApcs["adms_total_rank"], _ = strconv.Atoi(col)
			}
		}
		studentRes, err := g.Model("students").Data(student).Insert()
		if err != nil {
			return err
		}
		gsatScoreRes, err := g.Model("gsat_score").Data(gsatScore).Insert()
		if err != nil {
			return err
		}
		applyApcs["stu_id"], _ = studentRes.LastInsertId()
		applyApcs["gsat_score_id"], _ = gsatScoreRes.LastInsertId()
		_, err = g.Model("apply_apcs").Data(applyApcs).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}