package routers

import (
	"companies/controllers"
	"companies/db"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		dbService, err := db.NewServiceFromEnv()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		defer dbService.Close()
		if err := dbService.Health(); err != nil {
			c.JSON(500, gin.H{"error": "Database connection failed"})
			return
		}
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Authentication routes (no auth required)
	companyController := controllers.NewCompanyController()
	router.GET("/companies", companyController.GetCompanies)

	return router
}
