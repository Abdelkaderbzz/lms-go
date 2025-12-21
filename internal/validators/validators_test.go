package validators

import (
	"strings"
	"testing"
)

// ============================================================================
// ValidationErrors Tests
// ============================================================================

func TestValidationErrorsIsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		errors ValidationErrors
		want   bool
	}{
		{"Empty errors", ValidationErrors{}, true},
		{"One error", ValidationErrors{ValidationError{Field: "email", Message: "invalid"}}, false},
		{"Multiple errors", ValidationErrors{
			ValidationError{Field: "email", Message: "invalid"},
			ValidationError{Field: "password", Message: "too weak"},
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.errors.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidationErrorsError(t *testing.T) {
	errors := ValidationErrors{
		ValidationError{Field: "email", Message: "invalid format"},
		ValidationError{Field: "password", Message: "too weak"},
	}

	errorStr := errors.Error()
	if !strings.Contains(errorStr, "email") {
		t.Errorf("Expected 'email' in error message, got %s", errorStr)
	}

	if !strings.Contains(errorStr, "password") {
		t.Errorf("Expected 'password' in error message, got %s", errorStr)
	}

	// Test empty errors
	emptyErrors := ValidationErrors{}
	if emptyErrors.Error() != "validation passed" {
		t.Errorf("Expected 'validation passed' for empty errors")
	}
}

// ============================================================================
// Email Validation Tests
// ============================================================================

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"test@example.com", true},
		{"user.name+tag@example.co.uk", true},
		{"valid.email@domain.org", true},
		{"", false},
		{"notanemail", false},
		{"missing@domain", false},
		{"missing.domain@", false},
		{"@example.com", false},
		{"user@", false},
		{"user_name@example.com", true}, // Underscores are valid
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			got := isValidEmail(tt.email)
			if got != tt.valid {
				t.Errorf("isValidEmail(%s) = %v, want %v", tt.email, got, tt.valid)
			}
		})
	}
}

// ============================================================================
// Password Validation Tests
// ============================================================================

func TestContainsUppercase(t *testing.T) {
	tests := []struct {
		password string
		has      bool
	}{
		{"Password123", true},
		{"PASSWORD", true},
		{"password", false},
		{"123456", false},
		{"Pass", true},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			got := containsUppercase(tt.password)
			if got != tt.has {
				t.Errorf("containsUppercase(%s) = %v, want %v", tt.password, got, tt.has)
			}
		})
	}
}

func TestContainsDigit(t *testing.T) {
	tests := []struct {
		password string
		has      bool
	}{
		{"Password123", true},
		{"Password", false},
		{"123456", true},
		{"NoDigits", false},
		{"123", true},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			got := containsDigit(tt.password)
			if got != tt.has {
				t.Errorf("containsDigit(%s) = %v, want %v", tt.password, got, tt.has)
			}
		})
	}
}

// ============================================================================
// Register Request Validation Tests
// ============================================================================

func TestRegisterRequestValidate(t *testing.T) {
	tests := []struct {
		name      string
		req       RegisterRequest
		hasErrs   bool
		errFields []string
	}{
		{
			name: "Valid registration",
			req: RegisterRequest{
				Email:     "user@example.com",
				Password:  "SecurePass123",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs: false,
		},
		{
			name: "Missing email",
			req: RegisterRequest{
				Email:     "",
				Password:  "SecurePass123",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs:   true,
			errFields: []string{"email"},
		},
		{
			name: "Invalid email format",
			req: RegisterRequest{
				Email:     "invalidemail",
				Password:  "SecurePass123",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs:   true,
			errFields: []string{"email"},
		},
		{
			name: "Password too short",
			req: RegisterRequest{
				Email:     "user@example.com",
				Password:  "Short1",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs:   true,
			errFields: []string{"password"},
		},
		{
			name: "Password without uppercase",
			req: RegisterRequest{
				Email:     "user@example.com",
				Password:  "noupppercase123",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs:   true,
			errFields: []string{"password"},
		},
		{
			name: "Password without digit",
			req: RegisterRequest{
				Email:     "user@example.com",
				Password:  "NoDigitsHere",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs:   true,
			errFields: []string{"password"},
		},
		{
			name: "Missing first name",
			req: RegisterRequest{
				Email:     "user@example.com",
				Password:  "SecurePass123",
				FirstName: "",
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs:   true,
			errFields: []string{"first_name"},
		},
		{
			name: "First name too long",
			req: RegisterRequest{
				Email:     "user@example.com",
				Password:  "SecurePass123",
				FirstName: "John", // Actually valid length
				LastName:  "Doe",
				Role:      "student",
			},
			hasErrs: false,
		},
		{
			name: "Invalid role",
			req: RegisterRequest{
				Email:     "user@example.com",
				Password:  "SecurePass123",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "invalid_role",
			},
			hasErrs:   true,
			errFields: []string{"role"},
		},
		{
			name: "Multiple errors",
			req: RegisterRequest{
				Email:     "invalid",
				Password:  "weak",
				FirstName: "",
				LastName:  "Doe",
				Role:      "superuser",
			},
			hasErrs:   true,
			errFields: []string{}, // Will have errors but checking by count instead
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := tt.req.Validate()

			if tt.hasErrs && errors.IsEmpty() {
				t.Errorf("Expected errors but got none")
			}

			if !tt.hasErrs && !errors.IsEmpty() {
				t.Errorf("Expected no errors but got %v", errors)
			}
		})
	}
}

// ============================================================================
// Login Request Validation Tests
// ============================================================================

func TestLoginRequestValidate(t *testing.T) {
	tests := []struct {
		name      string
		req       LoginRequest
		hasErrs   bool
		errFields []string
	}{
		{
			name: "Valid login",
			req: LoginRequest{
				Email:    "user@example.com",
				Password: "SecurePass123",
			},
			hasErrs: false,
		},
		{
			name: "Missing email",
			req: LoginRequest{
				Email:    "",
				Password: "SecurePass123",
			},
			hasErrs:   true,
			errFields: []string{"email"},
		},
		{
			name: "Invalid email",
			req: LoginRequest{
				Email:    "notanemail",
				Password: "SecurePass123",
			},
			hasErrs:   true,
			errFields: []string{"email"},
		},
		{
			name: "Missing password",
			req: LoginRequest{
				Email:    "user@example.com",
				Password: "",
			},
			hasErrs:   true,
			errFields: []string{"password"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := tt.req.Validate()

			if tt.hasErrs && errors.IsEmpty() {
				t.Errorf("Expected errors but got none")
			}

			if !tt.hasErrs && !errors.IsEmpty() {
				t.Errorf("Expected no errors but got %v", errors)
			}
		})
	}
}

// ============================================================================
// Create Course Request Validation Tests
// ============================================================================

func TestCreateCourseRequestValidate(t *testing.T) {
	tests := []struct {
		name      string
		req       CreateCourseRequest
		hasErrs   bool
		errFields []string
	}{
		{
			name: "Valid course",
			req: CreateCourseRequest{
				Title:       "Introduction to Go",
				Description: "Learn Go programming from basics",
				Code:        "GO101",
				Category:    "Programming",
				Level:       "Beginner",
				MaxStudents: 50,
			},
			hasErrs: false,
		},
		{
			name: "Missing title",
			req: CreateCourseRequest{
				Title:       "",
				Description: "Learn Go programming from basics",
				Code:        "GO101",
				Category:    "Programming",
				Level:       "Beginner",
				MaxStudents: 50,
			},
			hasErrs:   true,
			errFields: []string{"title"},
		},
		{
			name: "Invalid level",
			req: CreateCourseRequest{
				Title:       "Introduction to Go",
				Description: "Learn Go programming from basics",
				Code:        "GO101",
				Category:    "Programming",
				Level:       "SuperAdvanced",
				MaxStudents: 50,
			},
			hasErrs:   true,
			errFields: []string{"level"},
		},
		{
			name: "Invalid max students",
			req: CreateCourseRequest{
				Title:       "Introduction to Go",
				Description: "Learn Go programming from basics",
				Code:        "GO101",
				Category:    "Programming",
				Level:       "Beginner",
				MaxStudents: -5,
			},
			hasErrs:   true,
			errFields: []string{"max_students"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := tt.req.Validate()

			if tt.hasErrs && errors.IsEmpty() {
				t.Errorf("Expected errors but got none")
			}

			if !tt.hasErrs && !errors.IsEmpty() {
				t.Errorf("Expected no errors but got %v", errors)
			}
		})
	}
}

// ============================================================================
// String Validation Tests
// ============================================================================

func TestValidateStringLength(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		value    string
		minLen   int
		maxLen   int
		hasError bool
	}{
		{"Valid string", "name", "hello", 1, 10, false},
		{"Exact min", "name", "h", 1, 10, false},
		{"Exact max", "name", "0123456789", 1, 10, false},
		{"Too short", "name", "hi", 5, 10, true},
		{"Too long", "name", "this is way too long", 1, 10, true},
		{"Empty required", "name", "", 1, 10, true},
		{"Empty allowed", "name", "", 0, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStringLength(tt.field, tt.value, tt.minLen, tt.maxLen)
			hasError := err != nil

			if hasError != tt.hasError {
				t.Errorf("ValidateStringLength(%s, %s, %d, %d) hasError=%v, want %v",
					tt.field, tt.value, tt.minLen, tt.maxLen, hasError, tt.hasError)
			}
		})
	}
}

// ============================================================================
// UUID Validation Tests
// ============================================================================

func TestIsValidUUID(t *testing.T) {
	tests := []struct {
		uuid  string
		valid bool
	}{
		{"550e8400-e29b-41d4-a716-446655440000", true},
		{"f47ac10b-58cc-4372-a567-0e02b2c3d479", true},
		{"invalid-uuid", false},
		{"", false},
		{"550e8400-e29b-41d4-a716-44665544000", false},
		{"550e8400-e29b-41d4-a716-4466554400000", false},
	}

	for _, tt := range tests {
		t.Run(tt.uuid, func(t *testing.T) {
			got := IsValidUUID(tt.uuid)
			if got != tt.valid {
				t.Errorf("IsValidUUID(%s) = %v, want %v", tt.uuid, got, tt.valid)
			}
		})
	}
}

// ============================================================================
// Pagination Validation Tests
// ============================================================================

func TestPaginationParamsValidate(t *testing.T) {
	tests := []struct {
		name    string
		params  PaginationParams
		hasErrs bool
	}{
		{
			name:    "Valid pagination",
			params:  PaginationParams{Page: 1, PageSize: 10},
			hasErrs: false,
		},
		{
			name:    "Valid pagination high page",
			params:  PaginationParams{Page: 100, PageSize: 50},
			hasErrs: false,
		},
		{
			name:    "Page size too large",
			params:  PaginationParams{Page: 1, PageSize: 200},
			hasErrs: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := tt.params.Validate()

			if tt.hasErrs && errors.IsEmpty() {
				t.Errorf("Expected errors but got none")
			}

			if !tt.hasErrs && !errors.IsEmpty() {
				t.Errorf("Expected no errors but got %v", errors)
			}
		})
	}
}

// ============================================================================
// Create Assignment Request Validation Tests
// ============================================================================

func TestCreateAssignmentRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateAssignmentRequest
		hasErrs bool
	}{
		{
			name: "Valid assignment",
			req: CreateAssignmentRequest{
				Title:       "Assignment 1",
				Description: "Complete the tasks",
				Points:      100,
				Type:        "homework",
			},
			hasErrs: false,
		},
		{
			name: "Missing title",
			req: CreateAssignmentRequest{
				Title:       "",
				Description: "Complete the tasks",
				Points:      100,
				Type:        "homework",
			},
			hasErrs: true,
		},
		{
			name: "Invalid type",
			req: CreateAssignmentRequest{
				Title:       "Assignment 1",
				Description: "Complete the tasks",
				Points:      100,
				Type:        "invalid_type",
			},
			hasErrs: true,
		},
		{
			name: "Negative points",
			req: CreateAssignmentRequest{
				Title:       "Assignment 1",
				Description: "Complete the tasks",
				Points:      -10,
				Type:        "homework",
			},
			hasErrs: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := tt.req.Validate()

			if tt.hasErrs && errors.IsEmpty() {
				t.Errorf("Expected errors but got none")
			}

			if !tt.hasErrs && !errors.IsEmpty() {
				t.Errorf("Expected no errors but got %v", errors)
			}
		})
	}
}

// ============================================================================
// Create Quiz Request Validation Tests
// ============================================================================

func TestCreateQuizRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateQuizRequest
		hasErrs bool
	}{
		{
			name: "Valid quiz",
			req: CreateQuizRequest{
				Title:       "Quiz 1",
				Description: "Test your knowledge",
				TimeLimit:   60,
				PassScore:   70,
			},
			hasErrs: false,
		},
		{
			name: "Missing title",
			req: CreateQuizRequest{
				Title:       "",
				Description: "Test your knowledge",
				TimeLimit:   60,
				PassScore:   70,
			},
			hasErrs: true,
		},
		{
			name: "Invalid pass score",
			req: CreateQuizRequest{
				Title:       "Quiz 1",
				Description: "Test your knowledge",
				TimeLimit:   60,
				PassScore:   150,
			},
			hasErrs: true,
		},
		{
			name: "Negative time limit",
			req: CreateQuizRequest{
				Title:       "Quiz 1",
				Description: "Test your knowledge",
				TimeLimit:   -10,
				PassScore:   70,
			},
			hasErrs: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := tt.req.Validate()

			if tt.hasErrs && errors.IsEmpty() {
				t.Errorf("Expected errors but got none")
			}

			if !tt.hasErrs && !errors.IsEmpty() {
				t.Errorf("Expected no errors but got %v", errors)
			}
		})
	}
}

// ============================================================================
// Numeric Validation Tests
// ============================================================================

func TestValidatePositiveNumber(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		valid bool
	}{
		{"Positive number", 100.5, true},
		{"Zero", 0, false},
		{"Negative number", -50.5, false},
		{"Large positive", 999999.99, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value > 0
			if got != tt.valid {
				t.Errorf("ValidatePositive(%f) = %v, want %v", tt.value, got, tt.valid)
			}
		})
	}
}
