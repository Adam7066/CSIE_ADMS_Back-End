// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ApplyGeneral is the golang structure for table apply_general.
type ApplyGeneral struct {
	Id                  int     `json:"id"                  ` //
	StuId               int     `json:"stuId"               ` //
	GsatScoreId         int     `json:"gsatScoreId"         ` //
	SchoolDeptCode      string  `json:"schoolDeptCode"      ` //
	GsatCalScore        float64 `json:"gsatCalScore"        ` //
	ProjectScore1       float64 `json:"projectScore1"       ` //
	ProjectTestScore    float64 `json:"projectTestScore"    ` //
	SelectionTotalScore float64 `json:"selectionTotalScore" ` //
	AdmsStatus          string  `json:"admsStatus"          ` //
	AdmsRank            int     `json:"admsRank"            ` //
	AdmsTotalRank       int     `json:"admsTotalRank"       ` //
}