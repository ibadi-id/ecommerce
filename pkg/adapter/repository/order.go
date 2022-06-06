package repository

import (
	"context"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/ent/order"
	"github/ibadi-id/ecommerce/ent/product"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	usecaseRepository "github/ibadi-id/ecommerce/pkg/usecase/repository"
)

type orderRepository struct {
	client *ent.Client
}

func NewOrderRepository(client *ent.Client) usecaseRepository.Order {
	return &orderRepository{client: client}
}

func (r *orderRepository) Get(ctx context.Context, id *int) (*model.Order, error) {
	u, err := r.client.Order.Query().Where(order.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *orderRepository) GetProduct(ctx context.Context, id *int) (*model.Product, error) {
	u, err := r.client.Product.Query().Where(product.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *orderRepository) List(ctx context.Context) ([]*model.Order, error) {
	u, err := r.client.Order.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *orderRepository) Create(ctx context.Context, input model.CreateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error) {
	client := WithTransactionalMutation(ctx)
	if client == nil {
		client = r.client
	}
	u, err := client.Order.
		Create().
		SetInput(input).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	for _, v := range items {
		_, err := client.
			OrderItem.
			Create().
			SetProductID(*v.ProductID).
			SetQuantity(v.Quantity).
			SetOrderID(u.ID).
			Save(ctx)
		if err != nil {
			return nil, model.NewDBError(err)
		}
	}

	return u, nil
}

func (r *orderRepository) Update(ctx context.Context, input model.UpdateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error) {
	u, err := r.client.Order.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}
