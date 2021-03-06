// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"github/ibadi-id/ecommerce/ent/order"
	"github/ibadi-id/ecommerce/ent/orderitem"
	"github/ibadi-id/ecommerce/ent/product"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int `json:"product_id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int `json:"order_id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges OrderItemEdges `json:"edges"`
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Order == nil {
			// The edge order was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[1] {
		if e.Product == nil {
			// The edge product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID, orderitem.FieldProductID, orderitem.FieldOrderID, orderitem.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case orderitem.FieldCreatedAt, orderitem.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		case orderitem.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				oi.ProductID = int(value.Int64)
			}
		case orderitem.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oi.OrderID = int(value.Int64)
			}
		case orderitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				oi.Quantity = int(value.Int64)
			}
		case orderitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oi.CreatedAt = value.Time
			}
		case orderitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oi.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryOrder queries the "order" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrder() *OrderQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrder(oi)
}

// QueryProduct queries the "product" edge of the OrderItem entity.
func (oi *OrderItem) QueryProduct() *ProductQuery {
	return (&OrderItemClient{config: oi.config}).QueryProduct(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return (&OrderItemClient{config: oi.config}).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderItem is not a transactional entity")
	}
	oi.config.driver = tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v", oi.ID))
	builder.WriteString(", product_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ProductID))
	builder.WriteString(", order_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderID))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", oi.Quantity))
	builder.WriteString(", created_at=")
	builder.WriteString(oi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(oi.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem

func (oi OrderItems) config(cfg config) {
	for _i := range oi {
		oi[_i].config = cfg
	}
}
