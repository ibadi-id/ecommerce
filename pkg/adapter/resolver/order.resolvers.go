package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/graph/generated"
	"github/ibadi-id/ecommerce/pkg/adapter/handler"
	"github/ibadi-id/ecommerce/pkg/util/datetime"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, input ent.CreateOrderInput, items []*ent.CreateOrderItemInput) (*ent.Order, error) {
	u, err := r.controller.Order.Create(ctx, input, items)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, input ent.UpdateOrderInput, items []*ent.CreateOrderItemInput) (*ent.Order, error) {
	u, err := r.controller.Order.Update(ctx, input, items)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *orderResolver) CreatedAt(ctx context.Context, obj *ent.Order) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *orderResolver) UpdatedAt(ctx context.Context, obj *ent.Order) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *orderItemResolver) CreatedAt(ctx context.Context, obj *ent.OrderItem) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *orderItemResolver) UpdatedAt(ctx context.Context, obj *ent.OrderItem) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *queryResolver) Order(ctx context.Context, id *int) (*ent.Order, error) {
	u, err := r.controller.Order.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*ent.Order, error) {
	u, err := r.controller.Order.List(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// OrderItem returns generated.OrderItemResolver implementation.
func (r *Resolver) OrderItem() generated.OrderItemResolver { return &orderItemResolver{r} }

type orderResolver struct{ *Resolver }
type orderItemResolver struct{ *Resolver }
