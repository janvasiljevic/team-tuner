// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"jv/team-tone-tuner/model/course"
	"jv/team-tone-tuner/model/group"
	"jv/team-tone-tuner/model/grouprun"
	"jv/team-tone-tuner/model/predicate"
	"jv/team-tone-tuner/model/user"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GroupRunQuery is the builder for querying GroupRun entities.
type GroupRunQuery struct {
	config
	ctx           *QueryContext
	order         []grouprun.OrderOption
	inters        []Interceptor
	predicates    []predicate.GroupRun
	withCreatedBy *UserQuery
	withGroups    *GroupQuery
	withCourse    *CourseQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupRunQuery builder.
func (grq *GroupRunQuery) Where(ps ...predicate.GroupRun) *GroupRunQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit the number of records to be returned by this query.
func (grq *GroupRunQuery) Limit(limit int) *GroupRunQuery {
	grq.ctx.Limit = &limit
	return grq
}

// Offset to start from.
func (grq *GroupRunQuery) Offset(offset int) *GroupRunQuery {
	grq.ctx.Offset = &offset
	return grq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (grq *GroupRunQuery) Unique(unique bool) *GroupRunQuery {
	grq.ctx.Unique = &unique
	return grq
}

// Order specifies how the records should be ordered.
func (grq *GroupRunQuery) Order(o ...grouprun.OrderOption) *GroupRunQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// QueryCreatedBy chains the current query on the "created_by" edge.
func (grq *GroupRunQuery) QueryCreatedBy() *UserQuery {
	query := (&UserClient{config: grq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprun.Table, grouprun.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, grouprun.CreatedByTable, grouprun.CreatedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroups chains the current query on the "groups" edge.
func (grq *GroupRunQuery) QueryGroups() *GroupQuery {
	query := (&GroupClient{config: grq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprun.Table, grouprun.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, grouprun.GroupsTable, grouprun.GroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCourse chains the current query on the "course" edge.
func (grq *GroupRunQuery) QueryCourse() *CourseQuery {
	query := (&CourseClient{config: grq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprun.Table, grouprun.FieldID, selector),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grouprun.CourseTable, grouprun.CourseColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupRun entity from the query.
// Returns a *NotFoundError when no GroupRun was found.
func (grq *GroupRunQuery) First(ctx context.Context) (*GroupRun, error) {
	nodes, err := grq.Limit(1).All(setContextOp(ctx, grq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grouprun.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (grq *GroupRunQuery) FirstX(ctx context.Context) *GroupRun {
	node, err := grq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupRun ID from the query.
// Returns a *NotFoundError when no GroupRun ID was found.
func (grq *GroupRunQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = grq.Limit(1).IDs(setContextOp(ctx, grq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grouprun.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (grq *GroupRunQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := grq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupRun entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupRun entity is found.
// Returns a *NotFoundError when no GroupRun entities are found.
func (grq *GroupRunQuery) Only(ctx context.Context) (*GroupRun, error) {
	nodes, err := grq.Limit(2).All(setContextOp(ctx, grq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grouprun.Label}
	default:
		return nil, &NotSingularError{grouprun.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (grq *GroupRunQuery) OnlyX(ctx context.Context) *GroupRun {
	node, err := grq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupRun ID in the query.
// Returns a *NotSingularError when more than one GroupRun ID is found.
// Returns a *NotFoundError when no entities are found.
func (grq *GroupRunQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = grq.Limit(2).IDs(setContextOp(ctx, grq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grouprun.Label}
	default:
		err = &NotSingularError{grouprun.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (grq *GroupRunQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := grq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupRuns.
func (grq *GroupRunQuery) All(ctx context.Context) ([]*GroupRun, error) {
	ctx = setContextOp(ctx, grq.ctx, "All")
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupRun, *GroupRunQuery]()
	return withInterceptors[[]*GroupRun](ctx, grq, qr, grq.inters)
}

// AllX is like All, but panics if an error occurs.
func (grq *GroupRunQuery) AllX(ctx context.Context) []*GroupRun {
	nodes, err := grq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupRun IDs.
func (grq *GroupRunQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if grq.ctx.Unique == nil && grq.path != nil {
		grq.Unique(true)
	}
	ctx = setContextOp(ctx, grq.ctx, "IDs")
	if err = grq.Select(grouprun.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (grq *GroupRunQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := grq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (grq *GroupRunQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, grq.ctx, "Count")
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, grq, querierCount[*GroupRunQuery](), grq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (grq *GroupRunQuery) CountX(ctx context.Context) int {
	count, err := grq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (grq *GroupRunQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, grq.ctx, "Exist")
	switch _, err := grq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (grq *GroupRunQuery) ExistX(ctx context.Context) bool {
	exist, err := grq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupRunQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (grq *GroupRunQuery) Clone() *GroupRunQuery {
	if grq == nil {
		return nil
	}
	return &GroupRunQuery{
		config:        grq.config,
		ctx:           grq.ctx.Clone(),
		order:         append([]grouprun.OrderOption{}, grq.order...),
		inters:        append([]Interceptor{}, grq.inters...),
		predicates:    append([]predicate.GroupRun{}, grq.predicates...),
		withCreatedBy: grq.withCreatedBy.Clone(),
		withGroups:    grq.withGroups.Clone(),
		withCourse:    grq.withCourse.Clone(),
		// clone intermediate query.
		sql:  grq.sql.Clone(),
		path: grq.path,
	}
}

// WithCreatedBy tells the query-builder to eager-load the nodes that are connected to
// the "created_by" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GroupRunQuery) WithCreatedBy(opts ...func(*UserQuery)) *GroupRunQuery {
	query := (&UserClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	grq.withCreatedBy = query
	return grq
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GroupRunQuery) WithGroups(opts ...func(*GroupQuery)) *GroupRunQuery {
	query := (&GroupClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	grq.withGroups = query
	return grq
}

// WithCourse tells the query-builder to eager-load the nodes that are connected to
// the "course" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GroupRunQuery) WithCourse(opts ...func(*CourseQuery)) *GroupRunQuery {
	query := (&CourseClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	grq.withCourse = query
	return grq
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
//	client.GroupRun.Query().
//		GroupBy(grouprun.FieldCreatedAt).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (grq *GroupRunQuery) GroupBy(field string, fields ...string) *GroupRunGroupBy {
	grq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupRunGroupBy{build: grq}
	grbuild.flds = &grq.ctx.Fields
	grbuild.label = grouprun.Label
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
//	client.GroupRun.Query().
//		Select(grouprun.FieldCreatedAt).
//		Scan(ctx, &v)
func (grq *GroupRunQuery) Select(fields ...string) *GroupRunSelect {
	grq.ctx.Fields = append(grq.ctx.Fields, fields...)
	sbuild := &GroupRunSelect{GroupRunQuery: grq}
	sbuild.label = grouprun.Label
	sbuild.flds, sbuild.scan = &grq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupRunSelect configured with the given aggregations.
func (grq *GroupRunQuery) Aggregate(fns ...AggregateFunc) *GroupRunSelect {
	return grq.Select().Aggregate(fns...)
}

func (grq *GroupRunQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range grq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, grq); err != nil {
				return err
			}
		}
	}
	for _, f := range grq.ctx.Fields {
		if !grouprun.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if grq.path != nil {
		prev, err := grq.path(ctx)
		if err != nil {
			return err
		}
		grq.sql = prev
	}
	return nil
}

func (grq *GroupRunQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupRun, error) {
	var (
		nodes       = []*GroupRun{}
		withFKs     = grq.withFKs
		_spec       = grq.querySpec()
		loadedTypes = [3]bool{
			grq.withCreatedBy != nil,
			grq.withGroups != nil,
			grq.withCourse != nil,
		}
	)
	if grq.withCreatedBy != nil || grq.withCourse != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, grouprun.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupRun).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupRun{config: grq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, grq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := grq.withCreatedBy; query != nil {
		if err := grq.loadCreatedBy(ctx, query, nodes, nil,
			func(n *GroupRun, e *User) { n.Edges.CreatedBy = e }); err != nil {
			return nil, err
		}
	}
	if query := grq.withGroups; query != nil {
		if err := grq.loadGroups(ctx, query, nodes,
			func(n *GroupRun) { n.Edges.Groups = []*Group{} },
			func(n *GroupRun, e *Group) { n.Edges.Groups = append(n.Edges.Groups, e) }); err != nil {
			return nil, err
		}
	}
	if query := grq.withCourse; query != nil {
		if err := grq.loadCourse(ctx, query, nodes, nil,
			func(n *GroupRun, e *Course) { n.Edges.Course = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (grq *GroupRunQuery) loadCreatedBy(ctx context.Context, query *UserQuery, nodes []*GroupRun, init func(*GroupRun), assign func(*GroupRun, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*GroupRun)
	for i := range nodes {
		if nodes[i].group_run_created_by == nil {
			continue
		}
		fk := *nodes[i].group_run_created_by
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_run_created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (grq *GroupRunQuery) loadGroups(ctx context.Context, query *GroupQuery, nodes []*GroupRun, init func(*GroupRun), assign func(*GroupRun, *Group)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*GroupRun)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Group(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(grouprun.GroupsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.group_run_groups
		if fk == nil {
			return fmt.Errorf(`foreign-key "group_run_groups" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "group_run_groups" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (grq *GroupRunQuery) loadCourse(ctx context.Context, query *CourseQuery, nodes []*GroupRun, init func(*GroupRun), assign func(*GroupRun, *Course)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*GroupRun)
	for i := range nodes {
		if nodes[i].course_group_runs == nil {
			continue
		}
		fk := *nodes[i].course_group_runs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(course.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "course_group_runs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (grq *GroupRunQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := grq.querySpec()
	_spec.Node.Columns = grq.ctx.Fields
	if len(grq.ctx.Fields) > 0 {
		_spec.Unique = grq.ctx.Unique != nil && *grq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, grq.driver, _spec)
}

func (grq *GroupRunQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(grouprun.Table, grouprun.Columns, sqlgraph.NewFieldSpec(grouprun.FieldID, field.TypeUUID))
	_spec.From = grq.sql
	if unique := grq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if grq.path != nil {
		_spec.Unique = true
	}
	if fields := grq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grouprun.FieldID)
		for i := range fields {
			if fields[i] != grouprun.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := grq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := grq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := grq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := grq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (grq *GroupRunQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(grq.driver.Dialect())
	t1 := builder.Table(grouprun.Table)
	columns := grq.ctx.Fields
	if len(columns) == 0 {
		columns = grouprun.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if grq.sql != nil {
		selector = grq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if grq.ctx.Unique != nil && *grq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range grq.predicates {
		p(selector)
	}
	for _, p := range grq.order {
		p(selector)
	}
	if offset := grq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := grq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupRunGroupBy is the group-by builder for GroupRun entities.
type GroupRunGroupBy struct {
	selector
	build *GroupRunQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GroupRunGroupBy) Aggregate(fns ...AggregateFunc) *GroupRunGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the selector query and scans the result into the given value.
func (grgb *GroupRunGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, grgb.build.ctx, "GroupBy")
	if err := grgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupRunQuery, *GroupRunGroupBy](ctx, grgb.build, grgb, grgb.build.inters, v)
}

func (grgb *GroupRunGroupBy) sqlScan(ctx context.Context, root *GroupRunQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(grgb.fns))
	for _, fn := range grgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*grgb.flds)+len(grgb.fns))
		for _, f := range *grgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*grgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupRunSelect is the builder for selecting fields of GroupRun entities.
type GroupRunSelect struct {
	*GroupRunQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (grs *GroupRunSelect) Aggregate(fns ...AggregateFunc) *GroupRunSelect {
	grs.fns = append(grs.fns, fns...)
	return grs
}

// Scan applies the selector query and scans the result into the given value.
func (grs *GroupRunSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, grs.ctx, "Select")
	if err := grs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupRunQuery, *GroupRunSelect](ctx, grs.GroupRunQuery, grs, grs.inters, v)
}

func (grs *GroupRunSelect) sqlScan(ctx context.Context, root *GroupRunQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(grs.fns))
	for _, fn := range grs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*grs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
