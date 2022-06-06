package repository

import (
	"context"
	"github/ibadi-id/ecommerce/pkg/entity/model"
)

// Customer is an interface of repository
type Customer interface {
	Get(ctx context.Context, id *int) (*model.Customer, error)
	List(ctx context.Context) ([]*model.Customer, error)
	Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error)
	Update(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error)
}
