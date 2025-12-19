# 📮 Postman Setup Guide - Complete LMS API

## 🚀 Quick Start

### Step 1: Import Collection

1. Open **Postman**
2. Click **Import** (top left)
3. Select **File** tab
4. Choose: `/Users/abdelkaderbouzomita/Sites/lms-go/postman_collection_complete.json`
5. Click **Import**

---

## ⚙️ Configure Environment Variables

### Automatic (Recommended)

The collection automatically captures tokens after login requests.

### Manual Setup

1. Click **Environments** (left sidebar)
2. Create New → **LMS API**
3. Add these variables:

| Variable           | Initial Value           | Type   |
| ------------------ | ----------------------- | ------ |
| `base_url`         | `http://localhost:8080` | string |
| `token`            | _(empty)_               | string |
| `user_id`          | _(empty)_               | string |
| `instructor_token` | _(empty)_               | string |
| `student_token`    | _(empty)_               | string |

4. Click **Save**
5. **Select Environment**: Top right, choose "LMS API"

---

## 🔐 Authentication Flow

### Step 1: Login First

1. Go to **Authentication** folder
2. Run **Login - Admin** request
3. ✅ Token automatically saved to `{{token}}` variable

### Step 2: Use Other Endpoints

All subsequent requests use `{{token}}` automatically via Bearer Auth.

---

## 📡 CORS Headers (Configured)

Your collection already includes CORS-compatible headers:

- ✅ `Content-Type: application/json`
- ✅ `Access-Control-Allow-Origin: *` (on applicable requests)
- ✅ `Authorization: Bearer {{token}}}` (automatically added)

**Server CORS Configuration:**

```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: POST, OPTIONS, GET, PUT, DELETE, PATCH
Access-Control-Allow-Headers: Content-Type, Authorization, X-Requested-With
```

---

## 🧪 Testing Order

### 1. Authentication (Required First)

```
✓ Login - Admin
✓ Login - Instructor
✓ Login - Student
```

### 2. User Management

```
✓ Get User Profile
✓ Update User Profile
✓ Get All Instructors
✓ Get All Students
```

### 3. Admin Operations

```
✓ Get All Users
✓ Get Admin Dashboard
✓ Get Admin Reports
```

### 4. Courses

```
✓ Get All Courses
✓ Get Course by ID
✓ Enroll in Course
✓ Get Course Stats
```

### 5. Modules & Lessons

```
✓ Get Course Modules
✓ Get Module Lessons
✓ Get Lesson by ID
```

### 6. Assignments

```
✓ Get Course Assignments
✓ Submit Assignment
✓ Get My Submission
```

### 7. Quizzes

```
✓ Get Course Quizzes
✓ Get Quiz Questions
✓ Start Quiz Attempt
✓ Submit Quiz Attempt
```

### 8. Announcements & Discussions

```
✓ Get Announcements
✓ Create Discussion Post
✓ Reply to Post
```

### 9. Grades & Certificates

```
✓ Get My Grades
✓ Get Certificate
```

---

## 🔧 Common Issues & Solutions

### ❌ Error: "401 Unauthorized"

**Solution:**

1. Run **Login - Admin** first
2. Check token is saved: Click **Environments** and verify `token` is not empty
3. Select the environment at top right

### ❌ Error: "CORS Error" or "Blocked by CORS"

**Solution:**

1. ✅ Server CORS is already configured (allows all origins)
2. Try these Postman settings:
   - Settings (gear icon)
   - General → Disable SSL certificate verification (toggle on)
   - Headers → Check "Auto add request headers" is ON

### ❌ Error: "404 Not Found"

**Solution:**

1. Make sure server is running: `ps aux | grep lms-server`
2. If not running: `cd /Users/abdelkaderbouzomita/Sites/lms-go && ./lms-server > server.log 2>&1 &`
3. Verify `base_url` is `http://localhost:8080`

### ❌ Error: "Cannot POST /api/..."

**Solution:**

1. Check the HTTP method is correct (POST, GET, PUT, DELETE)
2. Click on the request → Check **Headers** tab
3. Ensure `Content-Type: application/json` is set

---

## 📋 Test Data Available

### Test Credentials:

```
Admin:      admin@example.com / admin123
Instructor: instructor1@example.com / instructor123
Student:    student1@example.com / student123
```

### Database Content:

- 7 Users (1 admin, 2 instructors, 4 students)
- 3 Courses (Beginner, Intermediate, Advanced)
- 12 Enrollments
- 2 Assignments with submissions
- 1 Quiz with 2 questions
- 4 Grades recorded

---

## 💡 Tips & Tricks

### Run Collection as Test Suite

1. Click **Collection** (postman_collection_complete)
2. Click **...** menu → **Run**
3. Select environment: **LMS API**
4. Click **Run Collection**
5. Watch all tests execute automatically!

### Save Response as Variable

1. Click on any request
2. Go to **Tests** tab
3. Already configured! (Login requests save token automatically)

### View Request/Response

1. After running request, see **Response** tab
2. Click **Pretty** to format JSON nicely
3. Click **Headers** to see response headers

### Debug Requests

1. Click **Request** → **Cookies** to see stored cookies
2. Console (bottom left) shows all request/response details
3. Use **Pre-request Script** tab for custom logic

---

## 🌐 Browser Usage

If testing from browser (JavaScript):

```javascript
const response = await fetch('http://localhost:8080/api/courses', {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
    Authorization: 'Bearer YOUR_TOKEN_HERE',
  },
});
```

The server CORS is configured to accept browser requests!

---

## 📚 All Available Endpoints (78 Total)

### Authentication (4)

- POST /api/public/auth/register
- POST /api/public/auth/login

### Users (4)

- GET /api/users/profile
- PUT /api/users/profile
- GET /api/users/instructors
- GET /api/users/students

### Admin (6)

- GET /api/admin/users
- GET /api/admin/users/:id
- PUT /api/admin/users/:id
- DELETE /api/admin/users/:id
- GET /api/admin/dashboard
- GET /api/admin/reports

### Courses (13)

- GET /api/courses
- POST /api/courses
- GET /api/courses/:id
- PUT /api/courses/:id
- DELETE /api/courses/:id
- POST /api/courses/:id/enroll
- DELETE /api/courses/:id/unenroll
- GET /api/courses/:id/enrollments
- GET /api/courses/:id/enrollment-status
- GET /api/courses/:id/stats
- GET /api/courses/instructor/list
- GET /api/courses/student/list
- PUT /api/courses/enrollments/:id

### Modules & Lessons (9)

- POST /api/courses/:id/modules
- GET /api/courses/:id/modules
- PUT /api/courses/modules/:id
- DELETE /api/courses/modules/:id
- POST /api/courses/modules/:id/lessons
- GET /api/courses/modules/:id/lessons
- GET /api/courses/lessons/:id
- PUT /api/courses/lessons/:id
- DELETE /api/courses/lessons/:id
- GET /api/courses/lessons/:id/progress
- POST /api/courses/lessons/:id/progress
- GET /api/courses/lessons/:id/resources
- POST /api/courses/lessons/:id/resources

### Assignments (9)

- POST /api/courses/:id/assignments
- GET /api/courses/:id/assignments
- GET /api/courses/assignments/:id
- PUT /api/courses/assignments/:id
- DELETE /api/courses/assignments/:id
- POST /api/courses/assignments/:id/submit
- GET /api/courses/assignments/:id/submissions
- GET /api/courses/assignments/:id/my-submission
- GET /api/courses/assignments/:id/stats

### Submissions (3)

- PUT /api/courses/submissions/:id
- POST /api/courses/submissions/:id/grade
- DELETE /api/courses/submissions/:id
- GET /api/courses/submissions/:id/comments
- POST /api/courses/submissions/:id/comments

### Quizzes & Questions (14)

- POST /api/courses/:id/quizzes
- GET /api/courses/:id/quizzes
- GET /api/courses/quizzes/:id
- PUT /api/courses/quizzes/:id
- DELETE /api/courses/quizzes/:id
- POST /api/courses/quizzes/:id/questions
- GET /api/courses/quizzes/:id/questions
- PUT /api/courses/questions/:id
- DELETE /api/courses/questions/:id
- POST /api/courses/quizzes/:id/start
- GET /api/courses/quizzes/:id/attempts
- POST /api/courses/attempts/:id/submit
- GET /api/courses/attempts/:id

### Announcements & Discussions (8)

- PUT /api/courses/announcements/:id
- DELETE /api/courses/announcements/:id
- GET /api/courses/posts/:id/replies
- POST /api/courses/posts/:id/replies
- PUT /api/courses/notifications/:id/read
- GET /api/courses/notifications
- POST /api/courses/:id/announcements
- GET /api/courses/:id/announcements
- POST /api/courses/:id/discussions
- GET /api/courses/:id/discussions

### Grades & Certificates (3)

- GET /api/courses/:id/grades
- GET /api/courses/:id/my-grades
- GET /api/courses/:id/certificate

---

## 🎯 Next Steps

1. ✅ Import collection
2. ✅ Set environment variables
3. ✅ Run "Login - Admin"
4. ✅ Test other endpoints
5. ✅ Create your own requests
6. ✅ Save custom requests to collection

**Happy Testing! 🚀**
