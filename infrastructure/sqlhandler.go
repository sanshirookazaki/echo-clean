package infrastructure

import (
	//"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/sanshirookazaki/echo-clean/interfaces/database"
)

// SQLHandler .
type SQLHandler struct {
	Conn *gorm.DB
}

// NewSQLHandler .
func NewSQLHandler() *SQLHandler {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/task?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	defer db.close
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	return sqlHandler
}

func (handler *SQLHandler) Table(statement string) *SQLHandler {
	return handler.Conn.Table(statement)
}

func (handler *SQLHandler) Select(query interface{}, args ...interface{}) *SQLHandler {
	return handler.Conn.Select(query, args...)
}

func (handler *SQLHandler) Where(query interface{}, args ...interface{}) *SQLHandler {
	return handler.Conn.Where(query, args...)
}

func (handler *SQLHandler) Scan(dest interface{}) *SQLHandler {
	return handler.Conn.Scan(dest)
}

func (handler *SQLHandler) Create(value interface{}) *SQLHandler {
	return handler.Conn.Create(value)
}

func (handler *SQLHandler) Save(value interface{}) *SQLHandler {
	return handler.Conn.Save(value)
}

func (handler *SQLHandler) Delete(value interface{}, where ...interface{}) *SQLHandler {
	return handler.Conn.Save(value, where...)
}

func (handler *SQLHandler) Update(attrs ...interface{}) *SQLHandler {
	return handler.Conn.Update(attrs...)
}

func (handler *SQLHandler) Model(value interface{}) *SQLHandler {
	return handler.Conn.Model(value)
}
