package main

import (
	"fmt"
	"log"
	"time"

	"lms-go/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedDatabase seeds the database with test data
func SeedDatabase(db *gorm.DB) {
	fmt.Println("🌱 Starting database seeding...")

	// Create admin user
	adminUser := &models.User{
		Email:     "admin@example.com",
		Password:  hashPassword("admin123"),
		FirstName: "Admin",
		LastName:  "User",
		Role:      models.RoleAdmin,
		Phone:     "+1234567890",
		Bio:       "System administrator",
		Active:    true,
	}
	if err := db.Create(adminUser).Error; err != nil {
		log.Printf("Error creating admin user: %v", err)
	} else {
		fmt.Println("✓ Admin user created")
	}

	// Create instructor users
	instructors := []models.User{
		{
			Email:     "instructor1@example.com",
			Password:  hashPassword("instructor123"),
			FirstName: "John",
			LastName:  "Doe",
			Role:      models.RoleInstructor,
			Phone:     "+1234567891",
			Bio:       "Senior instructor with 10 years experience",
			Active:    true,
		},
		{
			Email:     "instructor2@example.com",
			Password:  hashPassword("instructor123"),
			FirstName: "Jane",
			LastName:  "Smith",
			Role:      models.RoleInstructor,
			Phone:     "+1234567892",
			Bio:       "Expert in web technologies",
			Active:    true,
		},
	}

	for _, instructor := range instructors {
		if err := db.Create(&instructor).Error; err != nil {
			log.Printf("Error creating instructor: %v", err)
		}
	}
	fmt.Println("✓ Instructor users created")

	// Create student users
	students := []models.User{
		{
			Email:     "student1@example.com",
			Password:  hashPassword("student123"),
			FirstName: "Alice",
			LastName:  "Johnson",
			Role:      models.RoleStudent,
			Phone:     "+1234567893",
			Bio:       "Enthusiastic learner",
			Active:    true,
		},
		{
			Email:     "student2@example.com",
			Password:  hashPassword("student123"),
			FirstName: "Bob",
			LastName:  "Wilson",
			Role:      models.RoleStudent,
			Phone:     "+1234567894",
			Bio:       "Passionate developer",
			Active:    true,
		},
		{
			Email:     "student3@example.com",
			Password:  hashPassword("student123"),
			FirstName: "Carol",
			LastName:  "Davis",
			Role:      models.RoleStudent,
			Phone:     "+1234567895",
			Bio:       "Full-time student",
			Active:    true,
		},
		{
			Email:     "student4@example.com",
			Password:  hashPassword("student123"),
			FirstName: "David",
			LastName:  "Miller",
			Role:      models.RoleStudent,
			Phone:     "+1234567896",
			Bio:       "Tech enthusiast",
			Active:    true,
		},
	}

	var studentIDs []string
	for _, student := range students {
		if err := db.Create(&student).Error; err != nil {
			log.Printf("Error creating student: %v", err)
		} else {
			studentIDs = append(studentIDs, student.ID)
		}
	}
	fmt.Println("✓ Student users created")

	// Get instructor ID for course creation
	var instructor models.User
	db.Where("role = ?", models.RoleInstructor).First(&instructor)

	// Create courses
	courses := []models.Course{
		{
			Title:       "Introduction to Go Programming",
			Description: "Learn the basics of Go programming language with hands-on examples",
			Code:        "GO-101",
			Category:    "Programming",
			Level:       "Beginner",
			Thumbnail:   "https://via.placeholder.com/300x200?text=Go+101",
			Status:      "active",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(0, 3, 0),
			MaxStudents: 50,
		},
		{
			Title:       "Advanced Go Concurrency",
			Description: "Master concurrent programming patterns in Go",
			Code:        "GO-301",
			Category:    "Programming",
			Level:       "Advanced",
			Thumbnail:   "https://via.placeholder.com/300x200?text=Go+Concurrency",
			Status:      "active",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(0, 3, 0),
			MaxStudents: 30,
		},
		{
			Title:       "Web Development with Go",
			Description: "Build RESTful APIs and web applications with Go",
			Code:        "GO-201",
			Category:    "Web Development",
			Level:       "Intermediate",
			Thumbnail:   "https://via.placeholder.com/300x200?text=Web+Dev",
			Status:      "active",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(0, 4, 0),
			MaxStudents: 40,
		},
	}

	var courseIDs []string
	for _, course := range courses {
		if err := db.Create(&course).Error; err != nil {
			log.Printf("Error creating course: %v", err)
		} else {
			courseIDs = append(courseIDs, course.ID)
			// Associate instructor with course
			if err := db.Model(&course).Association("Instructors").Append(&instructor); err != nil {
				log.Printf("Warning: Failed to add instructor to course: %v", err)
			}
		}
	}
	fmt.Println("✓ Courses created")

	// Create enrollments for students
	for _, studentID := range studentIDs {
		for _, courseID := range courseIDs {
			enrollment := &models.Enrollment{
				CourseID:   courseID,
				UserID:     studentID,
				Status:     "active",
				Progress:   0,
				EnrolledAt: time.Now(),
			}
			if err := db.Create(enrollment).Error; err != nil {
				log.Printf("Error creating enrollment: %v", err)
			}
		}
	}
	fmt.Println("✓ Enrollments created")

	// Create modules and lessons for first course
	if len(courseIDs) > 0 {
		modules := []models.Module{
			{
				CourseID:    courseIDs[0],
				Title:       "Module 1: Getting Started",
				Description: "Introduction and setup",
				Order:       1,
			},
			{
				CourseID:    courseIDs[0],
				Title:       "Module 2: Core Concepts",
				Description: "Fundamental concepts of Go",
				Order:       2,
			},
		}

		var moduleIDs []string
		for _, module := range modules {
			if err := db.Create(&module).Error; err != nil {
				log.Printf("Error creating module: %v", err)
			} else {
				moduleIDs = append(moduleIDs, module.ID)
			}
		}
		fmt.Println("✓ Modules created")

		// Create lessons
		if len(moduleIDs) > 0 {
			lessons := []models.Lesson{
				{
					ModuleID:    moduleIDs[0],
					Title:       "Lesson 1: Hello World",
					Description: "Your first Go program",
					Content:     "<p>Learn how to write and run your first Go program</p>",
					VideoURL:    "https://example.com/video1.mp4",
					Duration:    30,
					Order:       1,
					Published:   true,
				},
				{
					ModuleID:    moduleIDs[0],
					Title:       "Lesson 2: Variables and Types",
					Description: "Understanding data types",
					Content:     "<p>Learn about variables and different data types in Go</p>",
					VideoURL:    "https://example.com/video2.mp4",
					Duration:    45,
					Order:       2,
					Published:   true,
				},
				{
					ModuleID:    moduleIDs[1],
					Title:       "Lesson 3: Functions",
					Description: "Writing and using functions",
					Content:     "<p>Master function declarations and usage</p>",
					VideoURL:    "https://example.com/video3.mp4",
					Duration:    40,
					Order:       1,
					Published:   true,
				},
			}

			var lessonIDs []string
			for _, lesson := range lessons {
				if err := db.Create(&lesson).Error; err != nil {
					log.Printf("Error creating lesson: %v", err)
				} else {
					lessonIDs = append(lessonIDs, lesson.ID)
				}
			}
			fmt.Println("✓ Lessons created")

			// Create lesson progress for students
			for _, studentID := range studentIDs {
				for _, lessonID := range lessonIDs {
					progress := &models.LessonProgress{
						UserID:    studentID,
						LessonID:  lessonID,
						Completed: false,
						Progress:  0,
						LastView:  time.Now(),
					}
					if err := db.Create(progress).Error; err != nil {
						log.Printf("Error creating lesson progress: %v", err)
					}
				}
			}
			fmt.Println("✓ Lesson progress created")

			// Create resources
			if len(lessonIDs) > 0 {
				resources := []models.Resource{
					{
						LessonID: lessonIDs[0],
						CourseID: courseIDs[0],
						Type:     "pdf",
						Title:    "Hello World Guide",
						URL:      "https://example.com/hello-world.pdf",
						FileSize: 1024000,
					},
					{
						LessonID: lessonIDs[1],
						CourseID: courseIDs[0],
						Type:     "document",
						Title:    "Variables Cheat Sheet",
						URL:      "https://example.com/variables-cheatsheet.docx",
						FileSize: 512000,
					},
				}

				for _, resource := range resources {
					if err := db.Create(&resource).Error; err != nil {
						log.Printf("Error creating resource: %v", err)
					}
				}
				fmt.Println("✓ Resources created")
			}
		}

		// Create assignments
		assignments := []models.Assignment{
			{
				CourseID:    courseIDs[0],
				Title:       "Assignment 1: Basic Program",
				Description: "Write a simple Go program",
				DueDate:     time.Now().AddDate(0, 0, 7),
				Points:      100,
				Type:        "homework",
				Status:      "active",
			},
			{
				CourseID:    courseIDs[0],
				Title:       "Assignment 2: Function Implementation",
				Description: "Implement required functions",
				DueDate:     time.Now().AddDate(0, 0, 14),
				Points:      150,
				Type:        "project",
				Status:      "active",
			},
		}

		var assignmentIDs []string
		for _, assignment := range assignments {
			if err := db.Create(&assignment).Error; err != nil {
				log.Printf("Error creating assignment: %v", err)
			} else {
				assignmentIDs = append(assignmentIDs, assignment.ID)
			}
		}
		fmt.Println("✓ Assignments created")

		// Create submissions
		if len(assignmentIDs) > 0 && len(studentIDs) > 0 {
			for i, studentID := range studentIDs {
				submission := models.Submission{
					AssignmentID: assignmentIDs[0],
					UserID:       studentID,
					Content:      "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}",
					FileURL:      "https://example.com/submissions/student-" + studentID + ".go",
					SubmittedAt:  time.Now().AddDate(0, 0, -i),
					Status:       "submitted",
				}
				if err := db.Create(&submission).Error; err != nil {
					log.Printf("Error creating submission: %v", err)
				} else {
					// Create grade for submission
					grade := &models.Grade{
						UserID:       studentID,
						CourseID:     courseIDs[0],
						AssignmentID: assignmentIDs[0],
						SubmissionID: submission.ID,
						Points:       90 + float64(i*2),
						Feedback:     "Good work! Keep it up.",
						GradedBy:     instructor.ID,
						GradedAt:     time.Now(),
					}
					db.Create(grade)
				}
			}
			fmt.Println("✓ Submissions and grades created")
		}

		// Create quizzes
		quizzes := []models.Quiz{
			{
				CourseID:    courseIDs[0],
				CreatorID:   instructor.ID,
				Title:       "Quiz 1: Go Basics",
				Description: "Test your knowledge of Go basics",
				StartDate:   time.Now(),
				EndDate:     time.Now().AddDate(0, 0, 30),
				TimeLimit:   30,
				PassScore:   70,
				Shuffle:     true,
				Public:      true,
			},
		}

		var quizIDs []string
		for _, quiz := range quizzes {
			if err := db.Create(&quiz).Error; err != nil {
				log.Printf("Error creating quiz: %v", err)
			} else {
				quizIDs = append(quizIDs, quiz.ID)
			}
		}
		fmt.Println("✓ Quizzes created")

		// Create questions and options
		if len(quizIDs) > 0 {
			questions := []models.Question{
				{
					QuizID:   quizIDs[0],
					Type:     "multiple_choice",
					Question: "What is the correct way to declare a variable in Go?",
					Points:   10,
					Order:    1,
				},
				{
					QuizID:   quizIDs[0],
					Type:     "true_false",
					Question: "Go is a statically typed language",
					Points:   5,
					Order:    2,
				},
			}

			var questionIDs []string
			for _, question := range questions {
				if err := db.Create(&question).Error; err != nil {
					log.Printf("Error creating question: %v", err)
				} else {
					questionIDs = append(questionIDs, question.ID)
				}
			}

			// Create options for first question
			if len(questionIDs) > 0 {
				options := []models.Option{
					{
						QuestionID: questionIDs[0],
						Text:       "var x int = 5",
						IsCorrect:  true,
						Order:      1,
					},
					{
						QuestionID: questionIDs[0],
						Text:       "int x = 5",
						IsCorrect:  false,
						Order:      2,
					},
					{
						QuestionID: questionIDs[0],
						Text:       "declare x as 5",
						IsCorrect:  false,
						Order:      3,
					},
				}

				for _, option := range options {
					if err := db.Create(&option).Error; err != nil {
						log.Printf("Error creating option: %v", err)
					}
				}

				// Create options for second question
				if len(questionIDs) > 1 {
					options = []models.Option{
						{
							QuestionID: questionIDs[1],
							Text:       "True",
							IsCorrect:  true,
							Order:      1,
						},
						{
							QuestionID: questionIDs[1],
							Text:       "False",
							IsCorrect:  false,
							Order:      2,
						},
					}

					for _, option := range options {
						if err := db.Create(&option).Error; err != nil {
							log.Printf("Error creating option: %v", err)
						}
					}
				}
				fmt.Println("✓ Questions and options created")
			}

			// Create quiz attempts
			if len(quizIDs) > 0 && len(studentIDs) > 0 {
				for i, studentID := range studentIDs {
					attempt := &models.QuizAttempt{
						QuizID:    quizIDs[0],
						UserID:    studentID,
						StartedAt: time.Now().AddDate(0, 0, -i),
						Score:     75 + float64(i*3),
						Status:    "submitted",
					}
					now := time.Now()
					attempt.EndedAt = &now
					if err := db.Create(attempt).Error; err != nil {
						log.Printf("Error creating quiz attempt: %v", err)
					}
				}
				fmt.Println("✓ Quiz attempts created")
			}
		}

		// Create discussions
		discussions := []models.Discussion{
			{
				CourseID: courseIDs[0],
				Title:    "General Discussion",
				Content:  "Discuss course content and ask questions here",
			},
		}

		var discussionIDs []string
		for _, discussion := range discussions {
			if err := db.Create(&discussion).Error; err != nil {
				log.Printf("Error creating discussion: %v", err)
			} else {
				discussionIDs = append(discussionIDs, discussion.ID)
			}
		}
		fmt.Println("✓ Discussions created")

		// Create forum posts and replies
		if len(discussionIDs) > 0 && len(studentIDs) > 0 {
			posts := []models.ForumPost{
				{
					DiscussionID: discussionIDs[0],
					UserID:       studentIDs[0],
					Title:        "How to get started with Go?",
					Content:      "I'm new to Go and need some guidance on where to start.",
					Views:        15,
				},
				{
					DiscussionID: discussionIDs[0],
					UserID:       studentIDs[1],
					Title:        "Tips for learning concurrency",
					Content:      "What are the best practices for learning Go concurrency?",
					Views:        8,
				},
			}

			var postIDs []string
			for _, post := range posts {
				if err := db.Create(&post).Error; err != nil {
					log.Printf("Error creating forum post: %v", err)
				} else {
					postIDs = append(postIDs, post.ID)
				}
			}

			// Create forum replies
			if len(postIDs) > 0 {
				replies := []models.ForumReply{
					{
						PostID:  postIDs[0],
						UserID:  instructor.ID,
						Content: "Start with the official Go tour and documentation. It's very comprehensive!",
						Helpful: 5,
					},
					{
						PostID:  postIDs[1],
						UserID:  studentIDs[2],
						Content: "I found the book 'The Go Programming Language' very helpful for understanding concurrency patterns.",
						Helpful: 3,
					},
				}

				for _, reply := range replies {
					if err := db.Create(&reply).Error; err != nil {
						log.Printf("Error creating forum reply: %v", err)
					}
				}
				fmt.Println("✓ Forum posts and replies created")
			}
		}

		// Create announcements
		announcements := []models.Announcement{
			{
				CourseID:  courseIDs[0],
				CreatorID: instructor.ID,
				Title:     "Welcome to the Course!",
				Content:   "Welcome everyone! This course will help you master Go programming. Let's get started!",
				Important: true,
			},
			{
				CourseID:  courseIDs[0],
				CreatorID: instructor.ID,
				Title:     "Assignment 1 Due Date Extended",
				Content:   "Due to popular request, Assignment 1 due date has been extended to next Friday.",
				Important: true,
			},
		}

		for _, announcement := range announcements {
			if err := db.Create(&announcement).Error; err != nil {
				log.Printf("Error creating announcement: %v", err)
			}
		}
		fmt.Println("✓ Announcements created")

		// Create notifications
		if len(studentIDs) > 0 {
			notifications := []models.Notification{
				{
					UserID:    studentIDs[0],
					Type:      "announcement",
					Title:     "New announcement in Go 101",
					Message:   "Welcome to the Course!",
					Reference: courseIDs[0],
					Read:      false,
				},
				{
					UserID:    studentIDs[0],
					Type:      "assignment",
					Title:     "New assignment posted",
					Message:   "Assignment 1: Basic Program has been posted",
					Reference: assignmentIDs[0],
					Read:      false,
				},
			}

			for _, notification := range notifications {
				if err := db.Create(&notification).Error; err != nil {
					log.Printf("Error creating notification: %v", err)
				}
			}
			fmt.Println("✓ Notifications created")
		}
	}

	fmt.Println("\n✅ Database seeding completed successfully!")
	fmt.Println("📝 Test User Credentials:")
	fmt.Println("  Admin: admin@example.com / admin123")
	fmt.Println("  Instructor: instructor1@example.com / instructor123")
	fmt.Println("  Student: student1@example.com / student123")
	fmt.Println("")
}

// hashPassword hashes a password using bcrypt
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return password
	}
	return string(hash)
}
