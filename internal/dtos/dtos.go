package dtos

import "time"

// UserDTO represents user data transfer object
type UserDTO struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone"`
	Bio       string    `json:"bio"`
	Avatar    string    `json:"avatar"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CourseDTO represents course data transfer object
type CourseDTO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Code        string    `json:"code"`
	Category    string    `json:"category"`
	Level       string    `json:"level"`
	Status      string    `json:"status"`
	MaxStudents int       `json:"max_students"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// AssignmentDTO represents assignment data transfer object
type AssignmentDTO struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Points      float64   `json:"points"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SubmissionDTO represents submission data transfer object
type SubmissionDTO struct {
	ID           string    `json:"id"`
	AssignmentID string    `json:"assignment_id"`
	UserID       string    `json:"user_id"`
	Content      string    `json:"content"`
	FileURL      string    `json:"file_url"`
	SubmittedAt  time.Time `json:"submitted_at"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GradeDTO represents grade data transfer object
type GradeDTO struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	CourseID     string    `json:"course_id"`
	AssignmentID string    `json:"assignment_id"`
	Points       float64   `json:"points"`
	Feedback     string    `json:"feedback"`
	GradedAt     time.Time `json:"graded_at"`
}

// QuizDTO represents quiz data transfer object
type QuizDTO struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TimeLimit   int       `json:"time_limit"`
	PassScore   float64   `json:"pass_score"`
	Public      bool      `json:"public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// EnrollmentDTO represents enrollment data transfer object
type EnrollmentDTO struct {
	ID        string    `json:"id"`
	CourseID  string    `json:"course_id"`
	UserID    string    `json:"user_id"`
	Status    string    `json:"status"`
	Progress  float64   `json:"progress"`
	EnrolledAt time.Time `json:"enrolled_at"`
}

// NotificationDTO represents notification data transfer object
type NotificationDTO struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

// UpdateProfileRequest represents profile update request
type UpdateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Bio       string `json:"bio"`
	Avatar    string `json:"avatar"`
}

// GradeSubmissionRequest represents grade submission request
type GradeSubmissionRequest struct {
	Score    float64 `json:"score" binding:"required"`
	Feedback string  `json:"feedback"`
}

// CreateModuleRequest represents module creation request
type CreateModuleRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}

// CreateLessonRequest represents lesson creation request
type CreateLessonRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Content     string `json:"content"`
	VideoURL    string `json:"video_url"`
	Duration    int    `json:"duration"`
	Order       int    `json:"order"`
	Published   bool   `json:"published"`
}

// CreateResourceRequest represents resource creation request
type CreateResourceRequest struct {
	Type  string `json:"type" binding:"required"` // pdf, document, link, video
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required"`
}

// CreateAnnouncementRequest represents announcement creation request
type CreateAnnouncementRequest struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Important bool   `json:"important"`
}

// CreateDiscussionRequest represents discussion creation request
type CreateDiscussionRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateForumPostRequest represents forum post creation request
type CreateForumPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateForumReplyRequest represents forum reply creation request
type CreateForumReplyRequest struct {
	Content string `json:"content" binding:"required"`
}

// SubmitAssignmentRequest represents assignment submission request
type SubmitAssignmentRequest struct {
	Content string `json:"content"`
	FileURL string `json:"file_url"`
}

// CreateQuestionRequest represents question creation request
type CreateQuestionRequest struct {
	Type      string        `json:"type" binding:"required"`
	Question  string        `json:"question" binding:"required"`
	Points    float64       `json:"points" binding:"required"`
	Order     int           `json:"order"`
	Options   []OptionInput `json:"options"`
}

// OptionInput represents option input for questions
type OptionInput struct {
	Text      string `json:"text" binding:"required"`
	IsCorrect bool   `json:"is_correct"`
	Order     int    `json:"order"`
}

// SubmitQuizRequest represents quiz submission request
type SubmitQuizRequest struct {
	Answers []QuizAnswerInput `json:"answers" binding:"required"`
}

// QuizAnswerInput represents an answer input for quiz
type QuizAnswerInput struct {
	QuestionID string `json:"question_id" binding:"required"`
	SelectedID string `json:"selected_id"`
	TextAnswer string `json:"text_answer"`
}

// FilterParams represents common filter parameters
type FilterParams struct {
	Search   string `json:"search"`
	Status   string `json:"status"`
	Category string `json:"category"`
	Level    string `json:"level"`
	Sort     string `json:"sort"`
	Order    string `json:"order"` // asc or desc
}

// DashboardStatsDTO represents dashboard statistics
type DashboardStatsDTO struct {
	TotalUsers        int64   `json:"total_users"`
	TotalCourses      int64   `json:"total_courses"`
	TotalEnrollments  int64   `json:"total_enrollments"`
	TotalAssignments  int64   `json:"total_assignments"`
	AverageCompletion float64 `json:"average_completion"`
	ActiveCourses     int64   `json:"active_courses"`
	ActiveUsers       int64   `json:"active_users"`
}

// CourseStatsDTO represents course statistics
type CourseStatsDTO struct {
	CourseID          string  `json:"course_id"`
	TotalEnrollments  int64   `json:"total_enrollments"`
	TotalCompletions  int64   `json:"total_completions"`
	AverageProgress   float64 `json:"average_progress"`
	AverageGrade      float64 `json:"average_grade"`
	CompletionRate    float64 `json:"completion_rate"`
	LastUpdated       time.Time `json:"last_updated"`
}
