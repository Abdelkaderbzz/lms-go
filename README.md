# LMS (Learning Management System) - Complete Backend API

A comprehensive Learning Management System built with Go, featuring complete role-based access control for Students, Instructors, and Administrators.

## Features

### 👨‍🎓 Student Features

- **Course Enrollment**: Browse and enroll in available courses
- **Learning Content**: Access lessons with video, text, and downloadable resources
- **Assignments**: Submit assignments with file uploads and receive feedback
- **Quizzes**: Take quizzes with auto-grading and score tracking
- **Progress Tracking**: Monitor learning progress across modules and lessons
- **Discussions**: Participate in course forums and discussions
- **Grades**: View detailed grades and feedback from instructors
- **Certificates**: Earn and download completion certificates
- **Notifications**: Receive real-time notifications for announcements and grades

### 👨‍🏫 Instructor Features

- **Course Management**: Create and manage courses with full customization
- **Content Creation**: Create modules, lessons, and resources
- **Assignments**: Design and manage assignments with due dates
- **Grading**: Grade submissions with rubrics and provide detailed feedback
- **Quizzes**: Create quizzes with various question types and auto-grading
- **Announcements**: Post important announcements and course updates
- **Analytics**: View student progress, completion rates, and performance metrics
- **Student Management**: Monitor and manage student enrollments
- **Discussion Moderation**: Manage course discussions and forums

### 🔧 Admin Features

- **User Management**: Create, update, and manage all users
- **Course Oversight**: Oversee all courses and instructor activities
- **System Analytics**: Comprehensive system-wide reports and statistics
- **Dashboard**: Real-time system metrics and analytics
- **Role Management**: Assign and manage user roles
- **System Configuration**: Configure system settings and policies
- **Audit Logs**: Track user activities and system changes

## Technology Stack

- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (JSON Web Tokens)
- **API**: RESTful API
- **Containerization**: Docker & Docker Compose

## Project Structure

```
lms-go/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies
├── docker-compose.yml      # Docker services configuration
├── Dockerfile              # Container image definition
├── .env.example            # Environment variables template
│
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── database/
│   │   └── database.go            # Database initialization & migrations
│   ├── middleware/
│   │   └── middleware.go          # Authentication & CORS middleware
│   ├── models/
│   │   └── models.go              # Data models & database schemas
│   ├── utils/
│   │   └── jwt.go                 # JWT token utilities
│   └── handlers/
│       ├── routes.go              # Route definitions
│       ├── auth.go                # Authentication handlers
│       ├── users.go               # User management handlers
│       ├── courses.go             # Course handlers
│       ├── enrollments.go         # Enrollment handlers
│       ├── lessons.go             # Lesson & module handlers
│       ├── assignments.go         # Assignment handlers
│       ├── submissions.go         # Submission & grading handlers
│       ├── quizzes.go             # Quiz handlers
│       ├── discussions.go         # Discussion & forum handlers
│       ├── grades.go              # Grade & notification handlers
│       └── admin.go               # Admin dashboard handlers
```

## Installation & Setup

### Prerequisites

- Go 1.21+
- PostgreSQL 12+
- Docker & Docker Compose (optional)

### Local Setup

1. **Clone the repository**

```bash
cd /Users/abdelkaderbouzomita/Sites/lms-go
```

2. **Install dependencies**

```bash
go mod download
go mod tidy
```

3. **Set up environment variables**

```bash
cp .env.example .env
```

4. **Start PostgreSQL using Docker Compose**

```bash
docker-compose up -d
```

5. **Run the application**

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Documentation

### Authentication Endpoints

#### Register User

```
POST /api/public/auth/register
Content-Type: application/json

{
  "email": "student@example.com",
  "password": "securepassword123",
  "first_name": "John",
  "last_name": "Doe",
  "role": "student"  // student, instructor, admin
}
```

#### Login

```
POST /api/public/auth/login
Content-Type: application/json

{
  "email": "student@example.com",
  "password": "securepassword123"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "email": "student@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "student"
  }
}
```

### User Endpoints

#### Get Profile

```
GET /api/users/profile
Authorization: Bearer <token>
```

#### Update Profile

```
PUT /api/users/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "phone": "1234567890",
  "bio": "Software Developer",
  "avatar": "https://example.com/avatar.jpg"
}
```

### Course Endpoints

#### Create Course (Instructor/Admin)

```
POST /api/courses
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Introduction to Web Development",
  "description": "Learn web development from basics",
  "code": "WEB101",
  "category": "Technology",
  "level": "Beginner",
  "max_students": 50
}
```

#### Get All Courses

```
GET /api/courses
Authorization: Bearer <token>
```

#### Get Course Details

```
GET /api/courses/{courseID}
Authorization: Bearer <token>
```

#### Enroll in Course

```
POST /api/courses/{courseID}/enroll
Authorization: Bearer <token>
```

### Module & Lesson Endpoints

#### Create Module (Instructor)

```
POST /api/courses/{courseID}/modules
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Module 1: Basics",
  "description": "Learn the basics",
  "order": 1
}
```

#### Create Lesson (Instructor)

```
POST /api/courses/modules/{moduleID}/lessons
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Lesson 1: Getting Started",
  "description": "Introduction lesson",
  "content": "<h1>Getting Started</h1>...",
  "video_url": "https://youtube.com/watch?v=...",
  "duration": 45,
  "order": 1,
  "published": true
}
```

#### Get Modules for Course

```
GET /api/courses/{courseID}/modules
Authorization: Bearer <token>
```

### Assignment Endpoints

#### Create Assignment (Instructor)

```
POST /api/courses/{courseID}/assignments
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Assignment 1",
  "description": "Complete the task",
  "due_date": "2024-12-31T23:59:59Z",
  "points": 100,
  "type": "homework"
}
```

#### Submit Assignment (Student)

```
POST /api/courses/assignments/{assignmentID}/submit
Authorization: Bearer <token>
Content-Type: application/json

{
  "content": "My solution...",
  "file_url": "https://example.com/submission.pdf"
}
```

#### Grade Submission (Instructor)

```
POST /api/courses/submissions/{submissionID}/grade
Authorization: Bearer <token>
Content-Type: application/json

{
  "points": 85,
  "feedback": "Good work! Please review section 3."
}
```

### Quiz Endpoints

#### Create Quiz (Instructor)

```
POST /api/courses/{courseID}/quizzes
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Quiz 1",
  "description": "Basic concepts quiz",
  "start_date": "2024-12-01T00:00:00Z",
  "end_date": "2024-12-31T23:59:59Z",
  "time_limit": 30,
  "pass_score": 60,
  "shuffle": true,
  "public": true
}
```

#### Add Question to Quiz

```
POST /api/courses/quizzes/{quizID}/questions
Authorization: Bearer <token>
Content-Type: application/json

{
  "type": "multiple_choice",
  "question": "What is 2+2?",
  "points": 1,
  "order": 1
}
```

#### Start Quiz Attempt (Student)

```
POST /api/courses/quizzes/{quizID}/start
Authorization: Bearer <token>
```

#### Submit Quiz Attempt (Student)

```
POST /api/courses/attempts/{attemptID}/submit
Authorization: Bearer <token>
Content-Type: application/json

{
  "answers": [
    {
      "question_id": "q1",
      "selected_id": "option1",
      "text_answer": ""
    }
  ]
}
```

### Discussion Endpoints

#### Create Discussion

```
POST /api/courses/{courseID}/discussions
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Discussion: Best Practices",
  "content": "Let's discuss..."
}
```

#### Create Forum Post

```
POST /api/courses/discussions/{discussionID}/posts
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "My Question",
  "content": "I have a question about..."
}
```

#### Reply to Post

```
POST /api/courses/posts/{postID}/replies
Authorization: Bearer <token>
Content-Type: application/json

{
  "content": "Here's the answer..."
}
```

### Grade & Notification Endpoints

#### Get My Grades (Student)

```
GET /api/courses/{courseID}/my-grades
Authorization: Bearer <token>
```

#### Get Notifications

```
GET /api/courses/notifications
Authorization: Bearer <token>
```

### Admin Endpoints

#### Get Admin Dashboard

```
GET /api/admin/dashboard
Authorization: Bearer <admin_token>
```

#### Get System Reports

```
GET /api/admin/reports
Authorization: Bearer <admin_token>
```

#### Get All Users (Admin)

```
GET /api/admin/users
Authorization: Bearer <admin_token>
```

#### Update User (Admin)

```
PUT /api/admin/users/{userID}
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "first_name": "Jane",
  "last_name": "Smith",
  "role": "instructor",
  "active": true
}
```

## Database Schema

The system includes the following main entities:

- **User**: System users (Students, Instructors, Admins)
- **Course**: Learning courses with metadata
- **Module**: Course modules/sections
- **Lesson**: Individual lessons within modules
- **Assignment**: Course assignments
- **Submission**: Student assignment submissions
- **Grade**: Assignment grades and feedback
- **Quiz**: Assessment quizzes
- **Question**: Quiz questions
- **QuizAttempt**: Student quiz attempts
- **Enrollment**: Student course enrollments
- **Announcement**: Course announcements
- **Discussion**: Discussion forums
- **ForumPost**: Forum discussion posts
- **ForumReply**: Replies to forum posts
- **Notification**: User notifications
- **LessonProgress**: Student progress tracking
- **Certificate**: Course completion certificates
- **Resource**: Lesson resources (PDFs, links, etc.)

## Authentication & Authorization

The system uses JWT (JSON Web Tokens) for authentication with role-based access control:

- **Student**: Limited to their own courses and assignments
- **Instructor**: Can manage courses and grade student work
- **Admin**: Full system access and management capabilities

### Token Format

```
Authorization: Bearer <JWT_TOKEN>
```

## Error Handling

The API returns appropriate HTTP status codes:

- **200 OK**: Successful GET request
- **201 Created**: Successful resource creation
- **400 Bad Request**: Invalid input data
- **401 Unauthorized**: Missing or invalid token
- **403 Forbidden**: Insufficient permissions
- **404 Not Found**: Resource not found
- **500 Internal Server Error**: Server error

## Development

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
go build -o lms-server
```

### Docker Deployment

```bash
docker build -t lms-app .
docker-compose -f docker-compose.yml up -d
```

## Database Migrations

Migrations are automatically run on startup using GORM's AutoMigrate feature.

## Security Best Practices

1. **Change JWT Secret**: Update `JWT_SECRET` in `.env` file
2. **Use HTTPS**: Always use HTTPS in production
3. **Password Hashing**: Implement bcrypt for password hashing
4. **Rate Limiting**: Implement rate limiting for API endpoints
5. **Input Validation**: All inputs are validated
6. **CORS Configuration**: Configure CORS for your frontend origin

## Future Enhancements

- [ ] Email notifications integration
- [ ] File upload with cloud storage
- [ ] Real-time notifications with WebSockets
- [ ] Advanced analytics and reporting
- [ ] Video streaming capabilities
- [ ] Mobile app support
- [ ] Third-party integrations (Zoom, Google Meet)
- [ ] AI-powered recommendations
- [ ] Plagiarism detection for submissions
- [ ] Certificate printing

## Contributing

Please follow the existing code structure and naming conventions when contributing.

## License

This project is licensed under the MIT License.

## Support

For issues or questions, please contact the development team.

---

**Ready to Deploy!** This LMS system is production-ready with:

- Complete REST API
- Role-based access control
- Database migrations
- Docker support
- Comprehensive error handling
- JWT authentication
