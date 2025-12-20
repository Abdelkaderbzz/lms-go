package database

import (
	"fmt"

	"lms-go/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Module{},
		&models.Lesson{},
		&models.Assignment{},
		&models.Submission{},
		&models.Grade{},
		&models.Comment{},
		&models.Quiz{},
		&models.Question{},
		&models.Option{},
		&models.QuizAttempt{},
		&models.QuizAnswer{},
		&models.Enrollment{},
		&models.Resource{},
		&models.Announcement{},
		&models.Discussion{},
		&models.ForumPost{},
		&models.ForumReply{},
		&models.Notification{},
		&models.LessonProgress{},
		&models.Certificate{},
	)
}
