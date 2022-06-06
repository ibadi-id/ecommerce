package registry

import (
	"github/ibadi-id/ecommerce/pkg/adapter/controller"
	"github/ibadi-id/ecommerce/pkg/adapter/repository"
	"github/ibadi-id/ecommerce/pkg/usecase/usecase"
)

func (r *registry) NewOrderController() controller.Order {
	repo := repository.NewOrderRepository(r.client)
	u := usecase.NewOrderUsecase(repo)

	return controller.NewOrderController(u)
}
