package cmd

import (
	"CSIE_ADMS_Back-End/internal/consts"
	"CSIE_ADMS_Back-End/utility"
	"github.com/gogf/gf/v2/frame/g"
)

func InitTablesData() {
	initUsers()
	initGenders()
	initDegrees()
	initAdmissionMethods()
	initAdmissionGroups()
	initIdentityCategories()
	initCities()
	initAreas()
	initSchools()
}

func initUsers() {
	m := g.DB().Model("users")
	encodedHash, _ := utility.HashPWD("1234567890")
	_, err := m.Insert(g.Map{
		"email":    "admin@csie.adms",
		"password": encodedHash,
		"username": "admin",
		"role":     "admin",
	})
	utility.IfErrExit(err)
}

func initGenders() {
	_, err := g.Model("genders").Data(g.List{
		{"name": "男"},
		{"name": "女"},
	}).Insert()
	utility.IfErrExit(err)
}

func initDegrees() {
	_, err := g.Model("degrees").Data(g.List{
		{"name": "高中"},
		{"name": "大學"},
		{"name": "碩士"},
	}).Insert()
	utility.IfErrExit(err)
}

func initAdmissionMethods() {
	_, err := g.Model("admission_methods").Data(g.List{
		{"name": "特殊選才"},
		{"name": "繁星推薦"},
		{"name": "申請入學"},
		{"name": "分發入學"},
	}).Insert()
	utility.IfErrExit(err)
}

func initAdmissionGroups() {
	_, err := g.Model("admission_groups").Data(g.List{
		{"name": "一般組"},
		{"name": "APCS組"},
		{"name": "資安組"},
		{"name": "公費生"},
	}).Insert()
	utility.IfErrExit(err)
}

func initIdentityCategories() {
	_, err := g.Model("identity_categories").Data(g.List{
		{"name": "一般生"},
	}).Insert()
	utility.IfErrExit(err)
}

func initSchools() {
	m := g.Model("degrees")
	degreeId, err := m.Fields("id").Where("name", "高中").Value()
	utility.IfErrExit(err)
	highSchoolList := g.List{}
	for _, school := range consts.HighSchools {
		m = g.Model("cities")
		cityId, err := m.Fields("id").Where("name", school.City).Value()
		utility.IfErrExit(err)
		m = g.Model("areas")
		areaId, err := m.Fields("id").Where("name", school.Area).Value()
		utility.IfErrExit(err)
		highSchoolList = append(highSchoolList, g.Map{
			"degree_id":   degreeId,
			"name":        school.Name,
			"school_code": school.Code,
			"city_id":     cityId,
			"area_id":     areaId,
		})
	}
	_, err = g.Model("schools").Data(highSchoolList).Insert()
	utility.IfErrExit(err)
}

func initCities() {
	cityList := g.List{}
	for _, city := range consts.Cities {
		cityList = append(cityList, g.Map{
			"name": city,
		})
	}
	_, err := g.Model("cities").Data(cityList).Insert()
	utility.IfErrExit(err)
}

func initAreas() {
	_, err := g.Model("areas").Data(g.List{
		{"name": "北部"},
		{"name": "中部"},
		{"name": "南部"},
		{"name": "東部"},
		{"name": "離島"},
		{"name": "其他"},
	}).Insert()
	utility.IfErrExit(err)
}
