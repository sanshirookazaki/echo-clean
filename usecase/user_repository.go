package usecase

import "github.com/sanshirookazaki/echo-clean/domain"

type UserRepository interface {
	GetUserID(username, password string) int
	GetTaskAll(userid int) domain.Tasks
}
