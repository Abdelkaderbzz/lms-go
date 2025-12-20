package handlers

import (
	"net/http"
	"time"

	"lms-go/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SubmitAssignment submits an assignment
func SubmitAssignment(c *gin.Context, db *gorm.DB) {
	assignmentID := c.Param("assignmentID")
	userID := c.GetString("user_id")

	var req struct {
		Content string `json:"content"`
		FileURL string `json:"file_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submission := models.Submission{
		AssignmentID: assignmentID,
		UserID:       userID,
		Content:      req.Content,
		FileURL:      req.FileURL,
		SubmittedAt:  time.Now(),
		Status:       "submitted",
	}

	if err := db.Create(&submission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit assignment"})
		return
	}

	c.JSON(http.StatusCreated, submission)
}

// GetAssignmentSubmissions retrieves all submissions for an assignment
func GetAssignmentSubmissions(c *gin.Context, db *gorm.DB) {
	assignmentID := c.Param("assignmentID")

	var submissions []models.Submission
	if err := db.Where("assignment_id = ?", assignmentID).
		Preload("User").Preload("Grade").Find(&submissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve submissions"})
		return
	}

	c.JSON(http.StatusOK, submissions)
}

// GetStudentSubmission retrieves a student's submission
func GetStudentSubmission(c *gin.Context, db *gorm.DB) {
	assignmentID := c.Param("assignmentID")
	userID := c.GetString("user_id")

	var submission models.Submission
	if err := db.Where("assignment_id = ? AND user_id = ?", assignmentID, userID).
		Preload("Comments").Preload("Grade").First(&submission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	c.JSON(http.StatusOK, submission)
}

// UpdateSubmission updates a submission
func UpdateSubmission(c *gin.Context, db *gorm.DB) {
	submissionID := c.Param("id")
	userID := c.GetString("user_id")

	var req struct {
		Content string `json:"content"`
		FileURL string `json:"file_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify ownership
	var submission models.Submission
	if err := db.First(&submission, "id = ?", submissionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	if submission.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot update this submission"})
		return
	}

	if err := db.Model(&submission).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update submission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission updated successfully"})
}

// GradeSubmission grades a submission
func GradeSubmission(c *gin.Context, db *gorm.DB) {
	submissionID := c.Param("id")
	userID := c.GetString("user_id") // Instructor/Admin grading

	var req struct {
		Points   float64 `json:"points" binding:"required"`
		Feedback string  `json:"feedback"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var submission models.Submission
	if err := db.First(&submission, "id = ?", submissionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	// Create grade record
	grade := models.Grade{
		UserID:       submission.UserID,
		AssignmentID: submission.AssignmentID,
		SubmissionID: submissionID,
		Points:       req.Points,
		Feedback:     req.Feedback,
		GradedBy:     userID,
		GradedAt:     time.Now(),
	}

	if err := db.Create(&grade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to grade submission"})
		return
	}

	// Update submission status
	db.Model(&submission).Update("status", "graded")

	c.JSON(http.StatusCreated, grade)
}

// AddComment adds a comment to a submission
func AddComment(c *gin.Context, db *gorm.DB) {
	submissionID := c.Param("submissionID")
	userID := c.GetString("user_id")

	var req struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{
		SubmissionID: submissionID,
		UserID:       userID,
		Content:      req.Content,
	}

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetSubmissionComments retrieves comments for a submission
func GetSubmissionComments(c *gin.Context, db *gorm.DB) {
	submissionID := c.Param("submissionID")

	var comments []models.Comment
	if err := db.Where("submission_id = ?", submissionID).
		Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// DeleteSubmission deletes a submission
func DeleteSubmission(c *gin.Context, db *gorm.DB) {
	submissionID := c.Param("id")

	if err := db.Delete(&models.Submission{}, "id = ?", submissionID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete submission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission deleted successfully"})
}
