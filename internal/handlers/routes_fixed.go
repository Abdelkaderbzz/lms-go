package handlers

import (
	"lms-go/internal/middleware"
	"lms-go/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Public routes
	public := router.Group("/api/public")
	{
		public.POST("/auth/register", func(c *gin.Context) { Register(c, db) })
		public.POST("/auth/login", func(c *gin.Context) { Login(c, db) })
	}

	// Protected routes - require authentication
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		user := protected.Group("/users")
		{
			user.GET("/profile", func(c *gin.Context) { GetProfile(c, db) })
			user.PUT("/profile", func(c *gin.Context) { UpdateProfile(c, db) })
			user.GET("/instructors", func(c *gin.Context) { GetInstructors(c, db) })
			user.GET("/students", func(c *gin.Context) { GetStudents(c, db) })
		}

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware(models.RoleAdmin))
		{
			admin.GET("/users", func(c *gin.Context) { GetAllUsers(c, db) })
			admin.GET("/users/:id", func(c *gin.Context) { GetUserByID(c, db) })
			admin.PUT("/users/:id", func(c *gin.Context) { UpdateUser(c, db) })
			admin.DELETE("/users/:id", func(c *gin.Context) { DeleteUser(c, db) })
			admin.GET("/dashboard", func(c *gin.Context) { GetAdminDashboard(c, db) })
			admin.GET("/reports", func(c *gin.Context) { GetReports(c, db) })
		}

		// Course routes - organized by specificity (most specific first)
		courses := protected.Group("/courses")
		{
			// Static routes (no parameters)
			courses.GET("/instructor/list", middleware.RoleMiddleware(models.RoleInstructor), func(c *gin.Context) { GetInstructorCourses(c, db) })
			courses.GET("/student/list", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { GetStudentCourses(c, db) })
			courses.GET("/notifications", func(c *gin.Context) { GetNotifications(c, db) })
			
			// Lesson routes (all specific to /lessons/:id before going to assignments)
			courses.GET("/lessons/:id/progress", func(c *gin.Context) { GetLessonProgress(c, db) })
			courses.POST("/lessons/:id/progress", func(c *gin.Context) { UpdateLessonProgress(c, db) })
			courses.GET("/lessons/:id/resources", func(c *gin.Context) { GetLessonResources(c, db) })
			courses.POST("/lessons/:id/resources", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { AddResource(c, db) })
			courses.GET("/lessons/:id", func(c *gin.Context) { GetLessonByID(c, db) })
			courses.PUT("/lessons/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateLesson(c, db) })
			courses.DELETE("/lessons/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { DeleteLesson(c, db) })

			// Module routes with :id
			courses.POST("/modules/:id/lessons", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateLesson(c, db) })
			courses.GET("/modules/:id/lessons", func(c *gin.Context) { GetModuleLessons(c, db) })
			courses.PUT("/modules/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateModule(c, db) })
			courses.DELETE("/modules/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { DeleteModule(c, db) })

			// Assignment routes with specific :assignmentID first
			courses.POST("/assignments/:assignmentID/submit", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { SubmitAssignment(c, db) })
			courses.GET("/assignments/:assignmentID/submissions", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { GetAssignmentSubmissions(c, db) })
			courses.GET("/assignments/:assignmentID/my-submission", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { GetStudentSubmission(c, db) })
			
			// Assignment routes with :id
			courses.GET("/assignments/:id/stats", func(c *gin.Context) { GetAssignmentStats(c, db) })
			courses.GET("/assignments/:id", func(c *gin.Context) { GetAssignmentByID(c, db) })
			courses.PUT("/assignments/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateAssignment(c, db) })
			courses.DELETE("/assignments/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { DeleteAssignment(c, db) })

			// Submission routes with :id
			courses.PUT("/submissions/:id", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { UpdateSubmission(c, db) })
			courses.POST("/submissions/:id/grade", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { GradeSubmission(c, db) })
			courses.DELETE("/submissions/:id", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { DeleteSubmission(c, db) })
			courses.GET("/submissions/:id/comments", func(c *gin.Context) { GetSubmissionComments(c, db) })
			courses.POST("/submissions/:id/comments", func(c *gin.Context) { AddComment(c, db) })

			// Question routes with :id
			courses.PUT("/questions/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateQuestion(c, db) })
			courses.DELETE("/questions/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { DeleteQuestion(c, db) })

			// Quiz routes with specific :quizID first
			courses.POST("/quizzes/:quizID/questions", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateQuestion(c, db) })
			courses.GET("/quizzes/:quizID/questions", func(c *gin.Context) { GetQuizQuestions(c, db) })
			courses.POST("/quizzes/:quizID/start", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { StartQuizAttempt(c, db) })
			courses.GET("/quizzes/:quizID/attempts", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { GetStudentAttempts(c, db) })
			
			// Quiz routes with :id
			courses.GET("/quizzes/:id", func(c *gin.Context) { GetQuizByID(c, db) })
			courses.PUT("/quizzes/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateQuiz(c, db) })
			courses.DELETE("/quizzes/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { DeleteQuiz(c, db) })

			// Attempt routes with :id
			courses.POST("/attempts/:id/submit", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { SubmitQuizAttempt(c, db) })
			courses.GET("/attempts/:id", func(c *gin.Context) { GetAttemptDetail(c, db) })

			// Announcement routes with :id
			courses.PUT("/announcements/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateAnnouncement(c, db) })
			courses.DELETE("/announcements/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { DeleteAnnouncement(c, db) })

			// Post routes with :id
			courses.GET("/posts/:id/replies", func(c *gin.Context) { GetPostReplies(c, db) })
			courses.POST("/posts/:id/replies", func(c *gin.Context) { CreateForumReply(c, db) })

			// Notification routes with :id
			courses.PUT("/notifications/:id/read", func(c *gin.Context) { MarkNotificationRead(c, db) })

			// Enrollment routes with :id
			courses.PUT("/enrollments/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateEnrollmentStatus(c, db) })

			// Generic list and create
			courses.GET("", func(c *gin.Context) { GetAllCourses(c, db) })
			courses.POST("", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateCourse(c, db) })

			// Course-specific routes with :courseID (all must come after :id routes are defined with specific operations)
			courses.POST("/:courseID/enroll", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { EnrollStudent(c, db) })
			courses.DELETE("/:courseID/unenroll", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { UnenrollStudent(c, db) })
			courses.GET("/:courseID/enrollments", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { GetCourseEnrollments(c, db) })
			courses.GET("/:courseID/enrollment-status", func(c *gin.Context) { GetEnrollmentStatus(c, db) })
			courses.GET("/:courseID/stats", func(c *gin.Context) { GetEnrollmentStats(c, db) })
			courses.POST("/:courseID/modules", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateModule(c, db) })
			courses.GET("/:courseID/modules", func(c *gin.Context) { GetCourseModules(c, db) })
			courses.POST("/:courseID/assignments", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateAssignment(c, db) })
			courses.GET("/:courseID/assignments", func(c *gin.Context) { GetCourseAssignments(c, db) })
			courses.POST("/:courseID/quizzes", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateQuiz(c, db) })
			courses.GET("/:courseID/quizzes", func(c *gin.Context) { GetCourseQuizzes(c, db) })
			courses.POST("/:courseID/announcements", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateAnnouncement(c, db) })
			courses.GET("/:courseID/announcements", func(c *gin.Context) { GetCourseAnnouncements(c, db) })
			courses.POST("/:courseID/discussions", func(c *gin.Context) { CreateDiscussion(c, db) })
			courses.GET("/:courseID/discussions", func(c *gin.Context) { GetCourseDiscussions(c, db) })
			courses.GET("/:courseID/grades", func(c *gin.Context) { GetCourseGrades(c, db) })
			courses.GET("/:courseID/my-grades", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { GetStudentGrades(c, db) })
			courses.GET("/:courseID/certificate", middleware.RoleMiddleware(models.RoleStudent), func(c *gin.Context) { GetCertificate(c, db) })

			// Module routes with :moduleID
			courses.POST("/modules/:moduleID/lessons", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { CreateLesson(c, db) })
			courses.GET("/modules/:moduleID/lessons", func(c *gin.Context) { GetModuleLessons(c, db) })

			// Discussion routes with :discussionID
			courses.POST("/discussions/:discussionID/posts", func(c *gin.Context) { CreateForumPost(c, db) })
			courses.GET("/discussions/:discussionID/posts", func(c *gin.Context) { GetDiscussionPosts(c, db) })

			// Generic course operations (MUST BE LAST - after all specific operations)
			courses.GET("/:id", func(c *gin.Context) { GetCourseByID(c, db) })
			courses.PUT("/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { UpdateCourse(c, db) })
			courses.DELETE("/:id", middleware.RoleMiddleware(models.RoleInstructor, models.RoleAdmin), func(c *gin.Context) { DeleteCourse(c, db) })
		}
	}
}
