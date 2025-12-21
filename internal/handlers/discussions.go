package handlers

import (
	"net/http"

	"lms-go/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateAnnouncement creates a new announcement
func CreateAnnouncement(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")
	userID := c.GetString("user_id")

	var req struct {
		Title     string `json:"title" binding:"required"`
		Content   string `json:"content" binding:"required"`
		Important bool   `json:"important"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	announcement := models.Announcement{
		CourseID:  courseID,
		CreatorID: userID,
		Title:     req.Title,
		Content:   req.Content,
		Important: req.Important,
	}

	if err := db.Create(&announcement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create announcement"})
		return
	}

	c.JSON(http.StatusCreated, announcement)
}

// GetCourseAnnouncements retrieves all announcements for a course
func GetCourseAnnouncements(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var announcements []models.Announcement
	if err := db.Where("course_id = ?", courseID).
		Order("created_at DESC").Preload("User").Find(&announcements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve announcements"})
		return
	}

	c.JSON(http.StatusOK, announcements)
}

// UpdateAnnouncement updates an announcement
func UpdateAnnouncement(c *gin.Context, db *gorm.DB) {
	announcementID := c.Param("id")

	var req struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		Important bool   `json:"important"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&models.Announcement{}).Where("id = ?", announcementID).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update announcement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Announcement updated successfully"})
}

// DeleteAnnouncement deletes an announcement
func DeleteAnnouncement(c *gin.Context, db *gorm.DB) {
	announcementID := c.Param("id")

	if err := db.Delete(&models.Announcement{}, "id = ?", announcementID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete announcement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Announcement deleted successfully"})
}

// CreateDiscussion creates a new discussion
func CreateDiscussion(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discussion := models.Discussion{
		CourseID: courseID,
		Title:    req.Title,
		Content:  req.Content,
	}

	if err := db.Create(&discussion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create discussion"})
		return
	}

	c.JSON(http.StatusCreated, discussion)
}

// GetCourseDiscussions retrieves all discussions for a course
func GetCourseDiscussions(c *gin.Context, db *gorm.DB) {
	courseID := c.Param("courseID")

	var discussions []models.Discussion
	if err := db.Where("course_id = ?", courseID).
		Preload("Posts").Find(&discussions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve discussions"})
		return
	}

	c.JSON(http.StatusOK, discussions)
}

// CreateForumPost creates a new forum post
func CreateForumPost(c *gin.Context, db *gorm.DB) {
	discussionID := c.Param("discussionID")
	userID := c.GetString("user_id")

	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.ForumPost{
		DiscussionID: discussionID,
		UserID:       userID,
		Title:        req.Title,
		Content:      req.Content,
		Views:        0,
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetDiscussionPosts retrieves all posts in a discussion
func GetDiscussionPosts(c *gin.Context, db *gorm.DB) {
	discussionID := c.Param("discussionID")

	var posts []models.ForumPost
	if err := db.Where("discussion_id = ?", discussionID).
		Order("created_at DESC").Preload("User").Preload("Replies").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// CreateForumReply creates a reply to a forum post
func CreateForumReply(c *gin.Context, db *gorm.DB) {
	postID := c.Param("postID")
	userID := c.GetString("user_id")

	var req struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reply := models.ForumReply{
		PostID:  postID,
		UserID:  userID,
		Content: req.Content,
	}

	if err := db.Create(&reply).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reply"})
		return
	}

	c.JSON(http.StatusCreated, reply)
}

// GetPostReplies retrieves all replies for a post
func GetPostReplies(c *gin.Context, db *gorm.DB) {
	postID := c.Param("postID")

	var replies []models.ForumReply
	if err := db.Where("post_id = ?", postID).
		Preload("User").Find(&replies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve replies"})
		return
	}

	c.JSON(http.StatusOK, replies)
}
