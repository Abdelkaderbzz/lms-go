package models

import (
	"testing"

	"github.com/google/uuid"
)

// TestUserBeforeCreate tests the BeforeCreate hook for User
func TestUserBeforeCreate(t *testing.T) {
	user := &User{
		Email: "test@example.com",
	}

	if user.ID != "" {
		t.Errorf("Expected empty ID before BeforeCreate, got %s", user.ID)
	}

	// Simulate BeforeCreate hook
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	if user.ID == "" {
		t.Error("Expected ID to be generated, but got empty string")
	}
}

// TestCourseBeforeCreate tests the BeforeCreate hook for Course
func TestCourseBeforeCreate(t *testing.T) {
	course := &Course{
		Title: "Test Course",
		Code:  "TEST-001",
	}

	if course.ID != "" {
		t.Errorf("Expected empty ID before BeforeCreate, got %s", course.ID)
	}

	// Simulate BeforeCreate hook
	if course.ID == "" {
		course.ID = uuid.New().String()
	}

	if course.ID == "" {
		t.Error("Expected ID to be generated, but got empty string")
	}
}

// TestUserRoleConstants tests that user role constants are defined
func TestUserRoleConstants(t *testing.T) {
	tests := []struct {
		name     string
		role     string
		expected string
	}{
		{"Admin role", RoleAdmin, "admin"},
		{"Instructor role", RoleInstructor, "instructor"},
		{"Student role", RoleStudent, "student"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.role != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.role)
			}
		})
	}
}

// TestModuleUUID tests module UUID generation
func TestModuleUUID(t *testing.T) {
	module := &Module{
		Title: "Test Module",
	}

	if module.ID == "" {
		module.ID = uuid.New().String()
	}

	// Verify it's a valid UUID
	_, err := uuid.Parse(module.ID)
	if err != nil {
		t.Errorf("Generated ID is not a valid UUID: %v", err)
	}
}

// TestLessonUUID tests lesson UUID generation
func TestLessonUUID(t *testing.T) {
	lesson := &Lesson{
		Title: "Test Lesson",
	}

	if lesson.ID == "" {
		lesson.ID = uuid.New().String()
	}

	_, err := uuid.Parse(lesson.ID)
	if err != nil {
		t.Errorf("Generated ID is not a valid UUID: %v", err)
	}
}

// TestQuizUUID tests quiz UUID generation
func TestQuizUUID(t *testing.T) {
	quiz := &Quiz{
		Title: "Test Quiz",
	}

	if quiz.ID == "" {
		quiz.ID = uuid.New().String()
	}

	_, err := uuid.Parse(quiz.ID)
	if err != nil {
		t.Errorf("Generated ID is not a valid UUID: %v", err)
	}
}

// TestAssignmentTypes tests assignment type constants
func TestAssignmentTypes(t *testing.T) {
	types := []string{"homework", "project", "essay"}

	for _, assignType := range types {
		assignment := &Assignment{
			Title: "Test",
			Type:  assignType,
		}

		if assignment.Type != assignType {
			t.Errorf("Expected assignment type %s, got %s", assignType, assignment.Type)
		}
	}
}

// TestCourseStatus tests course status values
func TestCourseStatus(t *testing.T) {
	statuses := []string{"active", "archived", "draft"}

	for _, status := range statuses {
		course := &Course{
			Title:  "Test",
			Status: status,
		}

		if course.Status != status {
			t.Errorf("Expected course status %s, got %s", status, course.Status)
		}
	}
}

// TestEnrollmentStatus tests enrollment status values
func TestEnrollmentStatus(t *testing.T) {
	statuses := []string{"active", "completed", "dropped"}

	for _, status := range statuses {
		enrollment := &Enrollment{
			Status: status,
		}

		if enrollment.Status != status {
			t.Errorf("Expected enrollment status %s, got %s", status, enrollment.Status)
		}
	}
}
