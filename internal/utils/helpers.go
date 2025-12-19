package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashString creates a SHA256 hash of a string
func HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) bool {
	// Basic email validation - in production, use a more robust method
	if len(email) < 5 {
		return false
	}
	hasAt := false
	hasDot := false
	for i, c := range email {
		if c == '@' && !hasAt && i > 0 {
			hasAt = true
		} else if c == '.' && hasAt && i < len(email)-1 {
			hasDot = true
		}
	}
	return hasAt && hasDot
}

// ValidatePassword checks if password meets requirements
func ValidatePassword(password string) bool {
	// Minimum 6 characters
	return len(password) >= 6
}

// PaginationParams holds pagination information
type PaginationParams struct {
	Page    int `json:"page" query:"page"`
	PerPage int `json:"per_page" query:"per_page"`
}

// GetOffset calculates the offset for pagination
func (p *PaginationParams) GetOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 {
		p.PerPage = 10
	}
	if p.PerPage > 100 {
		p.PerPage = 100
	}
	return (p.Page - 1) * p.PerPage
}

// GetLimit returns the limit for pagination
func (p *PaginationParams) GetLimit() int {
	if p.PerPage <= 0 {
		p.PerPage = 10
	}
	if p.PerPage > 100 {
		p.PerPage = 100
	}
	return p.PerPage
}
