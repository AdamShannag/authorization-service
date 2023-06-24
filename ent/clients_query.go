// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/accesstokens"
	"authorization-service/ent/authorizecodes"
	"authorization-service/ent/clients"
	"authorization-service/ent/idsessions"
	"authorization-service/ent/pkces"
	"authorization-service/ent/predicate"
	"authorization-service/ent/refreshtokens"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ClientsQuery is the builder for querying Clients entities.
type ClientsQuery struct {
	config
	ctx               *QueryContext
	order             []clients.OrderOption
	inters            []Interceptor
	predicates        []predicate.Clients
	withAccessToken   *AccessTokensQuery
	withAuthorizeCode *AuthorizeCodesQuery
	withRefreshToken  *RefreshTokensQuery
	withIDSession     *IDSessionsQuery
	withPkce          *PKCESQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClientsQuery builder.
func (cq *ClientsQuery) Where(ps ...predicate.Clients) *ClientsQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ClientsQuery) Limit(limit int) *ClientsQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ClientsQuery) Offset(offset int) *ClientsQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ClientsQuery) Unique(unique bool) *ClientsQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ClientsQuery) Order(o ...clients.OrderOption) *ClientsQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryAccessToken chains the current query on the "access_token" edge.
func (cq *ClientsQuery) QueryAccessToken() *AccessTokensQuery {
	query := (&AccessTokensClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clients.Table, clients.FieldID, selector),
			sqlgraph.To(accesstokens.Table, accesstokens.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clients.AccessTokenTable, clients.AccessTokenColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthorizeCode chains the current query on the "authorize_code" edge.
func (cq *ClientsQuery) QueryAuthorizeCode() *AuthorizeCodesQuery {
	query := (&AuthorizeCodesClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clients.Table, clients.FieldID, selector),
			sqlgraph.To(authorizecodes.Table, authorizecodes.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clients.AuthorizeCodeTable, clients.AuthorizeCodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRefreshToken chains the current query on the "refresh_token" edge.
func (cq *ClientsQuery) QueryRefreshToken() *RefreshTokensQuery {
	query := (&RefreshTokensClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clients.Table, clients.FieldID, selector),
			sqlgraph.To(refreshtokens.Table, refreshtokens.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clients.RefreshTokenTable, clients.RefreshTokenColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIDSession chains the current query on the "id_session" edge.
func (cq *ClientsQuery) QueryIDSession() *IDSessionsQuery {
	query := (&IDSessionsClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clients.Table, clients.FieldID, selector),
			sqlgraph.To(idsessions.Table, idsessions.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clients.IDSessionTable, clients.IDSessionColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPkce chains the current query on the "pkce" edge.
func (cq *ClientsQuery) QueryPkce() *PKCESQuery {
	query := (&PKCESClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(clients.Table, clients.FieldID, selector),
			sqlgraph.To(pkces.Table, pkces.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, clients.PkceTable, clients.PkceColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Clients entity from the query.
// Returns a *NotFoundError when no Clients was found.
func (cq *ClientsQuery) First(ctx context.Context) (*Clients, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{clients.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ClientsQuery) FirstX(ctx context.Context) *Clients {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Clients ID from the query.
// Returns a *NotFoundError when no Clients ID was found.
func (cq *ClientsQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{clients.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ClientsQuery) FirstIDX(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Clients entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Clients entity is found.
// Returns a *NotFoundError when no Clients entities are found.
func (cq *ClientsQuery) Only(ctx context.Context) (*Clients, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{clients.Label}
	default:
		return nil, &NotSingularError{clients.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ClientsQuery) OnlyX(ctx context.Context) *Clients {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Clients ID in the query.
// Returns a *NotSingularError when more than one Clients ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ClientsQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{clients.Label}
	default:
		err = &NotSingularError{clients.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ClientsQuery) OnlyIDX(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClientsSlice.
func (cq *ClientsQuery) All(ctx context.Context) ([]*Clients, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Clients, *ClientsQuery]()
	return withInterceptors[[]*Clients](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ClientsQuery) AllX(ctx context.Context) []*Clients {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Clients IDs.
func (cq *ClientsQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(clients.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ClientsQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ClientsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ClientsQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ClientsQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ClientsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ClientsQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClientsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ClientsQuery) Clone() *ClientsQuery {
	if cq == nil {
		return nil
	}
	return &ClientsQuery{
		config:            cq.config,
		ctx:               cq.ctx.Clone(),
		order:             append([]clients.OrderOption{}, cq.order...),
		inters:            append([]Interceptor{}, cq.inters...),
		predicates:        append([]predicate.Clients{}, cq.predicates...),
		withAccessToken:   cq.withAccessToken.Clone(),
		withAuthorizeCode: cq.withAuthorizeCode.Clone(),
		withRefreshToken:  cq.withRefreshToken.Clone(),
		withIDSession:     cq.withIDSession.Clone(),
		withPkce:          cq.withPkce.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithAccessToken tells the query-builder to eager-load the nodes that are connected to
// the "access_token" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClientsQuery) WithAccessToken(opts ...func(*AccessTokensQuery)) *ClientsQuery {
	query := (&AccessTokensClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAccessToken = query
	return cq
}

// WithAuthorizeCode tells the query-builder to eager-load the nodes that are connected to
// the "authorize_code" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClientsQuery) WithAuthorizeCode(opts ...func(*AuthorizeCodesQuery)) *ClientsQuery {
	query := (&AuthorizeCodesClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAuthorizeCode = query
	return cq
}

// WithRefreshToken tells the query-builder to eager-load the nodes that are connected to
// the "refresh_token" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClientsQuery) WithRefreshToken(opts ...func(*RefreshTokensQuery)) *ClientsQuery {
	query := (&RefreshTokensClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withRefreshToken = query
	return cq
}

// WithIDSession tells the query-builder to eager-load the nodes that are connected to
// the "id_session" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClientsQuery) WithIDSession(opts ...func(*IDSessionsQuery)) *ClientsQuery {
	query := (&IDSessionsClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withIDSession = query
	return cq
}

// WithPkce tells the query-builder to eager-load the nodes that are connected to
// the "pkce" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClientsQuery) WithPkce(opts ...func(*PKCESQuery)) *ClientsQuery {
	query := (&PKCESClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withPkce = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClientSecret []byte `json:"client_secret,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Clients.Query().
//		GroupBy(clients.FieldClientSecret).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ClientsQuery) GroupBy(field string, fields ...string) *ClientsGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClientsGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = clients.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClientSecret []byte `json:"client_secret,omitempty"`
//	}
//
//	client.Clients.Query().
//		Select(clients.FieldClientSecret).
//		Scan(ctx, &v)
func (cq *ClientsQuery) Select(fields ...string) *ClientsSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ClientsSelect{ClientsQuery: cq}
	sbuild.label = clients.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClientsSelect configured with the given aggregations.
func (cq *ClientsQuery) Aggregate(fns ...AggregateFunc) *ClientsSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ClientsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !clients.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ClientsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Clients, error) {
	var (
		nodes       = []*Clients{}
		_spec       = cq.querySpec()
		loadedTypes = [5]bool{
			cq.withAccessToken != nil,
			cq.withAuthorizeCode != nil,
			cq.withRefreshToken != nil,
			cq.withIDSession != nil,
			cq.withPkce != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Clients).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Clients{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withAccessToken; query != nil {
		if err := cq.loadAccessToken(ctx, query, nodes,
			func(n *Clients) { n.Edges.AccessToken = []*AccessTokens{} },
			func(n *Clients, e *AccessTokens) { n.Edges.AccessToken = append(n.Edges.AccessToken, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withAuthorizeCode; query != nil {
		if err := cq.loadAuthorizeCode(ctx, query, nodes,
			func(n *Clients) { n.Edges.AuthorizeCode = []*AuthorizeCodes{} },
			func(n *Clients, e *AuthorizeCodes) { n.Edges.AuthorizeCode = append(n.Edges.AuthorizeCode, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withRefreshToken; query != nil {
		if err := cq.loadRefreshToken(ctx, query, nodes,
			func(n *Clients) { n.Edges.RefreshToken = []*RefreshTokens{} },
			func(n *Clients, e *RefreshTokens) { n.Edges.RefreshToken = append(n.Edges.RefreshToken, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withIDSession; query != nil {
		if err := cq.loadIDSession(ctx, query, nodes,
			func(n *Clients) { n.Edges.IDSession = []*IDSessions{} },
			func(n *Clients, e *IDSessions) { n.Edges.IDSession = append(n.Edges.IDSession, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withPkce; query != nil {
		if err := cq.loadPkce(ctx, query, nodes,
			func(n *Clients) { n.Edges.Pkce = []*PKCES{} },
			func(n *Clients, e *PKCES) { n.Edges.Pkce = append(n.Edges.Pkce, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ClientsQuery) loadAccessToken(ctx context.Context, query *AccessTokensQuery, nodes []*Clients, init func(*Clients), assign func(*Clients, *AccessTokens)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Clients)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AccessTokens(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(clients.AccessTokenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.clients_access_token
		if fk == nil {
			return fmt.Errorf(`foreign-key "clients_access_token" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "clients_access_token" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ClientsQuery) loadAuthorizeCode(ctx context.Context, query *AuthorizeCodesQuery, nodes []*Clients, init func(*Clients), assign func(*Clients, *AuthorizeCodes)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Clients)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AuthorizeCodes(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(clients.AuthorizeCodeColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.clients_authorize_code
		if fk == nil {
			return fmt.Errorf(`foreign-key "clients_authorize_code" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "clients_authorize_code" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ClientsQuery) loadRefreshToken(ctx context.Context, query *RefreshTokensQuery, nodes []*Clients, init func(*Clients), assign func(*Clients, *RefreshTokens)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Clients)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.RefreshTokens(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(clients.RefreshTokenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.clients_refresh_token
		if fk == nil {
			return fmt.Errorf(`foreign-key "clients_refresh_token" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "clients_refresh_token" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ClientsQuery) loadIDSession(ctx context.Context, query *IDSessionsQuery, nodes []*Clients, init func(*Clients), assign func(*Clients, *IDSessions)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Clients)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.IDSessions(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(clients.IDSessionColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.clients_id_session
		if fk == nil {
			return fmt.Errorf(`foreign-key "clients_id_session" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "clients_id_session" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ClientsQuery) loadPkce(ctx context.Context, query *PKCESQuery, nodes []*Clients, init func(*Clients), assign func(*Clients, *PKCES)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Clients)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PKCES(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(clients.PkceColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.clients_pkce
		if fk == nil {
			return fmt.Errorf(`foreign-key "clients_pkce" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "clients_pkce" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ClientsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ClientsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(clients.Table, clients.Columns, sqlgraph.NewFieldSpec(clients.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clients.FieldID)
		for i := range fields {
			if fields[i] != clients.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ClientsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(clients.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = clients.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClientsGroupBy is the group-by builder for Clients entities.
type ClientsGroupBy struct {
	selector
	build *ClientsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ClientsGroupBy) Aggregate(fns ...AggregateFunc) *ClientsGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ClientsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClientsQuery, *ClientsGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ClientsGroupBy) sqlScan(ctx context.Context, root *ClientsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ClientsSelect is the builder for selecting fields of Clients entities.
type ClientsSelect struct {
	*ClientsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ClientsSelect) Aggregate(fns ...AggregateFunc) *ClientsSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ClientsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClientsQuery, *ClientsSelect](ctx, cs.ClientsQuery, cs, cs.inters, v)
}

func (cs *ClientsSelect) sqlScan(ctx context.Context, root *ClientsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
