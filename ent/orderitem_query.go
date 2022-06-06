// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github/ibadi-id/ecommerce/ent/order"
	"github/ibadi-id/ecommerce/ent/orderitem"
	"github/ibadi-id/ecommerce/ent/predicate"
	"github/ibadi-id/ecommerce/ent/product"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderItemQuery is the builder for querying OrderItem entities.
type OrderItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderItem
	// eager-loading edges.
	withOrder   *OrderQuery
	withProduct *ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderItemQuery builder.
func (oiq *OrderItemQuery) Where(ps ...predicate.OrderItem) *OrderItemQuery {
	oiq.predicates = append(oiq.predicates, ps...)
	return oiq
}

// Limit adds a limit step to the query.
func (oiq *OrderItemQuery) Limit(limit int) *OrderItemQuery {
	oiq.limit = &limit
	return oiq
}

// Offset adds an offset step to the query.
func (oiq *OrderItemQuery) Offset(offset int) *OrderItemQuery {
	oiq.offset = &offset
	return oiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oiq *OrderItemQuery) Unique(unique bool) *OrderItemQuery {
	oiq.unique = &unique
	return oiq
}

// Order adds an order step to the query.
func (oiq *OrderItemQuery) Order(o ...OrderFunc) *OrderItemQuery {
	oiq.order = append(oiq.order, o...)
	return oiq
}

// QueryOrder chains the current query on the "order" edge.
func (oiq *OrderItemQuery) QueryOrder() *OrderQuery {
	query := &OrderQuery{config: oiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderitem.Table, orderitem.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderitem.OrderTable, orderitem.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (oiq *OrderItemQuery) QueryProduct() *ProductQuery {
	query := &ProductQuery{config: oiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderitem.Table, orderitem.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderitem.ProductTable, orderitem.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderItem entity from the query.
// Returns a *NotFoundError when no OrderItem was found.
func (oiq *OrderItemQuery) First(ctx context.Context) (*OrderItem, error) {
	nodes, err := oiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oiq *OrderItemQuery) FirstX(ctx context.Context) *OrderItem {
	node, err := oiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderItem ID from the query.
// Returns a *NotFoundError when no OrderItem ID was found.
func (oiq *OrderItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oiq *OrderItemQuery) FirstIDX(ctx context.Context) int {
	id, err := oiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderItem entity is found.
// Returns a *NotFoundError when no OrderItem entities are found.
func (oiq *OrderItemQuery) Only(ctx context.Context) (*OrderItem, error) {
	nodes, err := oiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderitem.Label}
	default:
		return nil, &NotSingularError{orderitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oiq *OrderItemQuery) OnlyX(ctx context.Context) *OrderItem {
	node, err := oiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderItem ID in the query.
// Returns a *NotSingularError when more than one OrderItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (oiq *OrderItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = &NotSingularError{orderitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oiq *OrderItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := oiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderItems.
func (oiq *OrderItemQuery) All(ctx context.Context) ([]*OrderItem, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oiq *OrderItemQuery) AllX(ctx context.Context) []*OrderItem {
	nodes, err := oiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderItem IDs.
func (oiq *OrderItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oiq.Select(orderitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oiq *OrderItemQuery) IDsX(ctx context.Context) []int {
	ids, err := oiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oiq *OrderItemQuery) Count(ctx context.Context) (int, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oiq *OrderItemQuery) CountX(ctx context.Context) int {
	count, err := oiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oiq *OrderItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oiq *OrderItemQuery) ExistX(ctx context.Context) bool {
	exist, err := oiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oiq *OrderItemQuery) Clone() *OrderItemQuery {
	if oiq == nil {
		return nil
	}
	return &OrderItemQuery{
		config:      oiq.config,
		limit:       oiq.limit,
		offset:      oiq.offset,
		order:       append([]OrderFunc{}, oiq.order...),
		predicates:  append([]predicate.OrderItem{}, oiq.predicates...),
		withOrder:   oiq.withOrder.Clone(),
		withProduct: oiq.withProduct.Clone(),
		// clone intermediate query.
		sql:    oiq.sql.Clone(),
		path:   oiq.path,
		unique: oiq.unique,
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrderItemQuery) WithOrder(opts ...func(*OrderQuery)) *OrderItemQuery {
	query := &OrderQuery{config: oiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oiq.withOrder = query
	return oiq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrderItemQuery) WithProduct(opts ...func(*ProductQuery)) *OrderItemQuery {
	query := &ProductQuery{config: oiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oiq.withProduct = query
	return oiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProductID int `json:"product_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderItem.Query().
//		GroupBy(orderitem.FieldProductID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oiq *OrderItemQuery) GroupBy(field string, fields ...string) *OrderItemGroupBy {
	group := &OrderItemGroupBy{config: oiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oiq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProductID int `json:"product_id,omitempty"`
//	}
//
//	client.OrderItem.Query().
//		Select(orderitem.FieldProductID).
//		Scan(ctx, &v)
//
func (oiq *OrderItemQuery) Select(fields ...string) *OrderItemSelect {
	oiq.fields = append(oiq.fields, fields...)
	return &OrderItemSelect{OrderItemQuery: oiq}
}

func (oiq *OrderItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oiq.fields {
		if !orderitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oiq.path != nil {
		prev, err := oiq.path(ctx)
		if err != nil {
			return err
		}
		oiq.sql = prev
	}
	return nil
}

func (oiq *OrderItemQuery) sqlAll(ctx context.Context) ([]*OrderItem, error) {
	var (
		nodes       = []*OrderItem{}
		_spec       = oiq.querySpec()
		loadedTypes = [2]bool{
			oiq.withOrder != nil,
			oiq.withProduct != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderItem{config: oiq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, oiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oiq.withOrder; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderItem)
		for i := range nodes {
			fk := nodes[i].OrderID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(order.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Order = n
			}
		}
	}

	if query := oiq.withProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderItem)
		for i := range nodes {
			fk := nodes[i].ProductID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(product.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Product = n
			}
		}
	}

	return nodes, nil
}

func (oiq *OrderItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oiq.querySpec()
	_spec.Node.Columns = oiq.fields
	if len(oiq.fields) > 0 {
		_spec.Unique = oiq.unique != nil && *oiq.unique
	}
	return sqlgraph.CountNodes(ctx, oiq.driver, _spec)
}

func (oiq *OrderItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oiq *OrderItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderitem.Table,
			Columns: orderitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderitem.FieldID,
			},
		},
		From:   oiq.sql,
		Unique: true,
	}
	if unique := oiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for i := range fields {
			if fields[i] != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oiq *OrderItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oiq.driver.Dialect())
	t1 := builder.Table(orderitem.Table)
	columns := oiq.fields
	if len(columns) == 0 {
		columns = orderitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oiq.sql != nil {
		selector = oiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oiq.unique != nil && *oiq.unique {
		selector.Distinct()
	}
	for _, p := range oiq.predicates {
		p(selector)
	}
	for _, p := range oiq.order {
		p(selector)
	}
	if offset := oiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderItemGroupBy is the group-by builder for OrderItem entities.
type OrderItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oigb *OrderItemGroupBy) Aggregate(fns ...AggregateFunc) *OrderItemGroupBy {
	oigb.fns = append(oigb.fns, fns...)
	return oigb
}

// Scan applies the group-by query and scans the result into the given value.
func (oigb *OrderItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oigb.path(ctx)
	if err != nil {
		return err
	}
	oigb.sql = query
	return oigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oigb *OrderItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oigb *OrderItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := oigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oigb *OrderItemGroupBy) StringX(ctx context.Context) string {
	v, err := oigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oigb *OrderItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := oigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oigb *OrderItemGroupBy) IntX(ctx context.Context) int {
	v, err := oigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oigb *OrderItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oigb *OrderItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oigb *OrderItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oigb *OrderItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := oigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oigb *OrderItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oigb.fields {
		if !orderitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oigb *OrderItemGroupBy) sqlQuery() *sql.Selector {
	selector := oigb.sql.Select()
	aggregation := make([]string, 0, len(oigb.fns))
	for _, fn := range oigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oigb.fields)+len(oigb.fns))
		for _, f := range oigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oigb.fields...)...)
}

// OrderItemSelect is the builder for selecting fields of OrderItem entities.
type OrderItemSelect struct {
	*OrderItemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ois *OrderItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ois.prepareQuery(ctx); err != nil {
		return err
	}
	ois.sql = ois.OrderItemQuery.sqlQuery(ctx)
	return ois.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ois *OrderItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ois.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ois *OrderItemSelect) StringsX(ctx context.Context) []string {
	v, err := ois.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ois.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ois *OrderItemSelect) StringX(ctx context.Context) string {
	v, err := ois.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ois *OrderItemSelect) IntsX(ctx context.Context) []int {
	v, err := ois.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ois.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ois *OrderItemSelect) IntX(ctx context.Context) int {
	v, err := ois.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ois *OrderItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ois.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ois.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ois *OrderItemSelect) Float64X(ctx context.Context) float64 {
	v, err := ois.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ois *OrderItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := ois.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ois.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ois *OrderItemSelect) BoolX(ctx context.Context) bool {
	v, err := ois.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ois *OrderItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ois.sql.Query()
	if err := ois.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}