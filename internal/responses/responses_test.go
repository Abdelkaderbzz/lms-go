package responses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// ============================================================================
// Helper Test Functions
// ============================================================================

// setupTestContext creates a test Gin context with a response recorder
func setupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	return ctx, w
}

// parseResponse parses the response body into a map
func parseResponse(body string, v interface{}) error {
	return json.Unmarshal([]byte(body), v)
}

// ============================================================================
// Success Response Tests
// ============================================================================

func TestSuccessResponse(t *testing.T) {
	ctx, w := setupTestContext()
	testData := map[string]interface{}{"id": "123", "name": "Test"}

	Success(ctx, "Success", testData)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["message"] != "Success" {
		t.Errorf("Expected message 'Success', got %s", response["message"])
	}

	if response["success"] != true {
		t.Errorf("Expected success=true, got %v", response["success"])
	}
}

func TestSuccessCreatedResponse(t *testing.T) {
	ctx, w := setupTestContext()
	testData := map[string]interface{}{"id": "123"}

	SuccessCreated(ctx, "Created", testData)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["code"] != 201.0 {
		t.Errorf("Expected code 201, got %v", response["code"])
	}
}

func TestSuccessPaginatedResponse(t *testing.T) {
	ctx, w := setupTestContext()
	items := []map[string]interface{}{
		{"id": "1", "name": "Item 1"},
		{"id": "2", "name": "Item 2"},
	}

	SuccessPaginated(ctx, items, 10, 1, 5)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["success"] != true {
		t.Errorf("Expected success=true")
	}
}

func TestNoContentResponse(t *testing.T) {
	ctx, w := setupTestContext()

	NoContent(ctx)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}
}

func TestAcceptedResponse(t *testing.T) {
	ctx, w := setupTestContext()
	testData := map[string]interface{}{"task_id": "123"}

	Accepted(ctx, "Request accepted", testData)

	if w.Code != http.StatusAccepted {
		t.Errorf("Expected status 202, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["code"] != 202.0 {
		t.Errorf("Expected code 202")
	}
}

// ============================================================================
// Error Response Tests
// ============================================================================

func TestBadRequestResponse(t *testing.T) {
	ctx, w := setupTestContext()
	errors := []map[string]interface{}{
		{"field": "email", "message": "invalid format"},
	}

	BadRequest(ctx, "Invalid input", errors)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["success"] != false {
		t.Errorf("Expected success=false")
	}

	if response["code"] != 400.0 {
		t.Errorf("Expected code 400")
	}
}

func TestUnauthorizedResponse(t *testing.T) {
	ctx, w := setupTestContext()

	Unauthorized(ctx, "Invalid credentials")

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["code"] != 401.0 {
		t.Errorf("Expected code 401")
	}
}

func TestForbiddenResponse(t *testing.T) {
	ctx, w := setupTestContext()

	Forbidden(ctx, "Access denied")

	if w.Code != http.StatusForbidden {
		t.Errorf("Expected status 403, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["code"] != 403.0 {
		t.Errorf("Expected code 403")
	}
}

func TestNotFoundResponse(t *testing.T) {
	ctx, w := setupTestContext()

	NotFound(ctx, "Resource not found")

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["code"] != 404.0 {
		t.Errorf("Expected code 404")
	}
}

func TestConflictResponse(t *testing.T) {
	ctx, w := setupTestContext()

	Conflict(ctx, "Resource already exists")

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["code"] != 409.0 {
		t.Errorf("Expected code 409")
	}
}

func TestInternalServerErrorResponse(t *testing.T) {
	ctx, w := setupTestContext()

	InternalServerError(ctx, "Server error", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["code"] != 500.0 {
		t.Errorf("Expected code 500")
	}
}

// ============================================================================
// Validation Error Response Tests
// ============================================================================

func TestValidationFailedResponse(t *testing.T) {
	ctx, w := setupTestContext()
	errors := []map[string]interface{}{
		{"field": "email", "message": "invalid format"},
		{"field": "password", "message": "too weak"},
	}

	ValidationFailed(ctx, errors)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response map[string]interface{}
	parseResponse(w.Body.String(), &response)

	if response["success"] != false {
		t.Errorf("Expected success=false")
	}
}

// ============================================================================
// Response Structure Tests
// ============================================================================

func TestAPIResponseStructure(t *testing.T) {
	ctx, w := setupTestContext()
	Success(ctx, "Test message", map[string]interface{}{"id": "123"})

	var response APIResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Message != "Test message" {
		t.Errorf("Expected message 'Test message'")
	}

	if !response.Success {
		t.Errorf("Expected success=true")
	}

	if response.Code != http.StatusOK {
		t.Errorf("Expected code 200")
	}
}

func TestPaginatedResponseStructure(t *testing.T) {
	ctx, w := setupTestContext()
	items := []interface{}{"item1", "item2"}
	SuccessPaginated(ctx, items, 20, 2, 10)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["data"] == nil {
		t.Error("Expected data field in response")
	}

	// Data contains the paginated response
	data := response["data"].(map[string]interface{})

	if data["items"] == nil {
		t.Error("Expected items field in data")
	}

	// Check total (could be int64 or float64 depending on JSON encoding)
	total := data["total"]
	if total != 20.0 && total != int64(20) {
		t.Errorf("Expected total=20, got %v", total)
	}

	// Check page (could be int64 or float64 depending on JSON encoding)
	page := data["page"]
	if page != 2.0 && page != int64(2) {
		t.Errorf("Expected page=2, got %v", page)
	}

	// Check page_size (could be int64 or float64 depending on JSON encoding)
	pageSize := data["page_size"]
	if pageSize != 10.0 && pageSize != int64(10) {
		t.Errorf("Expected page_size=10, got %v", pageSize)
	}
}

// ============================================================================
// Error Array Tests
// ============================================================================

func TestMultipleValidationErrors(t *testing.T) {
	ctx, w := setupTestContext()
	errors := []map[string]interface{}{
		{"field": "email", "message": "required"},
		{"field": "password", "message": "too short"},
		{"field": "role", "message": "invalid value"},
	}

	BadRequest(ctx, "Validation failed", errors)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["code"] != 400.0 {
		t.Error("Expected 400 status code")
	}

	if response["message"] != "Validation failed" {
		t.Error("Expected custom error message")
	}
}

// ============================================================================
// Response JSON Encoding Tests
// ============================================================================

func TestResponseJSONEncoding(t *testing.T) {
	tests := []struct {
		name       string
		createResp func(*gin.Context, *httptest.ResponseRecorder)
		checkCode  int
	}{
		{
			name: "Success response",
			createResp: func(ctx *gin.Context, w *httptest.ResponseRecorder) {
				Success(ctx, "Success", map[string]interface{}{"id": "1"})
			},
			checkCode: 200,
		},
		{
			name: "Created response",
			createResp: func(ctx *gin.Context, w *httptest.ResponseRecorder) {
				SuccessCreated(ctx, "Created", map[string]interface{}{"id": "1"})
			},
			checkCode: 201,
		},
		{
			name: "Bad request response",
			createResp: func(ctx *gin.Context, w *httptest.ResponseRecorder) {
				BadRequest(ctx, "Invalid", nil)
			},
			checkCode: 400,
		},
		{
			name: "Unauthorized response",
			createResp: func(ctx *gin.Context, w *httptest.ResponseRecorder) {
				Unauthorized(ctx, "Auth failed")
			},
			checkCode: 401,
		},
		{
			name: "Not found response",
			createResp: func(ctx *gin.Context, w *httptest.ResponseRecorder) {
				NotFound(ctx, "Not found")
			},
			checkCode: 404,
		},
		{
			name: "Server error response",
			createResp: func(ctx *gin.Context, w *httptest.ResponseRecorder) {
				InternalServerError(ctx, "Error", nil)
			},
			checkCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := setupTestContext()
			tt.createResp(ctx, w)

			if w.Code != tt.checkCode {
				t.Errorf("Expected code %d, got %d", tt.checkCode, w.Code)
			}

			// Verify valid JSON
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Response is not valid JSON: %v", err)
			}
		})
	}
}

// ============================================================================
// Response Content-Type Tests
// ============================================================================

func TestResponseContentType(t *testing.T) {
	ctx, w := setupTestContext()
	Success(ctx, "Test", nil)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" && contentType != "application/json" {
		t.Errorf("Expected JSON content-type, got %s", contentType)
	}
}

// ============================================================================
// Error Message Format Tests
// ============================================================================

func TestErrorMessageFormat(t *testing.T) {
	tests := []struct {
		name    string
		message string
		code    int
	}{
		{"Validation error", "Invalid email format", 400},
		{"Not found", "User not found", 404},
		{"Unauthorized", "Invalid credentials", 401},
		{"Server error", "Internal server error", 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, w := setupTestContext()

			if tt.code == 400 {
				BadRequest(ctx, tt.message, nil)
			} else if tt.code == 404 {
				NotFound(ctx, tt.message)
			} else if tt.code == 401 {
				Unauthorized(ctx, tt.message)
			} else if tt.code == 500 {
				InternalServerError(ctx, tt.message, nil)
			}

			if w.Code != tt.code {
				t.Errorf("Expected code %d, got %d", tt.code, w.Code)
			}

			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)

			if response["message"] != tt.message {
				t.Errorf("Expected message '%s', got '%v'", tt.message, response["message"])
			}
		})
	}
}
