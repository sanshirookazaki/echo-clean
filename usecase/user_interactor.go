package usecase

import "github.com/sanshirookazaki/echo-clean/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) GetUserID(username, password string) int {
	userid := interactor.UserRepository.GetUserID(username, password)
	return userid
}

func (interactor *UserInteractor) GetTaskAll(userid int) domain.Tasks {
	tasks := interactor.UserRepository.GetTaskAll(userid)
	return tasks
}
