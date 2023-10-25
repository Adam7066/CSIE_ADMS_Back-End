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

func (c *ControllerV1) UploadStarPlan(ctx context.Context, req *v1.UploadStarPlanReq) (res *v1.UploadStarPlanRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(v1.UploadStarPlanRes{
			Error: err.Error(),
		})
	}

	// Check if admin
	if !utility.IsAdmin(r.GetCtxVar("id").Int()) {
		r.Response.WriteJsonExit(v1.UploadStarPlanRes{
			Error: "Permission denied",
		})
	}

	// Get upload file
	if req.File == nil {
		r.Response.WriteJsonExit(v1.UploadStarPlanRes{
			Error: "No file uploaded",
		})
	}
	file := r.GetUploadFile("file")

	// Check file type
	fileType := file.Header.Get("content-type")
	if fileType != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		r.Response.WriteJsonExit(v1.UploadStarPlanRes{
			Error: "Only .xlsx file is allowed",
		})
	}

	// Save file
	filename, err := file.Save("./upload/", true)
	if err != nil {
		r.Response.WriteJsonExit(v1.UploadStarPlanRes{
			Error: err.Error(),
		})
	}

	// Handle file
	err = handleUploadedFile(req.Year, filename)
	if err != nil {
		r.Response.WriteJsonExit(v1.UploadStarPlanRes{
			Error: err.Error(),
		})
	}

	// Remove file if handled successfully
	os.Remove(fmt.Sprintf("./upload/%s", filename))

	r.Response.WriteJson(v1.UploadStarPlanRes{
		Data: "Success",
	})
	return
}

func handleUploadedFile(year int, filename string) error {
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
		admissionMethodId, _ := g.Model("admission_methods").Fields("id").Where("name", "繁星推薦").Value()
		student := g.Map{
			"degree_id":           degreeId,
			"admission_year":      year,
			"admission_method_id": admissionMethodId,
		}
		gsatScore := g.Map{}
		starPlan := g.Map{}
		for j, col := range row {
			col = strings.TrimSpace(col)
			if col == "" || col == "無成績" {
				col = "-1"
			}
			switch rows[0][j] {
			case "姓名":
				student["name"] = col
			case "性別":
				student["gender_id"], _ = g.Model("genders").Fields("id").Where("name", col).Value()
			case "學號":
				student["student_code"] = col
			case "身分別":
				student["identity_category_id"], _ = g.Model("identity_categories").Fields("id").Where("name", col).Value()
			case "推薦學校代碼":
				student["graduated_school_id"], _ = g.Model("schools").Fields("id").Where("school_code", col).Value()
			case "學測應試號碼":
				gsatScore["gsat_test_number"] = col
			case "學測國文級分":
				gsatScore["chinese"], _ = strconv.Atoi(col)
			case "學測英文級分":
				gsatScore["english"], _ = strconv.Atoi(col)
			case "學測數學級分":
				gsatScore["math"], _ = strconv.Atoi(col)
			case "學測自然級分":
				gsatScore["nature"], _ = strconv.Atoi(col)
			case "學測社會級分":
				gsatScore["society"], _ = strconv.Atoi(col)
			case "近三年內英聽最優等級":
				gsatScore["listening"] = col

			case "校系代碼":
				starPlan["school_dept_code"] = col
			case "推薦順位":
				starPlan["recommended_order"], _ = strconv.Atoi(col)
			case "錄取(通過篩選)志願序":
				starPlan["adms_order"], _ = strconv.Atoi(col)
			case "在校學業成績全校排名百分比":
				starPlan["school_ranking_percentage"], _ = strconv.Atoi(col)
			case "國文學業總平均成績全校排名百分比":
				starPlan["chinese_ranking_percentage"], _ = strconv.Atoi(col)
			case "英文學業總平均成績全校排名百分比":
				starPlan["english_ranking_percentage"], _ = strconv.Atoi(col)
			case "數學學業總平均成績全校排名百分比":
				starPlan["math_ranking_percentage"], _ = strconv.Atoi(col)
			case "物理學業總平均成績全校排名百分比":
				starPlan["physics_ranking_percentage"], _ = strconv.Atoi(col)
			case "化學學業總平均成績全校排名百分比":
				starPlan["chemistry_ranking_percentage"], _ = strconv.Atoi(col)
			case "生物學業總平均成績全校排名百分比":
				starPlan["biology_ranking_percentage"], _ = strconv.Atoi(col)
			case "地球科學學業總平均成績全校排名百分比":
				starPlan["earth_science_ranking_percentage"], _ = strconv.Atoi(col)
			case "歷史學業總平均成績全校排名百分比":
				starPlan["history_ranking_percentage"], _ = strconv.Atoi(col)
			case "地理學業總平均成績全校排名百分比":
				starPlan["geography_ranking_percentage"], _ = strconv.Atoi(col)
			case "公民與社會學業總平均成績全校排名百分比":
				starPlan["citizenship_ranking_percentage"], _ = strconv.Atoi(col)
			case "音樂學業總平均成績全校排名百分比":
				starPlan["music_ranking_percentage"], _ = strconv.Atoi(col)
			case "美術學業總平均成績全校排名百分比":
				starPlan["art_ranking_percentage"], _ = strconv.Atoi(col)
			case "舞蹈學業總平均成績全校排名百分比":
				starPlan["dance_ranking_percentage"], _ = strconv.Atoi(col)
			case "體育學業總平均成績全校排名百分比":
				starPlan["physical_ranking_percentage"], _ = strconv.Atoi(col)
			case "就讀科、學程、班別":
				starPlan["study_subject_course_class"] = col
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
		starPlan["stu_id"], _ = studentRes.LastInsertId()
		starPlan["gsat_score_id"], _ = gsatScoreRes.LastInsertId()
		_, err = g.Model("star_plan").Data(starPlan).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}
