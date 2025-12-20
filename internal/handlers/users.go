package handlers

import (
	"net/http"

	"lms-go/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllUsers retrieves all users (admin only)
func GetAllUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a specific user
func GetUserByID(c *gin.Context, db *gorm.DB) {
	userID := c.Param("id")

	var user models.User
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates user information (admin only)
func UpdateUser(c *gin.Context, db *gorm.DB) {
	userID := c.Param("id")

	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Role      string `json:"role"`
		Active    bool   `json:"active"`
		Email     string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.User{}).Where("id = ?", userID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user (admin only)
func DeleteUser(c *gin.Context, db *gorm.DB) {
	userID := c.Param("id")

	if err := db.Delete(&models.User{}, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetInstructors retrieves all instructors
func GetInstructors(c *gin.Context, db *gorm.DB) {
	var instructors []models.User

	if err := db.Where("role = ?", models.RoleInstructor).Find(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve instructors"})
		return
	}

	c.JSON(http.StatusOK, instructors)
}

// GetStudents retrieves all students
func GetStudents(c *gin.Context, db *gorm.DB) {
	var students []models.User

	if err := db.Where("role = ?", models.RoleStudent).Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
		return
	}

	c.JSON(http.StatusOK, students)
}
