package infrastructure

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sanshirookazaki/echo-clean/domain"
)

// SQLHandler .
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler .
func NewSQLHandler() *SQLHandler {
	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/task")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (h *SQLHandler) GetTaskAll(userid int) domain.Tasks {
	rows, err := h.Conn.Query("SELECT * FROM tasks WHERE Status = 0 and userid = " + strconv.Itoa(userid))
	var tasks domain.Tasks
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Userid, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, *t)
	}
	return tasks
}
