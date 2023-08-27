// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/schema"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BfiReportCreate is the builder for creating a BfiReport entity.
type BfiReportCreate struct {
	config
	mutation *BfiReportMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (brc *BfiReportCreate) SetCreatedAt(t time.Time) *BfiReportCreate {
	brc.mutation.SetCreatedAt(t)
	return brc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (brc *BfiReportCreate) SetNillableCreatedAt(t *time.Time) *BfiReportCreate {
	if t != nil {
		brc.SetCreatedAt(*t)
	}
	return brc
}

// SetUpdatedAt sets the "updated_at" field.
func (brc *BfiReportCreate) SetUpdatedAt(t time.Time) *BfiReportCreate {
	brc.mutation.SetUpdatedAt(t)
	return brc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (brc *BfiReportCreate) SetNillableUpdatedAt(t *time.Time) *BfiReportCreate {
	if t != nil {
		brc.SetUpdatedAt(*t)
	}
	return brc
}

// SetConscientiousness sets the "conscientiousness" field.
func (brc *BfiReportCreate) SetConscientiousness(sri schema.BfiReportItem) *BfiReportCreate {
	brc.mutation.SetConscientiousness(sri)
	return brc
}

// SetExtraversion sets the "extraversion" field.
func (brc *BfiReportCreate) SetExtraversion(sri schema.BfiReportItem) *BfiReportCreate {
	brc.mutation.SetExtraversion(sri)
	return brc
}

// SetAgreeableness sets the "agreeableness" field.
func (brc *BfiReportCreate) SetAgreeableness(sri schema.BfiReportItem) *BfiReportCreate {
	brc.mutation.SetAgreeableness(sri)
	return brc
}

// SetNeuroticism sets the "neuroticism" field.
func (brc *BfiReportCreate) SetNeuroticism(sri schema.BfiReportItem) *BfiReportCreate {
	brc.mutation.SetNeuroticism(sri)
	return brc
}

// SetOpenness sets the "openness" field.
func (brc *BfiReportCreate) SetOpenness(sri schema.BfiReportItem) *BfiReportCreate {
	brc.mutation.SetOpenness(sri)
	return brc
}

// SetID sets the "id" field.
func (brc *BfiReportCreate) SetID(u uuid.UUID) *BfiReportCreate {
	brc.mutation.SetID(u)
	return brc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (brc *BfiReportCreate) SetNillableID(u *uuid.UUID) *BfiReportCreate {
	if u != nil {
		brc.SetID(*u)
	}
	return brc
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (brc *BfiReportCreate) SetStudentID(id uuid.UUID) *BfiReportCreate {
	brc.mutation.SetStudentID(id)
	return brc
}

// SetNillableStudentID sets the "student" edge to the User entity by ID if the given value is not nil.
func (brc *BfiReportCreate) SetNillableStudentID(id *uuid.UUID) *BfiReportCreate {
	if id != nil {
		brc = brc.SetStudentID(*id)
	}
	return brc
}

// SetStudent sets the "student" edge to the User entity.
func (brc *BfiReportCreate) SetStudent(u *User) *BfiReportCreate {
	return brc.SetStudentID(u.ID)
}

// Mutation returns the BfiReportMutation object of the builder.
func (brc *BfiReportCreate) Mutation() *BfiReportMutation {
	return brc.mutation
}

// Save creates the BfiReport in the database.
func (brc *BfiReportCreate) Save(ctx context.Context) (*BfiReport, error) {
	brc.defaults()
	return withHooks(ctx, brc.sqlSave, brc.mutation, brc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (brc *BfiReportCreate) SaveX(ctx context.Context) *BfiReport {
	v, err := brc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (brc *BfiReportCreate) Exec(ctx context.Context) error {
	_, err := brc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (brc *BfiReportCreate) ExecX(ctx context.Context) {
	if err := brc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (brc *BfiReportCreate) defaults() {
	if _, ok := brc.mutation.CreatedAt(); !ok {
		v := bfireport.DefaultCreatedAt()
		brc.mutation.SetCreatedAt(v)
	}
	if _, ok := brc.mutation.UpdatedAt(); !ok {
		v := bfireport.DefaultUpdatedAt()
		brc.mutation.SetUpdatedAt(v)
	}
	if _, ok := brc.mutation.ID(); !ok {
		v := bfireport.DefaultID()
		brc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (brc *BfiReportCreate) check() error {
	if _, ok := brc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`model: missing required field "BfiReport.created_at"`)}
	}
	if _, ok := brc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`model: missing required field "BfiReport.updated_at"`)}
	}
	if _, ok := brc.mutation.Conscientiousness(); !ok {
		return &ValidationError{Name: "conscientiousness", err: errors.New(`model: missing required field "BfiReport.conscientiousness"`)}
	}
	if _, ok := brc.mutation.Extraversion(); !ok {
		return &ValidationError{Name: "extraversion", err: errors.New(`model: missing required field "BfiReport.extraversion"`)}
	}
	if _, ok := brc.mutation.Agreeableness(); !ok {
		return &ValidationError{Name: "agreeableness", err: errors.New(`model: missing required field "BfiReport.agreeableness"`)}
	}
	if _, ok := brc.mutation.Neuroticism(); !ok {
		return &ValidationError{Name: "neuroticism", err: errors.New(`model: missing required field "BfiReport.neuroticism"`)}
	}
	if _, ok := brc.mutation.Openness(); !ok {
		return &ValidationError{Name: "openness", err: errors.New(`model: missing required field "BfiReport.openness"`)}
	}
	return nil
}

func (brc *BfiReportCreate) sqlSave(ctx context.Context) (*BfiReport, error) {
	if err := brc.check(); err != nil {
		return nil, err
	}
	_node, _spec := brc.createSpec()
	if err := sqlgraph.CreateNode(ctx, brc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	brc.mutation.id = &_node.ID
	brc.mutation.done = true
	return _node, nil
}

func (brc *BfiReportCreate) createSpec() (*BfiReport, *sqlgraph.CreateSpec) {
	var (
		_node = &BfiReport{config: brc.config}
		_spec = sqlgraph.NewCreateSpec(bfireport.Table, sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = brc.conflict
	if id, ok := brc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := brc.mutation.CreatedAt(); ok {
		_spec.SetField(bfireport.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := brc.mutation.UpdatedAt(); ok {
		_spec.SetField(bfireport.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := brc.mutation.Conscientiousness(); ok {
		_spec.SetField(bfireport.FieldConscientiousness, field.TypeJSON, value)
		_node.Conscientiousness = value
	}
	if value, ok := brc.mutation.Extraversion(); ok {
		_spec.SetField(bfireport.FieldExtraversion, field.TypeJSON, value)
		_node.Extraversion = value
	}
	if value, ok := brc.mutation.Agreeableness(); ok {
		_spec.SetField(bfireport.FieldAgreeableness, field.TypeJSON, value)
		_node.Agreeableness = value
	}
	if value, ok := brc.mutation.Neuroticism(); ok {
		_spec.SetField(bfireport.FieldNeuroticism, field.TypeJSON, value)
		_node.Neuroticism = value
	}
	if value, ok := brc.mutation.Openness(); ok {
		_spec.SetField(bfireport.FieldOpenness, field.TypeJSON, value)
		_node.Openness = value
	}
	if nodes := brc.mutation.StudentIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BfiReport.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BfiReportUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (brc *BfiReportCreate) OnConflict(opts ...sql.ConflictOption) *BfiReportUpsertOne {
	brc.conflict = opts
	return &BfiReportUpsertOne{
		create: brc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BfiReport.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (brc *BfiReportCreate) OnConflictColumns(columns ...string) *BfiReportUpsertOne {
	brc.conflict = append(brc.conflict, sql.ConflictColumns(columns...))
	return &BfiReportUpsertOne{
		create: brc,
	}
}

type (
	// BfiReportUpsertOne is the builder for "upsert"-ing
	//  one BfiReport node.
	BfiReportUpsertOne struct {
		create *BfiReportCreate
	}

	// BfiReportUpsert is the "OnConflict" setter.
	BfiReportUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *BfiReportUpsert) SetUpdatedAt(v time.Time) *BfiReportUpsert {
	u.Set(bfireport.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BfiReportUpsert) UpdateUpdatedAt() *BfiReportUpsert {
	u.SetExcluded(bfireport.FieldUpdatedAt)
	return u
}

// SetConscientiousness sets the "conscientiousness" field.
func (u *BfiReportUpsert) SetConscientiousness(v schema.BfiReportItem) *BfiReportUpsert {
	u.Set(bfireport.FieldConscientiousness, v)
	return u
}

// UpdateConscientiousness sets the "conscientiousness" field to the value that was provided on create.
func (u *BfiReportUpsert) UpdateConscientiousness() *BfiReportUpsert {
	u.SetExcluded(bfireport.FieldConscientiousness)
	return u
}

// SetExtraversion sets the "extraversion" field.
func (u *BfiReportUpsert) SetExtraversion(v schema.BfiReportItem) *BfiReportUpsert {
	u.Set(bfireport.FieldExtraversion, v)
	return u
}

// UpdateExtraversion sets the "extraversion" field to the value that was provided on create.
func (u *BfiReportUpsert) UpdateExtraversion() *BfiReportUpsert {
	u.SetExcluded(bfireport.FieldExtraversion)
	return u
}

// SetAgreeableness sets the "agreeableness" field.
func (u *BfiReportUpsert) SetAgreeableness(v schema.BfiReportItem) *BfiReportUpsert {
	u.Set(bfireport.FieldAgreeableness, v)
	return u
}

// UpdateAgreeableness sets the "agreeableness" field to the value that was provided on create.
func (u *BfiReportUpsert) UpdateAgreeableness() *BfiReportUpsert {
	u.SetExcluded(bfireport.FieldAgreeableness)
	return u
}

// SetNeuroticism sets the "neuroticism" field.
func (u *BfiReportUpsert) SetNeuroticism(v schema.BfiReportItem) *BfiReportUpsert {
	u.Set(bfireport.FieldNeuroticism, v)
	return u
}

// UpdateNeuroticism sets the "neuroticism" field to the value that was provided on create.
func (u *BfiReportUpsert) UpdateNeuroticism() *BfiReportUpsert {
	u.SetExcluded(bfireport.FieldNeuroticism)
	return u
}

// SetOpenness sets the "openness" field.
func (u *BfiReportUpsert) SetOpenness(v schema.BfiReportItem) *BfiReportUpsert {
	u.Set(bfireport.FieldOpenness, v)
	return u
}

// UpdateOpenness sets the "openness" field to the value that was provided on create.
func (u *BfiReportUpsert) UpdateOpenness() *BfiReportUpsert {
	u.SetExcluded(bfireport.FieldOpenness)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BfiReport.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(bfireport.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BfiReportUpsertOne) UpdateNewValues() *BfiReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(bfireport.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(bfireport.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BfiReport.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BfiReportUpsertOne) Ignore() *BfiReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BfiReportUpsertOne) DoNothing() *BfiReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BfiReportCreate.OnConflict
// documentation for more info.
func (u *BfiReportUpsertOne) Update(set func(*BfiReportUpsert)) *BfiReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BfiReportUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BfiReportUpsertOne) SetUpdatedAt(v time.Time) *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BfiReportUpsertOne) UpdateUpdatedAt() *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetConscientiousness sets the "conscientiousness" field.
func (u *BfiReportUpsertOne) SetConscientiousness(v schema.BfiReportItem) *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetConscientiousness(v)
	})
}

// UpdateConscientiousness sets the "conscientiousness" field to the value that was provided on create.
func (u *BfiReportUpsertOne) UpdateConscientiousness() *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateConscientiousness()
	})
}

// SetExtraversion sets the "extraversion" field.
func (u *BfiReportUpsertOne) SetExtraversion(v schema.BfiReportItem) *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetExtraversion(v)
	})
}

// UpdateExtraversion sets the "extraversion" field to the value that was provided on create.
func (u *BfiReportUpsertOne) UpdateExtraversion() *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateExtraversion()
	})
}

// SetAgreeableness sets the "agreeableness" field.
func (u *BfiReportUpsertOne) SetAgreeableness(v schema.BfiReportItem) *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetAgreeableness(v)
	})
}

// UpdateAgreeableness sets the "agreeableness" field to the value that was provided on create.
func (u *BfiReportUpsertOne) UpdateAgreeableness() *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateAgreeableness()
	})
}

// SetNeuroticism sets the "neuroticism" field.
func (u *BfiReportUpsertOne) SetNeuroticism(v schema.BfiReportItem) *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetNeuroticism(v)
	})
}

// UpdateNeuroticism sets the "neuroticism" field to the value that was provided on create.
func (u *BfiReportUpsertOne) UpdateNeuroticism() *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateNeuroticism()
	})
}

// SetOpenness sets the "openness" field.
func (u *BfiReportUpsertOne) SetOpenness(v schema.BfiReportItem) *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetOpenness(v)
	})
}

// UpdateOpenness sets the "openness" field to the value that was provided on create.
func (u *BfiReportUpsertOne) UpdateOpenness() *BfiReportUpsertOne {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateOpenness()
	})
}

// Exec executes the query.
func (u *BfiReportUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for BfiReportCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BfiReportUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BfiReportUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: BfiReportUpsertOne.ID is not supported by MySQL driver. Use BfiReportUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BfiReportUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BfiReportCreateBulk is the builder for creating many BfiReport entities in bulk.
type BfiReportCreateBulk struct {
	config
	builders []*BfiReportCreate
	conflict []sql.ConflictOption
}

// Save creates the BfiReport entities in the database.
func (brcb *BfiReportCreateBulk) Save(ctx context.Context) ([]*BfiReport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(brcb.builders))
	nodes := make([]*BfiReport, len(brcb.builders))
	mutators := make([]Mutator, len(brcb.builders))
	for i := range brcb.builders {
		func(i int, root context.Context) {
			builder := brcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BfiReportMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, brcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = brcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, brcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, brcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (brcb *BfiReportCreateBulk) SaveX(ctx context.Context) []*BfiReport {
	v, err := brcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (brcb *BfiReportCreateBulk) Exec(ctx context.Context) error {
	_, err := brcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (brcb *BfiReportCreateBulk) ExecX(ctx context.Context) {
	if err := brcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BfiReport.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BfiReportUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (brcb *BfiReportCreateBulk) OnConflict(opts ...sql.ConflictOption) *BfiReportUpsertBulk {
	brcb.conflict = opts
	return &BfiReportUpsertBulk{
		create: brcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BfiReport.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (brcb *BfiReportCreateBulk) OnConflictColumns(columns ...string) *BfiReportUpsertBulk {
	brcb.conflict = append(brcb.conflict, sql.ConflictColumns(columns...))
	return &BfiReportUpsertBulk{
		create: brcb,
	}
}

// BfiReportUpsertBulk is the builder for "upsert"-ing
// a bulk of BfiReport nodes.
type BfiReportUpsertBulk struct {
	create *BfiReportCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BfiReport.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(bfireport.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BfiReportUpsertBulk) UpdateNewValues() *BfiReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(bfireport.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(bfireport.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BfiReport.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BfiReportUpsertBulk) Ignore() *BfiReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BfiReportUpsertBulk) DoNothing() *BfiReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BfiReportCreateBulk.OnConflict
// documentation for more info.
func (u *BfiReportUpsertBulk) Update(set func(*BfiReportUpsert)) *BfiReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BfiReportUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BfiReportUpsertBulk) SetUpdatedAt(v time.Time) *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BfiReportUpsertBulk) UpdateUpdatedAt() *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetConscientiousness sets the "conscientiousness" field.
func (u *BfiReportUpsertBulk) SetConscientiousness(v schema.BfiReportItem) *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetConscientiousness(v)
	})
}

// UpdateConscientiousness sets the "conscientiousness" field to the value that was provided on create.
func (u *BfiReportUpsertBulk) UpdateConscientiousness() *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateConscientiousness()
	})
}

// SetExtraversion sets the "extraversion" field.
func (u *BfiReportUpsertBulk) SetExtraversion(v schema.BfiReportItem) *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetExtraversion(v)
	})
}

// UpdateExtraversion sets the "extraversion" field to the value that was provided on create.
func (u *BfiReportUpsertBulk) UpdateExtraversion() *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateExtraversion()
	})
}

// SetAgreeableness sets the "agreeableness" field.
func (u *BfiReportUpsertBulk) SetAgreeableness(v schema.BfiReportItem) *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetAgreeableness(v)
	})
}

// UpdateAgreeableness sets the "agreeableness" field to the value that was provided on create.
func (u *BfiReportUpsertBulk) UpdateAgreeableness() *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateAgreeableness()
	})
}

// SetNeuroticism sets the "neuroticism" field.
func (u *BfiReportUpsertBulk) SetNeuroticism(v schema.BfiReportItem) *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetNeuroticism(v)
	})
}

// UpdateNeuroticism sets the "neuroticism" field to the value that was provided on create.
func (u *BfiReportUpsertBulk) UpdateNeuroticism() *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateNeuroticism()
	})
}

// SetOpenness sets the "openness" field.
func (u *BfiReportUpsertBulk) SetOpenness(v schema.BfiReportItem) *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.SetOpenness(v)
	})
}

// UpdateOpenness sets the "openness" field to the value that was provided on create.
func (u *BfiReportUpsertBulk) UpdateOpenness() *BfiReportUpsertBulk {
	return u.Update(func(s *BfiReportUpsert) {
		s.UpdateOpenness()
	})
}

// Exec executes the query.
func (u *BfiReportUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the BfiReportCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for BfiReportCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BfiReportUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
