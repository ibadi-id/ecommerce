package controller

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	"github/ibadi-id/ecommerce/pkg/usecase/usecase"
)

type product struct {
	productUsecase usecase.Product
}

// Product of interface
type Product interface {
	Get(ctx context.Context, id *int) (*model.Product, error)
	List(ctx context.Context) ([]*model.Product, error)
	Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error)
	Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error)
}

// NewProductController returns product controller
func NewProductController(uu usecase.Product) Product {
	return &product{productUsecase: uu}
}

func (u *product) Get(ctx context.Context, id *int) (*model.Product, error) {
	return u.productUsecase.Get(ctx, id)
}

func (u *product) List(ctx context.Context) ([]*model.Product, error) {
	return u.productUsecase.List(ctx)
}

func (u *product) Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	return u.productUsecase.Create(ctx, input)
}
func (u *product) Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	return u.productUsecase.Update(ctx, input)
}
