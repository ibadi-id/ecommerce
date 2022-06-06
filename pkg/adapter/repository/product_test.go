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

func TestProductRepository_List(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewProductRepository(client)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc []*model.Product, err error)
		assert  func(t *testing.T, uc []*model.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get product's list",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Product.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				products := []struct {
					name         string
					descriptions string
					sku          string
					price        int
					stock        int
				}{{name: "test", descriptions: "test produk 1", sku: "ABC123", price: 1000, stock: 10},
					{name: "test2", descriptions: "test produk 2", sku: "ABC124", price: 1000, stock: 10},
					{name: "test3", descriptions: "test produk 3", sku: "ABC125", price: 1000, stock: 10}}
				bulk := make([]*ent.ProductCreate, len(products))
				for i, u := range products {
					bulk[i] = client.Product.Create().SetName(u.name).
						SetDescriptions(u.descriptions).
						SetSku(u.sku).SetPrice(u.price).SetStock(u.stock)
				}

				_, err = client.Product.
					CreateBulk(bulk...).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
			},
			act: func(ctx context.Context, t *testing.T) (us []*model.Product, err error) {
				return repo.List(ctx)
			},
			assert: func(t *testing.T, got []*model.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 3, len(got))
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropProduct(t, client)
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

func TestProductRepository_Get(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewProductRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Product, err error)
		assert  func(t *testing.T, uc *model.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get one prodcut",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Product.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				cust, err := client.Product.Create().SetName("test").SetName("test").
					SetDescriptions("test product 1").
					SetSku("ABC123").SetPrice(10000).SetStock(10).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = cust.ID

			},
			act: func(ctx context.Context, t *testing.T) (us *model.Product, err error) {

				return repo.Get(ctx, &id)
			},
			assert: func(t *testing.T, got *model.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropProduct(t, client)
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

func TestProductRepository_Create(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewProductRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Product, err error)
		assert  func(t *testing.T, uc *model.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should create one product",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Product.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

			},
			act: func(ctx context.Context, t *testing.T) (us *model.Product, err error) {
				input := ent.CreateProductInput{
					Name:         "test",
					Descriptions: "test product 1",
					Sku:          "ABC123",
					Price:        10000,
					Stock:        10,
				}
				cust, err := repo.Create(ctx, input)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = cust.ID
				return cust, nil
			},
			assert: func(t *testing.T, got *model.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropProduct(t, client)
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

func TestProductRepository_Update(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewProductRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Product, err error)
		assert  func(t *testing.T, uc *model.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should update one product",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Product.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				product, err := client.Product.Create().SetName("test").SetName("test").
					SetDescriptions("test product 1").
					SetSku("ABC123").SetPrice(10000).SetStock(10).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = product.ID
			},
			act: func(ctx context.Context, t *testing.T) (us *model.Product, err error) {
				c := struct {
					name         string
					descriptions string
					sku          string
					price        int
					stock        int
				}{name: "tester", descriptions: "test produk 1", sku: "ABC123", price: 1000, stock: 10}
				input := ent.UpdateProductInput{
					ID:           id,
					Name:         &c.name,
					Descriptions: &c.descriptions,
					Sku:          &c.sku,
					Price:        &c.price,
					Stock:        &c.stock,
				}
				cust, err := repo.Update(ctx, input)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = cust.ID
				return cust, nil
			},
			assert: func(t *testing.T, got *model.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
				assert.Equal(t, got.Name, "tester")
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropProduct(t, client)
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
