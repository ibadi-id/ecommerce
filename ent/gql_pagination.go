// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github/ibadi-id/ecommerce/ent/customer"
	"github/ibadi-id/ecommerce/ent/order"
	"github/ibadi-id/ecommerce/ent/orderitem"
	"github/ibadi-id/ecommerce/ent/product"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// CustomerEdge is the edge representation of Customer.
type CustomerEdge struct {
	Node   *Customer `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// CustomerConnection is the connection containing edges to Customer.
type CustomerConnection struct {
	Edges      []*CustomerEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

// CustomerPaginateOption enables pagination customization.
type CustomerPaginateOption func(*customerPager) error

// WithCustomerOrder configures pagination ordering.
func WithCustomerOrder(order *CustomerOrder) CustomerPaginateOption {
	if order == nil {
		order = DefaultCustomerOrder
	}
	o := *order
	return func(pager *customerPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCustomerOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCustomerFilter configures pagination filter.
func WithCustomerFilter(filter func(*CustomerQuery) (*CustomerQuery, error)) CustomerPaginateOption {
	return func(pager *customerPager) error {
		if filter == nil {
			return errors.New("CustomerQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type customerPager struct {
	order  *CustomerOrder
	filter func(*CustomerQuery) (*CustomerQuery, error)
}

func newCustomerPager(opts []CustomerPaginateOption) (*customerPager, error) {
	pager := &customerPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCustomerOrder
	}
	return pager, nil
}

func (p *customerPager) applyFilter(query *CustomerQuery) (*CustomerQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *customerPager) toCursor(c *Customer) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *customerPager) applyCursors(query *CustomerQuery, after, before *Cursor) *CustomerQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultCustomerOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *customerPager) applyOrder(query *CustomerQuery, reverse bool) *CustomerQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultCustomerOrder.Field {
		query = query.Order(direction.orderFunc(DefaultCustomerOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Customer.
func (c *CustomerQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CustomerPaginateOption,
) (*CustomerConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCustomerPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}

	conn := &CustomerConnection{Edges: []*CustomerEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := c.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := c.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	c = pager.applyCursors(c, after, before)
	c = pager.applyOrder(c, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		c = c.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := c.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Customer
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Customer {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Customer {
			return nodes[i]
		}
	}

	conn.Edges = make([]*CustomerEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &CustomerEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// CustomerOrderField defines the ordering field of Customer.
type CustomerOrderField struct {
	field    string
	toCursor func(*Customer) Cursor
}

// CustomerOrder defines the ordering of Customer.
type CustomerOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *CustomerOrderField `json:"field"`
}

// DefaultCustomerOrder is the default ordering of Customer.
var DefaultCustomerOrder = &CustomerOrder{
	Direction: OrderDirectionAsc,
	Field: &CustomerOrderField{
		field: customer.FieldID,
		toCursor: func(c *Customer) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Customer into CustomerEdge.
func (c *Customer) ToEdge(order *CustomerOrder) *CustomerEdge {
	if order == nil {
		order = DefaultCustomerOrder
	}
	return &CustomerEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// OrderEdge is the edge representation of Order.
type OrderEdge struct {
	Node   *Order `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// OrderConnection is the connection containing edges to Order.
type OrderConnection struct {
	Edges      []*OrderEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

// OrderPaginateOption enables pagination customization.
type OrderPaginateOption func(*orderPager) error

// WithOrderOrder configures pagination ordering.
func WithOrderOrder(order *OrderOrder) OrderPaginateOption {
	if order == nil {
		order = DefaultOrderOrder
	}
	o := *order
	return func(pager *orderPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOrderOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOrderFilter configures pagination filter.
func WithOrderFilter(filter func(*OrderQuery) (*OrderQuery, error)) OrderPaginateOption {
	return func(pager *orderPager) error {
		if filter == nil {
			return errors.New("OrderQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type orderPager struct {
	order  *OrderOrder
	filter func(*OrderQuery) (*OrderQuery, error)
}

func newOrderPager(opts []OrderPaginateOption) (*orderPager, error) {
	pager := &orderPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOrderOrder
	}
	return pager, nil
}

func (p *orderPager) applyFilter(query *OrderQuery) (*OrderQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *orderPager) toCursor(o *Order) Cursor {
	return p.order.Field.toCursor(o)
}

func (p *orderPager) applyCursors(query *OrderQuery, after, before *Cursor) *OrderQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultOrderOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *orderPager) applyOrder(query *OrderQuery, reverse bool) *OrderQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultOrderOrder.Field {
		query = query.Order(direction.orderFunc(DefaultOrderOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Order.
func (o *OrderQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OrderPaginateOption,
) (*OrderConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOrderPager(opts)
	if err != nil {
		return nil, err
	}

	if o, err = pager.applyFilter(o); err != nil {
		return nil, err
	}

	conn := &OrderConnection{Edges: []*OrderEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := o.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := o.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	o = pager.applyCursors(o, after, before)
	o = pager.applyOrder(o, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		o = o.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		o = o.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := o.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Order
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Order {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Order {
			return nodes[i]
		}
	}

	conn.Edges = make([]*OrderEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &OrderEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// OrderOrderField defines the ordering field of Order.
type OrderOrderField struct {
	field    string
	toCursor func(*Order) Cursor
}

// OrderOrder defines the ordering of Order.
type OrderOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *OrderOrderField `json:"field"`
}

// DefaultOrderOrder is the default ordering of Order.
var DefaultOrderOrder = &OrderOrder{
	Direction: OrderDirectionAsc,
	Field: &OrderOrderField{
		field: order.FieldID,
		toCursor: func(o *Order) Cursor {
			return Cursor{ID: o.ID}
		},
	},
}

// ToEdge converts Order into OrderEdge.
func (o *Order) ToEdge(order *OrderOrder) *OrderEdge {
	if order == nil {
		order = DefaultOrderOrder
	}
	return &OrderEdge{
		Node:   o,
		Cursor: order.Field.toCursor(o),
	}
}

// OrderItemEdge is the edge representation of OrderItem.
type OrderItemEdge struct {
	Node   *OrderItem `json:"node"`
	Cursor Cursor     `json:"cursor"`
}

// OrderItemConnection is the connection containing edges to OrderItem.
type OrderItemConnection struct {
	Edges      []*OrderItemEdge `json:"edges"`
	PageInfo   PageInfo         `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

// OrderItemPaginateOption enables pagination customization.
type OrderItemPaginateOption func(*orderItemPager) error

// WithOrderItemOrder configures pagination ordering.
func WithOrderItemOrder(order *OrderItemOrder) OrderItemPaginateOption {
	if order == nil {
		order = DefaultOrderItemOrder
	}
	o := *order
	return func(pager *orderItemPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOrderItemOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOrderItemFilter configures pagination filter.
func WithOrderItemFilter(filter func(*OrderItemQuery) (*OrderItemQuery, error)) OrderItemPaginateOption {
	return func(pager *orderItemPager) error {
		if filter == nil {
			return errors.New("OrderItemQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type orderItemPager struct {
	order  *OrderItemOrder
	filter func(*OrderItemQuery) (*OrderItemQuery, error)
}

func newOrderItemPager(opts []OrderItemPaginateOption) (*orderItemPager, error) {
	pager := &orderItemPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOrderItemOrder
	}
	return pager, nil
}

func (p *orderItemPager) applyFilter(query *OrderItemQuery) (*OrderItemQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *orderItemPager) toCursor(oi *OrderItem) Cursor {
	return p.order.Field.toCursor(oi)
}

func (p *orderItemPager) applyCursors(query *OrderItemQuery, after, before *Cursor) *OrderItemQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultOrderItemOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *orderItemPager) applyOrder(query *OrderItemQuery, reverse bool) *OrderItemQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultOrderItemOrder.Field {
		query = query.Order(direction.orderFunc(DefaultOrderItemOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to OrderItem.
func (oi *OrderItemQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OrderItemPaginateOption,
) (*OrderItemConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOrderItemPager(opts)
	if err != nil {
		return nil, err
	}

	if oi, err = pager.applyFilter(oi); err != nil {
		return nil, err
	}

	conn := &OrderItemConnection{Edges: []*OrderItemEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := oi.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := oi.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	oi = pager.applyCursors(oi, after, before)
	oi = pager.applyOrder(oi, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		oi = oi.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		oi = oi.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := oi.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *OrderItem
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *OrderItem {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *OrderItem {
			return nodes[i]
		}
	}

	conn.Edges = make([]*OrderItemEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &OrderItemEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// OrderItemOrderField defines the ordering field of OrderItem.
type OrderItemOrderField struct {
	field    string
	toCursor func(*OrderItem) Cursor
}

// OrderItemOrder defines the ordering of OrderItem.
type OrderItemOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     *OrderItemOrderField `json:"field"`
}

// DefaultOrderItemOrder is the default ordering of OrderItem.
var DefaultOrderItemOrder = &OrderItemOrder{
	Direction: OrderDirectionAsc,
	Field: &OrderItemOrderField{
		field: orderitem.FieldID,
		toCursor: func(oi *OrderItem) Cursor {
			return Cursor{ID: oi.ID}
		},
	},
}

// ToEdge converts OrderItem into OrderItemEdge.
func (oi *OrderItem) ToEdge(order *OrderItemOrder) *OrderItemEdge {
	if order == nil {
		order = DefaultOrderItemOrder
	}
	return &OrderItemEdge{
		Node:   oi,
		Cursor: order.Field.toCursor(oi),
	}
}

// ProductEdge is the edge representation of Product.
type ProductEdge struct {
	Node   *Product `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// ProductConnection is the connection containing edges to Product.
type ProductConnection struct {
	Edges      []*ProductEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

// ProductPaginateOption enables pagination customization.
type ProductPaginateOption func(*productPager) error

// WithProductOrder configures pagination ordering.
func WithProductOrder(order *ProductOrder) ProductPaginateOption {
	if order == nil {
		order = DefaultProductOrder
	}
	o := *order
	return func(pager *productPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultProductOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithProductFilter configures pagination filter.
func WithProductFilter(filter func(*ProductQuery) (*ProductQuery, error)) ProductPaginateOption {
	return func(pager *productPager) error {
		if filter == nil {
			return errors.New("ProductQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type productPager struct {
	order  *ProductOrder
	filter func(*ProductQuery) (*ProductQuery, error)
}

func newProductPager(opts []ProductPaginateOption) (*productPager, error) {
	pager := &productPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultProductOrder
	}
	return pager, nil
}

func (p *productPager) applyFilter(query *ProductQuery) (*ProductQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *productPager) toCursor(pr *Product) Cursor {
	return p.order.Field.toCursor(pr)
}

func (p *productPager) applyCursors(query *ProductQuery, after, before *Cursor) *ProductQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultProductOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *productPager) applyOrder(query *ProductQuery, reverse bool) *ProductQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultProductOrder.Field {
		query = query.Order(direction.orderFunc(DefaultProductOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Product.
func (pr *ProductQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...ProductPaginateOption,
) (*ProductConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newProductPager(opts)
	if err != nil {
		return nil, err
	}

	if pr, err = pager.applyFilter(pr); err != nil {
		return nil, err
	}

	conn := &ProductConnection{Edges: []*ProductEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := pr.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := pr.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	pr = pager.applyCursors(pr, after, before)
	pr = pager.applyOrder(pr, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		pr = pr.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		pr = pr.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := pr.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Product
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Product {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Product {
			return nodes[i]
		}
	}

	conn.Edges = make([]*ProductEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &ProductEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// ProductOrderField defines the ordering field of Product.
type ProductOrderField struct {
	field    string
	toCursor func(*Product) Cursor
}

// ProductOrder defines the ordering of Product.
type ProductOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *ProductOrderField `json:"field"`
}

// DefaultProductOrder is the default ordering of Product.
var DefaultProductOrder = &ProductOrder{
	Direction: OrderDirectionAsc,
	Field: &ProductOrderField{
		field: product.FieldID,
		toCursor: func(pr *Product) Cursor {
			return Cursor{ID: pr.ID}
		},
	},
}

// ToEdge converts Product into ProductEdge.
func (pr *Product) ToEdge(order *ProductOrder) *ProductEdge {
	if order == nil {
		order = DefaultProductOrder
	}
	return &ProductEdge{
		Node:   pr,
		Cursor: order.Field.toCursor(pr),
	}
}
