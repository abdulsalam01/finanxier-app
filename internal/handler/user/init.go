package user

func New(
	userUsecase userUsecase,
	baseApp baseAppInitializerResource,
) *Handler {
	return &Handler{
		userUsecase: userUsecase,
		baseApp:     baseApp,
	}
}
