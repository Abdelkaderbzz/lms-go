# LMS API Testing Guide

## Quick Start

### 1. Setup & Installation

```bash
# Navigate to project
cd /Users/abdelkaderbouzomita/Sites/lms-go

# Copy environment file
cp .env.example .env

# Start PostgreSQL
docker-compose up -d

# Download dependencies
go mod download
go mod tidy

# Run the server
go run main.go
```

Server will be running at `http://localhost:8080`

## Testing the API

### Using cURL

#### 1. Register a Student Account

```bash
curl -X POST http://localhost:8080/api/public/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "student@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe",
    "role": "student"
  }'
```

#### 2. Register an Instructor Account

```bash
curl -X POST http://localhost:8080/api/public/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "instructor@example.com",
    "password": "password123",
    "first_name": "Jane",
    "last_name": "Smith",
    "role": "instructor"
  }'
```

#### 3. Login

```bash
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "student@example.com",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "student@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "student"
  }
}
```

Save the token for subsequent requests:
```bash
TOKEN="eyJhbGciOiJIUzI1NiIs..."
```

#### 4. Get User Profile

```bash
curl -X GET http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer $TOKEN"
```

#### 5. Create a Course (as Instructor)

```bash
curl -X POST http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Introduction to Web Development",
    "description": "Learn the basics of web development",
    "code": "WEB101",
    "category": "Technology",
    "level": "Beginner",
    "max_students": 50
  }'
```

#### 6. Get All Courses

```bash
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN"
```

#### 7. Enroll in a Course (as Student)

```bash
COURSE_ID="course-uuid-here"

curl -X POST "http://localhost:8080/api/courses/$COURSE_ID/enroll" \
  -H "Authorization: Bearer $TOKEN"
```

#### 8. Create a Module (as Instructor)

```bash
COURSE_ID="course-uuid-here"

curl -X POST "http://localhost:8080/api/courses/$COURSE_ID/modules" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Module 1: Getting Started",
    "description": "Learn the basics",
    "order": 1
  }'
```

#### 9. Create a Lesson (as Instructor)

```bash
MODULE_ID="module-uuid-here"

curl -X POST "http://localhost:8080/api/courses/modules/$MODULE_ID/lessons" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Lesson 1: HTML Basics",
    "description": "Learn HTML fundamentals",
    "content": "<h1>HTML Basics</h1><p>HTML is...</p>",
    "video_url": "https://youtube.com/watch?v=dQw4w9WgXcQ",
    "duration": 45,
    "order": 1,
    "published": true
  }'
```

#### 10. Create an Assignment (as Instructor)

```bash
COURSE_ID="course-uuid-here"

curl -X POST "http://localhost:8080/api/courses/$COURSE_ID/assignments" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Assignment 1: Create a Website",
    "description": "Create a simple website using HTML and CSS",
    "due_date": "2024-12-31T23:59:59Z",
    "points": 100,
    "type": "project"
  }'
```

#### 11. Submit Assignment (as Student)

```bash
ASSIGNMENT_ID="assignment-uuid-here"

curl -X POST "http://localhost:8080/api/courses/assignments/$ASSIGNMENT_ID/submit" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Here is my solution...",
    "file_url": "https://example.com/my-website.zip"
  }'
```

#### 12. Grade a Submission (as Instructor)

```bash
SUBMISSION_ID="submission-uuid-here"

curl -X POST "http://localhost:8080/api/courses/submissions/$SUBMISSION_ID/grade" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "points": 95,
    "feedback": "Excellent work! Great use of CSS. Consider adding more interactivity."
  }'
```

#### 13. Create a Quiz (as Instructor)

```bash
COURSE_ID="course-uuid-here"

curl -X POST "http://localhost:8080/api/courses/$COURSE_ID/quizzes" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Quiz 1: HTML Basics",
    "description": "Test your HTML knowledge",
    "start_date": "2024-12-01T00:00:00Z",
    "end_date": "2024-12-31T23:59:59Z",
    "time_limit": 30,
    "pass_score": 70,
    "shuffle": true,
    "public": true
  }'
```

#### 14. Add Quiz Question (as Instructor)

```bash
QUIZ_ID="quiz-uuid-here"

curl -X POST "http://localhost:8080/api/courses/quizzes/$QUIZ_ID/questions" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "multiple_choice",
    "question": "What does HTML stand for?",
    "points": 1,
    "order": 1
  }'
```

#### 15. Start Quiz (as Student)

```bash
QUIZ_ID="quiz-uuid-here"

curl -X POST "http://localhost:8080/api/courses/quizzes/$QUIZ_ID/start" \
  -H "Authorization: Bearer $TOKEN"
```

**Response includes attempt ID**

#### 16. Submit Quiz (as Student)

```bash
ATTEMPT_ID="attempt-uuid-here"

curl -X POST "http://localhost:8080/api/courses/attempts/$ATTEMPT_ID/submit" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "answers": [
      {
        "question_id": "q1",
        "selected_id": "opt1",
        "text_answer": ""
      }
    ]
  }'
```

#### 17. Create Announcement (as Instructor)

```bash
COURSE_ID="course-uuid-here"

curl -X POST "http://localhost:8080/api/courses/$COURSE_ID/announcements" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Important Notice",
    "content": "The assignment deadline has been extended to next Friday.",
    "important": true
  }'
```

#### 18. Create Discussion (as Student)

```bash
COURSE_ID="course-uuid-here"

curl -X POST "http://localhost:8080/api/courses/$COURSE_ID/discussions" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Discussion: Best Practices",
    "content": "Let'\''s discuss the best practices for web development."
  }'
```

#### 19. Get Grades (as Student)

```bash
COURSE_ID="course-uuid-here"

curl -X GET "http://localhost:8080/api/courses/$COURSE_ID/my-grades" \
  -H "Authorization: Bearer $TOKEN"
```

#### 20. Get Notifications

```bash
curl -X GET http://localhost:8080/api/courses/notifications \
  -H "Authorization: Bearer $TOKEN"
```

### Using Postman

1. Create a new collection "LMS API"
2. Add a variable `base_url` = `http://localhost:8080`
3. Add a variable `token` = (leave empty, will be filled after login)

**Login Request:**
```
POST {{base_url}}/api/public/auth/login
Content-Type: application/json

{
  "email": "student@example.com",
  "password": "password123"
}
```

In Tests tab, add:
```javascript
if (pm.response.code === 200) {
    pm.environment.set("token", pm.response.json().token);
}
```

4. Use `Authorization: Bearer {{token}}` in subsequent requests

## API Response Examples

### Success Response (201 Created)
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Introduction to Web Development",
  "description": "Learn the basics of web development",
  "code": "WEB101",
  "category": "Technology",
  "level": "Beginner",
  "status": "draft",
  "max_students": 50,
  "created_at": "2024-12-19T10:30:00Z",
  "updated_at": "2024-12-19T10:30:00Z"
}
```

### Error Response (400 Bad Request)
```json
{
  "error": "Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"
}
```

### Unauthorized Response (401)
```json
{
  "error": "Missing authorization header"
}
```

### Forbidden Response (403)
```json
{
  "error": "Insufficient permissions"
}
```

## Testing Scenarios

### Scenario 1: Complete Student Journey
1. Register as student
2. Login and get token
3. Browse available courses
4. Enroll in a course
5. View course modules and lessons
6. Update lesson progress
7. Submit assignments
8. Participate in discussions
9. Take quizzes
10. View grades and feedback
11. Download certificate (if completed)

### Scenario 2: Instructor Course Creation
1. Register as instructor
2. Login and get token
3. Create a new course
4. Create course modules
5. Add lessons with resources
6. Create assignments with due dates
7. Create quizzes
8. View enrolled students
9. Grade student submissions
10. View course analytics

### Scenario 3: Admin Management
1. Register as admin (requires system access)
2. View all users
3. View system dashboard
4. Generate reports
5. Manage user accounts
6. Monitor system health

## Troubleshooting

### Issue: 401 Unauthorized
- Check if token is included in headers
- Check if token is expired
- Regenerate token by logging in again

### Issue: 403 Forbidden
- Verify user role has permission for the action
- Students cannot grade submissions
- Instructors cannot manage other courses

### Issue: 404 Not Found
- Verify course/assignment/quiz ID is correct
- Check if resource exists in database
- Ensure correct route path

### Issue: 500 Internal Server Error
- Check server logs
- Verify database connection
- Check request body format

## Performance Testing

```bash
# Test with Apache Bench (50 requests, 10 concurrent)
ab -n 50 -c 10 http://localhost:8080/api/courses

# Test with wrk
wrk -t12 -c400 -d30s http://localhost:8080/api/courses
```

## Database Inspection

```bash
# Connect to PostgreSQL
psql postgres://user:password@localhost:5432/lms_db

# List all tables
\dt

# Query users
SELECT id, email, role, created_at FROM users LIMIT 10;

# Query courses
SELECT id, title, code, status FROM courses;

# Query enrollments
SELECT * FROM enrollments;
```

## Docker Commands

```bash
# View logs
docker-compose logs -f postgres

# Restart services
docker-compose restart

# Stop services
docker-compose down

# Remove volumes (WARNING: deletes data)
docker-compose down -v
```

## Next Steps

1. Implement password hashing with bcrypt
2. Add file upload functionality
3. Implement email notifications
4. Add rate limiting
5. Implement caching with Redis
6. Add comprehensive logging
7. Create frontend React application
8. Deploy to cloud (AWS, GCP, Azure)
