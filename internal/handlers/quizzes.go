package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lms-go/internal/models"
)

// CreateQuiz creates a new quiz
func CreateQuiz(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")
	userID := c.GetString("user_id")

	var req struct {
		Title       string    `json:"title" binding:"required"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
		TimeLimit   int       `json:"time_limit"`
		PassScore   float64   `json:"pass_score"`
		Shuffle     bool      `json:"shuffle"`
		Public      bool      `json:"public"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quiz := models.Quiz{
		CourseID:    courseID,
		CreatorID:   userID,
		Title:       req.Title,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		TimeLimit:   req.TimeLimit,
		PassScore:   req.PassScore,
		Shuffle:     req.Shuffle,
		Public:      req.Public,
	}

	if err := db.Create(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
		return
	}

	c.JSON(http.StatusCreated, quiz)
}

// GetCourseQuizzes retrieves all quizzes for a course
func GetCourseQuizzes(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var quizzes []models.Quiz
	if err := db.Where("course_id = ?", courseID).
		Preload("Creator").Find(&quizzes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve quizzes"})
		return
	}

	c.JSON(http.StatusOK, quizzes)
}

// GetQuizByID retrieves a specific quiz
func GetQuizByID(c *gin.Context, db *gorm.DB) {
	quizID := c.Param("id")

	var quiz models.Quiz
	if err := db.Preload("Questions.Options").Preload("Creator").
		First(&quiz, "id = ?", quizID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// UpdateQuiz updates quiz details
func UpdateQuiz(c *gin.Context, db *gorm.DB) {
	quizID := c.Param("id")

	var req struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
		TimeLimit   int       `json:"time_limit"`
		PassScore   float64   `json:"pass_score"`
		Shuffle     bool      `json:"shuffle"`
		Public      bool      `json:"public"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Quiz{}).Where("id = ?", quizID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quiz"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz updated successfully"})
}

// DeleteQuiz deletes a quiz
func DeleteQuiz(c *gin.Context, db *gorm.DB) {
	quizID := c.Param("id")

	if err := db.Delete(&models.Quiz{}, "id = ?", quizID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete quiz"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz deleted successfully"})
}

// CreateQuestion creates a quiz question
func CreateQuestion(c *gin.Context, db *gorm.DB) {
	quizID := c.Param("quizID")

	var req struct {
		Type    string  `json:"type" binding:"required"`
		Question string `json:"question" binding:"required"`
		Points  float64 `json:"points"`
		Order   int     `json:"order"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question := models.Question{
		QuizID:   quizID,
		Type:     req.Type,
		Question: req.Question,
		Points:   req.Points,
		Order:    req.Order,
	}

	if err := db.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	c.JSON(http.StatusCreated, question)
}

// GetQuizQuestions retrieves all questions for a quiz
func GetQuizQuestions(c *gin.Context, db *gorm.DB) {
	quizID := c.Param("quizID")

	var questions []models.Question
	if err := db.Where("quiz_id = ?", quizID).
		Order("order").Preload("Options").Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve questions"})
		return
	}

	c.JSON(http.StatusOK, questions)
}

// UpdateQuestion updates a question
func UpdateQuestion(c *gin.Context, db *gorm.DB) {
	questionID := c.Param("id")

	var req struct {
		Type     string  `json:"type"`
		Question string  `json:"question"`
		Points   float64 `json:"points"`
		Order    int     `json:"order"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Question{}).Where("id = ?", questionID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question updated successfully"})
}

// DeleteQuestion deletes a question
func DeleteQuestion(c *gin.Context, db *gorm.DB) {
	questionID := c.Param("id")

	if err := db.Delete(&models.Question{}, "id = ?", questionID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}

// StartQuizAttempt starts a new quiz attempt
func StartQuizAttempt(c *gin.Context, db *gorm.DB) {
	quizID := c.Param("quizID")
	userID := c.GetString("user_id")

	attempt := models.QuizAttempt{
		QuizID:    quizID,
		UserID:    userID,
		StartedAt: time.Now(),
		Status:    "in_progress",
	}

	if err := db.Create(&attempt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start quiz attempt"})
		return
	}

	c.JSON(http.StatusCreated, attempt)
}

// SubmitQuizAttempt submits a quiz attempt
func SubmitQuizAttempt(c *gin.Context, db *gorm.DB) {
	attemptID := c.Param("attemptID")

	var req struct {
		Answers []struct {
			QuestionID string `json:"question_id"`
			SelectedID string `json:"selected_id"`
			TextAnswer string `json:"text_answer"`
		} `json:"answers"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	if err := db.Model(&models.QuizAttempt{}).Where("id = ?", attemptID).
		Updates(map[string]interface{}{"ended_at": now, "status": "submitted"}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit quiz"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz submitted successfully"})
}

// GetStudentAttempts retrieves all attempts for a student
func GetStudentAttempts(c *gin.Context, db *gorm.DB) {
	quizID := c.Param("quizID")
	userID := c.GetString("user_id")

	var attempts []models.QuizAttempt
	if err := db.Where("quiz_id = ? AND user_id = ?", quizID, userID).
		Find(&attempts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attempts"})
		return
	}

	c.JSON(http.StatusOK, attempts)
}

// GetAttemptDetail retrieves attempt details
func GetAttemptDetail(c *gin.Context, db *gorm.DB) {
	attemptID := c.Param("id")

	var attempt models.QuizAttempt
	if err := db.Preload("Answers").Preload("Quiz").
		First(&attempt, "id = ?", attemptID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attempt not found"})
		return
	}

	c.JSON(http.StatusOK, attempt)
}
