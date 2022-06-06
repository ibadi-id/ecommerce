package repository

import (
	"context"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/ent/customer"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	usecaseRepository "github/ibadi-id/ecommerce/pkg/usecase/repository"
)

type customerRepository struct {
	client *ent.Client
}

func NewCustomerRepository(client *ent.Client) usecaseRepository.Customer {
	return &customerRepository{client: client}
}

func (r *customerRepository) Get(ctx context.Context, id *int) (*model.Customer, error) {
	u, err := r.client.Customer.Query().Where(customer.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *customerRepository) List(ctx context.Context) ([]*model.Customer, error) {
	u, err := r.client.Customer.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *customerRepository) Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	u, err := r.client.Customer.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (r *customerRepository) Update(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error) {
	u, err := r.client.Customer.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}
