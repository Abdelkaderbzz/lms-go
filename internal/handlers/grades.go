package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lms-go/internal/models"
)

// GetCourseGrades retrieves all grades for a course
func GetCourseGrades(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var grades []models.Grade
	if err := db.Where("course_id = ?", courseID).
		Preload("User").Preload("Assignment").Find(&grades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve grades"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

// GetStudentGrades retrieves grades for current student
func GetStudentGrades(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")
	userID := c.GetString("user_id")

	var grades []models.Grade
	if err := db.Where("course_id = ? AND user_id = ?", courseID, userID).
		Preload("Assignment").Find(&grades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve grades"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

// GetNotifications retrieves notifications for current user
func GetNotifications(c *gin.Context, db *gorm.DB) {
	userID := c.GetString("user_id")

	var notifications []models.Notification
	if err := db.Where("user_id = ?", userID).
		Order("created_at DESC").Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// MarkNotificationRead marks a notification as read
func MarkNotificationRead(c *gin.Context, db *gorm.DB) {
	notificationID := c.Param("id")

	if err := db.Model(&models.Notification{}).Where("id = ?", notificationID).
		Update("read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

// GetCertificate retrieves a course certificate for a student
func GetCertificate(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")
	userID := c.GetString("user_id")

	var certificate models.Certificate
	if err := db.Where("course_id = ? AND user_id = ?", courseID, userID).
		First(&certificate).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}

	c.JSON(http.StatusOK, certificate)
}
