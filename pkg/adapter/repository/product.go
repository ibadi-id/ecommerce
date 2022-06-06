package repository

import (
	"context"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/ent/product"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	usecaseRepository "github/ibadi-id/ecommerce/pkg/usecase/repository"
)

type productRepository struct {
	client *ent.Client
}

func NewProductRepository(client *ent.Client) usecaseRepository.Product {
	return &productRepository{client: client}
}

func (r *productRepository) Get(ctx context.Context, id *int) (*model.Product, error) {
	u, err := r.client.Product.Query().Where(product.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *productRepository) List(ctx context.Context) ([]*model.Product, error) {
	u, err := r.client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *productRepository) Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	u, err := r.client.Product.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (r *productRepository) Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	u, err := r.client.Product.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}
