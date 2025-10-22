package controllers

import (
	"companies/repositories"
	"companies/services"
	"companies/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyService *services.CompanyService
}

func NewCompanyController() *CompanyController {
	companyRepo := repositories.NewCompanyRepository()
	return &CompanyController{
		companyService: services.NewCompanyService(companyRepo),
	}
}

func (cc *CompanyController) GetCompanies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	companies, err := cc.companyService.GetCompanies(page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponseWithPagination(c, companies, page, pageSize)
}
