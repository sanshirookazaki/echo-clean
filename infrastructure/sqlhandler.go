package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanshirookazaki/echo-clean/interfaces/database"
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

func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

type SQLRow struct {
	Rows *sql.Rows
}

func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

func (r SQLRow) Close() error {
	return r.Rows.Close()
}
