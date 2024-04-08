package user

import "github.com/finanxier-app/cmd/app"

func New(
	userUsecase userUsecase,
	baseApp app.BaseAppInitializer,
) *Handler {
	return &Handler{
		userUsecase: userUsecase,
		baseApp:     baseApp,
	}
}
