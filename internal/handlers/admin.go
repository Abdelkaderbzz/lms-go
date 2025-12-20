package handlers

import (
	"net/http"

	"lms-go/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAdminDashboard retrieves admin dashboard statistics
func GetAdminDashboard(c *gin.Context, db *gorm.DB) {
	var stats struct {
		TotalUsers      int64 `json:"total_users"`
		TotalCourses    int64 `json:"total_courses"`
		TotalEnrollments int64 `json:"total_enrollments"`
		TotalAssignments int64 `json:"total_assignments"`
		ActiveStudents  int64 `json:"active_students"`
		ActiveCourses   int64 `json:"active_courses"`
	}

	db.Model(&models.User{}).Count(&stats.TotalUsers)
	db.Model(&models.Course{}).Count(&stats.TotalCourses)
	db.Model(&models.Enrollment{}).Count(&stats.TotalEnrollments)
	db.Model(&models.Assignment{}).Count(&stats.TotalAssignments)
	db.Model(&models.User{}).Where("role = ?", models.RoleStudent).Count(&stats.ActiveStudents)
	db.Model(&models.Course{}).Where("status = ?", "active").Count(&stats.ActiveCourses)

	c.JSON(http.StatusOK, stats)
}

// GetReports retrieves system reports
func GetReports(c *gin.Context, db *gorm.DB) {
	var enrollmentByStatus struct {
		Active    int64 `json:"active"`
		Completed int64 `json:"completed"`
		Dropped   int64 `json:"dropped"`
	}

	db.Model(&models.Enrollment{}).Where("status = ?", "active").Count(&enrollmentByStatus.Active)
	db.Model(&models.Enrollment{}).Where("status = ?", "completed").Count(&enrollmentByStatus.Completed)
	db.Model(&models.Enrollment{}).Where("status = ?", "dropped").Count(&enrollmentByStatus.Dropped)

	var usersByRole struct {
		Students    int64 `json:"students"`
		Instructors int64 `json:"instructors"`
		Admins      int64 `json:"admins"`
	}

	db.Model(&models.User{}).Where("role = ?", models.RoleStudent).Count(&usersByRole.Students)
	db.Model(&models.User{}).Where("role = ?", models.RoleInstructor).Count(&usersByRole.Instructors)
	db.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&usersByRole.Admins)

	c.JSON(http.StatusOK, gin.H{
		"enrollment_by_status": enrollmentByStatus,
		"users_by_role": usersByRole,
	})
}
