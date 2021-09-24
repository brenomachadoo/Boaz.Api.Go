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
	Add() bool
	Delete() bool
	Find(entity interface{}, newid int, err error) interface{}
	FindAll() []interface{}
	Update() bool
	Insert() bool
}
