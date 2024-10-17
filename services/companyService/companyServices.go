package Services

import (
	"bmachado/Boaz.Api.Go/domain/entities"
	"bmachado/Boaz.Api.Go/util"

	dbMysqlDrive "bmachado/Boaz.Api.Go/infra/database/mysqlll"
)

type companyService struct {
	companyServiceInterface
}

func NewCompanyService() companyService {
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
