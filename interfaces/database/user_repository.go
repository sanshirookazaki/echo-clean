package database

import (
	"strconv"

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

func (repo *UserRepository) GetUserNamePassword(userid int) (username, password string) {
	rows, err := repo.Query("SELECT username, password FROM users WHERE userid =" + strconv.Itoa(userid))
	if err != nil {
		xerrors.Errorf("Query error: %w", err)
	}
	defer rows.Close()
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserName, &u.Password); err != nil {
			xerrors.Errorf("user data scan error: %w", err)
		}
		username = u.UserName
		password = u.Password
	}
	return username, password
}
