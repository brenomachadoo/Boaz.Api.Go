package migrations

import (
	"bmachado/Boaz.Api.Go/domain/entities"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(entities.Company{})
}
