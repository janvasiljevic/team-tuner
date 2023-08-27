// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/predicate"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/schema"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BfiReportUpdate is the builder for updating BfiReport entities.
type BfiReportUpdate struct {
	config
	hooks    []Hook
	mutation *BfiReportMutation
}

// Where appends a list predicates to the BfiReportUpdate builder.
func (bru *BfiReportUpdate) Where(ps ...predicate.BfiReport) *BfiReportUpdate {
	bru.mutation.Where(ps...)
	return bru
}

// SetUpdatedAt sets the "updated_at" field.
func (bru *BfiReportUpdate) SetUpdatedAt(t time.Time) *BfiReportUpdate {
	bru.mutation.SetUpdatedAt(t)
	return bru
}

// SetConscientiousness sets the "conscientiousness" field.
func (bru *BfiReportUpdate) SetConscientiousness(sri schema.BfiReportItem) *BfiReportUpdate {
	bru.mutation.SetConscientiousness(sri)
	return bru
}

// SetExtraversion sets the "extraversion" field.
func (bru *BfiReportUpdate) SetExtraversion(sri schema.BfiReportItem) *BfiReportUpdate {
	bru.mutation.SetExtraversion(sri)
	return bru
}

// SetAgreeableness sets the "agreeableness" field.
func (bru *BfiReportUpdate) SetAgreeableness(sri schema.BfiReportItem) *BfiReportUpdate {
	bru.mutation.SetAgreeableness(sri)
	return bru
}

// SetNeuroticism sets the "neuroticism" field.
func (bru *BfiReportUpdate) SetNeuroticism(sri schema.BfiReportItem) *BfiReportUpdate {
	bru.mutation.SetNeuroticism(sri)
	return bru
}

// SetOpenness sets the "openness" field.
func (bru *BfiReportUpdate) SetOpenness(sri schema.BfiReportItem) *BfiReportUpdate {
	bru.mutation.SetOpenness(sri)
	return bru
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (bru *BfiReportUpdate) SetStudentID(id uuid.UUID) *BfiReportUpdate {
	bru.mutation.SetStudentID(id)
	return bru
}

// SetNillableStudentID sets the "student" edge to the User entity by ID if the given value is not nil.
func (bru *BfiReportUpdate) SetNillableStudentID(id *uuid.UUID) *BfiReportUpdate {
	if id != nil {
		bru = bru.SetStudentID(*id)
	}
	return bru
}

// SetStudent sets the "student" edge to the User entity.
func (bru *BfiReportUpdate) SetStudent(u *User) *BfiReportUpdate {
	return bru.SetStudentID(u.ID)
}

// Mutation returns the BfiReportMutation object of the builder.
func (bru *BfiReportUpdate) Mutation() *BfiReportMutation {
	return bru.mutation
}

// ClearStudent clears the "student" edge to the User entity.
func (bru *BfiReportUpdate) ClearStudent() *BfiReportUpdate {
	bru.mutation.ClearStudent()
	return bru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bru *BfiReportUpdate) Save(ctx context.Context) (int, error) {
	bru.defaults()
	return withHooks(ctx, bru.sqlSave, bru.mutation, bru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bru *BfiReportUpdate) SaveX(ctx context.Context) int {
	affected, err := bru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bru *BfiReportUpdate) Exec(ctx context.Context) error {
	_, err := bru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bru *BfiReportUpdate) ExecX(ctx context.Context) {
	if err := bru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bru *BfiReportUpdate) defaults() {
	if _, ok := bru.mutation.UpdatedAt(); !ok {
		v := bfireport.UpdateDefaultUpdatedAt()
		bru.mutation.SetUpdatedAt(v)
	}
}

func (bru *BfiReportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bfireport.Table, bfireport.Columns, sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID))
	if ps := bru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bru.mutation.UpdatedAt(); ok {
		_spec.SetField(bfireport.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bru.mutation.Conscientiousness(); ok {
		_spec.SetField(bfireport.FieldConscientiousness, field.TypeJSON, value)
	}
	if value, ok := bru.mutation.Extraversion(); ok {
		_spec.SetField(bfireport.FieldExtraversion, field.TypeJSON, value)
	}
	if value, ok := bru.mutation.Agreeableness(); ok {
		_spec.SetField(bfireport.FieldAgreeableness, field.TypeJSON, value)
	}
	if value, ok := bru.mutation.Neuroticism(); ok {
		_spec.SetField(bfireport.FieldNeuroticism, field.TypeJSON, value)
	}
	if value, ok := bru.mutation.Openness(); ok {
		_spec.SetField(bfireport.FieldOpenness, field.TypeJSON, value)
	}
	if bru.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   bfireport.StudentTable,
			Columns: []string{bfireport.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bru.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   bfireport.StudentTable,
			Columns: []string{bfireport.StudentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bfireport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bru.mutation.done = true
	return n, nil
}

// BfiReportUpdateOne is the builder for updating a single BfiReport entity.
type BfiReportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BfiReportMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (bruo *BfiReportUpdateOne) SetUpdatedAt(t time.Time) *BfiReportUpdateOne {
	bruo.mutation.SetUpdatedAt(t)
	return bruo
}

// SetConscientiousness sets the "conscientiousness" field.
func (bruo *BfiReportUpdateOne) SetConscientiousness(sri schema.BfiReportItem) *BfiReportUpdateOne {
	bruo.mutation.SetConscientiousness(sri)
	return bruo
}

// SetExtraversion sets the "extraversion" field.
func (bruo *BfiReportUpdateOne) SetExtraversion(sri schema.BfiReportItem) *BfiReportUpdateOne {
	bruo.mutation.SetExtraversion(sri)
	return bruo
}

// SetAgreeableness sets the "agreeableness" field.
func (bruo *BfiReportUpdateOne) SetAgreeableness(sri schema.BfiReportItem) *BfiReportUpdateOne {
	bruo.mutation.SetAgreeableness(sri)
	return bruo
}

// SetNeuroticism sets the "neuroticism" field.
func (bruo *BfiReportUpdateOne) SetNeuroticism(sri schema.BfiReportItem) *BfiReportUpdateOne {
	bruo.mutation.SetNeuroticism(sri)
	return bruo
}

// SetOpenness sets the "openness" field.
func (bruo *BfiReportUpdateOne) SetOpenness(sri schema.BfiReportItem) *BfiReportUpdateOne {
	bruo.mutation.SetOpenness(sri)
	return bruo
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (bruo *BfiReportUpdateOne) SetStudentID(id uuid.UUID) *BfiReportUpdateOne {
	bruo.mutation.SetStudentID(id)
	return bruo
}

// SetNillableStudentID sets the "student" edge to the User entity by ID if the given value is not nil.
func (bruo *BfiReportUpdateOne) SetNillableStudentID(id *uuid.UUID) *BfiReportUpdateOne {
	if id != nil {
		bruo = bruo.SetStudentID(*id)
	}
	return bruo
}

// SetStudent sets the "student" edge to the User entity.
func (bruo *BfiReportUpdateOne) SetStudent(u *User) *BfiReportUpdateOne {
	return bruo.SetStudentID(u.ID)
}

// Mutation returns the BfiReportMutation object of the builder.
func (bruo *BfiReportUpdateOne) Mutation() *BfiReportMutation {
	return bruo.mutation
}

// ClearStudent clears the "student" edge to the User entity.
func (bruo *BfiReportUpdateOne) ClearStudent() *BfiReportUpdateOne {
	bruo.mutation.ClearStudent()
	return bruo
}

// Where appends a list predicates to the BfiReportUpdate builder.
func (bruo *BfiReportUpdateOne) Where(ps ...predicate.BfiReport) *BfiReportUpdateOne {
	bruo.mutation.Where(ps...)
	return bruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bruo *BfiReportUpdateOne) Select(field string, fields ...string) *BfiReportUpdateOne {
	bruo.fields = append([]string{field}, fields...)
	return bruo
}

// Save executes the query and returns the updated BfiReport entity.
func (bruo *BfiReportUpdateOne) Save(ctx context.Context) (*BfiReport, error) {
	bruo.defaults()
	return withHooks(ctx, bruo.sqlSave, bruo.mutation, bruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bruo *BfiReportUpdateOne) SaveX(ctx context.Context) *BfiReport {
	node, err := bruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bruo *BfiReportUpdateOne) Exec(ctx context.Context) error {
	_, err := bruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bruo *BfiReportUpdateOne) ExecX(ctx context.Context) {
	if err := bruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bruo *BfiReportUpdateOne) defaults() {
	if _, ok := bruo.mutation.UpdatedAt(); !ok {
		v := bfireport.UpdateDefaultUpdatedAt()
		bruo.mutation.SetUpdatedAt(v)
	}
}

func (bruo *BfiReportUpdateOne) sqlSave(ctx context.Context) (_node *BfiReport, err error) {
	_spec := sqlgraph.NewUpdateSpec(bfireport.Table, bfireport.Columns, sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID))
	id, ok := bruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "BfiReport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bfireport.FieldID)
		for _, f := range fields {
			if !bfireport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != bfireport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bruo.mutation.UpdatedAt(); ok {
		_spec.SetField(bfireport.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bruo.mutation.Conscientiousness(); ok {
		_spec.SetField(bfireport.FieldConscientiousness, field.TypeJSON, value)
	}
	if value, ok := bruo.mutation.Extraversion(); ok {
		_spec.SetField(bfireport.FieldExtraversion, field.TypeJSON, value)
	}
	if value, ok := bruo.mutation.Agreeableness(); ok {
		_spec.SetField(bfireport.FieldAgreeableness, field.TypeJSON, value)
	}
	if value, ok := bruo.mutation.Neuroticism(); ok {
		_spec.SetField(bfireport.FieldNeuroticism, field.TypeJSON, value)
	}
	if value, ok := bruo.mutation.Openness(); ok {
		_spec.SetField(bfireport.FieldOpenness, field.TypeJSON, value)
	}
	if bruo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   bfireport.StudentTable,
			Columns: []string{bfireport.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bruo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   bfireport.StudentTable,
			Columns: []string{bfireport.StudentColumn},
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
	_node = &BfiReport{config: bruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bfireport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bruo.mutation.done = true
	return _node, nil
}