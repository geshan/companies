package services

import (
	"companies/models"
)

type CompanyRepositoryInterface interface {
	GetCompanies(page, pageSize int) ([]models.Company, error)
}

type CompanyService struct {
	repo CompanyRepositoryInterface
}

func NewCompanyService(repo CompanyRepositoryInterface) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) GetCompanies(page, pageSize int) ([]models.Company, error) {
	var companies []models.Company

	companies, err := s.repo.GetCompanies(page, pageSize)

	return companies, err
}
