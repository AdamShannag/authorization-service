// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/refreshtokens"
	"authorization-service/ent/request"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RefreshTokensCreate is the builder for creating a RefreshTokens entity.
type RefreshTokensCreate struct {
	config
	mutation *RefreshTokensMutation
	hooks    []Hook
}

// SetActive sets the "active" field.
func (rtc *RefreshTokensCreate) SetActive(b bool) *RefreshTokensCreate {
	rtc.mutation.SetActive(b)
	return rtc
}

// SetID sets the "id" field.
func (rtc *RefreshTokensCreate) SetID(s string) *RefreshTokensCreate {
	rtc.mutation.SetID(s)
	return rtc
}

// SetRequestIDID sets the "request_id" edge to the Request entity by ID.
func (rtc *RefreshTokensCreate) SetRequestIDID(id string) *RefreshTokensCreate {
	rtc.mutation.SetRequestIDID(id)
	return rtc
}

// SetNillableRequestIDID sets the "request_id" edge to the Request entity by ID if the given value is not nil.
func (rtc *RefreshTokensCreate) SetNillableRequestIDID(id *string) *RefreshTokensCreate {
	if id != nil {
		rtc = rtc.SetRequestIDID(*id)
	}
	return rtc
}

// SetRequestID sets the "request_id" edge to the Request entity.
func (rtc *RefreshTokensCreate) SetRequestID(r *Request) *RefreshTokensCreate {
	return rtc.SetRequestIDID(r.ID)
}

// Mutation returns the RefreshTokensMutation object of the builder.
func (rtc *RefreshTokensCreate) Mutation() *RefreshTokensMutation {
	return rtc.mutation
}

// Save creates the RefreshTokens in the database.
func (rtc *RefreshTokensCreate) Save(ctx context.Context) (*RefreshTokens, error) {
	return withHooks(ctx, rtc.sqlSave, rtc.mutation, rtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rtc *RefreshTokensCreate) SaveX(ctx context.Context) *RefreshTokens {
	v, err := rtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rtc *RefreshTokensCreate) Exec(ctx context.Context) error {
	_, err := rtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtc *RefreshTokensCreate) ExecX(ctx context.Context) {
	if err := rtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtc *RefreshTokensCreate) check() error {
	if _, ok := rtc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "RefreshTokens.active"`)}
	}
	return nil
}

func (rtc *RefreshTokensCreate) sqlSave(ctx context.Context) (*RefreshTokens, error) {
	if err := rtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RefreshTokens.ID type: %T", _spec.ID.Value)
		}
	}
	rtc.mutation.id = &_node.ID
	rtc.mutation.done = true
	return _node, nil
}

func (rtc *RefreshTokensCreate) createSpec() (*RefreshTokens, *sqlgraph.CreateSpec) {
	var (
		_node = &RefreshTokens{config: rtc.config}
		_spec = sqlgraph.NewCreateSpec(refreshtokens.Table, sqlgraph.NewFieldSpec(refreshtokens.FieldID, field.TypeString))
	)
	if id, ok := rtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rtc.mutation.Active(); ok {
		_spec.SetField(refreshtokens.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if nodes := rtc.mutation.RequestIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   refreshtokens.RequestIDTable,
			Columns: []string{refreshtokens.RequestIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.request_refresh_token = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RefreshTokensCreateBulk is the builder for creating many RefreshTokens entities in bulk.
type RefreshTokensCreateBulk struct {
	config
	builders []*RefreshTokensCreate
}

// Save creates the RefreshTokens entities in the database.
func (rtcb *RefreshTokensCreateBulk) Save(ctx context.Context) ([]*RefreshTokens, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rtcb.builders))
	nodes := make([]*RefreshTokens, len(rtcb.builders))
	mutators := make([]Mutator, len(rtcb.builders))
	for i := range rtcb.builders {
		func(i int, root context.Context) {
			builder := rtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RefreshTokensMutation)
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
					_, err = mutators[i+1].Mutate(root, rtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rtcb *RefreshTokensCreateBulk) SaveX(ctx context.Context) []*RefreshTokens {
	v, err := rtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rtcb *RefreshTokensCreateBulk) Exec(ctx context.Context) error {
	_, err := rtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtcb *RefreshTokensCreateBulk) ExecX(ctx context.Context) {
	if err := rtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
