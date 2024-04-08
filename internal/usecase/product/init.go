package product

func New(productRepo productResource) *productUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}
