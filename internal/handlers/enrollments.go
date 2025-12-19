package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lms-go/internal/models"
)

// EnrollStudent enrolls a student in a course
func EnrollStudent(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")
	userID := c.GetString("user_id")

	// Check if already enrolled
	var existing models.Enrollment
	if err := db.Where("course_id = ? AND user_id = ?", courseID, userID).
		First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Already enrolled in this course"})
		return
	}

	enrollment := models.Enrollment{
		CourseID:   courseID,
		UserID:     userID,
		Status:     "active",
		Progress:   0,
		EnrolledAt: time.Now(),
	}

	if err := db.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll in course"})
		return
	}

	c.JSON(http.StatusCreated, enrollment)
}

// UnenrollStudent removes a student from a course
func UnenrollStudent(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")
	userID := c.GetString("user_id")

	if err := db.Where("course_id = ? AND user_id = ?", courseID, userID).
		Delete(&models.Enrollment{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unenroll from course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully unenrolled from course"})
}

// GetCourseEnrollments retrieves all enrollments for a course
func GetCourseEnrollments(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var enrollments []models.Enrollment
	if err := db.Where("course_id = ?", courseID).
		Preload("User").Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve enrollments"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

// GetEnrollmentStatus gets enrollment status for current student
func GetEnrollmentStatus(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")
	userID := c.GetString("user_id")

	var enrollment models.Enrollment
	if err := db.Where("course_id = ? AND user_id = ?", courseID, userID).
		First(&enrollment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not enrolled"})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

// UpdateEnrollmentStatus updates enrollment status (admin/instructor only)
func UpdateEnrollmentStatus(c *gin.Context, db *gorm.DB) {
	enrollmentID := c.Param("id")

	var req struct {
		Status   string  `json:"status"`
		Progress float64 `json:"progress"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Enrollment{}).Where("id = ?", enrollmentID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment updated successfully"})
}

// GetEnrollmentStats gets course enrollment statistics
func GetEnrollmentStats(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var stats struct {
		Total     int64   `json:"total"`
		Active    int64   `json:"active"`
		Completed int64   `json:"completed"`
		Dropped   int64   `json:"dropped"`
		Average   float64 `json:"average_progress"`
	}

	db.Model(&models.Enrollment{}).Where("course_id = ?", courseID).Count(&stats.Total)
	db.Model(&models.Enrollment{}).Where("course_id = ? AND status = ?", courseID, "active").Count(&stats.Active)
	db.Model(&models.Enrollment{}).Where("course_id = ? AND status = ?", courseID, "completed").Count(&stats.Completed)
	db.Model(&models.Enrollment{}).Where("course_id = ? AND status = ?", courseID, "dropped").Count(&stats.Dropped)

	c.JSON(http.StatusOK, stats)
}
