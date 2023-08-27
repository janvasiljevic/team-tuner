// Code generated by ent, DO NOT EDIT.

package group

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeStudents holds the string denoting the students edge name in mutations.
	EdgeStudents = "students"
	// EdgeCourse holds the string denoting the course edge name in mutations.
	EdgeCourse = "course"
	// EdgeGroupRun holds the string denoting the group_run edge name in mutations.
	EdgeGroupRun = "group_run"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// StudentsTable is the table that holds the students relation/edge. The primary key declared below.
	StudentsTable = "group_students"
	// StudentsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	StudentsInverseTable = "users"
	// CourseTable is the table that holds the course relation/edge.
	CourseTable = "groups"
	// CourseInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CourseInverseTable = "courses"
	// CourseColumn is the table column denoting the course relation/edge.
	CourseColumn = "course_groups"
	// GroupRunTable is the table that holds the group_run relation/edge.
	GroupRunTable = "groups"
	// GroupRunInverseTable is the table name for the GroupRun entity.
	// It exists in this package in order to avoid circular dependency with the "grouprun" package.
	GroupRunInverseTable = "group_runs"
	// GroupRunColumn is the table column denoting the group_run relation/edge.
	GroupRunColumn = "group_run_groups"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "groups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"course_groups",
	"group_run_groups",
}

var (
	// StudentsPrimaryKey and StudentsColumn2 are the table columns denoting the
	// primary key for the students relation (M2M).
	StudentsPrimaryKey = []string{"group_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Group queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByStudentsCount orders the results by students count.
func ByStudentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStudentsStep(), opts...)
	}
}

// ByStudents orders the results by students terms.
func ByStudents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCourseField orders the results by course field.
func ByCourseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCourseStep(), sql.OrderByField(field, opts...))
	}
}

// ByGroupRunField orders the results by group_run field.
func ByGroupRunField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupRunStep(), sql.OrderByField(field, opts...))
	}
}
func newStudentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, StudentsTable, StudentsPrimaryKey...),
	)
}
func newCourseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CourseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
	)
}
func newGroupRunStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupRunInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GroupRunTable, GroupRunColumn),
	)
}
