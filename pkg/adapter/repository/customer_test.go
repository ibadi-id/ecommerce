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

func setup(t *testing.T) (client *ent.Client, teardown func()) {
	testutil.ReadConfig()
	c := testutil.NewDBClient(t)

	return c, func() {
		testutil.DropCustomer(t, c)
		testutil.DropOrder(t, c)
		testutil.DropOrderItem(t, c)
		testutil.DropProduct(t, c)
		defer c.Close()
	}
}

func TestCustomerRepository_List(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewCustomerRepository(client)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc []*model.Customer, err error)
		assert  func(t *testing.T, uc []*model.Customer, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get customer's list",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Customer.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				customers := []struct {
					name  string
					email string
					phone string
				}{{name: "test", email: "test@gmail.com", phone: "085155555"},
					{name: "test2", email: "test2@gmail.com", phone: "085155554"},
					{name: "test3", email: "test3@gmail.com", phone: "085155553"}}
				bulk := make([]*ent.CustomerCreate, len(customers))
				for i, u := range customers {
					bulk[i] = client.Customer.Create().SetName(u.name).SetEmail(u.email).SetPhone(u.phone)
				}

				_, err = client.Customer.
					CreateBulk(bulk...).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
			},
			act: func(ctx context.Context, t *testing.T) (us []*model.Customer, err error) {
				return repo.List(ctx)
			},
			assert: func(t *testing.T, got []*model.Customer, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 3, len(got))
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropCustomer(t, client)
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

func TestCustomerRepository_Get(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewCustomerRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Customer, err error)
		assert  func(t *testing.T, uc *model.Customer, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get one customer",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Customer.Delete().Exec(ctx)
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
				id = cust.ID

			},
			act: func(ctx context.Context, t *testing.T) (us *model.Customer, err error) {

				return repo.Get(ctx, &id)
			},
			assert: func(t *testing.T, got *model.Customer, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropCustomer(t, client)
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

func TestCustomerRepository_Create(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewCustomerRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Customer, err error)
		assert  func(t *testing.T, uc *model.Customer, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should create one customer",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Customer.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

			},
			act: func(ctx context.Context, t *testing.T) (us *model.Customer, err error) {
				input := ent.CreateCustomerInput{
					Name:  "test",
					Email: "test@gmail.com",
					Phone: "085155555",
				}
				cust, err := repo.Create(ctx, input)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = cust.ID
				return cust, nil
			},
			assert: func(t *testing.T, got *model.Customer, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropCustomer(t, client)
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

func TestCustomerRepository_Update(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewCustomerRepository(client)

	type args struct {
		ctx context.Context
	}

	var id int

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.Customer, err error)
		assert  func(t *testing.T, uc *model.Customer, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should update one customer",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.Customer.Delete().Exec(ctx)
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
				id = cust.ID
			},
			act: func(ctx context.Context, t *testing.T) (us *model.Customer, err error) {
				c := struct {
					name  string
					email string
					phone string
				}{name: "tester", email: "test@gmail.com", phone: "085155555"}
				input := ent.UpdateCustomerInput{
					ID:    id,
					Name:  &c.name,
					Email: &c.email,
					Phone: &c.phone,
				}
				cust, err := repo.Update(ctx, input)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				id = cust.ID
				return cust, nil
			},
			assert: func(t *testing.T, got *model.Customer, err error) {
				assert.Nil(t, err)
				assert.Equal(t, got.ID, id)
				assert.Equal(t, got.Name, "tester")
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropCustomer(t, client)
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
