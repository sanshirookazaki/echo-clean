package infrastructure

import (
	"database/sql"

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

