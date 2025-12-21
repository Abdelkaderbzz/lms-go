package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"lms-go/internal/database"
	"lms-go/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TestDB holds the test database connection
type TestDB struct {
	DB *gorm.DB
}

// SetupTestDB initializes a test database and runs migrations
func SetupTestDB(t *testing.T) *gorm.DB {
	// Use test database URL if provided, otherwise use default
	testDSN := os.Getenv("TEST_DATABASE_URL")
	if testDSN == "" {
		testDSN = "postgres://user:password@localhost:5432/lms_test_db"
	}

	db, err := gorm.Open(postgres.Open(testDSN), &gorm.Config{})
	if err != nil {
		t.Logf("Warning: Could not connect to test database: %v. Using in-memory fallback.", err)
		// Fallback to SQLite for testing if postgres is not available
		return nil
	}

	// Run migrations
	err = database.RunMigrations(db)
	if err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	// Clear all data before test
	db.Exec("TRUNCATE TABLE users CASCADE")

	return db
}

// TeardownTestDB cleans up the test database
func TeardownTestDB(t *testing.T, db *gorm.DB) {
	if db == nil {
		return
	}

	// Clean up all tables
	tables := []string{
		"notifications",
		"forum_replies",
		"forum_posts",
		"discussions",
		"announcements",
		"lesson_progress",
		"certificates",
		"quiz_answers",
		"quiz_attempts",
		"options",
		"questions",
		"comments",
		"grades",
		"submissions",
		"resources",
		"lessons",
		"modules",
		"assignments",
		"quizzes",
		"enrollments",
		"course_enrollments",
		"course_instructors",
		"courses",
		"users",
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)).Error; err != nil {
			t.Logf("Warning: Could not truncate table %s: %v", table, err)
		}
	}
}

// CreateTestUser creates a test user
func CreateTestUser(db *gorm.DB, email, password string, role string) *models.User {
	user := &models.User{
		Email:     email,
		Password:  password,
		FirstName: "Test",
		LastName:  "User",
		Role:      role,
		Phone:     "+1234567890",
		Active:    true,
	}

	if err := db.Create(user).Error; err != nil {
		log.Printf("Error creating test user: %v", err)
		return nil
	}

	return user
}

// CreateTestCourse creates a test course
func CreateTestCourse(db *gorm.DB, title, code string, instructor *models.User) *models.Course {
	course := &models.Course{
		Title:       title,
		Description: "Test course",
		Code:        code,
		Category:    "Test",
		Level:       "Beginner",
		Status:      "active",
		MaxStudents: 50,
	}

	if err := db.Create(course).Error; err != nil {
		log.Printf("Error creating test course: %v", err)
		return nil
	}

	if instructor != nil {
		if err := db.Model(course).Association("Instructors").Append(instructor); err != nil {
			log.Printf("Warning: Failed to add instructor to course: %v", err)
		}
	}

	return course
}

// CreateTestModule creates a test module
func CreateTestModule(db *gorm.DB, courseID, title string) *models.Module {
	module := &models.Module{
		CourseID:    courseID,
		Title:       title,
		Description: "Test module",
		Order:       1,
	}

	if err := db.Create(module).Error; err != nil {
		log.Printf("Error creating test module: %v", err)
		return nil
	}

	return module
}

// CreateTestLesson creates a test lesson
func CreateTestLesson(db *gorm.DB, moduleID, title string) *models.Lesson {
	lesson := &models.Lesson{
		ModuleID:    moduleID,
		Title:       title,
		Description: "Test lesson",
		Content:     "<p>Test content</p>",
		Duration:    30,
		Order:       1,
		Published:   true,
	}

	if err := db.Create(lesson).Error; err != nil {
		log.Printf("Error creating test lesson: %v", err)
		return nil
	}

	return lesson
}

// CreateTestAssignment creates a test assignment
func CreateTestAssignment(db *gorm.DB, courseID, title string) *models.Assignment {
	assignment := &models.Assignment{
		CourseID:    courseID,
		Title:       title,
		Description: "Test assignment",
		Points:      100,
		Type:        "homework",
		Status:      "active",
	}

	if err := db.Create(assignment).Error; err != nil {
		log.Printf("Error creating test assignment: %v", err)
		return nil
	}

	return assignment
}

// CreateTestSubmission creates a test submission
func CreateTestSubmission(db *gorm.DB, assignmentID, userID string) *models.Submission {
	submission := &models.Submission{
		AssignmentID: assignmentID,
		UserID:       userID,
		Content:      "Test submission",
		Status:       "submitted",
	}

	if err := db.Create(submission).Error; err != nil {
		log.Printf("Error creating test submission: %v", err)
		return nil
	}

	return submission
}

// CreateTestQuiz creates a test quiz
func CreateTestQuiz(db *gorm.DB, courseID, creatorID, title string) *models.Quiz {
	quiz := &models.Quiz{
		CourseID:    courseID,
		CreatorID:   creatorID,
		Title:       title,
		Description: "Test quiz",
		TimeLimit:   30,
		PassScore:   70,
		Public:      true,
	}

	if err := db.Create(quiz).Error; err != nil {
		log.Printf("Error creating test quiz: %v", err)
		return nil
	}

	return quiz
}

// CreateTestEnrollment creates a test enrollment
func CreateTestEnrollment(db *gorm.DB, courseID, userID string) *models.Enrollment {
	enrollment := &models.Enrollment{
		CourseID: courseID,
		UserID:   userID,
		Status:   "active",
		Progress: 0,
	}

	if err := db.Create(enrollment).Error; err != nil {
		log.Printf("Error creating test enrollment: %v", err)
		return nil
	}

	return enrollment
}
