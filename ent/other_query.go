// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/bug/ent/other"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OtherQuery is the builder for querying Other entities.
type OtherQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Other
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OtherQuery builder.
func (oq *OtherQuery) Where(ps ...predicate.Other) *OtherQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *OtherQuery) Limit(limit int) *OtherQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *OtherQuery) Offset(offset int) *OtherQuery {
	oq.offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OtherQuery) Unique(unique bool) *OtherQuery {
	oq.unique = &unique
	return oq
}

// Order adds an order step to the query.
func (oq *OtherQuery) Order(o ...OrderFunc) *OtherQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// First returns the first Other entity from the query.
// Returns a *NotFoundError when no Other was found.
func (oq *OtherQuery) First(ctx context.Context) (*Other, error) {
	nodes, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{other.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OtherQuery) FirstX(ctx context.Context) *Other {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Other ID from the query.
// Returns a *NotFoundError when no Other ID was found.
func (oq *OtherQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{other.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OtherQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Other entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Other entity is found.
// Returns a *NotFoundError when no Other entities are found.
func (oq *OtherQuery) Only(ctx context.Context) (*Other, error) {
	nodes, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{other.Label}
	default:
		return nil, &NotSingularError{other.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OtherQuery) OnlyX(ctx context.Context) *Other {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Other ID in the query.
// Returns a *NotSingularError when more than one Other ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OtherQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{other.Label}
	default:
		err = &NotSingularError{other.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OtherQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Others.
func (oq *OtherQuery) All(ctx context.Context) ([]*Other, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Other, *OtherQuery]()
	return withInterceptors[[]*Other](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OtherQuery) AllX(ctx context.Context) []*Other {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Other IDs.
func (oq *OtherQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(other.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OtherQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OtherQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OtherQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OtherQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OtherQuery) Exist(ctx context.Context) (bool, error) {
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OtherQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OtherQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OtherQuery) Clone() *OtherQuery {
	if oq == nil {
		return nil
	}
	return &OtherQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		predicates: append([]predicate.Other{}, oq.predicates...),
		// clone intermediate query.
		sql:    oq.sql.Clone(),
		path:   oq.path,
		unique: oq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Other.Query().
//		GroupBy(other.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OtherQuery) GroupBy(field string, fields ...string) *OtherGroupBy {
	oq.fields = append([]string{field}, fields...)
	grbuild := &OtherGroupBy{build: oq}
	grbuild.flds = &oq.fields
	grbuild.label = other.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Other.Query().
//		Select(other.FieldName).
//		Scan(ctx, &v)
func (oq *OtherQuery) Select(fields ...string) *OtherSelect {
	oq.fields = append(oq.fields, fields...)
	sbuild := &OtherSelect{OtherQuery: oq}
	sbuild.label = other.Label
	sbuild.flds, sbuild.scan = &oq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OtherSelect configured with the given aggregations.
func (oq *OtherQuery) Aggregate(fns ...AggregateFunc) *OtherSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OtherQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.fields {
		if !other.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OtherQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Other, error) {
	var (
		nodes = []*Other{}
		_spec = oq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Other).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Other{config: oq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oq *OtherQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.fields
	if len(oq.fields) > 0 {
		_spec.Unique = oq.unique != nil && *oq.unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OtherQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   other.Table,
			Columns: other.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: other.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if unique := oq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, other.FieldID)
		for i := range fields {
			if fields[i] != other.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OtherQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(other.Table)
	columns := oq.fields
	if len(columns) == 0 {
		columns = other.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.unique != nil && *oq.unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OtherGroupBy is the group-by builder for Other entities.
type OtherGroupBy struct {
	selector
	build *OtherQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OtherGroupBy) Aggregate(fns ...AggregateFunc) *OtherGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OtherGroupBy) Scan(ctx context.Context, v any) error {
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OtherGroupBy](ctx, ogb, ogb.build.inters, v)
}

func (ogb *OtherGroupBy) sqlScan(ctx context.Context, v any) error {
	selector := ogb.sqlQuery(ctx)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *OtherGroupBy) sqlQuery(ctx context.Context) *sql.Selector {
	selector := ogb.build.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(*ogb.flds...)...)
}

// OtherSelect is the builder for selecting fields of Other entities.
type OtherSelect struct {
	*OtherQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OtherSelect) Aggregate(fns ...AggregateFunc) *OtherSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OtherSelect) Scan(ctx context.Context, v any) error {
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OtherSelect](ctx, os, os.inters, v)
}

func (os *OtherSelect) sqlScan(ctx context.Context, v any) error {
	selector := os.OtherQuery.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
