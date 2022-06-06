package registry

import (
	"github/ibadi-id/ecommerce/pkg/adapter/controller"
	"github/ibadi-id/ecommerce/pkg/adapter/repository"
	"github/ibadi-id/ecommerce/pkg/usecase/usecase"
)

func (r *registry) NewCustomerController() controller.Customer {
	repo := repository.NewCustomerRepository(r.client)
	u := usecase.NewCustomerUsecase(repo)

	return controller.NewCustomerController(u)
}
