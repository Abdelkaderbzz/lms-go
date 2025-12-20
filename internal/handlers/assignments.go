package handlers

import (
	"net/http"
	"time"

	"lms-go/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateAssignment creates a new assignment
func CreateAssignment(c *gin.Context, db *gorm.DB) {
	var req struct {
		CourseID    string    `json:"course_id" binding:"required"`
		Title       string    `json:"title" binding:"required"`
		Description string    `json:"description"`
		DueDate     time.Time `json:"due_date" binding:"required"`
		Points      float64   `json:"points"`
		Type        string    `json:"type"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignment := models.Assignment{
		CourseID:    req.CourseID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Points:      req.Points,
		Type:        req.Type,
		Status:      "active",
	}

	if err := db.Create(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assignment"})
		return
	}

	c.JSON(http.StatusCreated, assignment)
}

// GetCourseAssignments retrieves all assignments for a course
func GetCourseAssignments(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var assignments []models.Assignment
	if err := db.Where("course_id = ?", courseID).Find(&assignments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve assignments"})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

// GetAssignmentByID retrieves a specific assignment
func GetAssignmentByID(c *gin.Context, db *gorm.DB) {
	assignmentID := c.Param("id")

	var assignment models.Assignment
	if err := db.Preload("Submissions").First(&assignment, "id = ?", assignmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}

	c.JSON(http.StatusOK, assignment)
}

// UpdateAssignment updates assignment details
func UpdateAssignment(c *gin.Context, db *gorm.DB) {
	assignmentID := c.Param("id")

	var req struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		DueDate     time.Time `json:"due_date"`
		Points      float64   `json:"points"`
		Status      string    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Assignment{}).Where("id = ?", assignmentID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update assignment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment updated successfully"})
}

// DeleteAssignment deletes an assignment
func DeleteAssignment(c *gin.Context, db *gorm.DB) {
	assignmentID := c.Param("id")

	if err := db.Delete(&models.Assignment{}, "id = ?", assignmentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete assignment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}

// GetAssignmentStats gets assignment statistics
func GetAssignmentStats(c *gin.Context, db *gorm.DB) {
	assignmentID := c.Param("id")

	var stats struct {
		TotalSubmissions int64   `json:"total_submissions"`
		SubmittedCount   int64   `json:"submitted_count"`
		AverageScore     float64 `json:"average_score"`
		HighestScore     float64 `json:"highest_score"`
		LowestScore      float64 `json:"lowest_score"`
	}

	db.Model(&models.Submission{}).Where("assignment_id = ?", assignmentID).Count(&stats.TotalSubmissions)
	db.Model(&models.Submission{}).Where("assignment_id = ? AND status IN ?", assignmentID, []string{"submitted", "graded"}).Count(&stats.SubmittedCount)

	c.JSON(http.StatusOK, stats)
}
