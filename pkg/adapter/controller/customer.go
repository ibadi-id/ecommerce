package controller

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
	"github/ibadi-id/ecommerce/pkg/usecase/usecase"
)

type customer struct {
	customerUsecase usecase.Customer
}

// Customer of interface
type Customer interface {
	Get(ctx context.Context, id *int) (*model.Customer, error)
	List(ctx context.Context) ([]*model.Customer, error)
	Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error)
	Update(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error)
}

// NewCustomerController returns customer controller
func NewCustomerController(uu usecase.Customer) Customer {
	return &customer{customerUsecase: uu}
}

func (u *customer) Get(ctx context.Context, id *int) (*model.Customer, error) {
	return u.customerUsecase.Get(ctx, id)
}

func (u *customer) List(ctx context.Context) ([]*model.Customer, error) {
	return u.customerUsecase.List(ctx)
}

func (u *customer) Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	return u.customerUsecase.Create(ctx, input)
}
func (u *customer) Update(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error) {
	return u.customerUsecase.Update(ctx, input)
}
