package database

import (
	"strconv"

	"github.com/sanshirookazaki/echo-clean/domain"
)

// SQLHandler .
type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) GetUserID(username, password string) int {
	password, err := PasswordHash(password)
	if err != nil {
		panic(err.Error())
	}
	rows, err := repo.Query("SELECT userid FROM users WHERE username = \"" + username + "\" and password = \"" + password + "\"")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var userid int
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserID); err != nil {
			panic(err.Error())
		}
		userid = u.UserID
	}
	return userid
}

func (repo *UserRepository) GetUserNamePassword(userid int) (username, password string) {
	rows, err := repo.Query("SELECT username, password FROM users WHERE userid =" + strconv.Itoa(userid))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserName, &u.Password); err != nil {
			panic(err.Error())
		}
		username = u.UserName
		password = u.Password
	}
	return username, password
}
