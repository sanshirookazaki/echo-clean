package usecase

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) GetUserID(username, password string) int {
	userid := interactor.UserRepository.GetUserID(username, password)
	return userid
}

func (interactor *UserInteractor) GetUserNamePassword(userid int) (username, password string) {
	username, password = interactor.UserRepository.GetUserNamePassword(userid)
	return username, password
}
