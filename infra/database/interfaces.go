package database

type Provider interface {
	ConnectionConfig()
	OpenConnection()
	ClosedConnection()
	GetDatabase()
	Migration()
}

type Crud interface {
	getDatabase()
	Add(entity interface{}) (interface{}, error)
	Delete() bool
	Find(entity interface{}, id int) (interface{}, error)
	FindAll() ([]interface{}, error)
	Update() bool
	Insert() bool
}
