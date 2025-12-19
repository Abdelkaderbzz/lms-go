package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lms-go/internal/models"
)

// CreateModule creates a new course module
func CreateModule(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Order       int    `json:"order"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module := models.Module{
		CourseID:    courseID,
		Title:       req.Title,
		Description: req.Description,
		Order:       req.Order,
	}

	if err := db.Create(&module).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create module"})
		return
	}

	c.JSON(http.StatusCreated, module)
}

// GetCourseModules retrieves all modules for a course
func GetCourseModules(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var modules []models.Module
	if err := db.Where("course_id = ?", courseID).
		Order("order").Preload("Lessons").Find(&modules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve modules"})
		return
	}

	c.JSON(http.StatusOK, modules)
}

// UpdateModule updates a module
func UpdateModule(c *gin.Context, db *gorm.DB) {
	moduleID := c.Param("id")

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Order       int    `json:"order"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Module{}).Where("id = ?", moduleID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update module"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Module updated successfully"})
}

// DeleteModule deletes a module
func DeleteModule(c *gin.Context, db *gorm.DB) {
	moduleID := c.Param("id")

	if err := db.Delete(&models.Module{}, "id = ?", moduleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete module"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Module deleted successfully"})
}

// CreateLesson creates a new lesson
func CreateLesson(c *gin.Context, db *gorm.DB) {
	moduleID := c.Param("moduleID")

	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Content     string `json:"content"`
		VideoURL    string `json:"video_url"`
		Duration    int    `json:"duration"`
		Order       int    `json:"order"`
		Published   bool   `json:"published"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson := models.Lesson{
		ModuleID:    moduleID,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		VideoURL:    req.VideoURL,
		Duration:    req.Duration,
		Order:       req.Order,
		Published:   req.Published,
	}

	if err := db.Create(&lesson).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lesson"})
		return
	}

	c.JSON(http.StatusCreated, lesson)
}

// GetModuleLessons retrieves all lessons for a module
func GetModuleLessons(c *gin.Context, db *gorm.DB) {
	moduleID := c.Param("moduleID")

	var lessons []models.Lesson
	if err := db.Where("module_id = ?", moduleID).
		Order("order").Preload("Resources").Find(&lessons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve lessons"})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

// GetLessonByID retrieves a specific lesson
func GetLessonByID(c *gin.Context, db *gorm.DB) {
	lessonID := c.Param("id")

	var lesson models.Lesson
	if err := db.Preload("Resources").First(&lesson, "id = ?", lessonID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// UpdateLesson updates a lesson
func UpdateLesson(c *gin.Context, db *gorm.DB) {
	lessonID := c.Param("id")

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Content     string `json:"content"`
		VideoURL    string `json:"video_url"`
		Duration    int    `json:"duration"`
		Order       int    `json:"order"`
		Published   bool   `json:"published"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Lesson{}).Where("id = ?", lessonID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lesson"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson updated successfully"})
}

// DeleteLesson deletes a lesson
func DeleteLesson(c *gin.Context, db *gorm.DB) {
	lessonID := c.Param("id")

	if err := db.Delete(&models.Lesson{}, "id = ?", lessonID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lesson"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}

// UpdateLessonProgress updates student's lesson progress
func UpdateLessonProgress(c *gin.Context, db *gorm.DB) {
	lessonID := c.Param("lessonID")
	userID := c.GetString("user_id")

	var req struct {
		Completed bool    `json:"completed"`
		Progress  float64 `json:"progress"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var progress models.LessonProgress
	result := db.Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&progress)

	if result.Error == gorm.ErrRecordNotFound {
		progress = models.LessonProgress{
			UserID:    userID,
			LessonID:  lessonID,
			Completed: req.Completed,
			Progress:  req.Progress,
		}
		if err := db.Create(&progress).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update progress"})
			return
		}
	} else {
		if err := db.Model(&progress).Updates(req).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update progress"})
			return
		}
	}

	c.JSON(http.StatusOK, progress)
}

// GetLessonProgress gets student's lesson progress
func GetLessonProgress(c *gin.Context, db *gorm.DB) {
	lessonID := c.Param("lessonID")
	userID := c.GetString("user_id")

	var progress models.LessonProgress
	if err := db.Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		First(&progress).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"progress": 0, "completed": false})
		return
	}

	c.JSON(http.StatusOK, progress)
}

// AddResource adds a resource to a lesson
func AddResource(c *gin.Context, db *gorm.DB) {
	lessonID := c.Param("lessonID")

	var req struct {
		Type     string `json:"type" binding:"required"`
		Title    string `json:"title" binding:"required"`
		URL      string `json:"url" binding:"required"`
		FileSize int64  `json:"file_size"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource := models.Resource{
		LessonID: lessonID,
		Type:     req.Type,
		Title:    req.Title,
		URL:      req.URL,
		FileSize: req.FileSize,
	}

	if err := db.Create(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add resource"})
		return
	}

	c.JSON(http.StatusCreated, resource)
}

// GetLessonResources retrieves all resources for a lesson
func GetLessonResources(c *gin.Context, db *gorm.DB) {
	lessonID := c.Param("lessonID")

	var resources []models.Resource
	if err := db.Where("lesson_id = ?", lessonID).Find(&resources).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve resources"})
		return
	}

	c.JSON(http.StatusOK, resources)
}
