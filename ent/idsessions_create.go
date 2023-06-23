// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/idsessions"
	"authorization-service/ent/request"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IDSessionsCreate is the builder for creating a IDSessions entity.
type IDSessionsCreate struct {
	config
	mutation *IDSessionsMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (isc *IDSessionsCreate) SetID(s string) *IDSessionsCreate {
	isc.mutation.SetID(s)
	return isc
}

// SetRequestIDID sets the "request_id" edge to the Request entity by ID.
func (isc *IDSessionsCreate) SetRequestIDID(id string) *IDSessionsCreate {
	isc.mutation.SetRequestIDID(id)
	return isc
}

// SetNillableRequestIDID sets the "request_id" edge to the Request entity by ID if the given value is not nil.
func (isc *IDSessionsCreate) SetNillableRequestIDID(id *string) *IDSessionsCreate {
	if id != nil {
		isc = isc.SetRequestIDID(*id)
	}
	return isc
}

// SetRequestID sets the "request_id" edge to the Request entity.
func (isc *IDSessionsCreate) SetRequestID(r *Request) *IDSessionsCreate {
	return isc.SetRequestIDID(r.ID)
}

// Mutation returns the IDSessionsMutation object of the builder.
func (isc *IDSessionsCreate) Mutation() *IDSessionsMutation {
	return isc.mutation
}

// Save creates the IDSessions in the database.
func (isc *IDSessionsCreate) Save(ctx context.Context) (*IDSessions, error) {
	return withHooks(ctx, isc.sqlSave, isc.mutation, isc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (isc *IDSessionsCreate) SaveX(ctx context.Context) *IDSessions {
	v, err := isc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (isc *IDSessionsCreate) Exec(ctx context.Context) error {
	_, err := isc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isc *IDSessionsCreate) ExecX(ctx context.Context) {
	if err := isc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (isc *IDSessionsCreate) check() error {
	return nil
}

func (isc *IDSessionsCreate) sqlSave(ctx context.Context) (*IDSessions, error) {
	if err := isc.check(); err != nil {
		return nil, err
	}
	_node, _spec := isc.createSpec()
	if err := sqlgraph.CreateNode(ctx, isc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected IDSessions.ID type: %T", _spec.ID.Value)
		}
	}
	isc.mutation.id = &_node.ID
	isc.mutation.done = true
	return _node, nil
}

func (isc *IDSessionsCreate) createSpec() (*IDSessions, *sqlgraph.CreateSpec) {
	var (
		_node = &IDSessions{config: isc.config}
		_spec = sqlgraph.NewCreateSpec(idsessions.Table, sqlgraph.NewFieldSpec(idsessions.FieldID, field.TypeString))
	)
	if id, ok := isc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := isc.mutation.RequestIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   idsessions.RequestIDTable,
			Columns: []string{idsessions.RequestIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.request_id_session = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IDSessionsCreateBulk is the builder for creating many IDSessions entities in bulk.
type IDSessionsCreateBulk struct {
	config
	builders []*IDSessionsCreate
}

// Save creates the IDSessions entities in the database.
func (iscb *IDSessionsCreateBulk) Save(ctx context.Context) ([]*IDSessions, error) {
	specs := make([]*sqlgraph.CreateSpec, len(iscb.builders))
	nodes := make([]*IDSessions, len(iscb.builders))
	mutators := make([]Mutator, len(iscb.builders))
	for i := range iscb.builders {
		func(i int, root context.Context) {
			builder := iscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IDSessionsMutation)
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
					_, err = mutators[i+1].Mutate(root, iscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iscb *IDSessionsCreateBulk) SaveX(ctx context.Context) []*IDSessions {
	v, err := iscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iscb *IDSessionsCreateBulk) Exec(ctx context.Context) error {
	_, err := iscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iscb *IDSessionsCreateBulk) ExecX(ctx context.Context) {
	if err := iscb.Exec(ctx); err != nil {
		panic(err)
	}
}
