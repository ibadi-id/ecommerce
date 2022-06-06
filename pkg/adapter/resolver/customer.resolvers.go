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

func (r *customerResolver) CreatedAt(ctx context.Context, obj *ent.Customer) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *customerResolver) UpdatedAt(ctx context.Context, obj *ent.Customer) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input ent.CreateCustomerInput) (*ent.Customer, error) {
	u, err := r.controller.Customer.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, input ent.UpdateCustomerInput) (*ent.Customer, error) {
	u, err := r.controller.Customer.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) Customer(ctx context.Context, id *int) (*ent.Customer, error) {
	u, err := r.controller.Customer.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*ent.Customer, error) {
	u, err := r.controller.Customer.List(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

type customerResolver struct{ *Resolver }
