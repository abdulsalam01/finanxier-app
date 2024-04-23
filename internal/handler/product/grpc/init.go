package grpc

func New(
	productUsecase productUsecase,
	baseApp baseAppInitializerResource,
) *Handler {
	return &Handler{
		productUsecase: productUsecase,
		baseApp:        baseApp,
	}
}
