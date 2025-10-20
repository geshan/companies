package models

import "time"

type Company struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description *string   `db:"description" json:"description,omitempty"`
	VisaSponsor bool      `db:"visa_sponsor" json:"visa_sponsor"`
	Website     *string   `db:"website" json:"website,omitempty"`
	LogoURL     *string   `db:"logo_url" json:"logo_url,omitempty"`
	JobsPage    *string   `db:"jobs_page" json:"jobs_page,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
