package main

import (
	"companies/db"
	"companies/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()

	dbService, err := db.NewServiceFromEnv()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbService.Close()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		if err := dbService.Health(); err != nil {
			c.JSON(500, gin.H{"error": "Database connection failed"})
			return
		}
		c.JSON(200, gin.H{"status": "healthy"})
	})

	router.GET("/companies", func(c *gin.Context) {
		rows, err := dbService.DB.Query(`SELECT 
    			id, name, description, visa_sponsor, website, logo_url, jobs_page, created_at, updated_at
			    FROM company LIMIT 2`)
		if err != nil {
			c.JSON(500, gin.H{"error": "Query execution failed, err." + err.Error()})
			return
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
				c.JSON(500, gin.H{"error": "Row scan failed, err." + err.Error()})
				return
			}
			companies = append(companies, company)
		}

		c.JSON(200, gin.H{"data": companies})
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
