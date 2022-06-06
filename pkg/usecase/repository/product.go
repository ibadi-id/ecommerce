package repository

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
)

// Product is an interface of repository
type Product interface {
	Get(ctx context.Context, id *int) (*model.Product, error)
	List(ctx context.Context) ([]*model.Product, error)
	Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error)
	Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error)
}
