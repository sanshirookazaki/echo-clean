package usecase

type UserRepository interface {
	GetUserID(username, password string) int
	GetUserNamePassword(userid int) (username, password string)
}
