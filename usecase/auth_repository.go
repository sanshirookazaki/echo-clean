package usecase

type AuthRepository interface {
	GetUserID(username, password string) int
	GetPassword(username string) string
	GetUserName(password string) string
	UserUniqueCheck(username string) string
	UserAdd(username, password string)
}
