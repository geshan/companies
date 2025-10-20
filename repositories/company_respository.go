package repositories

import (
	"companies/db"
	"companies/models"
	"log"
)

type CompanyRepository struct {
	dbService *db.Service
}

func NewCompanyRepository() *CompanyRepository {
	svc, err := db.NewServiceFromEnv()
	if err != nil {
		log.Fatalf("failed to initialize db service: %v", err)
	}
	return &CompanyRepository{
		dbService: svc,
	}
}

func (r *CompanyRepository) GetCompanies(page, pageSize int) ([]models.Company, error) {
	rows, err := r.dbService.DB.Query(`SELECT
        id, name, description, visa_sponsor, website, logo_url, jobs_page, created_at, updated_at
        FROM company LIMIT ? OFFSET ?`, pageSize, (page-1)*pageSize)
	if err != nil {
		log.Println("Query execution failed: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	companies := []models.Company{}
	for rows.Next() {
		var company models.Company
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Description,
			&company.VisaSponsor,
			&company.Website,
			&company.LogoURL,
			&company.JobsPage,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}
