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

func (repo *UserRepository) GetTaskAll(userid int) domain.Tasks {
	rows, err := repo.Query("SELECT * FROM tasks WHERE Status = 0 and userid = " + strconv.Itoa(userid))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var ts domain.Tasks
	var t domain.Task
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.UserID, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		ts = append(ts, t)
	}
	return ts
}
