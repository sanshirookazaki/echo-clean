package database

type SQLHandler interface {
	Table(string) SQLHandler
	Select(interface{}, ...interface{}) SQLHandler
	Where(interface{}, ...interface{}) SQLHandler
	Scan(interface{}) SQLHandler
	Create(interface{}) SQLHandler
}
