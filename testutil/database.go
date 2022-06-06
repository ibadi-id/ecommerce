package testutil

import (
	"context"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/ent/enttest"
	"github/ibadi-id/ecommerce/pkg/infrastructure/datastore"
	"testing"

	"entgo.io/ent/dialect"
)

// NewDBClient loads database for test.
func NewDBClient(t *testing.T) *ent.Client {
	d := datastore.New()
	return enttest.Open(t, dialect.MySQL, d)
}

// DropAll drops all data from database.
func DropAll(t *testing.T, client *ent.Client) {
	t.Log("drop data from database")
	DropOrder(t, client)
	DropProduct(t, client)
	DropOrderItem(t, client)
	DropCustomer(t, client)
	// DropTodo(t, client)
}

// DropCustomer drops data from customers.
func DropCustomer(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.Customer.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// DropProduct drops data from products.
func DropProduct(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.Product.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// DropOrder drops data from orders.
func DropOrder(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.Order.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// DropOrderItem drops data from order_items.
func DropOrderItem(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.OrderItem.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// DropTodo drops data from todos.
// func DropTodo(t *testing.T, client *ent.Client) {
// 	ctx := context.Background()
// 	_, err := client.Todo.Delete().Exec(ctx)

// 	if err != nil {
// 		t.Error(err)
// 		t.FailNow()
// 	}
// }
