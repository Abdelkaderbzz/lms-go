package services

import (
	"errors"

	"lms-go/internal/models"
	"lms-go/internal/validators"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService handles user business logic
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new user service
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// RegisterUser creates a new user with validation
func (s *UserService) RegisterUser(req validators.RegisterRequest) (*models.User, validators.ValidationErrors) {
	// Validate request
	if errs := req.Validate(); !errs.IsEmpty() {
		return nil, errs
	}

	// Check if email already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, validators.ValidationErrors{
			{
				Field:   "email",
				Message: "email already in use",
			},
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, validators.ValidationErrors{
			{
				Field:   "password",
				Message: "failed to process password",
			},
		}
	}

	// Create user
	user := &models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
		Active:    true,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, validators.ValidationErrors{
			{
				Field:   "email",
				Message: "failed to create user",
			},
		}
	}

	return user, nil
}

// GetUserByID retrieves a user by ID with caching support
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	if !validators.IsValidUUID(id) {
		return nil, errors.New("invalid user ID format")
	}

	var user models.User
	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUsersByRole retrieves users by role with pagination
func (s *UserService) GetUsersByRole(role string, page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := s.db.Where("role = ?", role)

	// Count total
	if err := query.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated data
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateUserProfile updates user profile information
func (s *UserService) UpdateUserProfile(userID string, updates map[string]interface{}) (*models.User, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Validate and sanitize updates
	allowedFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"phone":      true,
		"bio":        true,
		"avatar":     true,
	}

	validUpdates := make(map[string]interface{})
	for key, val := range updates {
		if allowedFields[key] {
			validUpdates[key] = val
		}
	}

	if err := s.db.Model(user).Updates(validUpdates).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser safely deletes a user
func (s *UserService) DeleteUser(userID string) error {
	if !validators.IsValidUUID(userID) {
		return errors.New("invalid user ID format")
	}

	// Check if user exists
	user, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Prevent deleting admin users (optional business logic)
	if user.Role == models.RoleAdmin {
		return errors.New("cannot delete admin users")
	}

	// Soft delete or hard delete based on your preference
	return s.db.Delete(user).Error
}

// VerifyPassword checks if password matches
func (s *UserService) VerifyPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}

// CourseService handles course business logic
type CourseService struct {
	db *gorm.DB
}

// NewCourseService creates a new course service
func NewCourseService(db *gorm.DB) *CourseService {
	return &CourseService{db: db}
}

// CreateCourse creates a new course with validation
func (s *CourseService) CreateCourse(req validators.CreateCourseRequest) (*models.Course, validators.ValidationErrors) {
	// Validate request
	if errs := req.Validate(); !errs.IsEmpty() {
		return nil, errs
	}

	// Check if course code already exists
	var existing models.Course
	if err := s.db.Where("code = ?", req.Code).First(&existing).Error; err == nil {
		return nil, validators.ValidationErrors{
			{
				Field:   "code",
				Message: "course code already exists",
			},
		}
	}

	course := &models.Course{
		Title:       req.Title,
		Description: req.Description,
		Code:        req.Code,
		Category:    req.Category,
		Level:       req.Level,
		MaxStudents: req.MaxStudents,
		Status:      "active",
	}

	if err := s.db.Create(course).Error; err != nil {
		return nil, validators.ValidationErrors{
			{
				Field:   "code",
				Message: "failed to create course",
			},
		}
	}

	return course, nil
}

// GetCourseByID retrieves a course by ID
func (s *CourseService) GetCourseByID(courseID string) (*models.Course, error) {
	if !validators.IsValidUUID(courseID) {
		return nil, errors.New("invalid course ID format")
	}

	var course models.Course
	if err := s.db.Preload("Instructors").Preload("Students").
		First(&course, "id = ?", courseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("course not found")
		}
		return nil, err
	}

	return &course, nil
}

// GetCoursesByLevel retrieves courses by level with pagination
func (s *CourseService) GetCoursesByLevel(level string, page, pageSize int) ([]models.Course, int64, error) {
	var courses []models.Course
	var total int64

	query := s.db.Where("level = ? AND status = ?", level, "active")

	// Count total
	if err := query.Model(&models.Course{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated data
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&courses).Error; err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}

// EnrollStudent enrolls a student in a course
func (s *CourseService) EnrollStudent(courseID, userID string) (*models.Enrollment, error) {
	// Validate IDs
	if !validators.IsValidUUID(courseID) || !validators.IsValidUUID(userID) {
		return nil, errors.New("invalid ID format")
	}

	// Check if course exists
	course, err := s.GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	// Check current enrollment count
	var enrollmentCount int64
	s.db.Model(&models.Enrollment{}).
		Where("course_id = ? AND status = ?", courseID, "active").
		Count(&enrollmentCount)

	if int(enrollmentCount) >= course.MaxStudents {
		return nil, errors.New("course is full")
	}

	// Check if already enrolled
	var existing models.Enrollment
	if err := s.db.Where("course_id = ? AND user_id = ?", courseID, userID).
		First(&existing).Error; err == nil {
		return nil, errors.New("student already enrolled in this course")
	}

	enrollment := &models.Enrollment{
		CourseID: courseID,
		UserID:   userID,
		Status:   "active",
		Progress: 0,
	}

	if err := s.db.Create(enrollment).Error; err != nil {
		return nil, err
	}

	return enrollment, nil
}

// GetEnrollmentStatus gets enrollment status for a student in a course
func (s *CourseService) GetEnrollmentStatus(courseID, userID string) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	if err := s.db.Where("course_id = ? AND user_id = ?", courseID, userID).
		First(&enrollment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("enrollment not found")
		}
		return nil, err
	}

	return &enrollment, nil
}

// AssignmentService handles assignment business logic
type AssignmentService struct {
	db *gorm.DB
}

// NewAssignmentService creates a new assignment service
func NewAssignmentService(db *gorm.DB) *AssignmentService {
	return &AssignmentService{db: db}
}

// CreateAssignment creates a new assignment with validation
func (s *AssignmentService) CreateAssignment(courseID string, req validators.CreateAssignmentRequest) (*models.Assignment, validators.ValidationErrors) {
	// Validate request
	if errs := req.Validate(); !errs.IsEmpty() {
		return nil, errs
	}

	assignment := &models.Assignment{
		CourseID:    courseID,
		Title:       req.Title,
		Description: req.Description,
		Points:      req.Points,
		Type:        req.Type,
		Status:      "active",
	}

	if err := s.db.Create(assignment).Error; err != nil {
		return nil, validators.ValidationErrors{
			{
				Field:   "title",
				Message: "failed to create assignment",
			},
		}
	}

	return assignment, nil
}

// GetAssignmentByID retrieves an assignment
func (s *AssignmentService) GetAssignmentByID(assignmentID string) (*models.Assignment, error) {
	if !validators.IsValidUUID(assignmentID) {
		return nil, errors.New("invalid assignment ID format")
	}

	var assignment models.Assignment
	if err := s.db.Preload("Submissions").First(&assignment, "id = ?", assignmentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("assignment not found")
		}
		return nil, err
	}

	return &assignment, nil
}

// GetCourseAssignments retrieves all assignments for a course
func (s *AssignmentService) GetCourseAssignments(courseID string, page, pageSize int) ([]models.Assignment, int64, error) {
	var assignments []models.Assignment
	var total int64

	query := s.db.Where("course_id = ?", courseID)

	if err := query.Model(&models.Assignment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Preload("Submissions").Find(&assignments).Error; err != nil {
		return nil, 0, err
	}

	return assignments, total, nil
}

// QuizService handles quiz business logic
type QuizService struct {
	db *gorm.DB
}

// NewQuizService creates a new quiz service
func NewQuizService(db *gorm.DB) *QuizService {
	return &QuizService{db: db}
}

// CreateQuiz creates a new quiz with validation
func (s *QuizService) CreateQuiz(courseID, creatorID string, req validators.CreateQuizRequest) (*models.Quiz, validators.ValidationErrors) {
	// Validate request
	if errs := req.Validate(); !errs.IsEmpty() {
		return nil, errs
	}

	quiz := &models.Quiz{
		CourseID:  courseID,
		CreatorID: creatorID,
		Title:     req.Title,
		Description: req.Description,
		TimeLimit: req.TimeLimit,
		PassScore: req.PassScore,
		Public:    true,
	}

	if err := s.db.Create(quiz).Error; err != nil {
		return nil, validators.ValidationErrors{
			{
				Field:   "title",
				Message: "failed to create quiz",
			},
		}
	}

	return quiz, nil
}

// GetQuizByID retrieves a quiz
func (s *QuizService) GetQuizByID(quizID string) (*models.Quiz, error) {
	if !validators.IsValidUUID(quizID) {
		return nil, errors.New("invalid quiz ID format")
	}

	var quiz models.Quiz
	if err := s.db.Preload("Questions.Options").First(&quiz, "id = ?", quizID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("quiz not found")
		}
		return nil, err
	}

	return &quiz, nil
}

// StartQuizAttempt creates a new quiz attempt
func (s *QuizService) StartQuizAttempt(quizID, userID string) (*models.QuizAttempt, error) {
	if !validators.IsValidUUID(quizID) || !validators.IsValidUUID(userID) {
		return nil, errors.New("invalid ID format")
	}

	// Check if quiz exists
	_, err := s.GetQuizByID(quizID)
	if err != nil {
		return nil, err
	}

	attempt := &models.QuizAttempt{
		QuizID:    quizID,
		UserID:    userID,
		Status:    "in_progress",
	}

	if err := s.db.Create(attempt).Error; err != nil {
		return nil, err
	}

	return attempt, nil
}

// GetQuizStats returns quiz statistics
func (s *QuizService) GetQuizStats(quizID string) (map[string]interface{}, error) {
	if !validators.IsValidUUID(quizID) {
		return nil, errors.New("invalid quiz ID format")
	}

	var total, submitted int64
	var average float64

	// Count total attempts
	s.db.Model(&models.QuizAttempt{}).
		Where("quiz_id = ?", quizID).
		Count(&total)

	// Count submitted attempts
	s.db.Model(&models.QuizAttempt{}).
		Where("quiz_id = ? AND status = ?", quizID, "submitted").
		Count(&submitted)

	// Calculate average score
	s.db.Model(&models.QuizAttempt{}).
		Where("quiz_id = ? AND status = ?", quizID, "submitted").
		Select("COALESCE(AVG(score), 0)").
		Row().Scan(&average)

	var submissionRate float64
	if total > 0 {
		submissionRate = (float64(submitted) / float64(total)) * 100
	}

	return map[string]interface{}{
		"total_attempts":    total,
		"submitted":         submitted,
		"average_score":     average,
		"submission_rate":   submissionRate,
	}, nil
}
