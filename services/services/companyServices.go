package services

import (
	"bmachado/Boaz.Api.Go/domain/entities"
	"bmachado/Boaz.Api.Go/util"

	dbMysqlDrive "bmachado/Boaz.Api.Go/infra/database/mysql"
)

type companyService struct {
	companyServiceInterface
}

func CompanyService() companyService {
	var CompanyService = companyService{}
	dbMysqlDrive.MysqlConfig().OpenConnection()
	return CompanyService
}

func (service companyService) Get(id int) (entities.Company, error) {
	var entityResult interface{}
	entityResult, errorResult := dbMysqlDrive.Crud().Find(entities.Company{}, id)

	util.LogInfo(entityResult)

	if errorResult != nil {
		util.LogFatal(errorResult.Error())
	}

	var convertObject = entityResult.(entities.Company)

	return convertObject, errorResult
}

func (service companyService) GetAll() ([]entities.Company, error) {
	var entityResult interface{}
	entityResult, errorResult := dbMysqlDrive.Crud().Crud.FindAll()

	util.LogInfo(entityResult)

	if errorResult != nil {
		util.LogFatal(errorResult.Error())
	}

	var convertObject = entityResult.([]entities.Company)

	return convertObject, errorResult
}

func (service companyService) Add(entity *entities.Company) (entities.Company, error) {
	var entityResult interface{}
	entityResult, errorResult := dbMysqlDrive.Crud().Add(entity)

	util.LogInfo(entityResult)

	if errorResult != nil {
		util.LogFatal(errorResult.Error())
	}

	var convertObject = entityResult.(entities.Company)

	return convertObject, errorResult
}
