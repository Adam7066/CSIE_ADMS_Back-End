// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyApcs is the golang structure of table apply_apcs for DAO operations like Where/Data.
type ApplyApcs struct {
	g.Meta              `orm:"table:apply_apcs, do:true"`
	Id                  interface{} //
	StuId               interface{} //
	GsatScoreId         interface{} //
	SchoolDeptCode      interface{} //
	GsatCalScore        interface{} //
	ProjectScore1       interface{} //
	ProjectScore2       interface{} //
	ProjectTestScore    interface{} //
	SelectionTotalScore interface{} //
	AdmsStatus          interface{} //
	AdmsRank            interface{} //
	AdmsTotalRank       interface{} //
}
