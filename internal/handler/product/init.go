package product

import "github.com/finanxier-app/cmd/app"

func New(
	productUsecase productUsecase,
	baseApp app.BaseAppInitializer,
) *Handler {
	return &Handler{
		productUsecase: productUsecase,
		baseApp:        baseApp,
	}
}
