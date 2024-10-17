package dbMysqlDrive

import (
	"fmt"
	"log"
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

func (crud crudGeneric) Find(entity interface{}, id int) (interface{}, error) {
	gormReult := db.First(&entity, id)

	err := gormReult.Error
	if err != nil {
		fmt.Println("Error: It was not possible find " + reflect.TypeOf(entity).String() + ": " + err.Error())
		return entity, err
	}
	return entity, err
}

func (crud crudGeneric) Add(entity interface{}) (interface{}, error) {
	err := db.Create(entity).Error
	if err != nil {
		log.Fatal(err.Error())
	}

	return entity, err
}
