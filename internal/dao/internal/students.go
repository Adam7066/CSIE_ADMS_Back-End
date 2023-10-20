// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StudentsDao is the data access object for table students.
type StudentsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns StudentsColumns // columns contains all the column names of Table for convenient usage.
}

// StudentsColumns defines and stores column names for table students.
type StudentsColumns struct {
	Id                 string //
	Name               string //
	GenderId           string //
	StudentCode        string //
	DegreeId           string //
	AdmissionMethodId  string //
	AdmissionGroupId   string //
	IdentityCategoryId string //
	GraduatedSchoolId  string //
	CreatedAt          string //
	UpdatedAt          string //
	DeletedAt          string //
}

// studentsColumns holds the columns for table students.
var studentsColumns = StudentsColumns{
	Id:                 "id",
	Name:               "name",
	GenderId:           "gender_id",
	StudentCode:        "student_code",
	DegreeId:           "degree_id",
	AdmissionMethodId:  "admission_method_id",
	AdmissionGroupId:   "admission_group_id",
	IdentityCategoryId: "identity_category_id",
	GraduatedSchoolId:  "graduated_school_id",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
}

// NewStudentsDao creates and returns a new DAO object for table data access.
func NewStudentsDao() *StudentsDao {
	return &StudentsDao{
		group:   "default",
		table:   "students",
		columns: studentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StudentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StudentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StudentsDao) Columns() StudentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StudentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StudentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StudentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
