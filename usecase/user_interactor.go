package usecase

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) GetUserID(username, password string) int {
	userid := interactor.UserRepository.GetUserID(username, password)
	return
}
