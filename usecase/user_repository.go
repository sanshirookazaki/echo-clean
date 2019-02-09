package usecase

type UserRepository interface {
	GetUserID(username, password string) int
}
