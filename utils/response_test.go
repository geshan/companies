package utils

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ErrorResponse(c, 400, "Bad Request")

	var resp Response
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if resp.Code != 400 {
		t.Errorf("expected code 400, got %d", resp.Code)
	}
	if resp.Message != "Bad Request" {
		t.Errorf("expected message 'Bad Request', got %s", resp.Message)
	}
	if resp.Data != nil {
		t.Errorf("expected data to be nil, got %v", resp.Data)
	}
}

func TestSuccessResponseWithPagination(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	data := []string{"a", "b"}
	SuccessResponseWithPagination(c, data, 2, 10)

	var resp PaginationResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if resp.Code != 200 {
		t.Errorf("expected code 200, got %d", resp.Code)
	}
	if resp.Message != "Success" {
		t.Errorf("expected message 'Success', got %s", resp.Message)
	}
	if resp.Pagination.Page != 2 {
		t.Errorf("expected page 2, got %d", resp.Pagination.Page)
	}
	if resp.Pagination.PageSize != 10 {
		t.Errorf("expected page_size 10, got %d", resp.Pagination.PageSize)
	}
	// Check data
	got, ok := resp.Data.([]interface{})
	if !ok || len(got) != 2 {
		t.Errorf("expected data length 2, got %v", resp.Data)
	}
}
