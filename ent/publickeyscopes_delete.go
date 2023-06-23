// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/predicate"
	"authorization-service/ent/publickeyscopes"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicKeyScopesDelete is the builder for deleting a PublicKeyScopes entity.
type PublicKeyScopesDelete struct {
	config
	hooks    []Hook
	mutation *PublicKeyScopesMutation
}

// Where appends a list predicates to the PublicKeyScopesDelete builder.
func (pksd *PublicKeyScopesDelete) Where(ps ...predicate.PublicKeyScopes) *PublicKeyScopesDelete {
	pksd.mutation.Where(ps...)
	return pksd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pksd *PublicKeyScopesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pksd.sqlExec, pksd.mutation, pksd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pksd *PublicKeyScopesDelete) ExecX(ctx context.Context) int {
	n, err := pksd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pksd *PublicKeyScopesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(publickeyscopes.Table, sqlgraph.NewFieldSpec(publickeyscopes.FieldID, field.TypeString))
	if ps := pksd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pksd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pksd.mutation.done = true
	return affected, err
}

// PublicKeyScopesDeleteOne is the builder for deleting a single PublicKeyScopes entity.
type PublicKeyScopesDeleteOne struct {
	pksd *PublicKeyScopesDelete
}

// Where appends a list predicates to the PublicKeyScopesDelete builder.
func (pksdo *PublicKeyScopesDeleteOne) Where(ps ...predicate.PublicKeyScopes) *PublicKeyScopesDeleteOne {
	pksdo.pksd.mutation.Where(ps...)
	return pksdo
}

// Exec executes the deletion query.
func (pksdo *PublicKeyScopesDeleteOne) Exec(ctx context.Context) error {
	n, err := pksdo.pksd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{publickeyscopes.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pksdo *PublicKeyScopesDeleteOne) ExecX(ctx context.Context) {
	if err := pksdo.Exec(ctx); err != nil {
		panic(err)
	}
}
