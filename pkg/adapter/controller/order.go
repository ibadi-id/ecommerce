package controller

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	"github/ibadi-id/ecommerce/pkg/usecase/usecase"
)

type order struct {
	orderUsecase usecase.Order
}

// Order of interface
type Order interface {
	Get(ctx context.Context, id *int) (*model.Order, error)
	List(ctx context.Context) ([]*model.Order, error)
	Create(ctx context.Context, input model.CreateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error)
	Update(ctx context.Context, input model.UpdateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error)
}

// NewOrderController returns order controller
func NewOrderController(uu usecase.Order) Order {
	return &order{orderUsecase: uu}
}

func (u *order) Get(ctx context.Context, id *int) (*model.Order, error) {
	return u.orderUsecase.Get(ctx, id)
}

func (u *order) List(ctx context.Context) ([]*model.Order, error) {
	return u.orderUsecase.List(ctx)
}

func (u *order) Create(ctx context.Context, input model.CreateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error) {
	return u.orderUsecase.Create(ctx, input, items)
}
func (u *order) Update(ctx context.Context, input model.UpdateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error) {
	return u.orderUsecase.Update(ctx, input, items)
}
