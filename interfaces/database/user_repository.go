package database

import (
	"golang.org/x/xerrors"

	"github.com/sanshirookazaki/echo-clean/domain"
)

// SQLHandler .
type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) GetUserID(username, password string) int {
	password, err := PasswordHash(password)
	if err != nil {
		xerrors.Errorf("PasswordHash error: %w", err)
	}
	var u domain.User
	repo.Table("users").Select("userid").Where("username = \"" + username + "\" and password \"" + password + "\"").Scan(&u)
	return u.Userid
}

func (repo *UserRepository) GetUserNamePassword(userid int) (username, password string) {
	var u domain.User
	var uid string
	repo.Table("users").Select("username, password").Where("userid = ?", uid).Scan(&u)
	username = u.Username
	password = u.Password

	return username, password
}
