package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

type PaginationResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func SuccessResponseWithPagination(c *gin.Context, data interface{}, page int, pageSize int) {
	c.JSON(200, PaginationResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
		Pagination: Pagination{
			Page:     page,
			PageSize: pageSize,
		},
	})
}
