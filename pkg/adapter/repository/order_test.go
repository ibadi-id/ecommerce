package repository_test

import (
	"context"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/pkg/adapter/repository"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	"github/ibadi-id/ecommerce/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderRepository_List(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewOrderRepository(client)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc []*model.Order, err error)
		assert  func(t *testing.T, uc []*model.Order, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get order's list",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Order.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				cust, err := client.Customer.Create().SetName("test").SetEmail("test@gmail.com").SetPhone("085155555").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				orders := []struct {
					customer_id      int
					shipping_address string
					amount           int
				}{
					{
						customer_id:      cust.ID,
						shipping_address: "Jakarta Pusat",
						amount:           10000,
					},
					{
						customer_id:      cust.ID,
						shipping_address: "Jakarta Timur",
						amount:           20000,
					},
					{
						customer_id:      cust.ID,
						shipping_address: "Jakarta Barat",
						amount:           30000,
					},
				}
				bulk := make([]*ent.OrderCreate, len(orders))
				for i, u := range orders {
					bulk[i] = client.Order.Create().SetCustomerID(u.customer_id).
						SetAmount(u.amount).
						SetShippingAddress(u.shipping_address)
				}

				_, err = client.Order.
					CreateBulk(bulk...).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
			},
			act: func(ctx context.Context, t *testing.T) (us []*model.Order, err error) {
				return repo.List(ctx)
			},
			assert: func(t *testing.T, got []*model.Order, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 3, len(got))
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropOrder(t, client)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}

func TestOrderRepository_Get(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewOrderRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Order, err error)
		assert  func(t *testing.T, uc *model.Order, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get one prodcut",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Order.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				cust, err := client.Customer.Create().SetName("test").SetEmail("test@gmail.com").SetPhone("085155555").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				ord, err := client.Order.Create().SetCustomerID(cust.ID).
					SetAmount(10000).
					SetShippingAddress("Jakarta Pusat").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = ord.ID

			},
			act: func(ctx context.Context, t *testing.T) (us *model.Order, err error) {

				return repo.Get(ctx, &id)
			},
			assert: func(t *testing.T, got *model.Order, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropOrder(t, client)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}

func TestOrderRepository_Create(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewOrderRepository(client)

	type args struct {
		ctx context.Context
	}

	var orderID int
	var inputOrder ent.CreateOrderInput
	var items []*ent.CreateOrderItemInput
	var item ent.CreateOrderItemInput

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Order, err error)
		assert  func(t *testing.T, uc *model.Order, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should create one order",
			arrange: func(t *testing.T) {
				ctx := context.Background()

				_, err := client.Order.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				cust, err := client.Customer.Create().SetName("test").SetEmail("test1234@gmail.com").SetPhone("085155555").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				order, err := client.Product.Create().SetName("test").SetName("test").
					SetDescriptions("test order 1").
					SetSku("ABC123").SetPrice(10000).SetStock(10).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				inputOrder = ent.CreateOrderInput{
					ShippingAddress: "Jakarta Pusat",
					Amount:          10000,
					CustomerID:      &cust.ID,
				}

				item = ent.CreateOrderItemInput{
					Quantity:  1,
					ProductID: &order.ID,
				}
				items = []*ent.CreateOrderItemInput{
					&item,
				}

			},
			act: func(ctx context.Context, t *testing.T) (us *model.Order, err error) {
				order, err := repo.Create(ctx, inputOrder, items)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				orderID = order.ID
				return order, nil
			},
			assert: func(t *testing.T, got *model.Order, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, orderID)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropOrder(t, client)
				testutil.DropProduct(t, client)
				testutil.DropCustomer(t, client)
				testutil.DropOrderItem(t, client)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}

func TestOrderRepository_Update(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewOrderRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int
	var items []*ent.CreateOrderItemInput
	// var item ent.CreateOrderItemInput

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Order, err error)
		assert  func(t *testing.T, uc *model.Order, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should update one order",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Order.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				cust, err := client.Customer.Create().SetName("test").SetEmail("test@gmail.com").SetPhone("085155555").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				ord, err := client.Order.Create().SetCustomerID(cust.ID).
					SetAmount(10000).
					SetShippingAddress("Jakarta Pusat").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = ord.ID
			},
			act: func(ctx context.Context, t *testing.T) (us *model.Order, err error) {
				new_address := "Jakarta Barat"
				input := ent.UpdateOrderInput{
					ID:              id,
					ShippingAddress: &new_address,
				}
				cust, err := repo.Update(ctx, input, items)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = cust.ID
				return cust, nil
			},
			assert: func(t *testing.T, got *model.Order, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
				assert.Equal(t, got.ShippingAddress, "Jakarta Barat")
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropOrder(t, client)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}
