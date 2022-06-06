package registry

import (
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/pkg/adapter/controller"
)

type registry struct {
	client *ent.Client
}

// Registry is an interface of registry
type Registry interface {
	NewController() controller.Controller
}

// New registers entire controller with dependencies
func New(client *ent.Client) Registry {
	return &registry{
		client: client,
	}
}

// NewController generates controllers
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		Customer: r.NewCustomerController(),
		Product:  r.NewProductController(),
		Order:    r.NewOrderController(),
	}
}
