package database

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"

	"github.com/sanshirookazaki/echo-clean/domain"
)

type AuthRepository struct {
	SQLHandler
}

type Pass struct {
	Password string
	Userid   int
}

func (repo *AuthRepository) GetUserID(username, password string) int {
	var u Pass
	repo.Table("users").Select("userid").Where("username = ? and password = ?", username, password).Scan(&u)
	userid := u.Userid
	return userid
}

func (repo *AuthRepository) GetPassword(username string) string {
	//var u domain.User
	var u Pass
	repo.Table("users").Select("password").Where("username = ?", username).Scan(&u)
	password := u.Password
	return password
}

func (repo *AuthRepository) GetUserName(password string) string {
	var u domain.User
	repo.Table("users").Select("username").Where("password = \"" + password + "\"").Scan(&u)
	username := u.Username
	return username
}

func (repo *AuthRepository) UserUniqueCheck(username string) string {
	var u domain.User
	repo.Table("users").Select("username").Where("username = \"" + username + "\"").Scan(&u)
	existUser := u.Username
	return existUser
}

func (repo *AuthRepository) UserAdd(username, password string) {
	password, err := PasswordHash(password)
	if err != nil {
		xerrors.Errorf("PasswordHash error: %w", err)
	}
	//_, err = repo.Query("INSERT INTO users (username, password) VALUES ( \"" + username + "\" , \"" + password + "\" )")

	u := domain.User{Username: username, Password: password}
	repo.Table("users").Create(&u)
	repo.Save(&u)
}

func PasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func PasswordVerify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
