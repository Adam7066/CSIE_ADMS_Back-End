// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GsatScore is the golang structure of table gsat_score for DAO operations like Where/Data.
type GsatScore struct {
	g.Meta         `orm:"table:gsat_score, do:true"`
	Id             interface{} //
	GsatTestNumber interface{} //
	Chinese        interface{} //
	English        interface{} //
	Math           interface{} //
	Nature         interface{} //
	Society        interface{} //
	Listening      interface{} //
	SumScore       interface{} //
}
