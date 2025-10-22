package services

import (
	"companies/models"
	"errors"
	"reflect"
	"testing"
	"time"
)

func strPtr(s string) *string { return &s }

type mockCompanyRepository struct {
	mockData []models.Company
	err      error
}

func (m *mockCompanyRepository) GetCompanies(page, pageSize int) ([]models.Company, error) {
	return m.mockData, m.err
}

func TestGetCompanies_Success(t *testing.T) {
	createdAt1, _ := time.Parse(time.RFC3339, "2020-06-28T23:28:22Z")
	updatedAt1, _ := time.Parse(time.RFC3339, "2020-07-28T02:00:43Z")
	createdAt2, _ := time.Parse(time.RFC3339, "2020-06-28T23:28:22Z")
	updatedAt2, _ := time.Parse(time.RFC3339, "2020-07-28T02:01:05Z")

	mockData := []models.Company{
		{
			ID:          1,
			Name:        "tyro",
			Description: strPtr("Launching in 2003, Tyro has grown to become Australia's largest EFTPOS."),
			VisaSponsor: true,
			Website:     strPtr("https://www.tyro.com"),
			LogoURL:     strPtr("https://raw.githubusercontent.com/autechjobs/assets/master/images/company-logos/tyro.png"),
			JobsPage:    strPtr("https://jobs.lever.co/tyro/"),
			CreatedAt:   createdAt1,
			UpdatedAt:   updatedAt1,
		},
		{
			ID:          2,
			Name:        "SafetyCulture",
			Description: strPtr("SafetyCulture is an Australian-based, international tech scale-up. We create SaaS solutions."),
			VisaSponsor: true,
			Website:     strPtr("https://safetyculture.com.au"),
			LogoURL:     strPtr("https://raw.githubusercontent.com/autechjobs/assets/master/images/company-logos/safety-culture.png"),
			JobsPage:    strPtr("https://safetyculture.com/jobs/"),
			CreatedAt:   createdAt2,
			UpdatedAt:   updatedAt2,
		},
	}
	mockRepo := &mockCompanyRepository{mockData: mockData}
	service := NewCompanyService(mockRepo)

	companies, err := service.GetCompanies(1, 10)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(companies, mockData) {
		t.Fatalf("expected %v, got %v", mockData, companies)
	}
}

func TestGetCompanies_Error(t *testing.T) {
	mockErr := errors.New("database failure")
	mockRepo := &mockCompanyRepository{err: mockErr}
	service := NewCompanyService(mockRepo)

	_, err := service.GetCompanies(1, 10)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != mockErr {
		t.Fatalf("expected %v, got %v", mockErr, err)
	}
}
