// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 100},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "phone", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "shipping_address", Type: field.TypeString},
		{Name: "amount", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_customers_orders",
				Columns:    []*schema.Column{OrdersColumns[5]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderItemsColumns holds the columns for the "order_items" table.
	OrderItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "order_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// OrderItemsTable holds the schema information for the "order_items" table.
	OrderItemsTable = &schema.Table{
		Name:       "order_items",
		Columns:    OrderItemsColumns,
		PrimaryKey: []*schema.Column{OrderItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_items_orders_items",
				Columns:    []*schema.Column{OrderItemsColumns[4]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "order_items_products_order_items",
				Columns:    []*schema.Column{OrderItemsColumns[5]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "descriptions", Type: field.TypeString},
		{Name: "sku", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt},
		{Name: "stock", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CustomersTable,
		OrdersTable,
		OrderItemsTable,
		ProductsTable,
	}
)

func init() {
	OrdersTable.ForeignKeys[0].RefTable = CustomersTable
	OrderItemsTable.ForeignKeys[0].RefTable = OrdersTable
	OrderItemsTable.ForeignKeys[1].RefTable = ProductsTable
}
