// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/pkces"
	"authorization-service/ent/predicate"
	"authorization-service/ent/request"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PKCESUpdate is the builder for updating PKCES entities.
type PKCESUpdate struct {
	config
	hooks    []Hook
	mutation *PKCESMutation
}

// Where appends a list predicates to the PKCESUpdate builder.
func (pu *PKCESUpdate) Where(ps ...predicate.PKCES) *PKCESUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetRequestIDID sets the "request_id" edge to the Request entity by ID.
func (pu *PKCESUpdate) SetRequestIDID(id string) *PKCESUpdate {
	pu.mutation.SetRequestIDID(id)
	return pu
}

// SetNillableRequestIDID sets the "request_id" edge to the Request entity by ID if the given value is not nil.
func (pu *PKCESUpdate) SetNillableRequestIDID(id *string) *PKCESUpdate {
	if id != nil {
		pu = pu.SetRequestIDID(*id)
	}
	return pu
}

// SetRequestID sets the "request_id" edge to the Request entity.
func (pu *PKCESUpdate) SetRequestID(r *Request) *PKCESUpdate {
	return pu.SetRequestIDID(r.ID)
}

// Mutation returns the PKCESMutation object of the builder.
func (pu *PKCESUpdate) Mutation() *PKCESMutation {
	return pu.mutation
}

// ClearRequestID clears the "request_id" edge to the Request entity.
func (pu *PKCESUpdate) ClearRequestID() *PKCESUpdate {
	pu.mutation.ClearRequestID()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PKCESUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PKCESUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PKCESUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PKCESUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PKCESUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pkces.Table, pkces.Columns, sqlgraph.NewFieldSpec(pkces.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.RequestIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pkces.RequestIDTable,
			Columns: []string{pkces.RequestIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RequestIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pkces.RequestIDTable,
			Columns: []string{pkces.RequestIDColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pkces.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PKCESUpdateOne is the builder for updating a single PKCES entity.
type PKCESUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PKCESMutation
}

// SetRequestIDID sets the "request_id" edge to the Request entity by ID.
func (puo *PKCESUpdateOne) SetRequestIDID(id string) *PKCESUpdateOne {
	puo.mutation.SetRequestIDID(id)
	return puo
}

// SetNillableRequestIDID sets the "request_id" edge to the Request entity by ID if the given value is not nil.
func (puo *PKCESUpdateOne) SetNillableRequestIDID(id *string) *PKCESUpdateOne {
	if id != nil {
		puo = puo.SetRequestIDID(*id)
	}
	return puo
}

// SetRequestID sets the "request_id" edge to the Request entity.
func (puo *PKCESUpdateOne) SetRequestID(r *Request) *PKCESUpdateOne {
	return puo.SetRequestIDID(r.ID)
}

// Mutation returns the PKCESMutation object of the builder.
func (puo *PKCESUpdateOne) Mutation() *PKCESMutation {
	return puo.mutation
}

// ClearRequestID clears the "request_id" edge to the Request entity.
func (puo *PKCESUpdateOne) ClearRequestID() *PKCESUpdateOne {
	puo.mutation.ClearRequestID()
	return puo
}

// Where appends a list predicates to the PKCESUpdate builder.
func (puo *PKCESUpdateOne) Where(ps ...predicate.PKCES) *PKCESUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PKCESUpdateOne) Select(field string, fields ...string) *PKCESUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated PKCES entity.
func (puo *PKCESUpdateOne) Save(ctx context.Context) (*PKCES, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PKCESUpdateOne) SaveX(ctx context.Context) *PKCES {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PKCESUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PKCESUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PKCESUpdateOne) sqlSave(ctx context.Context) (_node *PKCES, err error) {
	_spec := sqlgraph.NewUpdateSpec(pkces.Table, pkces.Columns, sqlgraph.NewFieldSpec(pkces.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PKCES.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pkces.FieldID)
		for _, f := range fields {
			if !pkces.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pkces.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.RequestIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pkces.RequestIDTable,
			Columns: []string{pkces.RequestIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RequestIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pkces.RequestIDTable,
			Columns: []string{pkces.RequestIDColumn},
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
	_node = &PKCES{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pkces.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
