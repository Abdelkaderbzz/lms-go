package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// User roles
const (
	RoleAdmin      = "admin"
	RoleInstructor = "instructor"
	RoleStudent    = "student"
)

// User represents a system user
type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"` // admin, instructor, student
	Phone     string    `json:"phone"`
	Bio       string    `json:"bio"`
	Avatar    string    `json:"avatar"`
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Courses          []Course       `gorm:"many2many:course_instructors;" json:"-"`
	EnrolledCourses  []Course       `gorm:"many2many:course_enrollments;" json:"-"`
	Submissions      []Submission   `gorm:"foreignKey:UserID" json:"-"`
	Comments         []Comment      `gorm:"foreignKey:UserID" json:"-"`
	Grades           []Grade        `gorm:"foreignKey:UserID" json:"-"`
	Announcements    []Announcement `gorm:"foreignKey:CreatorID" json:"-"`
	NotificationRead []Notification `gorm:"foreignKey:UserID" json:"-"`
	CreatedQuizzes   []Quiz         `gorm:"foreignKey:CreatorID" json:"-"`
	QuizAttempts     []QuizAttempt  `gorm:"foreignKey:UserID" json:"-"`
	ForumPosts       []ForumPost    `gorm:"foreignKey:UserID" json:"-"`
	ForumReplies     []ForumReply   `gorm:"foreignKey:UserID" json:"-"`
}

// Course represents a learning course
type Course struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Code        string    `gorm:"uniqueIndex" json:"code"`
	Category    string    `json:"category"`
	Level       string    `json:"level"` // Beginner, Intermediate, Advanced
	Thumbnail   string    `json:"thumbnail"`
	Status      string    `json:"status"` // active, archived, draft
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	MaxStudents int       `json:"max_students"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relations
	Instructors   []User         `gorm:"many2many:course_instructors;" json:"instructors"`
	Students      []User         `gorm:"many2many:course_enrollments;" json:"students"`
	Modules       []Module       `json:"modules"`
	Assignments   []Assignment   `json:"assignments"`
	Quizzes       []Quiz         `json:"quizzes"`
	Announcements []Announcement `json:"announcements"`
	Discussions   []Discussion   `json:"discussions"`
	Resources     []Resource     `json:"resources"`
	Enrollments   []Enrollment   `json:"enrollments"`
	Grades        []Grade        `json:"grades"`
}

// Module represents a course module
type Module struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Course  Course   `gorm:"foreignKey:CourseID" json:"-"`
	Lessons []Lesson `gorm:"foreignKey:ModuleID" json:"lessons"`
}

// Lesson represents course content
type Lesson struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	ModuleID    string    `json:"module_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"` // HTML content
	VideoURL    string    `json:"video_url"`
	Duration    int       `json:"duration"` // in minutes
	Order       int       `json:"order"`
	Published   bool      `json:"published" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Module    Module           `gorm:"foreignKey:ModuleID" json:"-"`
	Resources []Resource       `gorm:"foreignKey:LessonID" json:"resources"`
	Progress  []LessonProgress `gorm:"foreignKey:LessonID" json:"progress"`
}

// Assignment represents a course assignment
type Assignment struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Points      float64   `json:"points"`
	Type        string    `json:"type"`   // homework, project, essay
	Status      string    `json:"status"` // active, closed
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Course      Course         `gorm:"foreignKey:CourseID" json:"-"`
	Submissions []Submission   `gorm:"foreignKey:AssignmentID" json:"submissions"`
	Rubric      datatypes.JSON `json:"rubric"` // Grading rubric
}

// Submission represents student assignment submission
type Submission struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	AssignmentID string    `json:"assignment_id"`
	UserID       string    `json:"user_id"`
	Content      string    `json:"content"`
	FileURL      string    `json:"file_url"`
	SubmittedAt  time.Time `json:"submitted_at"`
	Status       string    `json:"status"` // submitted, graded, late
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Assignment Assignment `gorm:"foreignKey:AssignmentID" json:"-"`
	User       User       `gorm:"foreignKey:UserID" json:"-"`
	Comments   []Comment  `gorm:"foreignKey:SubmissionID" json:"comments"`
}

// Grade represents grading information
type Grade struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	UserID       string    `json:"user_id"`
	CourseID     string    `json:"course_id"`
	AssignmentID string    `json:"assignment_id"`
	SubmissionID string    `json:"submission_id"`
	Points       float64   `json:"points"`
	Feedback     string    `json:"feedback"`
	GradedBy     string    `json:"graded_by"`
	GradedAt     time.Time `json:"graded_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	User       User       `gorm:"foreignKey:UserID" json:"-"`
	Course     Course     `gorm:"foreignKey:CourseID" json:"-"`
	Assignment Assignment `gorm:"foreignKey:AssignmentID" json:"-"`
	Submission Submission `gorm:"foreignKey:SubmissionID" json:"-"`
	Grader     User       `json:"-" gorm:"foreignKey:GradedBy"`
}

// Comment represents submission comments
type Comment struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	SubmissionID string    `json:"submission_id"`
	UserID       string    `json:"user_id"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Submission Submission `gorm:"foreignKey:SubmissionID" json:"-"`
	User       User       `gorm:"foreignKey:UserID" json:"-"`
}

// Quiz represents assessment quizzes
type Quiz struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	CourseID    string    `json:"course_id"`
	CreatorID   string    `json:"creator_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	TimeLimit   int       `json:"time_limit"` // minutes
	PassScore   float64   `json:"pass_score"` // percentage
	Shuffle     bool      `json:"shuffle"`
	Public      bool      `json:"public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Course    Course        `gorm:"foreignKey:CourseID" json:"-"`
	Creator   User          `gorm:"foreignKey:CreatorID" json:"-"`
	Questions []Question    `gorm:"foreignKey:QuizID" json:"questions"`
	Attempts  []QuizAttempt `gorm:"foreignKey:QuizID" json:"attempts"`
}

// Question represents quiz questions
type Question struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	QuizID    string    `json:"quiz_id"`
	Type      string    `json:"type"` // multiple_choice, true_false, essay, matching
	Question  string    `json:"question"`
	Points    float64   `json:"points"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Quiz    Quiz     `gorm:"foreignKey:QuizID" json:"-"`
	Options []Option `gorm:"foreignKey:QuestionID" json:"options"`
}

// Option represents question options
type Option struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	QuestionID string    `json:"question_id"`
	Text       string    `json:"text"`
	IsCorrect  bool      `json:"is_correct"`
	Order      int       `json:"order"`
	CreatedAt  time.Time `json:"created_at"`

	Question Question `gorm:"foreignKey:QuestionID" json:"-"`
}

// QuizAttempt represents student quiz attempts
type QuizAttempt struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	QuizID    string     `json:"quiz_id"`
	UserID    string     `json:"user_id"`
	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at"`
	Score     float64    `json:"score"`
	Status    string     `json:"status"` // in_progress, submitted, reviewed
	CreatedAt time.Time  `json:"created_at"`

	Quiz    Quiz         `gorm:"foreignKey:QuizID" json:"-"`
	User    User         `gorm:"foreignKey:UserID" json:"-"`
	Answers []QuizAnswer `gorm:"foreignKey:QuizAttemptID" json:"answers"`
}

// QuizAnswer represents student answers
type QuizAnswer struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	QuizAttemptID string    `json:"quiz_attempt_id"`
	QuestionID    string    `json:"question_id"`
	SelectedID    string    `json:"selected_id"` // For multiple choice/matching
	TextAnswer    string    `json:"text_answer"` // For essay/short answer
	Points        float64   `json:"points"`
	CreatedAt     time.Time `json:"created_at"`

	QuizAttempt QuizAttempt `gorm:"foreignKey:QuizAttemptID" json:"-"`
	Question    Question    `gorm:"foreignKey:QuestionID" json:"-"`
}

// Enrollment represents course enrollment
type Enrollment struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	CourseID   string    `json:"course_id"`
	UserID     string    `json:"user_id"`
	Status     string    `json:"status"` // active, completed, dropped
	Progress   float64   `json:"progress"`
	EnrolledAt time.Time `json:"enrolled_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Course Course `gorm:"foreignKey:CourseID" json:"-"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
}

// Resource represents course resources
type Resource struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	LessonID  string    `json:"lesson_id"`
	CourseID  string    `json:"course_id"`
	Type      string    `json:"type"` // pdf, document, link, video
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	FileSize  int64     `json:"file_size"`
	CreatedAt time.Time `json:"created_at"`

	Lesson Lesson `gorm:"foreignKey:LessonID" json:"-"`
	Course Course `gorm:"foreignKey:CourseID" json:"-"`
}

// Announcement represents course announcements
type Announcement struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CourseID  string    `json:"course_id"`
	CreatorID string    `json:"creator_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Important bool      `json:"important"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Course  Course `gorm:"foreignKey:CourseID" json:"-"`
	Creator User   `gorm:"foreignKey:CreatorID" json:"-"`
}

// Discussion represents course discussions
type Discussion struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CourseID  string    `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Course Course      `gorm:"foreignKey:CourseID" json:"-"`
	Posts  []ForumPost `gorm:"foreignKey:DiscussionID" json:"posts"`
}

// ForumPost represents forum posts
type ForumPost struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	DiscussionID string    `json:"discussion_id"`
	UserID       string    `json:"user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Views        int       `json:"views"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Discussion Discussion   `gorm:"foreignKey:DiscussionID" json:"-"`
	User       User         `gorm:"foreignKey:UserID" json:"-"`
	Replies    []ForumReply `gorm:"foreignKey:PostID" json:"replies"`
}

// ForumReply represents forum replies
type ForumReply struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	Helpful   int       `json:"helpful"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Post ForumPost `gorm:"foreignKey:PostID" json:"-"`
	User User      `gorm:"foreignKey:UserID" json:"-"`
}

// Notification represents system notifications
type Notification struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"` // assignment, grade, announcement, message
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Reference string    `json:"reference"` // ID of related item
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}

// LessonProgress tracks student progress
type LessonProgress struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	LessonID  string    `json:"lesson_id"`
	Completed bool      `json:"completed"`
	Progress  float64   `json:"progress"` // percentage
	LastView  time.Time `json:"last_view"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User   `gorm:"foreignKey:UserID" json:"-"`
	Lesson Lesson `gorm:"foreignKey:LessonID" json:"-"`
}

// Certificate represents course completion certificates
type Certificate struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	CourseID  string     `json:"course_id"`
	UserID    string     `json:"user_id"`
	Code      string     `gorm:"uniqueIndex" json:"code"`
	IssuedAt  time.Time  `json:"issued_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`

	Course Course `gorm:"foreignKey:CourseID" json:"-"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
}

// BeforeCreate hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

// Apply to all other models
func (c *Course) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

func (m *Module) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}

func (l *Lesson) BeforeCreate(tx *gorm.DB) error {
	if l.ID == "" {
		l.ID = uuid.New().String()
	}
	return nil
}

func (a *Assignment) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}

func (s *Submission) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (g *Grade) BeforeCreate(tx *gorm.DB) error {
	if g.ID == "" {
		g.ID = uuid.New().String()
	}
	return nil
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) error {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	return nil
}

func (qu *Question) BeforeCreate(tx *gorm.DB) error {
	if qu.ID == "" {
		qu.ID = uuid.New().String()
	}
	return nil
}

func (o *Option) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return nil
}

func (qa *QuizAttempt) BeforeCreate(tx *gorm.DB) error {
	if qa.ID == "" {
		qa.ID = uuid.New().String()
	}
	return nil
}

func (qas *QuizAnswer) BeforeCreate(tx *gorm.DB) error {
	if qas.ID == "" {
		qas.ID = uuid.New().String()
	}
	return nil
}

func (e *Enrollment) BeforeCreate(tx *gorm.DB) error {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return nil
}

func (r *Resource) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

func (an *Announcement) BeforeCreate(tx *gorm.DB) error {
	if an.ID == "" {
		an.ID = uuid.New().String()
	}
	return nil
}

func (d *Discussion) BeforeCreate(tx *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.New().String()
	}
	return nil
}

func (fp *ForumPost) BeforeCreate(tx *gorm.DB) error {
	if fp.ID == "" {
		fp.ID = uuid.New().String()
	}
	return nil
}

func (fr *ForumReply) BeforeCreate(tx *gorm.DB) error {
	if fr.ID == "" {
		fr.ID = uuid.New().String()
	}
	return nil
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == "" {
		n.ID = uuid.New().String()
	}
	return nil
}

func (lp *LessonProgress) BeforeCreate(tx *gorm.DB) error {
	if lp.ID == "" {
		lp.ID = uuid.New().String()
	}
	return nil
}

func (c *Certificate) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}
