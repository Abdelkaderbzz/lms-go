# LMS Go Backend - Testing Guide

This document provides a comprehensive guide for testing the LMS Go Backend project.

## Quick Start

### Running All Tests

```bash
make test
```

### Running Tests with Coverage

```bash
make test-coverage
```

This will generate a coverage report in `coverage.html`.

### Seeding the Database

```bash
make seed
```

This will populate your database with test data including:

- Admin user (admin@example.com / admin123)
- Instructor users (instructor1@example.com / instructor123)
- Student users (student1@example.com / student123)
- 3 sample courses
- Multiple modules, lessons, and assignments
- Quiz questions and options
- Enrollments, submissions, and grades
- Discussion posts and forum replies
- Announcements and notifications

## Test Structure

### Unit Tests

The project includes comprehensive unit tests for:

#### 1. **Models** (`internal/models/models_test.go`)

- UUID generation with BeforeCreate hooks
- User role constants validation
- Model creation and validation
- Course, Module, Lesson, Quiz, and Assignment models
- Course status and enrollment status validation

Tests cover:

```bash
go test -v ./internal/models
```

#### 2. **Utilities** (`internal/utils/jwt_test.go`)

- JWT token generation
- Token validation
- Token expiration handling
- Claims structure validation

Tests cover:

```bash
go test -v ./internal/utils
```

#### 3. **Configuration** (`internal/config/config_test.go`)

- Configuration loading with defaults
- Custom environment variables
- Environment variable helpers

Tests cover:

```bash
go test -v ./internal/config
```

#### 4. **Handlers** (`internal/handlers/auth_test.go`, `internal/handlers/handlers_test.go`)

- Authentication request structures
- Request/Response validation
- Login and registration logic
- Course, enrollment, assignment, quiz structures
- Grade and submission validation

Tests cover:

```bash
go test -v ./internal/handlers
```

#### 5. **Middleware** (`internal/middleware/middleware_test.go`)

- CORS middleware configuration
- Error handling middleware
- Authentication middleware
- Role-based access control (RBAC)
- Middleware chaining
- CORS preflight requests

Tests cover:

```bash
go test -v ./internal/middleware
```

### Test Helpers

A comprehensive test helper file is available at `internal/test/helpers.go` that provides:

- `SetupTestDB(t *testing.T) *gorm.DB` - Initialize test database
- `TeardownTestDB(t *testing.T, db *gorm.DB)` - Clean up test database
- `CreateTestUser()` - Create test users
- `CreateTestCourse()` - Create test courses
- `CreateTestModule()` - Create test modules
- `CreateTestLesson()` - Create test lessons
- `CreateTestAssignment()` - Create test assignments
- `CreateTestSubmission()` - Create test submissions
- `CreateTestQuiz()` - Create test quizzes
- `CreateTestEnrollment()` - Create test enrollments

## Running Specific Tests

### Run tests in specific package

```bash
go test -v ./internal/models
go test -v ./internal/utils
go test -v ./internal/config
go test -v ./internal/handlers
go test -v ./internal/middleware
```

### Run specific test

```bash
go test -v -run TestUserBeforeCreate ./internal/models
go test -v -run TestGenerateToken ./internal/utils
```

### Run tests with verbose output

```bash
go test -v -cover ./...
```

### Run tests with race detection

```bash
go test -race ./...
```

## Seeded Data Overview

After running `make seed`, the database contains:

### Users

- **Admin**: admin@example.com (password: admin123)
- **Instructors**:
  - instructor1@example.com (password: instructor123)
  - instructor2@example.com (password: instructor123)
- **Students** (4 total):
  - student1@example.com through student4@example.com (password: student123)

### Courses (3)

1. **Introduction to Go Programming** (GO-101)

   - Beginner level
   - 3-month duration
   - Max 50 students

2. **Advanced Go Concurrency** (GO-301)

   - Advanced level
   - 3-month duration
   - Max 30 students

3. **Web Development with Go** (GO-201)
   - Intermediate level
   - 4-month duration
   - Max 40 students

### Course Content

Each course includes:

- 2-3 Modules
- 3 Lessons per module
- 2 Assignments
- 1 Quiz with 2 questions
- Announcements
- Discussion forums with posts and replies
- Student enrollments and submissions
- Grades for submissions

## Testing Workflow

### 1. Setup Test Database

```bash
export TEST_DATABASE_URL="postgres://user:password@localhost:5432/lms_test_db"
```

### 2. Run Unit Tests

```bash
make test
```

### 3. Check Coverage

```bash
make test-coverage
```

This opens the coverage report at `coverage.html`.

### 4. Seed Database for Integration Testing

```bash
make seed
```

### 5. Start Application

```bash
make run
```

### 6. Test Endpoints

Use Postman or curl to test API endpoints:

```bash
# Login as student
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"student1@example.com","password":"student123"}'

# Get all courses
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer <token>"
```

## Coverage Goals

The test suite aims for:

- **Model tests**: 100% coverage of BeforeCreate hooks and constants
- **Utility tests**: 100% coverage of JWT and config functions
- **Handler tests**: Core structure and validation tests
- **Middleware tests**: CORS, auth, error handling, and RBAC

## Continuous Integration

For CI/CD pipelines, use:

```bash
# Run all tests with coverage
go test -v -coverprofile=coverage.out ./...

# Generate coverage report
go tool cover -html=coverage.out -o coverage.html

# Run with race detection
go test -race ./...

# Format and lint check
go fmt ./...
golangci-lint run ./...
```

## Common Issues and Solutions

### Issue: Tests fail to connect to database

**Solution**: Ensure PostgreSQL is running and TEST_DATABASE_URL is set correctly

```bash
export TEST_DATABASE_URL="postgres://user:password@localhost:5432/lms_test_db"
```

### Issue: Tests timeout

**Solution**: Increase test timeout

```bash
go test -timeout 30s ./...
```

### Issue: Port already in use

**Solution**: Run tests with a different port

```bash
export PORT=8081
make test
```

### Issue: Seeding fails

**Solution**: Ensure migrations ran successfully and database is accessible

```bash
go run main.go migrate
make seed
```

## Next Steps

To further improve test coverage:

1. Add database integration tests for CRUD operations
2. Add end-to-end API tests with authentication
3. Add performance benchmarks for critical paths
4. Add stress testing for concurrent operations
5. Add security testing for RBAC and authorization

## Resources

- [Go Testing Documentation](https://golang.org/doc/effective_go#testing)
- [Gin Testing Guide](https://github.com/gin-gonic/gin#testing)
- [GORM Testing](https://gorm.io/docs/index.html)
- [testify/assert Package](https://github.com/stretchr/testify)
