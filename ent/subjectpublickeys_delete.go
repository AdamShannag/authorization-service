// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/predicate"
	"authorization-service/ent/subjectpublickeys"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubjectPublicKeysDelete is the builder for deleting a SubjectPublicKeys entity.
type SubjectPublicKeysDelete struct {
	config
	hooks    []Hook
	mutation *SubjectPublicKeysMutation
}

// Where appends a list predicates to the SubjectPublicKeysDelete builder.
func (spkd *SubjectPublicKeysDelete) Where(ps ...predicate.SubjectPublicKeys) *SubjectPublicKeysDelete {
	spkd.mutation.Where(ps...)
	return spkd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spkd *SubjectPublicKeysDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, spkd.sqlExec, spkd.mutation, spkd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (spkd *SubjectPublicKeysDelete) ExecX(ctx context.Context) int {
	n, err := spkd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spkd *SubjectPublicKeysDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(subjectpublickeys.Table, sqlgraph.NewFieldSpec(subjectpublickeys.FieldID, field.TypeString))
	if ps := spkd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spkd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	spkd.mutation.done = true
	return affected, err
}

// SubjectPublicKeysDeleteOne is the builder for deleting a single SubjectPublicKeys entity.
type SubjectPublicKeysDeleteOne struct {
	spkd *SubjectPublicKeysDelete
}

// Where appends a list predicates to the SubjectPublicKeysDelete builder.
func (spkdo *SubjectPublicKeysDeleteOne) Where(ps ...predicate.SubjectPublicKeys) *SubjectPublicKeysDeleteOne {
	spkdo.spkd.mutation.Where(ps...)
	return spkdo
}

// Exec executes the deletion query.
func (spkdo *SubjectPublicKeysDeleteOne) Exec(ctx context.Context) error {
	n, err := spkdo.spkd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subjectpublickeys.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spkdo *SubjectPublicKeysDeleteOne) ExecX(ctx context.Context) {
	if err := spkdo.Exec(ctx); err != nil {
		panic(err)
	}
}