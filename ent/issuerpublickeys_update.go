// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/issuerpublickeys"
	"authorization-service/ent/predicate"
	"authorization-service/ent/subjectpublickeys"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IssuerPublicKeysUpdate is the builder for updating IssuerPublicKeys entities.
type IssuerPublicKeysUpdate struct {
	config
	hooks    []Hook
	mutation *IssuerPublicKeysMutation
}

// Where appends a list predicates to the IssuerPublicKeysUpdate builder.
func (ipku *IssuerPublicKeysUpdate) Where(ps ...predicate.IssuerPublicKeys) *IssuerPublicKeysUpdate {
	ipku.mutation.Where(ps...)
	return ipku
}

// SetSubjectPublicKeyID sets the "subject_public_key" edge to the SubjectPublicKeys entity by ID.
func (ipku *IssuerPublicKeysUpdate) SetSubjectPublicKeyID(id string) *IssuerPublicKeysUpdate {
	ipku.mutation.SetSubjectPublicKeyID(id)
	return ipku
}

// SetNillableSubjectPublicKeyID sets the "subject_public_key" edge to the SubjectPublicKeys entity by ID if the given value is not nil.
func (ipku *IssuerPublicKeysUpdate) SetNillableSubjectPublicKeyID(id *string) *IssuerPublicKeysUpdate {
	if id != nil {
		ipku = ipku.SetSubjectPublicKeyID(*id)
	}
	return ipku
}

// SetSubjectPublicKey sets the "subject_public_key" edge to the SubjectPublicKeys entity.
func (ipku *IssuerPublicKeysUpdate) SetSubjectPublicKey(s *SubjectPublicKeys) *IssuerPublicKeysUpdate {
	return ipku.SetSubjectPublicKeyID(s.ID)
}

// Mutation returns the IssuerPublicKeysMutation object of the builder.
func (ipku *IssuerPublicKeysUpdate) Mutation() *IssuerPublicKeysMutation {
	return ipku.mutation
}

// ClearSubjectPublicKey clears the "subject_public_key" edge to the SubjectPublicKeys entity.
func (ipku *IssuerPublicKeysUpdate) ClearSubjectPublicKey() *IssuerPublicKeysUpdate {
	ipku.mutation.ClearSubjectPublicKey()
	return ipku
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ipku *IssuerPublicKeysUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ipku.sqlSave, ipku.mutation, ipku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ipku *IssuerPublicKeysUpdate) SaveX(ctx context.Context) int {
	affected, err := ipku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ipku *IssuerPublicKeysUpdate) Exec(ctx context.Context) error {
	_, err := ipku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipku *IssuerPublicKeysUpdate) ExecX(ctx context.Context) {
	if err := ipku.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ipku *IssuerPublicKeysUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(issuerpublickeys.Table, issuerpublickeys.Columns, sqlgraph.NewFieldSpec(issuerpublickeys.FieldID, field.TypeString))
	if ps := ipku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ipku.mutation.SubjectPublicKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   issuerpublickeys.SubjectPublicKeyTable,
			Columns: []string{issuerpublickeys.SubjectPublicKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectpublickeys.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ipku.mutation.SubjectPublicKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   issuerpublickeys.SubjectPublicKeyTable,
			Columns: []string{issuerpublickeys.SubjectPublicKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectpublickeys.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ipku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issuerpublickeys.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ipku.mutation.done = true
	return n, nil
}

// IssuerPublicKeysUpdateOne is the builder for updating a single IssuerPublicKeys entity.
type IssuerPublicKeysUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IssuerPublicKeysMutation
}

// SetSubjectPublicKeyID sets the "subject_public_key" edge to the SubjectPublicKeys entity by ID.
func (ipkuo *IssuerPublicKeysUpdateOne) SetSubjectPublicKeyID(id string) *IssuerPublicKeysUpdateOne {
	ipkuo.mutation.SetSubjectPublicKeyID(id)
	return ipkuo
}

// SetNillableSubjectPublicKeyID sets the "subject_public_key" edge to the SubjectPublicKeys entity by ID if the given value is not nil.
func (ipkuo *IssuerPublicKeysUpdateOne) SetNillableSubjectPublicKeyID(id *string) *IssuerPublicKeysUpdateOne {
	if id != nil {
		ipkuo = ipkuo.SetSubjectPublicKeyID(*id)
	}
	return ipkuo
}

// SetSubjectPublicKey sets the "subject_public_key" edge to the SubjectPublicKeys entity.
func (ipkuo *IssuerPublicKeysUpdateOne) SetSubjectPublicKey(s *SubjectPublicKeys) *IssuerPublicKeysUpdateOne {
	return ipkuo.SetSubjectPublicKeyID(s.ID)
}

// Mutation returns the IssuerPublicKeysMutation object of the builder.
func (ipkuo *IssuerPublicKeysUpdateOne) Mutation() *IssuerPublicKeysMutation {
	return ipkuo.mutation
}

// ClearSubjectPublicKey clears the "subject_public_key" edge to the SubjectPublicKeys entity.
func (ipkuo *IssuerPublicKeysUpdateOne) ClearSubjectPublicKey() *IssuerPublicKeysUpdateOne {
	ipkuo.mutation.ClearSubjectPublicKey()
	return ipkuo
}

// Where appends a list predicates to the IssuerPublicKeysUpdate builder.
func (ipkuo *IssuerPublicKeysUpdateOne) Where(ps ...predicate.IssuerPublicKeys) *IssuerPublicKeysUpdateOne {
	ipkuo.mutation.Where(ps...)
	return ipkuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ipkuo *IssuerPublicKeysUpdateOne) Select(field string, fields ...string) *IssuerPublicKeysUpdateOne {
	ipkuo.fields = append([]string{field}, fields...)
	return ipkuo
}

// Save executes the query and returns the updated IssuerPublicKeys entity.
func (ipkuo *IssuerPublicKeysUpdateOne) Save(ctx context.Context) (*IssuerPublicKeys, error) {
	return withHooks(ctx, ipkuo.sqlSave, ipkuo.mutation, ipkuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ipkuo *IssuerPublicKeysUpdateOne) SaveX(ctx context.Context) *IssuerPublicKeys {
	node, err := ipkuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ipkuo *IssuerPublicKeysUpdateOne) Exec(ctx context.Context) error {
	_, err := ipkuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipkuo *IssuerPublicKeysUpdateOne) ExecX(ctx context.Context) {
	if err := ipkuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ipkuo *IssuerPublicKeysUpdateOne) sqlSave(ctx context.Context) (_node *IssuerPublicKeys, err error) {
	_spec := sqlgraph.NewUpdateSpec(issuerpublickeys.Table, issuerpublickeys.Columns, sqlgraph.NewFieldSpec(issuerpublickeys.FieldID, field.TypeString))
	id, ok := ipkuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IssuerPublicKeys.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ipkuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, issuerpublickeys.FieldID)
		for _, f := range fields {
			if !issuerpublickeys.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != issuerpublickeys.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ipkuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ipkuo.mutation.SubjectPublicKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   issuerpublickeys.SubjectPublicKeyTable,
			Columns: []string{issuerpublickeys.SubjectPublicKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectpublickeys.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ipkuo.mutation.SubjectPublicKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   issuerpublickeys.SubjectPublicKeyTable,
			Columns: []string{issuerpublickeys.SubjectPublicKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectpublickeys.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IssuerPublicKeys{config: ipkuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ipkuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{issuerpublickeys.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ipkuo.mutation.done = true
	return _node, nil
}
