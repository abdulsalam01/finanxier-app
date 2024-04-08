package user

func New(userRepo userResource) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
