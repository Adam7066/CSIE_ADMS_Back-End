package rank

import (
	"CSIE_ADMS_Back-End/api/rank/v1"
	"CSIE_ADMS_Back-End/internal/model/entity"
	"CSIE_ADMS_Back-End/utility"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"
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
	err = handleUploadSemesterRankFile(req.Year, req.Semester, filename)
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

func handleUploadSemesterRankFile(year, semester int, filename string) error {
	path := fmt.Sprintf("./upload/%s", filename)
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, sheetName := range f.GetSheetList() {
		grade, err := strconv.Atoi(sheetName[0:1])
		if err != nil {
			return err
		}
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return err
		}
		for i, row := range rows {
			if i == 0 {
				continue
			}
			// Get studentCode
			var studentCode string
			for j, col := range row {
				col := strings.TrimSpace(col)
				if rows[0][j] == "學號" {
					studentCode = col
					break
				}
			}

			// Check if student exists in db
			students := make([]entity.Students, 0)
			cnt := 0
			err := g.Model("students").Where("student_code", studentCode).ScanAndCount(&students, &cnt, true)
			if err != nil {
				return err
			}

			rank := g.Map{
				"year":     year,
				"semester": semester,
			}
			if cnt != 0 {
				rank["stu_id"] = students[0].Id
				// Check if semester rank exists in db
				semesterRanks := make([]entity.SemesterRank, 0)
				err := g.Model("semester_rank").
					Where("stu_id=? AND year=? AND semester=?", students[0].Id, year, semester).
					ScanAndCount(&semesterRanks, &cnt, true)
				if err != nil {
					return err
				}
				if cnt != 0 {
					continue
				}
			} else {
				degreeId, _ := g.Model("degrees").Fields("id").Where("name", "大學").Value()
				student := g.Map{
					"student_code":   studentCode,
					"degree_id":      degreeId,
					"admission_year": year - grade + 1,
				}
				for j, col := range row {
					col := strings.TrimSpace(col)
					switch rows[0][j] {
					case "姓名":
						student["name"] = col
					case "入學類別":
						if col == "考試分發" {
							col = "分發入學"
						} else if col == "個人申請" {
							col = "申請入學"
						}
						student["admission_method_id"], _ = g.Model("admission_methods").Fields("id").Where("name", col).Value()
					case "身分":
						student["identity_category_id"], _ = g.Model("identity_categories").Fields("id").Where("name", col).Value()
					}
				}
				studentRes, err := g.Model("students").Data(student).Insert()
				if err != nil {
					return err
				}
				rank["stu_id"], _ = studentRes.LastInsertId()
			}

			for j, col := range row {
				col := strings.TrimSpace(col)
				switch rows[0][j] {
				case "總積分":
					rank["total_points"], _ = strconv.ParseFloat(col, 64)
				case "總修學分數":
					rank["total_credits"], _ = strconv.Atoi(col)
				case "實得學分數":
					rank["earned_credits"], _ = strconv.Atoi(col)
				case "不及格學分數":
					rank["failed_credits"], _ = strconv.Atoi(col)
				case "平均成績":
					rank["average_score"], _ = strconv.ParseFloat(col, 64)
				case "系排名":
					rank["department_rank"], _ = strconv.Atoi(col)
				case "系排名百分比":
					rank["department_rank_percent"], _ = strconv.ParseFloat(col, 64)
				case "目前學籍狀態":
					rank["current_student_status"] = col
				}
			}

			_, err = g.Model("semester_rank").Data(rank).Insert()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
