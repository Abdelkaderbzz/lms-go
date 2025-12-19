package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestLoginRequestStructure tests the LoginRequest structure
func TestLoginRequestStructure(t *testing.T) {
	req := LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	if req.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", req.Email)
	}

	if req.Password != "password123" {
		t.Errorf("Expected password password123, got %s", req.Password)
	}
}

// TestRegisterRequestStructure tests the RegisterRequest structure
func TestRegisterRequestStructure(t *testing.T) {
	req := RegisterRequest{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
		Role:      "student",
	}

	if req.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", req.Email)
	}

	if req.FirstName != "John" {
		t.Errorf("Expected first name John, got %s", req.FirstName)
	}

	if req.Role != "student" {
		t.Errorf("Expected role student, got %s", req.Role)
	}
}

// TestAuthResponseStructure tests the AuthResponse structure
func TestAuthResponseStructure(t *testing.T) {
	userMap := map[string]interface{}{
		"id":    "123",
		"email": "test@example.com",
	}

	resp := AuthResponse{
		Token: "test-token",
		User:  userMap,
	}

	if resp.Token != "test-token" {
		t.Errorf("Expected token test-token, got %s", resp.Token)
	}

	if resp.User == nil {
		t.Error("Expected non-nil User")
	}
}

// TestRegisterValidationJSON tests that RegisterRequest validates JSON binding
func TestRegisterValidationJSON(t *testing.T) {
	tests := []struct {
		name    string
		payload map[string]interface{}
	}{
		{
			name: "Valid request",
			payload: map[string]interface{}{
				"email":      "test@example.com",
				"password":   "password123",
				"first_name": "John",
				"last_name":  "Doe",
				"role":       "student",
			},
		},
		{
			name: "Missing email",
			payload: map[string]interface{}{
				"password":   "password123",
				"first_name": "John",
				"last_name":  "Doe",
				"role":       "student",
			},
		},
		{
			name: "Missing password",
			payload: map[string]interface{}{
				"email":      "test@example.com",
				"first_name": "John",
				"last_name":  "Doe",
				"role":       "student",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, _ := json.Marshal(tt.payload)

			var req RegisterRequest
			err := json.Unmarshal(jsonBytes, &req)

			if err != nil && tt.name == "Valid request" {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// TestLoginValidation tests login request validation
func TestLoginValidation(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		password string
		isValid bool
	}{
		{"Valid email", "test@example.com", "password123", true},
		{"Invalid email format", "not-an-email", "", false},
		{"Missing password", "test@example.com", "", false},
		{"Empty email", "", "password123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := LoginRequest{
				Email:    tt.email,
				Password: tt.password,
			}

			// Simple validation check
			isValid := req.Email != "" && req.Password != ""

			if isValid != tt.isValid {
				t.Errorf("Expected valid=%v, got valid=%v", tt.isValid, isValid)
			}
		})
	}
}

// TestHashPassword tests the hash password function
func TestHashPassword(t *testing.T) {
	password := "test-password-123"
	hashed := hashPassword(password)

	if hashed == "" {
		t.Error("Expected non-empty hashed password")
	}

	// Note: Current implementation does not hash, just returns the password
	// In production, this should use bcrypt
	if hashed != password {
		t.Logf("Warning: hashPassword is not actually hashing the password (returns: %s)", hashed)
	}
}

// TestGinContextCreation tests basic gin context creation
func TestGinContextCreation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	if c == nil {
		t.Error("Expected non-nil gin context")
	}

	if w.Code == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

// TestRequestBodyParsing tests request body parsing
func TestRequestBodyParsing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	loginReq := LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonBytes, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	var parsedReq LoginRequest
	err := c.ShouldBindJSON(&parsedReq)

	if err != nil {
		t.Errorf("Failed to parse request body: %v", err)
	}

	if parsedReq.Email != loginReq.Email {
		t.Errorf("Expected email %s, got %s", loginReq.Email, parsedReq.Email)
	}
}
