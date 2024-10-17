package services

import "bmachado/Boaz.Api.Go/domain/entities"

type companyServiceInterface interface {
	Add(company *entities.Company) (*entities.Company, error)
	Delete(id string) error
	Update(company *entities.Company)
	GetAll() ([]*entities.Company, error)
	Get(id int) (*entities.Company, error)
}
