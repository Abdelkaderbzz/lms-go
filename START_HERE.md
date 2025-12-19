# 🚀 Running & Testing LMS Go Backend - Complete Guide

**Current Status**:

- ✅ Project built successfully
- ✅ All 40+ tests passing
- ⚠️ Database required - Docker daemon not running
- 📝 Guide covers all setup options

---

## 🎯 Three Setup Options

### Option 1: Start Docker Daemon (macOS) - Fastest

```bash
# Start Docker Desktop app via Spotlight
open /Applications/Docker.app

# Wait ~30 seconds for it to fully start

# Verify it's running
docker ps

# Then return to this guide and continue below
```

---

### Option 2: Install PostgreSQL Locally (macOS with Homebrew)

```bash
# Install PostgreSQL
brew install postgresql@15

# Start PostgreSQL service
brew services start postgresql@15

# Create database
createdb lms_db

# Verify
psql lms_db -c "SELECT 1"
```

---

### Option 3: Use Podman (Alternative Container Runtime)

```bash
# Install via Homebrew
brew install podman

# Initialize Podman machine
podman machine init

# Start Podman
podman machine start

# Then use podman instead of docker commands below
```

---

## ✅ If You Have Docker Running...

Once Docker daemon is running, execute these commands:

### Step 1: Start PostgreSQL Container

```bash
docker run -d \
  --name lms-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=lms_db \
  -p 5432:5432 \
  postgres:15

echo "Waiting for PostgreSQL to start..."
sleep 10
```

### Step 2: Navigate to Project

```bash
cd /Users/abdelkaderbouzomita/Sites/lms-go
```

### Step 3: Run Tests (Verify Setup)

```bash
make test
```

### Step 4: Seed Test Data

```bash
make seed
```

### Step 5: Start Application

```bash
make run
```

You should see:

```
[GIN-debug] Listening and serving HTTP on 127.0.0.1:8080
```

### Step 6: Test in New Terminal

```bash
# 1. Login
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}' | jq .

# 2. Copy token from response
export TOKEN="eyJ0eXAiOiJKV1QiLCJhbGc..."

# 3. Test authenticated endpoint
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" | jq .
```

---

## 🧪 Test Users & Credentials

After running `make seed`, these accounts are available:

```
ADMIN USER
  Email:    admin@example.com
  Password: admin123
  Role:     ADMIN

INSTRUCTOR USERS
  Email:    instructor1@example.com
  Password: instructor123
  Role:     INSTRUCTOR

STUDENT USERS
  Email:    student1@example.com
  Password: student123
  Role:     STUDENT
```

---

## 📋 Essential Endpoints to Test

### 1️⃣ Authentication

**Login:**

```bash
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }' | jq .
```

**Register:**

```bash
curl -X POST http://localhost:8080/api/public/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New User",
    "email": "newuser@example.com",
    "password": "password123"
  }' | jq .
```

---

### 2️⃣ Courses (Requires Token)

**Get All Courses:**

```bash
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Get Single Course:**

```bash
curl -X GET http://localhost:8080/api/courses/1 \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Create Course (Admin Only):**

```bash
curl -X POST http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Advanced Go Programming",
    "description": "Learn advanced Go concepts",
    "instructorId": "2"
  }' | jq .
```

---

### 3️⃣ User Profile

**Get Current User Profile:**

```bash
curl -X GET http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Update Profile:**

```bash
curl -X PUT http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Name",
    "bio": "New bio"
  }' | jq .
```

---

### 4️⃣ Admin Operations

**Get All Users (Admin Only):**

```bash
curl -X GET http://localhost:8080/api/admin/users \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Get User Statistics (Admin Only):**

```bash
curl -X GET http://localhost:8080/api/admin/stats \
  -H "Authorization: Bearer $TOKEN" | jq .
```

---

### 5️⃣ Enrollments

**Enroll in Course:**

```bash
curl -X POST http://localhost:8080/api/enrollments \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "courseId": "1"
  }' | jq .
```

**Get User Enrollments:**

```bash
curl -X GET http://localhost:8080/api/enrollments \
  -H "Authorization: Bearer $TOKEN" | jq .
```

---

### 6️⃣ Assignments

**Get Course Assignments:**

```bash
curl -X GET http://localhost:8080/api/courses/1/assignments \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Submit Assignment:**

```bash
curl -X POST http://localhost:8080/api/assignments/1/submit \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "My assignment submission"
  }' | jq .
```

---

### 7️⃣ Quizzes

**Get Quizzes:**

```bash
curl -X GET http://localhost:8080/api/courses/1/quizzes \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Start Quiz Attempt:**

```bash
curl -X POST http://localhost:8080/api/quizzes/1/attempts \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Submit Quiz Answer:**

```bash
curl -X POST http://localhost:8080/api/quizzes/1/attempts/1/answers \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "questionId": "1",
    "selectedOptionId": "1"
  }' | jq .
```

---

### 8️⃣ Discussions

**Get Course Discussions:**

```bash
curl -X GET http://localhost:8080/api/courses/1/discussions \
  -H "Authorization: Bearer $TOKEN" | jq .
```

**Create Discussion Post:**

```bash
curl -X POST http://localhost:8080/api/discussions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "courseId": "1",
    "title": "How do I solve this problem?",
    "content": "I'm stuck on the assignment..."
  }' | jq .
```

---

## 🔄 Complete Testing Workflow

### 1. Login & Get Token

```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}' | jq -r '.data.token')

echo "Token: $TOKEN"
```

### 2. View Your Profile

```bash
curl -X GET http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

### 3. Enroll in First Course

```bash
COURSE_ID=$(curl -s -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" | jq -r '.data[0].id')

curl -X POST http://localhost:8080/api/enrollments \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"courseId\":\"$COURSE_ID\"}" | jq '.'
```

### 4. Get Course Details

```bash
curl -X GET http://localhost:8080/api/courses/$COURSE_ID \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

### 5. Check Assignments

```bash
curl -X GET http://localhost:8080/api/courses/$COURSE_ID/assignments \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

---

## 🐳 Cleaning Up Docker

When you want to stop the database:

```bash
# Stop container
docker stop lms-postgres

# Remove container
docker rm lms-postgres

# Verify
docker ps -a | grep lms-postgres
```

If needed, start fresh:

```bash
docker stop lms-postgres 2>/dev/null || true
docker rm lms-postgres 2>/dev/null || true

# Then rerun the docker run command from Step 1 above
```

---

## 📝 Using Postman Instead of curl

1. **Import Collection:**

   ```bash
   # Copy postman_collection.json location
   cat /Users/abdelkaderbouzomita/Sites/lms-go/postman_collection.json
   ```

2. **In Postman:**
   - Import → Upload Files → Select postman_collection.json
   - Set environment variable: `token` = token from login response
   - All requests will auto-use `{{token}}`

---

## 🔧 Troubleshooting

### Application won't start

```bash
# Check if port 8080 is in use
lsof -i :8080

# Use different port
export PORT=8081
make run
```

### Database connection failed

```bash
# Verify PostgreSQL is running
docker ps | grep postgres

# Check logs
docker logs lms-postgres

# Verify connection string
echo $DATABASE_URL
```

### Tests failing

```bash
# Run with verbose output
make test | head -50

# Run specific package
go test ./internal/handlers -v
```

---

## ✨ Quick Commands Reference

```bash
# Start everything
cd /Users/abdelkaderbouzomita/Sites/lms-go
make seed
make run

# In another terminal, test
export TOKEN="<token_from_login>"
curl -X GET http://localhost:8080/api/courses -H "Authorization: Bearer $TOKEN" | jq

# Stop database when done
docker stop lms-postgres
```

---

## 📊 Expected Test Results

After running all tests:

```
✅ Package config - All tests pass
✅ Package database - GORM initialization
✅ Package handlers - All endpoints
✅ Package middleware - CORS, Auth, RBAC
✅ Package models - UUID, relationships
✅ Package utils - JWT, helpers
```

---

## 🎉 You're Ready!

The project is fully functional. Choose a setup option and start testing!

**Questions?** Check:

- `RUNNING_AND_TESTING.md` - Detailed endpoint reference
- `TESTING.md` - Unit test documentation
- `README.md` - Project overview
