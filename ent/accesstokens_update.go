// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/accesstokens"
	"authorization-service/ent/clients"
	"authorization-service/ent/predicate"
	"authorization-service/ent/session"
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"golang.org/x/text/language"
)

// AccessTokensUpdate is the builder for updating AccessTokens entities.
type AccessTokensUpdate struct {
	config
	hooks    []Hook
	mutation *AccessTokensMutation
}

// Where appends a list predicates to the AccessTokensUpdate builder.
func (atu *AccessTokensUpdate) Where(ps ...predicate.AccessTokens) *AccessTokensUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetRequestID sets the "request_id" field.
func (atu *AccessTokensUpdate) SetRequestID(s string) *AccessTokensUpdate {
	atu.mutation.SetRequestID(s)
	return atu
}

// SetRequestedAt sets the "requestedAt" field.
func (atu *AccessTokensUpdate) SetRequestedAt(t time.Time) *AccessTokensUpdate {
	atu.mutation.SetRequestedAt(t)
	return atu
}

// SetScopes sets the "scopes" field.
func (atu *AccessTokensUpdate) SetScopes(s []string) *AccessTokensUpdate {
	atu.mutation.SetScopes(s)
	return atu
}

// AppendScopes appends s to the "scopes" field.
func (atu *AccessTokensUpdate) AppendScopes(s []string) *AccessTokensUpdate {
	atu.mutation.AppendScopes(s)
	return atu
}

// SetGrantedScopes sets the "granted_scopes" field.
func (atu *AccessTokensUpdate) SetGrantedScopes(s []string) *AccessTokensUpdate {
	atu.mutation.SetGrantedScopes(s)
	return atu
}

// AppendGrantedScopes appends s to the "granted_scopes" field.
func (atu *AccessTokensUpdate) AppendGrantedScopes(s []string) *AccessTokensUpdate {
	atu.mutation.AppendGrantedScopes(s)
	return atu
}

// SetRequestedAudience sets the "requested_audience" field.
func (atu *AccessTokensUpdate) SetRequestedAudience(s []string) *AccessTokensUpdate {
	atu.mutation.SetRequestedAudience(s)
	return atu
}

// AppendRequestedAudience appends s to the "requested_audience" field.
func (atu *AccessTokensUpdate) AppendRequestedAudience(s []string) *AccessTokensUpdate {
	atu.mutation.AppendRequestedAudience(s)
	return atu
}

// SetGrantedAudience sets the "granted_audience" field.
func (atu *AccessTokensUpdate) SetGrantedAudience(s []string) *AccessTokensUpdate {
	atu.mutation.SetGrantedAudience(s)
	return atu
}

// AppendGrantedAudience appends s to the "granted_audience" field.
func (atu *AccessTokensUpdate) AppendGrantedAudience(s []string) *AccessTokensUpdate {
	atu.mutation.AppendGrantedAudience(s)
	return atu
}

// SetForm sets the "form" field.
func (atu *AccessTokensUpdate) SetForm(u url.Values) *AccessTokensUpdate {
	atu.mutation.SetForm(u)
	return atu
}

// SetLang sets the "lang" field.
func (atu *AccessTokensUpdate) SetLang(l language.Tag) *AccessTokensUpdate {
	atu.mutation.SetLang(l)
	return atu
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (atu *AccessTokensUpdate) SetNillableLang(l *language.Tag) *AccessTokensUpdate {
	if l != nil {
		atu.SetLang(*l)
	}
	return atu
}

// ClearLang clears the value of the "lang" field.
func (atu *AccessTokensUpdate) ClearLang() *AccessTokensUpdate {
	atu.mutation.ClearLang()
	return atu
}

// SetClientIDID sets the "client_id" edge to the Clients entity by ID.
func (atu *AccessTokensUpdate) SetClientIDID(id string) *AccessTokensUpdate {
	atu.mutation.SetClientIDID(id)
	return atu
}

// SetNillableClientIDID sets the "client_id" edge to the Clients entity by ID if the given value is not nil.
func (atu *AccessTokensUpdate) SetNillableClientIDID(id *string) *AccessTokensUpdate {
	if id != nil {
		atu = atu.SetClientIDID(*id)
	}
	return atu
}

// SetClientID sets the "client_id" edge to the Clients entity.
func (atu *AccessTokensUpdate) SetClientID(c *Clients) *AccessTokensUpdate {
	return atu.SetClientIDID(c.ID)
}

// SetSessionIDID sets the "session_id" edge to the Session entity by ID.
func (atu *AccessTokensUpdate) SetSessionIDID(id string) *AccessTokensUpdate {
	atu.mutation.SetSessionIDID(id)
	return atu
}

// SetNillableSessionIDID sets the "session_id" edge to the Session entity by ID if the given value is not nil.
func (atu *AccessTokensUpdate) SetNillableSessionIDID(id *string) *AccessTokensUpdate {
	if id != nil {
		atu = atu.SetSessionIDID(*id)
	}
	return atu
}

// SetSessionID sets the "session_id" edge to the Session entity.
func (atu *AccessTokensUpdate) SetSessionID(s *Session) *AccessTokensUpdate {
	return atu.SetSessionIDID(s.ID)
}

// Mutation returns the AccessTokensMutation object of the builder.
func (atu *AccessTokensUpdate) Mutation() *AccessTokensMutation {
	return atu.mutation
}

// ClearClientID clears the "client_id" edge to the Clients entity.
func (atu *AccessTokensUpdate) ClearClientID() *AccessTokensUpdate {
	atu.mutation.ClearClientID()
	return atu
}

// ClearSessionID clears the "session_id" edge to the Session entity.
func (atu *AccessTokensUpdate) ClearSessionID() *AccessTokensUpdate {
	atu.mutation.ClearSessionID()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AccessTokensUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AccessTokensUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AccessTokensUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AccessTokensUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atu *AccessTokensUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesstokens.Table, accesstokens.Columns, sqlgraph.NewFieldSpec(accesstokens.FieldID, field.TypeString))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.RequestID(); ok {
		_spec.SetField(accesstokens.FieldRequestID, field.TypeString, value)
	}
	if value, ok := atu.mutation.RequestedAt(); ok {
		_spec.SetField(accesstokens.FieldRequestedAt, field.TypeTime, value)
	}
	if value, ok := atu.mutation.Scopes(); ok {
		_spec.SetField(accesstokens.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := atu.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldScopes, value)
		})
	}
	if value, ok := atu.mutation.GrantedScopes(); ok {
		_spec.SetField(accesstokens.FieldGrantedScopes, field.TypeJSON, value)
	}
	if value, ok := atu.mutation.AppendedGrantedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldGrantedScopes, value)
		})
	}
	if value, ok := atu.mutation.RequestedAudience(); ok {
		_spec.SetField(accesstokens.FieldRequestedAudience, field.TypeJSON, value)
	}
	if value, ok := atu.mutation.AppendedRequestedAudience(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldRequestedAudience, value)
		})
	}
	if value, ok := atu.mutation.GrantedAudience(); ok {
		_spec.SetField(accesstokens.FieldGrantedAudience, field.TypeJSON, value)
	}
	if value, ok := atu.mutation.AppendedGrantedAudience(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldGrantedAudience, value)
		})
	}
	if value, ok := atu.mutation.Form(); ok {
		_spec.SetField(accesstokens.FieldForm, field.TypeJSON, value)
	}
	if value, ok := atu.mutation.Lang(); ok {
		_spec.SetField(accesstokens.FieldLang, field.TypeJSON, value)
	}
	if atu.mutation.LangCleared() {
		_spec.ClearField(accesstokens.FieldLang, field.TypeJSON)
	}
	if atu.mutation.ClientIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.ClientIDTable,
			Columns: []string{accesstokens.ClientIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clients.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.ClientIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.ClientIDTable,
			Columns: []string{accesstokens.ClientIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clients.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.SessionIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.SessionIDTable,
			Columns: []string{accesstokens.SessionIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.SessionIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.SessionIDTable,
			Columns: []string{accesstokens.SessionIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AccessTokensUpdateOne is the builder for updating a single AccessTokens entity.
type AccessTokensUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessTokensMutation
}

// SetRequestID sets the "request_id" field.
func (atuo *AccessTokensUpdateOne) SetRequestID(s string) *AccessTokensUpdateOne {
	atuo.mutation.SetRequestID(s)
	return atuo
}

// SetRequestedAt sets the "requestedAt" field.
func (atuo *AccessTokensUpdateOne) SetRequestedAt(t time.Time) *AccessTokensUpdateOne {
	atuo.mutation.SetRequestedAt(t)
	return atuo
}

// SetScopes sets the "scopes" field.
func (atuo *AccessTokensUpdateOne) SetScopes(s []string) *AccessTokensUpdateOne {
	atuo.mutation.SetScopes(s)
	return atuo
}

// AppendScopes appends s to the "scopes" field.
func (atuo *AccessTokensUpdateOne) AppendScopes(s []string) *AccessTokensUpdateOne {
	atuo.mutation.AppendScopes(s)
	return atuo
}

// SetGrantedScopes sets the "granted_scopes" field.
func (atuo *AccessTokensUpdateOne) SetGrantedScopes(s []string) *AccessTokensUpdateOne {
	atuo.mutation.SetGrantedScopes(s)
	return atuo
}

// AppendGrantedScopes appends s to the "granted_scopes" field.
func (atuo *AccessTokensUpdateOne) AppendGrantedScopes(s []string) *AccessTokensUpdateOne {
	atuo.mutation.AppendGrantedScopes(s)
	return atuo
}

// SetRequestedAudience sets the "requested_audience" field.
func (atuo *AccessTokensUpdateOne) SetRequestedAudience(s []string) *AccessTokensUpdateOne {
	atuo.mutation.SetRequestedAudience(s)
	return atuo
}

// AppendRequestedAudience appends s to the "requested_audience" field.
func (atuo *AccessTokensUpdateOne) AppendRequestedAudience(s []string) *AccessTokensUpdateOne {
	atuo.mutation.AppendRequestedAudience(s)
	return atuo
}

// SetGrantedAudience sets the "granted_audience" field.
func (atuo *AccessTokensUpdateOne) SetGrantedAudience(s []string) *AccessTokensUpdateOne {
	atuo.mutation.SetGrantedAudience(s)
	return atuo
}

// AppendGrantedAudience appends s to the "granted_audience" field.
func (atuo *AccessTokensUpdateOne) AppendGrantedAudience(s []string) *AccessTokensUpdateOne {
	atuo.mutation.AppendGrantedAudience(s)
	return atuo
}

// SetForm sets the "form" field.
func (atuo *AccessTokensUpdateOne) SetForm(u url.Values) *AccessTokensUpdateOne {
	atuo.mutation.SetForm(u)
	return atuo
}

// SetLang sets the "lang" field.
func (atuo *AccessTokensUpdateOne) SetLang(l language.Tag) *AccessTokensUpdateOne {
	atuo.mutation.SetLang(l)
	return atuo
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (atuo *AccessTokensUpdateOne) SetNillableLang(l *language.Tag) *AccessTokensUpdateOne {
	if l != nil {
		atuo.SetLang(*l)
	}
	return atuo
}

// ClearLang clears the value of the "lang" field.
func (atuo *AccessTokensUpdateOne) ClearLang() *AccessTokensUpdateOne {
	atuo.mutation.ClearLang()
	return atuo
}

// SetClientIDID sets the "client_id" edge to the Clients entity by ID.
func (atuo *AccessTokensUpdateOne) SetClientIDID(id string) *AccessTokensUpdateOne {
	atuo.mutation.SetClientIDID(id)
	return atuo
}

// SetNillableClientIDID sets the "client_id" edge to the Clients entity by ID if the given value is not nil.
func (atuo *AccessTokensUpdateOne) SetNillableClientIDID(id *string) *AccessTokensUpdateOne {
	if id != nil {
		atuo = atuo.SetClientIDID(*id)
	}
	return atuo
}

// SetClientID sets the "client_id" edge to the Clients entity.
func (atuo *AccessTokensUpdateOne) SetClientID(c *Clients) *AccessTokensUpdateOne {
	return atuo.SetClientIDID(c.ID)
}

// SetSessionIDID sets the "session_id" edge to the Session entity by ID.
func (atuo *AccessTokensUpdateOne) SetSessionIDID(id string) *AccessTokensUpdateOne {
	atuo.mutation.SetSessionIDID(id)
	return atuo
}

// SetNillableSessionIDID sets the "session_id" edge to the Session entity by ID if the given value is not nil.
func (atuo *AccessTokensUpdateOne) SetNillableSessionIDID(id *string) *AccessTokensUpdateOne {
	if id != nil {
		atuo = atuo.SetSessionIDID(*id)
	}
	return atuo
}

// SetSessionID sets the "session_id" edge to the Session entity.
func (atuo *AccessTokensUpdateOne) SetSessionID(s *Session) *AccessTokensUpdateOne {
	return atuo.SetSessionIDID(s.ID)
}

// Mutation returns the AccessTokensMutation object of the builder.
func (atuo *AccessTokensUpdateOne) Mutation() *AccessTokensMutation {
	return atuo.mutation
}

// ClearClientID clears the "client_id" edge to the Clients entity.
func (atuo *AccessTokensUpdateOne) ClearClientID() *AccessTokensUpdateOne {
	atuo.mutation.ClearClientID()
	return atuo
}

// ClearSessionID clears the "session_id" edge to the Session entity.
func (atuo *AccessTokensUpdateOne) ClearSessionID() *AccessTokensUpdateOne {
	atuo.mutation.ClearSessionID()
	return atuo
}

// Where appends a list predicates to the AccessTokensUpdate builder.
func (atuo *AccessTokensUpdateOne) Where(ps ...predicate.AccessTokens) *AccessTokensUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AccessTokensUpdateOne) Select(field string, fields ...string) *AccessTokensUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AccessTokens entity.
func (atuo *AccessTokensUpdateOne) Save(ctx context.Context) (*AccessTokens, error) {
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AccessTokensUpdateOne) SaveX(ctx context.Context) *AccessTokens {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AccessTokensUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AccessTokensUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atuo *AccessTokensUpdateOne) sqlSave(ctx context.Context) (_node *AccessTokens, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesstokens.Table, accesstokens.Columns, sqlgraph.NewFieldSpec(accesstokens.FieldID, field.TypeString))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccessTokens.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesstokens.FieldID)
		for _, f := range fields {
			if !accesstokens.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accesstokens.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.RequestID(); ok {
		_spec.SetField(accesstokens.FieldRequestID, field.TypeString, value)
	}
	if value, ok := atuo.mutation.RequestedAt(); ok {
		_spec.SetField(accesstokens.FieldRequestedAt, field.TypeTime, value)
	}
	if value, ok := atuo.mutation.Scopes(); ok {
		_spec.SetField(accesstokens.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := atuo.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldScopes, value)
		})
	}
	if value, ok := atuo.mutation.GrantedScopes(); ok {
		_spec.SetField(accesstokens.FieldGrantedScopes, field.TypeJSON, value)
	}
	if value, ok := atuo.mutation.AppendedGrantedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldGrantedScopes, value)
		})
	}
	if value, ok := atuo.mutation.RequestedAudience(); ok {
		_spec.SetField(accesstokens.FieldRequestedAudience, field.TypeJSON, value)
	}
	if value, ok := atuo.mutation.AppendedRequestedAudience(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldRequestedAudience, value)
		})
	}
	if value, ok := atuo.mutation.GrantedAudience(); ok {
		_spec.SetField(accesstokens.FieldGrantedAudience, field.TypeJSON, value)
	}
	if value, ok := atuo.mutation.AppendedGrantedAudience(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, accesstokens.FieldGrantedAudience, value)
		})
	}
	if value, ok := atuo.mutation.Form(); ok {
		_spec.SetField(accesstokens.FieldForm, field.TypeJSON, value)
	}
	if value, ok := atuo.mutation.Lang(); ok {
		_spec.SetField(accesstokens.FieldLang, field.TypeJSON, value)
	}
	if atuo.mutation.LangCleared() {
		_spec.ClearField(accesstokens.FieldLang, field.TypeJSON)
	}
	if atuo.mutation.ClientIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.ClientIDTable,
			Columns: []string{accesstokens.ClientIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clients.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.ClientIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.ClientIDTable,
			Columns: []string{accesstokens.ClientIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clients.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.SessionIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.SessionIDTable,
			Columns: []string{accesstokens.SessionIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.SessionIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstokens.SessionIDTable,
			Columns: []string{accesstokens.SessionIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AccessTokens{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
