package Services

import (
	"bmachado/Boaz.Api.Go/domain/entities"
)

type companyService struct {
	companyServiceInterface
}

var CompanyService companyService = companyService{}

func (service companyService) Add(company *entities.Company) (*entities.Company, error) {
	var err error
	return company, err
}
