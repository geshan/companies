package services

import (
	"companies/models"
	"companies/repositories"
)

type CompanyService struct {
	companyRepository *repositories.CompanyRepository
}

func NewCompanyService() *CompanyService {
	return &CompanyService{
		companyRepository: repositories.NewCompanyRepository(),
	}
}

func (s *CompanyService) GetCompanies(page, pageSize int) ([]models.Company, error) {
	var companies []models.Company

	companies, err := s.companyRepository.GetCompanies(page, pageSize)

	return companies, err
}
