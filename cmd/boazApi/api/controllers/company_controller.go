package controllers

import (
	"bmachado/Boaz.Api.Go/domain/entities"
	dbMysqlDrive "bmachado/Boaz.Api.Go/infra/database/Mysql"
	Services "bmachado/Boaz.Api.Go/services/companyService"
	"bmachado/Boaz.Api.Go/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID not integer",
		})
		return
	}

	companyResult, errorService := Services.NewCompanyService().Get(newid)
	if errorService != nil {
		util.LogFatal(errorService)
	}
	util.LogInfo(companyResult)

	// db := dbMysqlDrive.GetDatabase()

	// var empresa entities.Company
	// err = db.First(empresa, newid).Error
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"Error": "It was not possible find company: " + err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(200, companyResult)
}

func AddCompany(c *gin.Context) {
	db := dbMysqlDrive.GetDatabase()

	var companyEntity *entities.Company

	err := c.ShouldBindJSON(companyEntity)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error JSON convert: " + err.Error(),
		})
		return
	}

	err = db.Create(companyEntity).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Erro register company: " + err.Error(),
		})
		return
	}

	c.JSON(200, companyEntity)
}

func GetCompanies(c *gin.Context) {
	db := dbMysqlDrive.GetDatabase()
	var empresas []entities.Company
	err := db.Find(empresas).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Error find company: " + err.Error(),
		})

		return
	}

	c.JSON(200, empresas)
}
