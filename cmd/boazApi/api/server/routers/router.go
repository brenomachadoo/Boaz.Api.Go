package routers

import (
	"bmachado/Boaz.Api.Go/cmd/boazApi/api/controllers"

	"github.com/gin-gonic/gin"
)

func ContigureRoutesCompany(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		empresas := main.Group("company")
		{
			empresas.GET("/:id", controllers.GetCompany)
			empresas.GET("/", controllers.GetCompanies)
			empresas.POST("/", controllers.AddCompany)
		}
	}

	return router
}
