// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Schools is the golang structure of table schools for DAO operations like Where/Data.
type Schools struct {
	g.Meta     `orm:"table:schools, do:true"`
	Id         interface{} //
	DegreeId   interface{} //
	Name       interface{} //
	SchoolCode interface{} //
	CityId     interface{} //
	AreaId     interface{} //
}
