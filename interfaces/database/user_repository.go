package database

import (
	"github.com/sanshirookazaki/echo-clean/domain"
)

// SQLHandler .
type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) GetUserID(username, password string) int {
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
