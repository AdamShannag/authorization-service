// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/pkces"
	"authorization-service/ent/predicate"
	"authorization-service/ent/request"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PKCESQuery is the builder for querying PKCES entities.
type PKCESQuery struct {
	config
	ctx           *QueryContext
	order         []pkces.OrderOption
	inters        []Interceptor
	predicates    []predicate.PKCES
	withRequestID *RequestQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PKCESQuery builder.
func (pq *PKCESQuery) Where(ps ...predicate.PKCES) *PKCESQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PKCESQuery) Limit(limit int) *PKCESQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PKCESQuery) Offset(offset int) *PKCESQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PKCESQuery) Unique(unique bool) *PKCESQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PKCESQuery) Order(o ...pkces.OrderOption) *PKCESQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryRequestID chains the current query on the "request_id" edge.
func (pq *PKCESQuery) QueryRequestID() *RequestQuery {
	query := (&RequestClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pkces.Table, pkces.FieldID, selector),
			sqlgraph.To(request.Table, request.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, pkces.RequestIDTable, pkces.RequestIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PKCES entity from the query.
// Returns a *NotFoundError when no PKCES was found.
func (pq *PKCESQuery) First(ctx context.Context) (*PKCES, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pkces.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PKCESQuery) FirstX(ctx context.Context) *PKCES {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PKCES ID from the query.
// Returns a *NotFoundError when no PKCES ID was found.
func (pq *PKCESQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pkces.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PKCESQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PKCES entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PKCES entity is found.
// Returns a *NotFoundError when no PKCES entities are found.
func (pq *PKCESQuery) Only(ctx context.Context) (*PKCES, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pkces.Label}
	default:
		return nil, &NotSingularError{pkces.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PKCESQuery) OnlyX(ctx context.Context) *PKCES {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PKCES ID in the query.
// Returns a *NotSingularError when more than one PKCES ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PKCESQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pkces.Label}
	default:
		err = &NotSingularError{pkces.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PKCESQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PKCESs.
func (pq *PKCESQuery) All(ctx context.Context) ([]*PKCES, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PKCES, *PKCESQuery]()
	return withInterceptors[[]*PKCES](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PKCESQuery) AllX(ctx context.Context) []*PKCES {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PKCES IDs.
func (pq *PKCESQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(pkces.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PKCESQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PKCESQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PKCESQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PKCESQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PKCESQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PKCESQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PKCESQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PKCESQuery) Clone() *PKCESQuery {
	if pq == nil {
		return nil
	}
	return &PKCESQuery{
		config:        pq.config,
		ctx:           pq.ctx.Clone(),
		order:         append([]pkces.OrderOption{}, pq.order...),
		inters:        append([]Interceptor{}, pq.inters...),
		predicates:    append([]predicate.PKCES{}, pq.predicates...),
		withRequestID: pq.withRequestID.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithRequestID tells the query-builder to eager-load the nodes that are connected to
// the "request_id" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PKCESQuery) WithRequestID(opts ...func(*RequestQuery)) *PKCESQuery {
	query := (&RequestClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withRequestID = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (pq *PKCESQuery) GroupBy(field string, fields ...string) *PKCESGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PKCESGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = pkces.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (pq *PKCESQuery) Select(fields ...string) *PKCESSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PKCESSelect{PKCESQuery: pq}
	sbuild.label = pkces.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PKCESSelect configured with the given aggregations.
func (pq *PKCESQuery) Aggregate(fns ...AggregateFunc) *PKCESSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PKCESQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !pkces.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PKCESQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PKCES, error) {
	var (
		nodes       = []*PKCES{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withRequestID != nil,
		}
	)
	if pq.withRequestID != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pkces.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PKCES).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PKCES{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withRequestID; query != nil {
		if err := pq.loadRequestID(ctx, query, nodes, nil,
			func(n *PKCES, e *Request) { n.Edges.RequestID = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PKCESQuery) loadRequestID(ctx context.Context, query *RequestQuery, nodes []*PKCES, init func(*PKCES), assign func(*PKCES, *Request)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PKCES)
	for i := range nodes {
		if nodes[i].request_pkce == nil {
			continue
		}
		fk := *nodes[i].request_pkce
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(request.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "request_pkce" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PKCESQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PKCESQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pkces.Table, pkces.Columns, sqlgraph.NewFieldSpec(pkces.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pkces.FieldID)
		for i := range fields {
			if fields[i] != pkces.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PKCESQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pkces.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = pkces.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PKCESGroupBy is the group-by builder for PKCES entities.
type PKCESGroupBy struct {
	selector
	build *PKCESQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PKCESGroupBy) Aggregate(fns ...AggregateFunc) *PKCESGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PKCESGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PKCESQuery, *PKCESGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PKCESGroupBy) sqlScan(ctx context.Context, root *PKCESQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PKCESSelect is the builder for selecting fields of PKCES entities.
type PKCESSelect struct {
	*PKCESQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PKCESSelect) Aggregate(fns ...AggregateFunc) *PKCESSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PKCESSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PKCESQuery, *PKCESSelect](ctx, ps.PKCESQuery, ps, ps.inters, v)
}

func (ps *PKCESSelect) sqlScan(ctx context.Context, root *PKCESQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
