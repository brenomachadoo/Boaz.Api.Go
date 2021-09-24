package dbMysqlDrive

import (
	"fmt"
	"reflect"

	"bmachado/Boaz.Api.Go/infra/database"

	"gorm.io/gorm"
)

type crudGeneric struct {
	Crud database.Crud
}

var db *gorm.DB

func Crud() crudGeneric {
	db = GetDatabase()
	return crudGeneric{}
}

func (crud crudGeneric) Find(entity interface{}, newid int, err error) interface{} {
	gormReult := db.First(&entity, newid)
	if gormReult.Error != nil {
		err = gormReult.Error
		fmt.Println("Error: It was not possible find " + reflect.TypeOf(entity).String() + ": " + err.Error())
		return entity
	}
	return entity
}
