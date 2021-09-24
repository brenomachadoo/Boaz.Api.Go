package database

type DatabaseTypeEnum int8

const (
	DatabaseTypeMysql   DatabaseTypeEnum = 0
	DatabaseTypePostgre                  = 1
)
