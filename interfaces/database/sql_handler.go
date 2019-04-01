package database

import "github.com/jinzhu/gorm"

type SQLHandler interface {
	Table(string) *gorm.DB
	Select(interface{}, ...interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	Scan(interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}, ...interface{}) *gorm.DB
	Update(...interface{}) *gorm.DB
	Model(interface{}) *gorm.DB
}
