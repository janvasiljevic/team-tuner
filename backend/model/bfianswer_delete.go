// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"jv/team-tone-tuner/model/bfianswer"
	"jv/team-tone-tuner/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BfiAnswerDelete is the builder for deleting a BfiAnswer entity.
type BfiAnswerDelete struct {
	config
	hooks    []Hook
	mutation *BfiAnswerMutation
}

// Where appends a list predicates to the BfiAnswerDelete builder.
func (bad *BfiAnswerDelete) Where(ps ...predicate.BfiAnswer) *BfiAnswerDelete {
	bad.mutation.Where(ps...)
	return bad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bad *BfiAnswerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bad.sqlExec, bad.mutation, bad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bad *BfiAnswerDelete) ExecX(ctx context.Context) int {
	n, err := bad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bad *BfiAnswerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bfianswer.Table, sqlgraph.NewFieldSpec(bfianswer.FieldID, field.TypeUUID))
	if ps := bad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bad.mutation.done = true
	return affected, err
}

// BfiAnswerDeleteOne is the builder for deleting a single BfiAnswer entity.
type BfiAnswerDeleteOne struct {
	bad *BfiAnswerDelete
}

// Where appends a list predicates to the BfiAnswerDelete builder.
func (bado *BfiAnswerDeleteOne) Where(ps ...predicate.BfiAnswer) *BfiAnswerDeleteOne {
	bado.bad.mutation.Where(ps...)
	return bado
}

// Exec executes the deletion query.
func (bado *BfiAnswerDeleteOne) Exec(ctx context.Context) error {
	n, err := bado.bad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bfianswer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bado *BfiAnswerDeleteOne) ExecX(ctx context.Context) {
	if err := bado.Exec(ctx); err != nil {
		panic(err)
	}
}
