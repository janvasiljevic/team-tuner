// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"jv/team-tone-tuner/model/course"
	"jv/team-tone-tuner/model/group"
	"jv/team-tone-tuner/model/grouprun"
	"jv/team-tone-tuner/model/predicate"
	"jv/team-tone-tuner/model/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GroupUpdate) SetUpdatedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// AddStudentIDs adds the "students" edge to the User entity by IDs.
func (gu *GroupUpdate) AddStudentIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.AddStudentIDs(ids...)
	return gu
}

// AddStudents adds the "students" edges to the User entity.
func (gu *GroupUpdate) AddStudents(u ...*User) *GroupUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddStudentIDs(ids...)
}

// SetCourseID sets the "course" edge to the Course entity by ID.
func (gu *GroupUpdate) SetCourseID(id uuid.UUID) *GroupUpdate {
	gu.mutation.SetCourseID(id)
	return gu
}

// SetCourse sets the "course" edge to the Course entity.
func (gu *GroupUpdate) SetCourse(c *Course) *GroupUpdate {
	return gu.SetCourseID(c.ID)
}

// SetGroupRunID sets the "group_run" edge to the GroupRun entity by ID.
func (gu *GroupUpdate) SetGroupRunID(id uuid.UUID) *GroupUpdate {
	gu.mutation.SetGroupRunID(id)
	return gu
}

// SetGroupRun sets the "group_run" edge to the GroupRun entity.
func (gu *GroupUpdate) SetGroupRun(g *GroupRun) *GroupUpdate {
	return gu.SetGroupRunID(g.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearStudents clears all "students" edges to the User entity.
func (gu *GroupUpdate) ClearStudents() *GroupUpdate {
	gu.mutation.ClearStudents()
	return gu
}

// RemoveStudentIDs removes the "students" edge to User entities by IDs.
func (gu *GroupUpdate) RemoveStudentIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.RemoveStudentIDs(ids...)
	return gu
}

// RemoveStudents removes "students" edges to User entities.
func (gu *GroupUpdate) RemoveStudents(u ...*User) *GroupUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveStudentIDs(ids...)
}

// ClearCourse clears the "course" edge to the Course entity.
func (gu *GroupUpdate) ClearCourse() *GroupUpdate {
	gu.mutation.ClearCourse()
	return gu
}

// ClearGroupRun clears the "group_run" edge to the GroupRun entity.
func (gu *GroupUpdate) ClearGroupRun() *GroupUpdate {
	gu.mutation.ClearGroupRun()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GroupUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := group.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if v, ok := gu.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Group.name": %w`, err)}
		}
	}
	if _, ok := gu.mutation.CourseID(); gu.mutation.CourseCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Group.course"`)
	}
	if _, ok := gu.mutation.GroupRunID(); gu.mutation.GroupRunCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Group.group_run"`)
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if gu.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.StudentsTable,
			Columns: group.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !gu.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.StudentsTable,
			Columns: group.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.StudentsTable,
			Columns: group.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.CourseTable,
			Columns: []string{group.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.CourseTable,
			Columns: []string{group.CourseColumn},
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
	if gu.mutation.GroupRunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.GroupRunTable,
			Columns: []string{group.GroupRunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GroupRunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.GroupRunTable,
			Columns: []string{group.GroupRunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GroupUpdateOne) SetUpdatedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// AddStudentIDs adds the "students" edge to the User entity by IDs.
func (guo *GroupUpdateOne) AddStudentIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.AddStudentIDs(ids...)
	return guo
}

// AddStudents adds the "students" edges to the User entity.
func (guo *GroupUpdateOne) AddStudents(u ...*User) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddStudentIDs(ids...)
}

// SetCourseID sets the "course" edge to the Course entity by ID.
func (guo *GroupUpdateOne) SetCourseID(id uuid.UUID) *GroupUpdateOne {
	guo.mutation.SetCourseID(id)
	return guo
}

// SetCourse sets the "course" edge to the Course entity.
func (guo *GroupUpdateOne) SetCourse(c *Course) *GroupUpdateOne {
	return guo.SetCourseID(c.ID)
}

// SetGroupRunID sets the "group_run" edge to the GroupRun entity by ID.
func (guo *GroupUpdateOne) SetGroupRunID(id uuid.UUID) *GroupUpdateOne {
	guo.mutation.SetGroupRunID(id)
	return guo
}

// SetGroupRun sets the "group_run" edge to the GroupRun entity.
func (guo *GroupUpdateOne) SetGroupRun(g *GroupRun) *GroupUpdateOne {
	return guo.SetGroupRunID(g.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearStudents clears all "students" edges to the User entity.
func (guo *GroupUpdateOne) ClearStudents() *GroupUpdateOne {
	guo.mutation.ClearStudents()
	return guo
}

// RemoveStudentIDs removes the "students" edge to User entities by IDs.
func (guo *GroupUpdateOne) RemoveStudentIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.RemoveStudentIDs(ids...)
	return guo
}

// RemoveStudents removes "students" edges to User entities.
func (guo *GroupUpdateOne) RemoveStudents(u ...*User) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveStudentIDs(ids...)
}

// ClearCourse clears the "course" edge to the Course entity.
func (guo *GroupUpdateOne) ClearCourse() *GroupUpdateOne {
	guo.mutation.ClearCourse()
	return guo
}

// ClearGroupRun clears the "group_run" edge to the GroupRun entity.
func (guo *GroupUpdateOne) ClearGroupRun() *GroupUpdateOne {
	guo.mutation.ClearGroupRun()
	return guo
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	guo.defaults()
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GroupUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := group.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if v, ok := guo.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Group.name": %w`, err)}
		}
	}
	if _, ok := guo.mutation.CourseID(); guo.mutation.CourseCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Group.course"`)
	}
	if _, ok := guo.mutation.GroupRunID(); guo.mutation.GroupRunCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Group.group_run"`)
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if guo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.StudentsTable,
			Columns: group.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !guo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.StudentsTable,
			Columns: group.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.StudentsTable,
			Columns: group.StudentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.CourseTable,
			Columns: []string{group.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.CourseTable,
			Columns: []string{group.CourseColumn},
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
	if guo.mutation.GroupRunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.GroupRunTable,
			Columns: []string{group.GroupRunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GroupRunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.GroupRunTable,
			Columns: []string{group.GroupRunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
