package database

import "github.com/sanshirookazaki/echo-clean/domain"

type AuthRepository struct {
	SQLHandler
}

func (repo *AuthRepository) GetPassword(username string) string {
	rows, err := repo.Query("SELECT password FROM users WHERE username = \"" + username + "\"")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var password string
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.Password); err != nil {
			panic(err.Error())
		}
		password = u.Password
	}
	return password
}

func (repo *AuthRepository) GetUserName(password string) string {
	rows, err := repo.Query("SELECT username FROM users WHERE password = \"" + password + "\"")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var username string
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserName); err != nil {
			panic(err.Error())
		}
		username = u.UserName
	}
	return username
}

func (repo *AuthRepository) UserUniqueCheck(username string) string {
	rows, err := repo.Query("SELECT username FROM users WHERE username = \"" + username + "\"")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var existUser string
	var u domain.User
	for rows.Next() {
		if err := rows.Scan(&u.UserName); err != nil {
			panic(err.Error())
		}
		existUser = u.UserName
	}
	return existUser
}

func (repo *AuthRepository) UserAdd(username, password string) {
	_, err := repo.Query("INSERT INTO users (username, password) VALUES ( \"" + username + "\" , \"" + password + "\" )")
	if err != nil {
		panic(err.Error())
	}
}
