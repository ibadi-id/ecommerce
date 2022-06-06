package usecase

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	"github/ibadi-id/ecommerce/pkg/usecase/repository"
)

type product struct {
	productRepository repository.Product
}

// Product of usecase.
type Product interface {
	Get(ctx context.Context, id *int) (*model.Product, error)
	List(ctx context.Context) ([]*model.Product, error)
	Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error)
	Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error)
}

// NewProductUsecase returns product usecase.
func NewProductUsecase(r repository.Product) Product {
	return &product{productRepository: r}
}

func (u *product) Get(ctx context.Context, id *int) (*model.Product, error) {
	return u.productRepository.Get(ctx, id)
}

func (u *product) Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	return u.productRepository.Create(ctx, input)
}

func (u *product) List(ctx context.Context) ([]*model.Product, error) {
	return u.productRepository.List(ctx)
}

func (u *product) Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	return u.productRepository.Update(ctx, input)
}
