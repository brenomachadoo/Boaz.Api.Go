package companyServices

import "bmachado/Boaz.Api.Go/domain/entities"

type CompanyService interface {
	Add(company *entities.Company) error
	Delete(id string) error
	Update(company *entities.Company)
	GetAll() ([]*entities.Company, error)
	Get(id int) (*entities.Company, error)
}
