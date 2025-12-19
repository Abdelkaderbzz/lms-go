package utils

import (
	"testing"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

// TestGenerateToken tests token generation
func TestGenerateToken(t *testing.T) {
	userID := "test-user-id"
	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(userID, email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Error("Expected non-empty token")
	}
}

// TestValidateToken tests token validation
func TestValidateToken(t *testing.T) {
	userID := "test-user-id"
	email := "test@example.com"
	role := "student"

	// Generate a token
	token, err := GenerateToken(userID, email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate the token
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected user ID %s, got %s", userID, claims.UserID)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

// TestValidateInvalidToken tests validation of invalid token
func TestValidateInvalidToken(t *testing.T) {
	invalidToken := "invalid.token.here"

	_, err := ValidateToken(invalidToken)
	if err == nil {
		t.Error("Expected error when validating invalid token")
	}
}

// TestTokenExpiration tests token expiration
func TestTokenExpiration(t *testing.T) {
	userID := "test-user-id"
	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(userID, email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	// Check that token has expiration time
	if claims.RegisteredClaims.ExpiresAt == nil {
		t.Error("Expected token to have expiration time")
	}

	// Check that token will expire in future (approximately 24 hours)
	if claims.RegisteredClaims.ExpiresAt.Time.Before(time.Now()) {
		t.Error("Expected token to not be expired")
	}
}

// TestClaimsStructure tests Claims structure
func TestClaimsStructure(t *testing.T) {
	claims := &Claims{
		UserID: "123",
		Email:  "test@example.com",
		Role:   "admin",
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	if claims.UserID != "123" {
		t.Errorf("Expected UserID 123, got %s", claims.UserID)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", claims.Email)
	}

	if claims.Role != "admin" {
		t.Errorf("Expected role admin, got %s", claims.Role)
	}
}
