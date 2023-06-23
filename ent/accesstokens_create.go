// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/accesstokens"
	"authorization-service/ent/request"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessTokensCreate is the builder for creating a AccessTokens entity.
type AccessTokensCreate struct {
	config
	mutation *AccessTokensMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (atc *AccessTokensCreate) SetID(s string) *AccessTokensCreate {
	atc.mutation.SetID(s)
	return atc
}

// SetRequestIDID sets the "request_id" edge to the Request entity by ID.
func (atc *AccessTokensCreate) SetRequestIDID(id string) *AccessTokensCreate {
	atc.mutation.SetRequestIDID(id)
	return atc
}

// SetNillableRequestIDID sets the "request_id" edge to the Request entity by ID if the given value is not nil.
func (atc *AccessTokensCreate) SetNillableRequestIDID(id *string) *AccessTokensCreate {
	if id != nil {
		atc = atc.SetRequestIDID(*id)
	}
	return atc
}

// SetRequestID sets the "request_id" edge to the Request entity.
func (atc *AccessTokensCreate) SetRequestID(r *Request) *AccessTokensCreate {
	return atc.SetRequestIDID(r.ID)
}

// Mutation returns the AccessTokensMutation object of the builder.
func (atc *AccessTokensCreate) Mutation() *AccessTokensMutation {
	return atc.mutation
}

// Save creates the AccessTokens in the database.
func (atc *AccessTokensCreate) Save(ctx context.Context) (*AccessTokens, error) {
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AccessTokensCreate) SaveX(ctx context.Context) *AccessTokens {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AccessTokensCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AccessTokensCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AccessTokensCreate) check() error {
	return nil
}

func (atc *AccessTokensCreate) sqlSave(ctx context.Context) (*AccessTokens, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AccessTokens.ID type: %T", _spec.ID.Value)
		}
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *AccessTokensCreate) createSpec() (*AccessTokens, *sqlgraph.CreateSpec) {
	var (
		_node = &AccessTokens{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(accesstokens.Table, sqlgraph.NewFieldSpec(accesstokens.FieldID, field.TypeString))
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := atc.mutation.RequestIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   accesstokens.RequestIDTable,
			Columns: []string{accesstokens.RequestIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.request_access_token = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccessTokensCreateBulk is the builder for creating many AccessTokens entities in bulk.
type AccessTokensCreateBulk struct {
	config
	builders []*AccessTokensCreate
}

// Save creates the AccessTokens entities in the database.
func (atcb *AccessTokensCreateBulk) Save(ctx context.Context) ([]*AccessTokens, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AccessTokens, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccessTokensMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AccessTokensCreateBulk) SaveX(ctx context.Context) []*AccessTokens {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AccessTokensCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AccessTokensCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}