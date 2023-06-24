// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/accesstokens"
	"authorization-service/ent/authorizecodes"
	"authorization-service/ent/idsessions"
	"authorization-service/ent/pkces"
	"authorization-service/ent/refreshtokens"
	"authorization-service/ent/session"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetExpiresAt sets the "expires_at" field.
func (sc *SessionCreate) SetExpiresAt(m map[string]time.Time) *SessionCreate {
	sc.mutation.SetExpiresAt(m)
	return sc
}

// SetUsername sets the "username" field.
func (sc *SessionCreate) SetUsername(s string) *SessionCreate {
	sc.mutation.SetUsername(s)
	return sc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (sc *SessionCreate) SetNillableUsername(s *string) *SessionCreate {
	if s != nil {
		sc.SetUsername(*s)
	}
	return sc
}

// SetSubject sets the "subject" field.
func (sc *SessionCreate) SetSubject(s string) *SessionCreate {
	sc.mutation.SetSubject(s)
	return sc
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (sc *SessionCreate) SetNillableSubject(s *string) *SessionCreate {
	if s != nil {
		sc.SetSubject(*s)
	}
	return sc
}

// SetExtra sets the "extra" field.
func (sc *SessionCreate) SetExtra(m map[string]interface{}) *SessionCreate {
	sc.mutation.SetExtra(m)
	return sc
}

// SetSession sets the "session" field.
func (sc *SessionCreate) SetSession(a any) *SessionCreate {
	sc.mutation.SetSession(a)
	return sc
}

// SetID sets the "id" field.
func (sc *SessionCreate) SetID(s string) *SessionCreate {
	sc.mutation.SetID(s)
	return sc
}

// AddAccessTokenIDs adds the "access_token" edge to the AccessTokens entity by IDs.
func (sc *SessionCreate) AddAccessTokenIDs(ids ...string) *SessionCreate {
	sc.mutation.AddAccessTokenIDs(ids...)
	return sc
}

// AddAccessToken adds the "access_token" edges to the AccessTokens entity.
func (sc *SessionCreate) AddAccessToken(a ...*AccessTokens) *SessionCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAccessTokenIDs(ids...)
}

// AddAuthorizeCodeIDs adds the "authorize_code" edge to the AuthorizeCodes entity by IDs.
func (sc *SessionCreate) AddAuthorizeCodeIDs(ids ...string) *SessionCreate {
	sc.mutation.AddAuthorizeCodeIDs(ids...)
	return sc
}

// AddAuthorizeCode adds the "authorize_code" edges to the AuthorizeCodes entity.
func (sc *SessionCreate) AddAuthorizeCode(a ...*AuthorizeCodes) *SessionCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAuthorizeCodeIDs(ids...)
}

// AddRefreshTokenIDs adds the "refresh_token" edge to the RefreshTokens entity by IDs.
func (sc *SessionCreate) AddRefreshTokenIDs(ids ...string) *SessionCreate {
	sc.mutation.AddRefreshTokenIDs(ids...)
	return sc
}

// AddRefreshToken adds the "refresh_token" edges to the RefreshTokens entity.
func (sc *SessionCreate) AddRefreshToken(r ...*RefreshTokens) *SessionCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sc.AddRefreshTokenIDs(ids...)
}

// AddIDSessionIDs adds the "id_session" edge to the IDSessions entity by IDs.
func (sc *SessionCreate) AddIDSessionIDs(ids ...string) *SessionCreate {
	sc.mutation.AddIDSessionIDs(ids...)
	return sc
}

// AddIDSession adds the "id_session" edges to the IDSessions entity.
func (sc *SessionCreate) AddIDSession(i ...*IDSessions) *SessionCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return sc.AddIDSessionIDs(ids...)
}

// AddPkceIDs adds the "pkce" edge to the PKCES entity by IDs.
func (sc *SessionCreate) AddPkceIDs(ids ...string) *SessionCreate {
	sc.mutation.AddPkceIDs(ids...)
	return sc
}

// AddPkce adds the "pkce" edges to the PKCES entity.
func (sc *SessionCreate) AddPkce(p ...*PKCES) *SessionCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPkceIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "Session.expires_at"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Session.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(session.Table, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeJSON, value)
		_node.ExpiresAt = value
	}
	if value, ok := sc.mutation.Username(); ok {
		_spec.SetField(session.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := sc.mutation.Subject(); ok {
		_spec.SetField(session.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := sc.mutation.Extra(); ok {
		_spec.SetField(session.FieldExtra, field.TypeJSON, value)
		_node.Extra = value
	}
	if value, ok := sc.mutation.Session(); ok {
		_spec.SetField(session.FieldSession, field.TypeJSON, value)
		_node.Session = value
	}
	if nodes := sc.mutation.AccessTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AccessTokenTable,
			Columns: []string{session.AccessTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accesstokens.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.AuthorizeCodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AuthorizeCodeTable,
			Columns: []string{session.AuthorizeCodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authorizecodes.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.RefreshTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.RefreshTokenTable,
			Columns: []string{session.RefreshTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtokens.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.IDSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.IDSessionTable,
			Columns: []string{session.IDSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(idsessions.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PkceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.PkceTable,
			Columns: []string{session.PkceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pkces.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
