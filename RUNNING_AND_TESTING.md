# LMS Go Backend - Running & Testing Guide

## 🚀 Current Status

✅ **Server Status**: RUNNING

- **Process ID**: 54051
- **Port**: 8080
- **URL**: http://localhost:8080
- **Database**: PostgreSQL (running in Docker on port 5432)
- **Authentication**: JWT Token-based
- **Test Date**: December 19, 2025

---

## 🛠️ Quick Start

### Check Server Status

```bash
# Check if server is running
lsof -i :8080

# View server logs
tail -f server.log
```

### Restart the Server

```bash
# Build the application
cd /Users/abdelkaderbouzomita/Sites/lms-go
CGO_ENABLED=0 go build -o lms-server .

# Stop any running instance
pkill -f lms-server

# Start the server
./lms-server > server.log 2>&1 &

# Verify it's running
lsof -i :8080
```

---

## � Test Credentials

### Admin Account

- **Email**: `admin@example.com`
- **Password**: `admin123`
- **Role**: admin
- **Access**: Full system access

### Instructor Account

- **Email**: `instructor1@example.com`
- **Password**: `instructor123`
- **Role**: instructor
- **Access**: Course management, grading, assignment creation

### Student Account

- **Email**: `student1@example.com`
- **Password**: `student123`
- **Role**: student
- **Access**: Course enrollment, assignment submission, quiz taking

---

## 🔑 Authentication Flow

### Step 2: Run Tests

```bash
make test
```

**Expected output:**

```
ok      lms-go/internal/config          0.5s
ok      lms-go/internal/handlers        1.0s
ok      lms-go/internal/middleware      0.8s
ok      lms-go/internal/models          1.4s
ok      lms-go/internal/utils           1.8s

PASS - All 40+ tests passing
```

### Step 3: Start the Application

```bash
make run
```

**Expected output:**

```
Server running on port 8080
```

### Step 4: Test an Endpoint

In a new terminal:

```bash
# Login as student
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"student1@example.com","password":"student123"}'
```

**Expected response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "...",
    "email": "student1@example.com",
    "first_name": "Alice",
    "last_name": "Johnson",
    "role": "student"
  }
}
```

---

## 🌐 API Endpoints to Test

### Authentication (No Token Required)

#### 1. **Register User**

```bash
curl -X POST http://localhost:8080/api/public/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe",
    "role": "student"
  }'
```

#### 2. **Login**

```bash
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'
```

Save the token from response for authenticated requests.

---

### User Endpoints (Requires Token)

#### 3. **Get User Profile**

```bash
curl -X GET http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 4. **Update User Profile**

```bash
curl -X PUT http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Updated",
    "last_name": "Name",
    "bio": "Updated bio"
  }'
```

#### 5. **Get All Instructors**

```bash
curl -X GET http://localhost:8080/api/users/instructors \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 6. **Get All Students**

```bash
curl -X GET http://localhost:8080/api/users/students \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

### Course Endpoints (Requires Token)

#### 7. **Get All Courses**

```bash
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 8. **Get Course by ID**

```bash
curl -X GET http://localhost:8080/api/courses/COURSE_ID \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 9. **Create Course (Instructor/Admin only)**

```bash
curl -X POST http://localhost:8080/api/courses \
  -H "Authorization: Bearer INSTRUCTOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Go Course",
    "description": "Learn advanced Go",
    "code": "GO-401",
    "category": "Programming",
    "level": "Advanced",
    "max_students": 40
  }'
```

#### 10. **Enroll in Course (Student only)**

```bash
curl -X POST http://localhost:8080/api/courses/COURSE_ID/enroll \
  -H "Authorization: Bearer STUDENT_TOKEN" \
  -H "Content-Type: application/json"
```

#### 11. **Get Student's Courses**

```bash
curl -X GET http://localhost:8080/api/courses/student/list \
  -H "Authorization: Bearer STUDENT_TOKEN"
```

#### 12. **Get Instructor's Courses**

```bash
curl -X GET http://localhost:8080/api/courses/instructor/list \
  -H "Authorization: Bearer INSTRUCTOR_TOKEN"
```

---

### Module & Lesson Endpoints

#### 13. **Get Course Modules**

```bash
curl -X GET http://localhost:8080/api/courses/COURSE_ID/modules \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 14. **Get Module Lessons**

```bash
curl -X GET http://localhost:8080/api/courses/modules/MODULE_ID/lessons \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 15. **Get Lesson Details**

```bash
curl -X GET http://localhost:8080/api/courses/lessons/LESSON_ID \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 16. **Update Lesson Progress**

```bash
curl -X POST http://localhost:8080/api/courses/lessons/LESSON_ID/progress \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "completed": true,
    "progress": 100
  }'
```

---

### Assignment Endpoints

#### 17. **Get Course Assignments**

```bash
curl -X GET http://localhost:8080/api/courses/COURSE_ID/assignments \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 18. **Get Assignment Details**

```bash
curl -X GET http://localhost:8080/api/courses/assignments/ASSIGNMENT_ID \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 19. **Submit Assignment (Student)**

```bash
curl -X POST http://localhost:8080/api/courses/assignments/ASSIGNMENT_ID/submit \
  -H "Authorization: Bearer STUDENT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "My submission content",
    "file_url": "https://example.com/my-file.pdf"
  }'
```

#### 20. **Get Assignment Submissions (Instructor)**

```bash
curl -X GET http://localhost:8080/api/courses/assignments/ASSIGNMENT_ID/submissions \
  -H "Authorization: Bearer INSTRUCTOR_TOKEN"
```

---

### Quiz Endpoints

#### 21. **Get Course Quizzes**

```bash
curl -X GET http://localhost:8080/api/courses/COURSE_ID/quizzes \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 22. **Get Quiz Details**

```bash
curl -X GET http://localhost:8080/api/courses/quizzes/QUIZ_ID \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 23. **Start Quiz Attempt**

```bash
curl -X POST http://localhost:8080/api/courses/quizzes/QUIZ_ID/attempt \
  -H "Authorization: Bearer STUDENT_TOKEN" \
  -H "Content-Type: application/json"
```

#### 24. **Submit Quiz Answer**

```bash
curl -X POST http://localhost:8080/api/courses/quizzes/attempts/ATTEMPT_ID/answer \
  -H "Authorization: Bearer STUDENT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "question_id": "QUESTION_ID",
    "selected_id": "OPTION_ID"
  }'
```

---

### Discussion/Forum Endpoints

#### 25. **Get Course Discussions**

```bash
curl -X GET http://localhost:8080/api/courses/COURSE_ID/discussions \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 26. **Get Discussion Posts**

```bash
curl -X GET http://localhost:8080/api/courses/discussions/DISCUSSION_ID/posts \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 27. **Create Forum Post**

```bash
curl -X POST http://localhost:8080/api/courses/discussions/DISCUSSION_ID/posts \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "How to get started?",
    "content": "I am new to Go and need guidance"
  }'
```

---

### Admin Endpoints (Admin only)

#### 28. **Get All Users**

```bash
curl -X GET http://localhost:8080/api/admin/users \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

#### 29. **Get User by ID**

```bash
curl -X GET http://localhost:8080/api/admin/users/USER_ID \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

#### 30. **Update User**

```bash
curl -X PUT http://localhost:8080/api/admin/users/USER_ID \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Updated",
    "last_name": "User",
    "active": true
  }'
```

#### 31. **Delete User**

```bash
curl -X DELETE http://localhost:8080/api/admin/users/USER_ID \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

#### 32. **Get Admin Dashboard**

```bash
curl -X GET http://localhost:8080/api/admin/dashboard \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

---

## 🧪 Using Postman (Alternative to curl)

1. **Import Collection**

   - File: `postman_collection.json`
   - Postman → Import → Select file

2. **Set Environment Variables**

   - Create variable: `base_url` = `http://localhost:8080`
   - Create variable: `token` = (leave empty initially)

3. **Login First**

   - Run "Login" request
   - Copy token from response
   - Set `token` variable = copied token

4. **Run Requests**
   - All requests will use `{{base_url}}` and `{{token}}`
   - Replace IDs as needed

---

## 🔍 Testing Workflow

### 1. Initial Setup (First Time)

```bash
# Seed database
make seed

# Run tests
make test

# Start server
make run
```

### 2. In New Terminal - Test Endpoints

```bash
# Login
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}'

# Save token
export TOKEN="<token_from_response>"

# Get courses
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN"
```

### 3. Monitor Logs

- Server logs appear in main terminal
- Check for errors or unexpected behavior

### 4. Stop Server

```bash
# Press Ctrl+C in server terminal
```

### 5. Run Tests Again

```bash
make test
```

---

## 🛠️ Useful Endpoints for Development

### Quick Test Flow

**Step 1: Get Token**

```bash
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"student1@example.com","password":"student123"}' | jq '.token'
```

**Step 2: Get Courses**

```bash
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

**Step 3: Enroll in First Course**

```bash
COURSE_ID=$(curl -s http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" | jq -r '.[0].id')

curl -X POST http://localhost:8080/api/courses/$COURSE_ID/enroll \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

**Step 4: Get Student Courses**

```bash
curl -X GET http://localhost:8080/api/courses/student/list \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

---

## 📊 Test Data Available

After seeding, you have:

### Users

- Admin: admin@example.com
- Instructor 1: instructor1@example.com
- Instructor 2: instructor2@example.com
- Students: student1-4@example.com (all same password: student123)

### Courses

1. GO-101: Introduction to Go Programming (Beginner)
2. GO-201: Web Development with Go (Intermediate)
3. GO-301: Advanced Go Concurrency (Advanced)

### Content

- Each course has 2-3 modules
- Each module has 3 lessons
- 2 assignments per course
- 1 quiz per course with 2+ questions
- Discussion forums with sample posts

---

## 🐛 Troubleshooting

### Server Won't Start

**Error: "Failed to connect to database"**

```bash
# Solution: Check PostgreSQL is running
psql --version
createdb lms_db  # Create database if missing

# Or set DATABASE_URL
export DATABASE_URL="postgres://user:password@localhost:5432/lms_db"
make run
```

**Error: "Port 8080 already in use"**

```bash
# Solution: Use different port
export PORT=8081
make run
```

### Login Returns Error

**Error: "Email already registered"**

- Try with different email: newuser@example.com

**Error: "Invalid credentials"**

- Verify email and password match (use seed credentials)
- Check database was seeded: `make seed`

### Token Expired

**Error: "Invalid token"**

- Get new token with login request
- Copy full token value (with quotes if present)

### Database Errors

**Error: "Database doesn't exist"**

```bash
# Create database
createdb lms_db

# Re-seed
make seed
```

**Error: "Column doesn't exist"**

```bash
# Restart migrations
dropdb lms_db
createdb lms_db
make seed
```

---

## 📝 Sample Testing Script

Create `test_endpoints.sh`:

```bash
#!/bin/bash

BASE_URL="http://localhost:8080"
ADMIN_EMAIL="admin@example.com"
ADMIN_PASS="admin123"

echo "🧪 Testing LMS Go Backend"
echo "=========================="

# Login
echo "1️⃣  Logging in..."
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$ADMIN_EMAIL\",\"password\":\"$ADMIN_PASS\"}")

TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.token')
echo "✅ Token: ${TOKEN:0:20}..."

# Get Courses
echo "2️⃣  Getting courses..."
curl -s -X GET $BASE_URL/api/courses \
  -H "Authorization: Bearer $TOKEN" | jq '.[] | {id, title, code}'

# Get Users (Admin)
echo "3️⃣  Getting users..."
curl -s -X GET $BASE_URL/api/admin/users \
  -H "Authorization: Bearer $TOKEN" | jq '.[] | {id, email, role}'

echo "✅ All tests completed!"
```

Run with:

```bash
chmod +x test_endpoints.sh
./test_endpoints.sh
```

---

## 🚀 Next Steps

1. **Test Basic Endpoints**

   - Start server with `make run`
   - Test login and course listing

2. **Test CRUD Operations**

   - Create new course
   - Enroll in course
   - Submit assignment

3. **Test Role-Based Access**

   - Try admin endpoints with student token
   - Verify access denied errors

4. **Test All Roles**

   - Test with admin token
   - Test with instructor token
   - Test with student token

5. **Integration Testing**
   - Create complete workflow
   - Test course completion
   - Test grading system

---

## 📞 Support

For detailed testing guide, see: **TESTING.md**  
For API reference, see: **QUICK_REFERENCE.md**  
For quick commands, see: **QUICK_START.md**

---

**Happy testing! 🎉**
