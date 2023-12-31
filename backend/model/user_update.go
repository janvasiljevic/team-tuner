// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"jv/team-tone-tuner/model/bfianswer"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/course"
	"jv/team-tone-tuner/model/group"
	"jv/team-tone-tuner/model/predicate"
	"jv/team-tone-tuner/model/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetGithubUsername sets the "github_username" field.
func (uu *UserUpdate) SetGithubUsername(s string) *UserUpdate {
	uu.mutation.SetGithubUsername(s)
	return uu
}

// SetUniveresityID sets the "univeresity_id" field.
func (uu *UserUpdate) SetUniveresityID(s string) *UserUpdate {
	uu.mutation.SetUniveresityID(s)
	return uu
}

// SetNillableUniveresityID sets the "univeresity_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUniveresityID(s *string) *UserUpdate {
	if s != nil {
		uu.SetUniveresityID(*s)
	}
	return uu
}

// ClearUniveresityID clears the value of the "univeresity_id" field.
func (uu *UserUpdate) ClearUniveresityID() *UserUpdate {
	uu.mutation.ClearUniveresityID()
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(u user.Role) *UserUpdate {
	uu.mutation.SetRole(u)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(u *user.Role) *UserUpdate {
	if u != nil {
		uu.SetRole(*u)
	}
	return uu
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (uu *UserUpdate) AddCourseIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddCourseIDs(ids...)
	return uu
}

// AddCourses adds the "courses" edges to the Course entity.
func (uu *UserUpdate) AddCourses(c ...*Course) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCourseIDs(ids...)
}

// SetBfiReportID sets the "bfi_report" edge to the BfiReport entity by ID.
func (uu *UserUpdate) SetBfiReportID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetBfiReportID(id)
	return uu
}

// SetNillableBfiReportID sets the "bfi_report" edge to the BfiReport entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableBfiReportID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetBfiReportID(*id)
	}
	return uu
}

// SetBfiReport sets the "bfi_report" edge to the BfiReport entity.
func (uu *UserUpdate) SetBfiReport(b *BfiReport) *UserUpdate {
	return uu.SetBfiReportID(b.ID)
}

// AddBfiAnswerIDs adds the "bfi_answers" edge to the BfiAnswer entity by IDs.
func (uu *UserUpdate) AddBfiAnswerIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddBfiAnswerIDs(ids...)
	return uu
}

// AddBfiAnswers adds the "bfi_answers" edges to the BfiAnswer entity.
func (uu *UserUpdate) AddBfiAnswers(b ...*BfiAnswer) *UserUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.AddBfiAnswerIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the Group entity.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (uu *UserUpdate) ClearCourses() *UserUpdate {
	uu.mutation.ClearCourses()
	return uu
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (uu *UserUpdate) RemoveCourseIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveCourseIDs(ids...)
	return uu
}

// RemoveCourses removes "courses" edges to Course entities.
func (uu *UserUpdate) RemoveCourses(c ...*Course) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCourseIDs(ids...)
}

// ClearBfiReport clears the "bfi_report" edge to the BfiReport entity.
func (uu *UserUpdate) ClearBfiReport() *UserUpdate {
	uu.mutation.ClearBfiReport()
	return uu
}

// ClearBfiAnswers clears all "bfi_answers" edges to the BfiAnswer entity.
func (uu *UserUpdate) ClearBfiAnswers() *UserUpdate {
	uu.mutation.ClearBfiAnswers()
	return uu
}

// RemoveBfiAnswerIDs removes the "bfi_answers" edge to BfiAnswer entities by IDs.
func (uu *UserUpdate) RemoveBfiAnswerIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveBfiAnswerIDs(ids...)
	return uu
}

// RemoveBfiAnswers removes "bfi_answers" edges to BfiAnswer entities.
func (uu *UserUpdate) RemoveBfiAnswers(b ...*BfiAnswer) *UserUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.RemoveBfiAnswerIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to Group entities.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`model: validator failed for field "User.role": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.GithubUsername(); ok {
		_spec.SetField(user.FieldGithubUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.UniveresityID(); ok {
		_spec.SetField(user.FieldUniveresityID, field.TypeString, value)
	}
	if uu.mutation.UniveresityIDCleared() {
		_spec.ClearField(user.FieldUniveresityID, field.TypeString)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if uu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.CoursesTable,
			Columns: user.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !uu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.CoursesTable,
			Columns: user.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.CoursesTable,
			Columns: user.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.BfiReportCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.BfiReportTable,
			Columns: []string{user.BfiReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.BfiReportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.BfiReportTable,
			Columns: []string{user.BfiReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.BfiAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BfiAnswersTable,
			Columns: []string{user.BfiAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfianswer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedBfiAnswersIDs(); len(nodes) > 0 && !uu.mutation.BfiAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BfiAnswersTable,
			Columns: []string{user.BfiAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfianswer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.BfiAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BfiAnswersTable,
			Columns: []string{user.BfiAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfianswer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetGithubUsername sets the "github_username" field.
func (uuo *UserUpdateOne) SetGithubUsername(s string) *UserUpdateOne {
	uuo.mutation.SetGithubUsername(s)
	return uuo
}

// SetUniveresityID sets the "univeresity_id" field.
func (uuo *UserUpdateOne) SetUniveresityID(s string) *UserUpdateOne {
	uuo.mutation.SetUniveresityID(s)
	return uuo
}

// SetNillableUniveresityID sets the "univeresity_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUniveresityID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUniveresityID(*s)
	}
	return uuo
}

// ClearUniveresityID clears the value of the "univeresity_id" field.
func (uuo *UserUpdateOne) ClearUniveresityID() *UserUpdateOne {
	uuo.mutation.ClearUniveresityID()
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(u user.Role) *UserUpdateOne {
	uuo.mutation.SetRole(u)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(u *user.Role) *UserUpdateOne {
	if u != nil {
		uuo.SetRole(*u)
	}
	return uuo
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (uuo *UserUpdateOne) AddCourseIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddCourseIDs(ids...)
	return uuo
}

// AddCourses adds the "courses" edges to the Course entity.
func (uuo *UserUpdateOne) AddCourses(c ...*Course) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCourseIDs(ids...)
}

// SetBfiReportID sets the "bfi_report" edge to the BfiReport entity by ID.
func (uuo *UserUpdateOne) SetBfiReportID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetBfiReportID(id)
	return uuo
}

// SetNillableBfiReportID sets the "bfi_report" edge to the BfiReport entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBfiReportID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetBfiReportID(*id)
	}
	return uuo
}

// SetBfiReport sets the "bfi_report" edge to the BfiReport entity.
func (uuo *UserUpdateOne) SetBfiReport(b *BfiReport) *UserUpdateOne {
	return uuo.SetBfiReportID(b.ID)
}

// AddBfiAnswerIDs adds the "bfi_answers" edge to the BfiAnswer entity by IDs.
func (uuo *UserUpdateOne) AddBfiAnswerIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddBfiAnswerIDs(ids...)
	return uuo
}

// AddBfiAnswers adds the "bfi_answers" edges to the BfiAnswer entity.
func (uuo *UserUpdateOne) AddBfiAnswers(b ...*BfiAnswer) *UserUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.AddBfiAnswerIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (uuo *UserUpdateOne) ClearCourses() *UserUpdateOne {
	uuo.mutation.ClearCourses()
	return uuo
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (uuo *UserUpdateOne) RemoveCourseIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveCourseIDs(ids...)
	return uuo
}

// RemoveCourses removes "courses" edges to Course entities.
func (uuo *UserUpdateOne) RemoveCourses(c ...*Course) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCourseIDs(ids...)
}

// ClearBfiReport clears the "bfi_report" edge to the BfiReport entity.
func (uuo *UserUpdateOne) ClearBfiReport() *UserUpdateOne {
	uuo.mutation.ClearBfiReport()
	return uuo
}

// ClearBfiAnswers clears all "bfi_answers" edges to the BfiAnswer entity.
func (uuo *UserUpdateOne) ClearBfiAnswers() *UserUpdateOne {
	uuo.mutation.ClearBfiAnswers()
	return uuo
}

// RemoveBfiAnswerIDs removes the "bfi_answers" edge to BfiAnswer entities by IDs.
func (uuo *UserUpdateOne) RemoveBfiAnswerIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveBfiAnswerIDs(ids...)
	return uuo
}

// RemoveBfiAnswers removes "bfi_answers" edges to BfiAnswer entities.
func (uuo *UserUpdateOne) RemoveBfiAnswers(b ...*BfiAnswer) *UserUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.RemoveBfiAnswerIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`model: validator failed for field "User.role": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.GithubUsername(); ok {
		_spec.SetField(user.FieldGithubUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UniveresityID(); ok {
		_spec.SetField(user.FieldUniveresityID, field.TypeString, value)
	}
	if uuo.mutation.UniveresityIDCleared() {
		_spec.ClearField(user.FieldUniveresityID, field.TypeString)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if uuo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.CoursesTable,
			Columns: user.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !uuo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.CoursesTable,
			Columns: user.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.CoursesTable,
			Columns: user.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.BfiReportCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.BfiReportTable,
			Columns: []string{user.BfiReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.BfiReportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   user.BfiReportTable,
			Columns: []string{user.BfiReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.BfiAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BfiAnswersTable,
			Columns: []string{user.BfiAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfianswer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedBfiAnswersIDs(); len(nodes) > 0 && !uuo.mutation.BfiAnswersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BfiAnswersTable,
			Columns: []string{user.BfiAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfianswer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.BfiAnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BfiAnswersTable,
			Columns: []string{user.BfiAnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bfianswer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
