package usecase

type AuthInteractor struct {
	AuthRepository AuthRepository
}

func (interactor *AuthInteractor) GetPassword(username string) string {
	password := interactor.AuthRepository.GetPassword(username)
	return password
}

func (interactor *AuthInteractor) GetUserName(password string) string {
	username := interactor.AuthRepository.GetUserName(password)
	return username
}

func (interactor *AuthInteractor) UserUniqueCheck(username string) string {
	existuser := interactor.AuthRepository.UserUniqueCheck(username)
	return existuser
}

func (interactor *AuthInteractor) UserAdd(username, password string) {
	interactor.AuthRepository.UserAdd(username, password)
}
