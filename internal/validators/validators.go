package validators

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors is a slice of validation errors
type ValidationErrors []ValidationError

// Error implements the error interface
func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return "validation passed"
	}
	messages := make([]string, len(ve))
	for i, e := range ve {
		messages[i] = fmt.Sprintf("%s: %s", e.Field, e.Message)
	}
	return strings.Join(messages, "; ")
}

// IsEmpty checks if there are any errors
func (ve ValidationErrors) IsEmpty() bool {
	return len(ve) == 0
}

// RegisterRequest validation
type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

func (r RegisterRequest) Validate() ValidationErrors {
	var errs ValidationErrors

	// Email validation
	if !isValidEmail(r.Email) {
		errs = append(errs, ValidationError{
			Field:   "email",
			Message: "invalid email format",
		})
	}

	// Password validation
	if len(r.Password) < 8 {
		errs = append(errs, ValidationError{
			Field:   "password",
			Message: "password must be at least 8 characters",
		})
	}
	if !containsUppercase(r.Password) {
		errs = append(errs, ValidationError{
			Field:   "password",
			Message: "password must contain at least one uppercase letter",
		})
	}
	if !containsDigit(r.Password) {
		errs = append(errs, ValidationError{
			Field:   "password",
			Message: "password must contain at least one digit",
		})
	}

	// FirstName validation
	if len(strings.TrimSpace(r.FirstName)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "first_name",
			Message: "first name is required",
		})
	}
	if len(r.FirstName) > 100 {
		errs = append(errs, ValidationError{
			Field:   "first_name",
			Message: "first name cannot exceed 100 characters",
		})
	}

	// LastName validation
	if len(strings.TrimSpace(r.LastName)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "last_name",
			Message: "last name is required",
		})
	}
	if len(r.LastName) > 100 {
		errs = append(errs, ValidationError{
			Field:   "last_name",
			Message: "last name cannot exceed 100 characters",
		})
	}

	// Role validation
	validRoles := map[string]bool{"admin": true, "instructor": true, "student": true}
	if !validRoles[r.Role] {
		errs = append(errs, ValidationError{
			Field:   "role",
			Message: "invalid role. must be admin, instructor, or student",
		})
	}

	return errs
}

// LoginRequest validation
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r LoginRequest) Validate() ValidationErrors {
	var errs ValidationErrors

	if !isValidEmail(r.Email) {
		errs = append(errs, ValidationError{
			Field:   "email",
			Message: "invalid email format",
		})
	}

	if len(r.Password) == 0 {
		errs = append(errs, ValidationError{
			Field:   "password",
			Message: "password is required",
		})
	}

	return errs
}

// CreateCourseRequest validation
type CreateCourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Level       string `json:"level" binding:"required"`
	MaxStudents int    `json:"max_students" binding:"required"`
}

func (r CreateCourseRequest) Validate() ValidationErrors {
	var errs ValidationErrors

	if len(strings.TrimSpace(r.Title)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "title",
			Message: "title is required",
		})
	}
	if len(r.Title) > 255 {
		errs = append(errs, ValidationError{
			Field:   "title",
			Message: "title cannot exceed 255 characters",
		})
	}

	if len(strings.TrimSpace(r.Description)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "description",
			Message: "description is required",
		})
	}
	if len(r.Description) > 2000 {
		errs = append(errs, ValidationError{
			Field:   "description",
			Message: "description cannot exceed 2000 characters",
		})
	}

	if len(strings.TrimSpace(r.Code)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "code",
			Message: "course code is required",
		})
	}
	if len(r.Code) > 50 {
		errs = append(errs, ValidationError{
			Field:   "code",
			Message: "course code cannot exceed 50 characters",
		})
	}

	validLevels := map[string]bool{"Beginner": true, "Intermediate": true, "Advanced": true}
	if !validLevels[r.Level] {
		errs = append(errs, ValidationError{
			Field:   "level",
			Message: "invalid level. must be Beginner, Intermediate, or Advanced",
		})
	}

	if r.MaxStudents <= 0 {
		errs = append(errs, ValidationError{
			Field:   "max_students",
			Message: "max students must be greater than 0",
		})
	}
	if r.MaxStudents > 10000 {
		errs = append(errs, ValidationError{
			Field:   "max_students",
			Message: "max students cannot exceed 10000",
		})
	}

	return errs
}

// CreateAssignmentRequest validation
type CreateAssignmentRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Points      float64 `json:"points" binding:"required"`
	Type        string  `json:"type" binding:"required"`
}

func (r CreateAssignmentRequest) Validate() ValidationErrors {
	var errs ValidationErrors

	if len(strings.TrimSpace(r.Title)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "title",
			Message: "title is required",
		})
	}
	if len(r.Title) > 255 {
		errs = append(errs, ValidationError{
			Field:   "title",
			Message: "title cannot exceed 255 characters",
		})
	}

	if len(strings.TrimSpace(r.Description)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "description",
			Message: "description is required",
		})
	}

	if r.Points <= 0 {
		errs = append(errs, ValidationError{
			Field:   "points",
			Message: "points must be greater than 0",
		})
	}
	if r.Points > 1000 {
		errs = append(errs, ValidationError{
			Field:   "points",
			Message: "points cannot exceed 1000",
		})
	}

	validTypes := map[string]bool{"homework": true, "project": true, "essay": true, "quiz": true}
	if !validTypes[r.Type] {
		errs = append(errs, ValidationError{
			Field:   "type",
			Message: "invalid type. must be homework, project, essay, or quiz",
		})
	}

	return errs
}

// CreateQuizRequest validation
type CreateQuizRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	TimeLimit   int     `json:"time_limit" binding:"required"`
	PassScore   float64 `json:"pass_score" binding:"required"`
}

func (r CreateQuizRequest) Validate() ValidationErrors {
	var errs ValidationErrors

	if len(strings.TrimSpace(r.Title)) == 0 {
		errs = append(errs, ValidationError{
			Field:   "title",
			Message: "title is required",
		})
	}
	if len(r.Title) > 255 {
		errs = append(errs, ValidationError{
			Field:   "title",
			Message: "title cannot exceed 255 characters",
		})
	}

	if r.TimeLimit <= 0 {
		errs = append(errs, ValidationError{
			Field:   "time_limit",
			Message: "time limit must be greater than 0 minutes",
		})
	}
	if r.TimeLimit > 480 {
		errs = append(errs, ValidationError{
			Field:   "time_limit",
			Message: "time limit cannot exceed 480 minutes (8 hours)",
		})
	}

	if r.PassScore < 0 || r.PassScore > 100 {
		errs = append(errs, ValidationError{
			Field:   "pass_score",
			Message: "pass score must be between 0 and 100",
		})
	}

	return errs
}

// Helper functions
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func containsUppercase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// Pagination validation
type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (p *PaginationParams) Validate() ValidationErrors {
	var errs ValidationErrors

	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		errs = append(errs, ValidationError{
			Field:   "page_size",
			Message: "page size cannot exceed 100",
		})
		p.PageSize = 100
	}

	return errs
}

// GetPaginationParams extracts pagination from query
func GetPaginationParams(c *gin.Context) PaginationParams {
	var params PaginationParams
	if err := c.ShouldBindQuery(&params); err != nil {
		params.Page = 1
		params.PageSize = 10
	}
	// Validate and normalize pagination params
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 10
	}
	return params
}

// IDValidator validates UUID format
func IsValidUUID(id string) bool {
	const uuidRegex = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
	re := regexp.MustCompile(uuidRegex)
	return re.MatchString(strings.ToLower(id))
}

// StringValidator validates string length
func ValidateStringLength(field, value string, minLen, maxLen int) *ValidationError {
	if len(strings.TrimSpace(value)) < minLen {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be at least %d characters", minLen),
		}
	}
	if len(value) > maxLen {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("cannot exceed %d characters", maxLen),
		}
	}
	return nil
}
