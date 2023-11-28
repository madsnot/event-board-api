package usecase

type Usecase struct {
	Auth  *AuthUsecase
	User  *UserUsecase
	Event *EventUsecase
}

func NewUsecase(auth *AuthUsecase, user *UserUsecase, event *EventUsecase) Usecase {
	return Usecase{
		Auth:  auth,
		User:  user,
		Event: event,
	}
}
