// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/clients"
	"authorization-service/ent/request"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ClientsCreate is the builder for creating a Clients entity.
type ClientsCreate struct {
	config
	mutation *ClientsMutation
	hooks    []Hook
}

// SetClientSecret sets the "client_secret" field.
func (cc *ClientsCreate) SetClientSecret(b []byte) *ClientsCreate {
	cc.mutation.SetClientSecret(b)
	return cc
}

// SetRotatedSecrets sets the "rotated_secrets" field.
func (cc *ClientsCreate) SetRotatedSecrets(u [][]uint8) *ClientsCreate {
	cc.mutation.SetRotatedSecrets(u)
	return cc
}

// SetRedirectUris sets the "redirect_uris" field.
func (cc *ClientsCreate) SetRedirectUris(s []string) *ClientsCreate {
	cc.mutation.SetRedirectUris(s)
	return cc
}

// SetGrantTypes sets the "grant_types" field.
func (cc *ClientsCreate) SetGrantTypes(s []string) *ClientsCreate {
	cc.mutation.SetGrantTypes(s)
	return cc
}

// SetResponseTypes sets the "response_types" field.
func (cc *ClientsCreate) SetResponseTypes(s []string) *ClientsCreate {
	cc.mutation.SetResponseTypes(s)
	return cc
}

// SetScopes sets the "scopes" field.
func (cc *ClientsCreate) SetScopes(s []string) *ClientsCreate {
	cc.mutation.SetScopes(s)
	return cc
}

// SetAudience sets the "audience" field.
func (cc *ClientsCreate) SetAudience(s []string) *ClientsCreate {
	cc.mutation.SetAudience(s)
	return cc
}

// SetPublic sets the "public" field.
func (cc *ClientsCreate) SetPublic(b bool) *ClientsCreate {
	cc.mutation.SetPublic(b)
	return cc
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (cc *ClientsCreate) SetNillablePublic(b *bool) *ClientsCreate {
	if b != nil {
		cc.SetPublic(*b)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ClientsCreate) SetID(s string) *ClientsCreate {
	cc.mutation.SetID(s)
	return cc
}

// AddRequestIDs adds the "requests" edge to the Request entity by IDs.
func (cc *ClientsCreate) AddRequestIDs(ids ...string) *ClientsCreate {
	cc.mutation.AddRequestIDs(ids...)
	return cc
}

// AddRequests adds the "requests" edges to the Request entity.
func (cc *ClientsCreate) AddRequests(r ...*Request) *ClientsCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRequestIDs(ids...)
}

// Mutation returns the ClientsMutation object of the builder.
func (cc *ClientsCreate) Mutation() *ClientsMutation {
	return cc.mutation
}

// Save creates the Clients in the database.
func (cc *ClientsCreate) Save(ctx context.Context) (*Clients, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClientsCreate) SaveX(ctx context.Context) *Clients {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClientsCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClientsCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClientsCreate) check() error {
	if _, ok := cc.mutation.ClientSecret(); !ok {
		return &ValidationError{Name: "client_secret", err: errors.New(`ent: missing required field "Clients.client_secret"`)}
	}
	if _, ok := cc.mutation.RotatedSecrets(); !ok {
		return &ValidationError{Name: "rotated_secrets", err: errors.New(`ent: missing required field "Clients.rotated_secrets"`)}
	}
	if _, ok := cc.mutation.RedirectUris(); !ok {
		return &ValidationError{Name: "redirect_uris", err: errors.New(`ent: missing required field "Clients.redirect_uris"`)}
	}
	if _, ok := cc.mutation.GrantTypes(); !ok {
		return &ValidationError{Name: "grant_types", err: errors.New(`ent: missing required field "Clients.grant_types"`)}
	}
	if _, ok := cc.mutation.ResponseTypes(); !ok {
		return &ValidationError{Name: "response_types", err: errors.New(`ent: missing required field "Clients.response_types"`)}
	}
	if _, ok := cc.mutation.Scopes(); !ok {
		return &ValidationError{Name: "scopes", err: errors.New(`ent: missing required field "Clients.scopes"`)}
	}
	return nil
}

func (cc *ClientsCreate) sqlSave(ctx context.Context) (*Clients, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Clients.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ClientsCreate) createSpec() (*Clients, *sqlgraph.CreateSpec) {
	var (
		_node = &Clients{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(clients.Table, sqlgraph.NewFieldSpec(clients.FieldID, field.TypeString))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.ClientSecret(); ok {
		_spec.SetField(clients.FieldClientSecret, field.TypeBytes, value)
		_node.ClientSecret = value
	}
	if value, ok := cc.mutation.RotatedSecrets(); ok {
		_spec.SetField(clients.FieldRotatedSecrets, field.TypeJSON, value)
		_node.RotatedSecrets = value
	}
	if value, ok := cc.mutation.RedirectUris(); ok {
		_spec.SetField(clients.FieldRedirectUris, field.TypeJSON, value)
		_node.RedirectUris = value
	}
	if value, ok := cc.mutation.GrantTypes(); ok {
		_spec.SetField(clients.FieldGrantTypes, field.TypeJSON, value)
		_node.GrantTypes = value
	}
	if value, ok := cc.mutation.ResponseTypes(); ok {
		_spec.SetField(clients.FieldResponseTypes, field.TypeJSON, value)
		_node.ResponseTypes = value
	}
	if value, ok := cc.mutation.Scopes(); ok {
		_spec.SetField(clients.FieldScopes, field.TypeJSON, value)
		_node.Scopes = value
	}
	if value, ok := cc.mutation.Audience(); ok {
		_spec.SetField(clients.FieldAudience, field.TypeJSON, value)
		_node.Audience = value
	}
	if value, ok := cc.mutation.Public(); ok {
		_spec.SetField(clients.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if nodes := cc.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clients.RequestsTable,
			Columns: []string{clients.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClientsCreateBulk is the builder for creating many Clients entities in bulk.
type ClientsCreateBulk struct {
	config
	builders []*ClientsCreate
}

// Save creates the Clients entities in the database.
func (ccb *ClientsCreateBulk) Save(ctx context.Context) ([]*Clients, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Clients, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClientsMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClientsCreateBulk) SaveX(ctx context.Context) []*Clients {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClientsCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClientsCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
