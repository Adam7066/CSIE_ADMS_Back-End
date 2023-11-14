// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyGeneral is the golang structure of table apply_general for DAO operations like Where/Data.
type ApplyGeneral struct {
	g.Meta              `orm:"table:apply_general, do:true"`
	Id                  interface{} //
	StuId               interface{} //
	GsatScoreId         interface{} //
	SchoolDeptCode      interface{} //
	GsatCalScore        interface{} //
	ProjectScore1       interface{} //
	ProjectTestScore    interface{} //
	SelectionTotalScore interface{} //
	AdmsStatus          interface{} //
	AdmsRank            interface{} //
	AdmsTotalRank       interface{} //
}
