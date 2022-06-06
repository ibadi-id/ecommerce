// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github/ibadi-id/ecommerce/ent/customer"
	"github/ibadi-id/ecommerce/ent/order"
	"github/ibadi-id/ecommerce/ent/orderitem"
	"github/ibadi-id/ecommerce/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetCustomerID sets the "customer_id" field.
func (ou *OrderUpdate) SetCustomerID(i int) *OrderUpdate {
	ou.mutation.SetCustomerID(i)
	return ou
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCustomerID(i *int) *OrderUpdate {
	if i != nil {
		ou.SetCustomerID(*i)
	}
	return ou
}

// ClearCustomerID clears the value of the "customer_id" field.
func (ou *OrderUpdate) ClearCustomerID() *OrderUpdate {
	ou.mutation.ClearCustomerID()
	return ou
}

// SetShippingAddress sets the "shipping_address" field.
func (ou *OrderUpdate) SetShippingAddress(s string) *OrderUpdate {
	ou.mutation.SetShippingAddress(s)
	return ou
}

// SetAmount sets the "amount" field.
func (ou *OrderUpdate) SetAmount(i int) *OrderUpdate {
	ou.mutation.ResetAmount()
	ou.mutation.SetAmount(i)
	return ou
}

// AddAmount adds i to the "amount" field.
func (ou *OrderUpdate) AddAmount(i int) *OrderUpdate {
	ou.mutation.AddAmount(i)
	return ou
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ou *OrderUpdate) SetCustomer(c *Customer) *OrderUpdate {
	return ou.SetCustomerID(c.ID)
}

// AddItemIDs adds the "items" edge to the OrderItem entity by IDs.
func (ou *OrderUpdate) AddItemIDs(ids ...int) *OrderUpdate {
	ou.mutation.AddItemIDs(ids...)
	return ou
}

// AddItems adds the "items" edges to the OrderItem entity.
func (ou *OrderUpdate) AddItems(o ...*OrderItem) *OrderUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddItemIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (ou *OrderUpdate) ClearCustomer() *OrderUpdate {
	ou.mutation.ClearCustomer()
	return ou
}

// ClearItems clears all "items" edges to the OrderItem entity.
func (ou *OrderUpdate) ClearItems() *OrderUpdate {
	ou.mutation.ClearItems()
	return ou
}

// RemoveItemIDs removes the "items" edge to OrderItem entities by IDs.
func (ou *OrderUpdate) RemoveItemIDs(ids ...int) *OrderUpdate {
	ou.mutation.RemoveItemIDs(ids...)
	return ou
}

// RemoveItems removes "items" edges to OrderItem entities.
func (ou *OrderUpdate) RemoveItems(o ...*OrderItem) *OrderUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.Amount(); ok {
		if err := order.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Order.amount": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.ShippingAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldShippingAddress,
		})
	}
	if value, ok := ou.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAmount,
		})
	}
	if value, ok := ou.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAmount,
		})
	}
	if ou.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ItemsTable,
			Columns: []string{order.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedItemsIDs(); len(nodes) > 0 && !ou.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ItemsTable,
			Columns: []string{order.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ItemsTable,
			Columns: []string{order.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetCustomerID sets the "customer_id" field.
func (ouo *OrderUpdateOne) SetCustomerID(i int) *OrderUpdateOne {
	ouo.mutation.SetCustomerID(i)
	return ouo
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCustomerID(i *int) *OrderUpdateOne {
	if i != nil {
		ouo.SetCustomerID(*i)
	}
	return ouo
}

// ClearCustomerID clears the value of the "customer_id" field.
func (ouo *OrderUpdateOne) ClearCustomerID() *OrderUpdateOne {
	ouo.mutation.ClearCustomerID()
	return ouo
}

// SetShippingAddress sets the "shipping_address" field.
func (ouo *OrderUpdateOne) SetShippingAddress(s string) *OrderUpdateOne {
	ouo.mutation.SetShippingAddress(s)
	return ouo
}

// SetAmount sets the "amount" field.
func (ouo *OrderUpdateOne) SetAmount(i int) *OrderUpdateOne {
	ouo.mutation.ResetAmount()
	ouo.mutation.SetAmount(i)
	return ouo
}

// AddAmount adds i to the "amount" field.
func (ouo *OrderUpdateOne) AddAmount(i int) *OrderUpdateOne {
	ouo.mutation.AddAmount(i)
	return ouo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ouo *OrderUpdateOne) SetCustomer(c *Customer) *OrderUpdateOne {
	return ouo.SetCustomerID(c.ID)
}

// AddItemIDs adds the "items" edge to the OrderItem entity by IDs.
func (ouo *OrderUpdateOne) AddItemIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.AddItemIDs(ids...)
	return ouo
}

// AddItems adds the "items" edges to the OrderItem entity.
func (ouo *OrderUpdateOne) AddItems(o ...*OrderItem) *OrderUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddItemIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (ouo *OrderUpdateOne) ClearCustomer() *OrderUpdateOne {
	ouo.mutation.ClearCustomer()
	return ouo
}

// ClearItems clears all "items" edges to the OrderItem entity.
func (ouo *OrderUpdateOne) ClearItems() *OrderUpdateOne {
	ouo.mutation.ClearItems()
	return ouo
}

// RemoveItemIDs removes the "items" edge to OrderItem entities by IDs.
func (ouo *OrderUpdateOne) RemoveItemIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.RemoveItemIDs(ids...)
	return ouo
}

// RemoveItems removes "items" edges to OrderItem entities.
func (ouo *OrderUpdateOne) RemoveItems(o ...*OrderItem) *OrderUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveItemIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.Amount(); ok {
		if err := order.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Order.amount": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.ShippingAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldShippingAddress,
		})
	}
	if value, ok := ouo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAmount,
		})
	}
	if value, ok := ouo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAmount,
		})
	}
	if ouo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.CustomerTable,
			Columns: []string{order.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ItemsTable,
			Columns: []string{order.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !ouo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ItemsTable,
			Columns: []string{order.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ItemsTable,
			Columns: []string{order.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}