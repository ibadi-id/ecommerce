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

func (r *mutationResolver) CreateProduct(ctx context.Context, input ent.CreateProductInput) (*ent.Product, error) {
	u, err := r.controller.Product.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input ent.UpdateProductInput) (*ent.Product, error) {
	u, err := r.controller.Product.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *productResolver) CreatedAt(ctx context.Context, obj *ent.Product) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *productResolver) UpdatedAt(ctx context.Context, obj *ent.Product) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *queryResolver) Product(ctx context.Context, id *int) (*ent.Product, error) {
	u, err := r.controller.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*ent.Product, error) {
	u, err := r.controller.Product.List(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
