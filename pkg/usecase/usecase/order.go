package usecase

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	"github/ibadi-id/ecommerce/pkg/usecase/repository"
)

type order struct {
	orderRepository repository.Order
}

// Order of usecase.
type Order interface {
	Get(ctx context.Context, id *int) (*model.Order, error)
	List(ctx context.Context) ([]*model.Order, error)
	Create(ctx context.Context, input model.CreateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error)
	Update(ctx context.Context, input model.UpdateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error)
}

// NewOrderUsecase returns order usecase.
func NewOrderUsecase(r repository.Order) Order {
	return &order{orderRepository: r}
}

func (u *order) Get(ctx context.Context, id *int) (*model.Order, error) {
	return u.orderRepository.Get(ctx, id)
}

func (u *order) Create(ctx context.Context, input model.CreateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error) {
	var total_amount int
	for _, v := range items {
		p, err := u.orderRepository.GetProduct(ctx, v.ProductID)
		if err != nil {
			return nil, model.NewDBError(err)
		}
		total_amount += p.Price * v.Quantity
	}

	input.Amount = total_amount
	return u.orderRepository.Create(ctx, input, items)
}

func (u *order) List(ctx context.Context) ([]*model.Order, error) {
	return u.orderRepository.List(ctx)
}

func (u *order) Update(ctx context.Context, input model.UpdateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error) {
	return u.orderRepository.Update(ctx, input, items)
}
