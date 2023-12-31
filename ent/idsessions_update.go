// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/idsessions"
	"authorization-service/ent/predicate"
	"authorization-service/ent/request"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IDSessionsUpdate is the builder for updating IDSessions entities.
type IDSessionsUpdate struct {
	config
	hooks    []Hook
	mutation *IDSessionsMutation
}

// Where appends a list predicates to the IDSessionsUpdate builder.
func (isu *IDSessionsUpdate) Where(ps ...predicate.IDSessions) *IDSessionsUpdate {
	isu.mutation.Where(ps...)
	return isu
}

// SetRequestIDID sets the "request_id" edge to the Request entity by ID.
func (isu *IDSessionsUpdate) SetRequestIDID(id string) *IDSessionsUpdate {
	isu.mutation.SetRequestIDID(id)
	return isu
}

// SetNillableRequestIDID sets the "request_id" edge to the Request entity by ID if the given value is not nil.
func (isu *IDSessionsUpdate) SetNillableRequestIDID(id *string) *IDSessionsUpdate {
	if id != nil {
		isu = isu.SetRequestIDID(*id)
	}
	return isu
}

// SetRequestID sets the "request_id" edge to the Request entity.
func (isu *IDSessionsUpdate) SetRequestID(r *Request) *IDSessionsUpdate {
	return isu.SetRequestIDID(r.ID)
}

// Mutation returns the IDSessionsMutation object of the builder.
func (isu *IDSessionsUpdate) Mutation() *IDSessionsMutation {
	return isu.mutation
}

// ClearRequestID clears the "request_id" edge to the Request entity.
func (isu *IDSessionsUpdate) ClearRequestID() *IDSessionsUpdate {
	isu.mutation.ClearRequestID()
	return isu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (isu *IDSessionsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, isu.sqlSave, isu.mutation, isu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (isu *IDSessionsUpdate) SaveX(ctx context.Context) int {
	affected, err := isu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (isu *IDSessionsUpdate) Exec(ctx context.Context) error {
	_, err := isu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isu *IDSessionsUpdate) ExecX(ctx context.Context) {
	if err := isu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (isu *IDSessionsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(idsessions.Table, idsessions.Columns, sqlgraph.NewFieldSpec(idsessions.FieldID, field.TypeString))
	if ps := isu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if isu.mutation.RequestIDCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isu.mutation.RequestIDIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, isu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{idsessions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	isu.mutation.done = true
	return n, nil
}

// IDSessionsUpdateOne is the builder for updating a single IDSessions entity.
type IDSessionsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IDSessionsMutation
}

// SetRequestIDID sets the "request_id" edge to the Request entity by ID.
func (isuo *IDSessionsUpdateOne) SetRequestIDID(id string) *IDSessionsUpdateOne {
	isuo.mutation.SetRequestIDID(id)
	return isuo
}

// SetNillableRequestIDID sets the "request_id" edge to the Request entity by ID if the given value is not nil.
func (isuo *IDSessionsUpdateOne) SetNillableRequestIDID(id *string) *IDSessionsUpdateOne {
	if id != nil {
		isuo = isuo.SetRequestIDID(*id)
	}
	return isuo
}

// SetRequestID sets the "request_id" edge to the Request entity.
func (isuo *IDSessionsUpdateOne) SetRequestID(r *Request) *IDSessionsUpdateOne {
	return isuo.SetRequestIDID(r.ID)
}

// Mutation returns the IDSessionsMutation object of the builder.
func (isuo *IDSessionsUpdateOne) Mutation() *IDSessionsMutation {
	return isuo.mutation
}

// ClearRequestID clears the "request_id" edge to the Request entity.
func (isuo *IDSessionsUpdateOne) ClearRequestID() *IDSessionsUpdateOne {
	isuo.mutation.ClearRequestID()
	return isuo
}

// Where appends a list predicates to the IDSessionsUpdate builder.
func (isuo *IDSessionsUpdateOne) Where(ps ...predicate.IDSessions) *IDSessionsUpdateOne {
	isuo.mutation.Where(ps...)
	return isuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (isuo *IDSessionsUpdateOne) Select(field string, fields ...string) *IDSessionsUpdateOne {
	isuo.fields = append([]string{field}, fields...)
	return isuo
}

// Save executes the query and returns the updated IDSessions entity.
func (isuo *IDSessionsUpdateOne) Save(ctx context.Context) (*IDSessions, error) {
	return withHooks(ctx, isuo.sqlSave, isuo.mutation, isuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (isuo *IDSessionsUpdateOne) SaveX(ctx context.Context) *IDSessions {
	node, err := isuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (isuo *IDSessionsUpdateOne) Exec(ctx context.Context) error {
	_, err := isuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isuo *IDSessionsUpdateOne) ExecX(ctx context.Context) {
	if err := isuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (isuo *IDSessionsUpdateOne) sqlSave(ctx context.Context) (_node *IDSessions, err error) {
	_spec := sqlgraph.NewUpdateSpec(idsessions.Table, idsessions.Columns, sqlgraph.NewFieldSpec(idsessions.FieldID, field.TypeString))
	id, ok := isuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IDSessions.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := isuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, idsessions.FieldID)
		for _, f := range fields {
			if !idsessions.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != idsessions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := isuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if isuo.mutation.RequestIDCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isuo.mutation.RequestIDIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IDSessions{config: isuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, isuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{idsessions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	isuo.mutation.done = true
	return _node, nil
}
