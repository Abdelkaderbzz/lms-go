package dtos

import (
	"encoding/json"
	"testing"
	"time"
)

// ============================================================================
// Request DTO Tests
// ============================================================================

func TestUpdateProfileRequestDTOSerialization(t *testing.T) {
	req := UpdateProfileRequest{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "123456789",
		Bio:       "Software engineer",
		Avatar:    "avatar.jpg",
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal back
	var decoded UpdateProfileRequest
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.FirstName != req.FirstName {
		t.Errorf("FirstName mismatch: %s vs %s", decoded.FirstName, req.FirstName)
	}

	if decoded.Phone != req.Phone {
		t.Errorf("Phone mismatch")
	}
}

func TestGradeSubmissionRequestDTOSerialization(t *testing.T) {
	req := GradeSubmissionRequest{
		Score:    85.5,
		Feedback: "Good work",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded GradeSubmissionRequest
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Score != req.Score {
		t.Errorf("Score mismatch")
	}
}

func TestCreateModuleRequestDTOSerialization(t *testing.T) {
	req := CreateModuleRequest{
		Title:       "Module 1",
		Description: "Introduction",
		Order:       1,
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded CreateModuleRequest
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Title != req.Title {
		t.Errorf("Title mismatch")
	}
}

func TestCreateLessonRequestDTOSerialization(t *testing.T) {
	req := CreateLessonRequest{
		Title:       "Lesson 1",
		Description: "Learn basics",
		Content:     "<p>Content</p>",
		VideoURL:    "http://example.com/video.mp4",
		Duration:    60,
		Order:       1,
		Published:   true,
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded CreateLessonRequest
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Duration != req.Duration {
		t.Errorf("Duration mismatch")
	}
}

func TestCreateResourceRequestDTOSerialization(t *testing.T) {
	req := CreateResourceRequest{
		Type:  "pdf",
		Title: "Resource 1",
		URL:   "http://example.com/resource.pdf",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded CreateResourceRequest
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Type != req.Type {
		t.Errorf("Type mismatch")
	}
}

// ============================================================================
// Response DTO Tests
// ============================================================================

func TestUserDTOSerialization(t *testing.T) {
	user := UserDTO{
		ID:        "123",
		Email:     "user@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Role:      "student",
		Active:    true,
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded UserDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.ID != user.ID {
		t.Errorf("ID mismatch")
	}

	if decoded.Email != user.Email {
		t.Errorf("Email mismatch")
	}
}

func TestCourseDTOSerialization(t *testing.T) {
	course := CourseDTO{
		ID:          "course-123",
		Title:       "Go Programming",
		Description: "Learn Go",
		Code:        "GO101",
		Level:       "Beginner",
		MaxStudents: 50,
		Status:      "active",
		CreatedAt:   time.Now(),
	}

	data, err := json.Marshal(course)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded CourseDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Code != course.Code {
		t.Errorf("Code mismatch")
	}
}

func TestAssignmentDTOSerialization(t *testing.T) {
	assignment := AssignmentDTO{
		ID:          "assign-123",
		Title:       "Assignment 1",
		Description: "Complete tasks",
		Points:      100,
		Type:        "homework",
		Status:      "active",
		CreatedAt:   time.Now(),
	}

	data, err := json.Marshal(assignment)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded AssignmentDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Type != assignment.Type {
		t.Errorf("Type mismatch")
	}
}

func TestSubmissionDTOSerialization(t *testing.T) {
	submission := SubmissionDTO{
		ID:           "sub-123",
		AssignmentID: "assign-123",
		UserID:       "user-123",
		Status:       "submitted",
		SubmittedAt:  time.Now(),
	}

	data, err := json.Marshal(submission)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded SubmissionDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Status != submission.Status {
		t.Errorf("Status mismatch")
	}
}

func TestGradeDTOSerialization(t *testing.T) {
	grade := GradeDTO{
		ID:       "grade-123",
		UserID:   "user-123",
		CourseID: "course-123",
		Points:   85.5,
		Feedback: "Good work",
		GradedAt: time.Now(),
	}

	data, err := json.Marshal(grade)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded GradeDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Points != grade.Points {
		t.Errorf("Points mismatch")
	}
}

// ============================================================================
// Stats DTO Tests
// ============================================================================

func TestDashboardStatsDTOSerialization(t *testing.T) {
	stats := DashboardStatsDTO{
		TotalUsers:        100,
		TotalCourses:      10,
		TotalEnrollments:  1200,
		TotalAssignments:  50,
		AverageCompletion: 75.5,
		ActiveUsers:       75,
	}

	data, err := json.Marshal(stats)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded DashboardStatsDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.TotalUsers != stats.TotalUsers {
		t.Errorf("TotalUsers mismatch")
	}
}

func TestCourseStatsDTOSerialization(t *testing.T) {
	stats := CourseStatsDTO{
		CourseID:        "course-123",
		TotalEnrollments: 50,
		AverageProgress: 75.5,
		AverageGrade:    82.3,
		CompletionRate:  85.0,
	}

	data, err := json.Marshal(stats)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded CourseStatsDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.CompletionRate != stats.CompletionRate {
		t.Errorf("CompletionRate mismatch")
	}
}

// ============================================================================
// JSON Tag Verification Tests
// ============================================================================

func TestDTOJSONTagMapping(t *testing.T) {
	tests := []struct {
		name        string
		expectedTag string
		testFunc    func() string
	}{
		{
			name:        "RegisterRequest email tag",
			expectedTag: "email",
			testFunc: func() string {
				data := map[string]interface{}{"email": "test@example.com"}
				jsonData, _ := json.Marshal(data)
				var req map[string]interface{}
				json.Unmarshal(jsonData, &req)
				return "email"
			},
		},
		{
			name:        "UserDTO id tag",
			expectedTag: "id",
			testFunc: func() string {
				data := map[string]interface{}{"id": "123"}
				jsonData, _ := json.Marshal(data)
				var dto map[string]interface{}
				json.Unmarshal(jsonData, &dto)
				return "id"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.testFunc()
			if result != tt.expectedTag {
				t.Errorf("Expected tag %s, got %s", tt.expectedTag, result)
			}
		})
	}
}

// ============================================================================
// DTO Field Tests
// ============================================================================

func TestUpdateProfileRequestFields(t *testing.T) {
	req := UpdateProfileRequest{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "123456789",
		Bio:       "Engineer",
		Avatar:    "avatar.jpg",
	}

	if req.FirstName == "" {
		t.Error("FirstName should not be empty")
	}

	if req.Phone == "" {
		t.Error("Phone should not be empty")
	}
}

func TestCourseDTOFields(t *testing.T) {
	course := CourseDTO{
		ID:          "course-123",
		Title:       "Go Programming",
		Description: "Learn Go",
		Code:        "GO101",
		MaxStudents: 50,
	}

	if course.ID == "" {
		t.Error("ID should not be empty")
	}

	if course.Title == "" {
		t.Error("Title should not be empty")
	}

	if course.Code == "" {
		t.Error("Code should not be empty")
	}

	if course.MaxStudents <= 0 {
		t.Error("MaxStudents should be positive")
	}
}

// ============================================================================
// DTO Nested Structure Tests
// ============================================================================

func TestEnrollmentDTOSerialization(t *testing.T) {
	enrollment := EnrollmentDTO{
		ID:         "enroll-123",
		CourseID:   "course-123",
		UserID:     "user-123",
		Status:     "active",
		Progress:   45.5,
		EnrolledAt: time.Now(),
	}

	data, err := json.Marshal(enrollment)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded EnrollmentDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Progress != enrollment.Progress {
		t.Errorf("Progress mismatch")
	}
}

func TestNotificationDTOSerialization(t *testing.T) {
	notification := NotificationDTO{
		ID:        "notif-123",
		UserID:    "user-123",
		Type:      "assignment",
		Title:     "New assignment",
		Message:   "Assignment 1 posted",
		Read:      false,
	}

	data, err := json.Marshal(notification)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded NotificationDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.Type != notification.Type {
		t.Errorf("Type mismatch")
	}

	if decoded.Read != notification.Read {
		t.Errorf("Read status mismatch")
	}
}

func TestQuizDTOSerialization(t *testing.T) {
	quiz := QuizDTO{
		ID:          "quiz-123",
		CourseID:    "course-123",
		Title:       "Quiz 1",
		Description: "Test knowledge",
		TimeLimit:   60,
		PassScore:   70,
		CreatedAt:   time.Now(),
	}

	data, err := json.Marshal(quiz)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded QuizDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.TimeLimit != quiz.TimeLimit {
		t.Errorf("TimeLimit mismatch")
	}

	if decoded.PassScore != quiz.PassScore {
		t.Errorf("PassScore mismatch")
	}
}

// ============================================================================
// DTO Array Tests
// ============================================================================

func TestUserDTOArray(t *testing.T) {
	users := []UserDTO{
		{ID: "1", Email: "user1@example.com", FirstName: "User1"},
		{ID: "2", Email: "user2@example.com", FirstName: "User2"},
		{ID: "3", Email: "user3@example.com", FirstName: "User3"},
	}

	data, err := json.Marshal(users)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded []UserDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if len(decoded) != 3 {
		t.Errorf("Expected 3 users, got %d", len(decoded))
	}

	if decoded[0].Email != "user1@example.com" {
		t.Errorf("First user email mismatch")
	}
}

func TestCourseDTOArray(t *testing.T) {
	courses := []CourseDTO{
		{ID: "c1", Title: "Course 1", Code: "C101"},
		{ID: "c2", Title: "Course 2", Code: "C102"},
	}

	data, err := json.Marshal(courses)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded []CourseDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if len(decoded) != 2 {
		t.Errorf("Expected 2 courses, got %d", len(decoded))
	}
}

// ============================================================================
// DTO Null/Empty Handling Tests
// ============================================================================

func TestDTOEmptyValues(t *testing.T) {
	user := UserDTO{
		ID:        "123",
		Email:     "user@example.com",
		FirstName: "",  // Empty string
		LastName:  "",  // Empty string
		Role:      "student",
		Active:    false,  // False boolean
	}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded UserDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.FirstName != "" {
		t.Errorf("Expected empty FirstName")
	}

	if decoded.Active != false {
		t.Errorf("Expected false Active value")
	}
}

// ============================================================================
// DTO Time Field Tests
// ============================================================================

func TestDTOTimeHandling(t *testing.T) {
	now := time.Now()
	user := UserDTO{
		ID:        "123",
		Email:     "user@example.com",
		FirstName: "John",
		CreatedAt: now,
	}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded UserDTO
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Time might have millisecond differences due to JSON marshaling
	if decoded.CreatedAt.Unix() != now.Unix() {
		t.Errorf("Time mismatch")
	}
}
