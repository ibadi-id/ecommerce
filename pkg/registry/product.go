package registry

import (
	"github/ibadi-id/ecommerce/pkg/adapter/controller"
	"github/ibadi-id/ecommerce/pkg/adapter/repository"
	"github/ibadi-id/ecommerce/pkg/usecase/usecase"
)

func (r *registry) NewProductController() controller.Product {
	repo := repository.NewProductRepository(r.client)
	u := usecase.NewProductUsecase(repo)

	return controller.NewProductController(u)
}
