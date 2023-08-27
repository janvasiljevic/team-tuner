// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/predicate"
	"jv/team-tone-tuner/model/user"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BfiReportQuery is the builder for querying BfiReport entities.
type BfiReportQuery struct {
	config
	ctx         *QueryContext
	order       []bfireport.OrderOption
	inters      []Interceptor
	predicates  []predicate.BfiReport
	withStudent *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BfiReportQuery builder.
func (brq *BfiReportQuery) Where(ps ...predicate.BfiReport) *BfiReportQuery {
	brq.predicates = append(brq.predicates, ps...)
	return brq
}

// Limit the number of records to be returned by this query.
func (brq *BfiReportQuery) Limit(limit int) *BfiReportQuery {
	brq.ctx.Limit = &limit
	return brq
}

// Offset to start from.
func (brq *BfiReportQuery) Offset(offset int) *BfiReportQuery {
	brq.ctx.Offset = &offset
	return brq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (brq *BfiReportQuery) Unique(unique bool) *BfiReportQuery {
	brq.ctx.Unique = &unique
	return brq
}

// Order specifies how the records should be ordered.
func (brq *BfiReportQuery) Order(o ...bfireport.OrderOption) *BfiReportQuery {
	brq.order = append(brq.order, o...)
	return brq
}

// QueryStudent chains the current query on the "student" edge.
func (brq *BfiReportQuery) QueryStudent() *UserQuery {
	query := (&UserClient{config: brq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := brq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := brq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bfireport.Table, bfireport.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, bfireport.StudentTable, bfireport.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(brq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BfiReport entity from the query.
// Returns a *NotFoundError when no BfiReport was found.
func (brq *BfiReportQuery) First(ctx context.Context) (*BfiReport, error) {
	nodes, err := brq.Limit(1).All(setContextOp(ctx, brq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bfireport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (brq *BfiReportQuery) FirstX(ctx context.Context) *BfiReport {
	node, err := brq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BfiReport ID from the query.
// Returns a *NotFoundError when no BfiReport ID was found.
func (brq *BfiReportQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = brq.Limit(1).IDs(setContextOp(ctx, brq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bfireport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (brq *BfiReportQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := brq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BfiReport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BfiReport entity is found.
// Returns a *NotFoundError when no BfiReport entities are found.
func (brq *BfiReportQuery) Only(ctx context.Context) (*BfiReport, error) {
	nodes, err := brq.Limit(2).All(setContextOp(ctx, brq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bfireport.Label}
	default:
		return nil, &NotSingularError{bfireport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (brq *BfiReportQuery) OnlyX(ctx context.Context) *BfiReport {
	node, err := brq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BfiReport ID in the query.
// Returns a *NotSingularError when more than one BfiReport ID is found.
// Returns a *NotFoundError when no entities are found.
func (brq *BfiReportQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = brq.Limit(2).IDs(setContextOp(ctx, brq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bfireport.Label}
	default:
		err = &NotSingularError{bfireport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (brq *BfiReportQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := brq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BfiReports.
func (brq *BfiReportQuery) All(ctx context.Context) ([]*BfiReport, error) {
	ctx = setContextOp(ctx, brq.ctx, "All")
	if err := brq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BfiReport, *BfiReportQuery]()
	return withInterceptors[[]*BfiReport](ctx, brq, qr, brq.inters)
}

// AllX is like All, but panics if an error occurs.
func (brq *BfiReportQuery) AllX(ctx context.Context) []*BfiReport {
	nodes, err := brq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BfiReport IDs.
func (brq *BfiReportQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if brq.ctx.Unique == nil && brq.path != nil {
		brq.Unique(true)
	}
	ctx = setContextOp(ctx, brq.ctx, "IDs")
	if err = brq.Select(bfireport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (brq *BfiReportQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := brq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (brq *BfiReportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, brq.ctx, "Count")
	if err := brq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, brq, querierCount[*BfiReportQuery](), brq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (brq *BfiReportQuery) CountX(ctx context.Context) int {
	count, err := brq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (brq *BfiReportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, brq.ctx, "Exist")
	switch _, err := brq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (brq *BfiReportQuery) ExistX(ctx context.Context) bool {
	exist, err := brq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BfiReportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (brq *BfiReportQuery) Clone() *BfiReportQuery {
	if brq == nil {
		return nil
	}
	return &BfiReportQuery{
		config:      brq.config,
		ctx:         brq.ctx.Clone(),
		order:       append([]bfireport.OrderOption{}, brq.order...),
		inters:      append([]Interceptor{}, brq.inters...),
		predicates:  append([]predicate.BfiReport{}, brq.predicates...),
		withStudent: brq.withStudent.Clone(),
		// clone intermediate query.
		sql:  brq.sql.Clone(),
		path: brq.path,
	}
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (brq *BfiReportQuery) WithStudent(opts ...func(*UserQuery)) *BfiReportQuery {
	query := (&UserClient{config: brq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	brq.withStudent = query
	return brq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BfiReport.Query().
//		GroupBy(bfireport.FieldCreatedAt).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (brq *BfiReportQuery) GroupBy(field string, fields ...string) *BfiReportGroupBy {
	brq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BfiReportGroupBy{build: brq}
	grbuild.flds = &brq.ctx.Fields
	grbuild.label = bfireport.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.BfiReport.Query().
//		Select(bfireport.FieldCreatedAt).
//		Scan(ctx, &v)
func (brq *BfiReportQuery) Select(fields ...string) *BfiReportSelect {
	brq.ctx.Fields = append(brq.ctx.Fields, fields...)
	sbuild := &BfiReportSelect{BfiReportQuery: brq}
	sbuild.label = bfireport.Label
	sbuild.flds, sbuild.scan = &brq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BfiReportSelect configured with the given aggregations.
func (brq *BfiReportQuery) Aggregate(fns ...AggregateFunc) *BfiReportSelect {
	return brq.Select().Aggregate(fns...)
}

func (brq *BfiReportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range brq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, brq); err != nil {
				return err
			}
		}
	}
	for _, f := range brq.ctx.Fields {
		if !bfireport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if brq.path != nil {
		prev, err := brq.path(ctx)
		if err != nil {
			return err
		}
		brq.sql = prev
	}
	return nil
}

func (brq *BfiReportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BfiReport, error) {
	var (
		nodes       = []*BfiReport{}
		_spec       = brq.querySpec()
		loadedTypes = [1]bool{
			brq.withStudent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BfiReport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BfiReport{config: brq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, brq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := brq.withStudent; query != nil {
		if err := brq.loadStudent(ctx, query, nodes, nil,
			func(n *BfiReport, e *User) { n.Edges.Student = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (brq *BfiReportQuery) loadStudent(ctx context.Context, query *UserQuery, nodes []*BfiReport, init func(*BfiReport), assign func(*BfiReport, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*BfiReport)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bfireport.StudentColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.bfi_report_student
		if fk == nil {
			return fmt.Errorf(`foreign-key "bfi_report_student" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "bfi_report_student" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (brq *BfiReportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := brq.querySpec()
	_spec.Node.Columns = brq.ctx.Fields
	if len(brq.ctx.Fields) > 0 {
		_spec.Unique = brq.ctx.Unique != nil && *brq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, brq.driver, _spec)
}

func (brq *BfiReportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bfireport.Table, bfireport.Columns, sqlgraph.NewFieldSpec(bfireport.FieldID, field.TypeUUID))
	_spec.From = brq.sql
	if unique := brq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if brq.path != nil {
		_spec.Unique = true
	}
	if fields := brq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bfireport.FieldID)
		for i := range fields {
			if fields[i] != bfireport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := brq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := brq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := brq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := brq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (brq *BfiReportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(brq.driver.Dialect())
	t1 := builder.Table(bfireport.Table)
	columns := brq.ctx.Fields
	if len(columns) == 0 {
		columns = bfireport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if brq.sql != nil {
		selector = brq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if brq.ctx.Unique != nil && *brq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range brq.predicates {
		p(selector)
	}
	for _, p := range brq.order {
		p(selector)
	}
	if offset := brq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := brq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BfiReportGroupBy is the group-by builder for BfiReport entities.
type BfiReportGroupBy struct {
	selector
	build *BfiReportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (brgb *BfiReportGroupBy) Aggregate(fns ...AggregateFunc) *BfiReportGroupBy {
	brgb.fns = append(brgb.fns, fns...)
	return brgb
}

// Scan applies the selector query and scans the result into the given value.
func (brgb *BfiReportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, brgb.build.ctx, "GroupBy")
	if err := brgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BfiReportQuery, *BfiReportGroupBy](ctx, brgb.build, brgb, brgb.build.inters, v)
}

func (brgb *BfiReportGroupBy) sqlScan(ctx context.Context, root *BfiReportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(brgb.fns))
	for _, fn := range brgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*brgb.flds)+len(brgb.fns))
		for _, f := range *brgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*brgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := brgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BfiReportSelect is the builder for selecting fields of BfiReport entities.
type BfiReportSelect struct {
	*BfiReportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (brs *BfiReportSelect) Aggregate(fns ...AggregateFunc) *BfiReportSelect {
	brs.fns = append(brs.fns, fns...)
	return brs
}

// Scan applies the selector query and scans the result into the given value.
func (brs *BfiReportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, brs.ctx, "Select")
	if err := brs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BfiReportQuery, *BfiReportSelect](ctx, brs.BfiReportQuery, brs, brs.inters, v)
}

func (brs *BfiReportSelect) sqlScan(ctx context.Context, root *BfiReportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(brs.fns))
	for _, fn := range brs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*brs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := brs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
