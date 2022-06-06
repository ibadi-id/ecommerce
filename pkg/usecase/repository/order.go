package repository

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
)

// Order is an interface of repository
type Order interface {
	Get(ctx context.Context, id *int) (*model.Order, error)
	List(ctx context.Context) ([]*model.Order, error)
	Create(ctx context.Context, input model.CreateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error)
	Update(ctx context.Context, input model.UpdateOrderInput, items []*model.CreateOrderItemInput) (*model.Order, error)
	GetProduct(ctx context.Context, id *int) (*model.Product, error)
}
