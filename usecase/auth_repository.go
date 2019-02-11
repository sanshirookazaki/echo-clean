package usecase

type AuthRepository interface {
	GetPassword(username string) string
	GetUserName(password string) string
	UserUniqueCheck(username string) string
	UserAdd(username, password string)
}
