package handlers

import (
	"net/http"

	"lms-go/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCourse creates a new course (instructor/admin only)
func CreateCourse(c *gin.Context, db *gorm.DB) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Code        string `json:"code" binding:"required"`
		Category    string `json:"category"`
		Level       string `json:"level"`
		Thumbnail   string `json:"thumbnail"`
		MaxStudents int    `json:"max_students"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_ = c.GetString("user_id") // Get current user ID for audit purposes

	course := models.Course{
		Title:       req.Title,
		Description: req.Description,
		Code:        req.Code,
		Category:    req.Category,
		Level:       req.Level,
		Thumbnail:   req.Thumbnail,
		MaxStudents: req.MaxStudents,
		Status:      "draft",
	}

	if err := db.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	// Add instructor to course
	db.Model(&course).Association("Instructors").Append(&models.User{})

	c.JSON(http.StatusCreated, course)
}

// GetAllCourses retrieves all courses
func GetAllCourses(c *gin.Context, db *gorm.DB) {
	var courses []models.Course

	if err := db.Preload("Instructors").Preload("Students").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetCourseByID retrieves a specific course
func GetCourseByID(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("id")

	var course models.Course
	if err := db.Preload("Instructors").Preload("Students").Preload("Modules").
		Preload("Assignments").Preload("Quizzes").Preload("Announcements").
		First(&course, "id = ?", courseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// UpdateCourse updates course information
func UpdateCourse(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("id")

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Level       string `json:"level"`
		Thumbnail   string `json:"thumbnail"`
		Status      string `json:"status"`
		MaxStudents int    `json:"max_students"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Course{}).Where("id = ?", courseID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

// DeleteCourse deletes a course
func DeleteCourse(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("id")

	if err := db.Delete(&models.Course{}, "id = ?", courseID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// GetInstructorCourses retrieves courses for a specific instructor
func GetInstructorCourses(c *gin.Context, db *gorm.DB) {
	instructorID := c.GetString("user_id")

	var courses []models.Course
	if err := db.Joins("JOIN course_instructors ON course_instructors.course_id = courses.id").
		Where("course_instructors.user_id = ?", instructorID).
		Preload("Students").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetStudentCourses retrieves courses for a specific student
func GetStudentCourses(c *gin.Context, db *gorm.DB) {
	studentID := c.GetString("user_id")

	var enrollments []models.Enrollment
	if err := db.Where("user_id = ?", studentID).
		Preload("Course").Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}
