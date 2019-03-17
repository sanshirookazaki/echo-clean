package database

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"

	"github.com/sanshirookazaki/echo-clean/domain"
)

type AuthRepository struct {
	SQLHandler
}

func (repo *AuthRepository) GetUserID(username, password string) int {
	rows, err := repo.Query("SELECT userid FROM users WHERE username = \"" + username + "\" and password = \"" + password + "\"")
	if err != nil {
		xerrors.Errorf("Query error: %w", err)
	}
	defer rows.Close()
	var userid int
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserID); err != nil {
			xerrors.Errorf("user data scan error: %w", err)
		}
		userid = u.UserID
	}
	return userid
}

func (repo *AuthRepository) GetPassword(username string) string {
	rows, err := repo.Query("SELECT password FROM users WHERE username = \"" + username + "\"")
	if err != nil {
		xerrors.Errorf("Query error: %w", err)
	}
	defer rows.Close()
	var password string
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.Password); err != nil {
			xerrors.Errorf("user data scan error: %w", err)
		}
		password = u.Password
	}
	return password
}

func (repo *AuthRepository) GetUserName(password string) string {
	rows, err := repo.Query("SELECT username FROM users WHERE password = \"" + password + "\"")
	if err != nil {
		xerrors.Errorf("Query error: %w", err)
	}
	defer rows.Close()
	var username string
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserName); err != nil {
			xerrors.Errorf("user data scan error: %w", err)
		}
		username = u.UserName
	}
	return username
}

func (repo *AuthRepository) UserUniqueCheck(username string) string {
	rows, err := repo.Query("SELECT username FROM users WHERE username = \"" + username + "\"")
	if err != nil {
		xerrors.Errorf("Query error: %w", err)
	}
	defer rows.Close()
	var existUser string
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserName); err != nil {
			xerrors.Errorf("user data scan error: %w", err)
		}
		existUser = u.UserName
	}
	return existUser
}

func (repo *AuthRepository) UserAdd(username, password string) {
	password, err := PasswordHash(password)
	if err != nil {
		xerrors.Errorf("PasswordHash error: %w", err)
	}
	_, err = repo.Query("INSERT INTO users (username, password) VALUES ( \"" + username + "\" , \"" + password + "\" )")
	if err != nil {
		xerrors.Errorf("Query error: %w", err)
	}
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
