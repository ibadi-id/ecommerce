package usecase

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	"github/ibadi-id/ecommerce/pkg/usecase/repository"
)

type customer struct {
	customerRepository repository.Customer
}

// Customer of usecase.
type Customer interface {
	Get(ctx context.Context, id *int) (*model.Customer, error)
	List(ctx context.Context) ([]*model.Customer, error)
	Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error)
	Update(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error)
}

// NewCustomerUsecase returns customer usecase.
func NewCustomerUsecase(r repository.Customer) Customer {
	return &customer{customerRepository: r}
}

func (u *customer) Get(ctx context.Context, id *int) (*model.Customer, error) {
	return u.customerRepository.Get(ctx, id)
}

func (u *customer) Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	return u.customerRepository.Create(ctx, input)
}

func (u *customer) List(ctx context.Context) ([]*model.Customer, error) {
	return u.customerRepository.List(ctx)
}

func (u *customer) Update(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error) {
	return u.customerRepository.Update(ctx, input)
}
