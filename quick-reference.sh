#!/bin/bash
# LMS Go Backend - Quick Reference Commands
# Usage: Source this file or copy individual commands

# ================ CONFIGURATION ================
BASE_URL="http://localhost:8080"
ADMIN_EMAIL="admin@example.com"
ADMIN_PASS="admin123"
INSTRUCTOR_EMAIL="instructor1@example.com"
INSTRUCTOR_PASS="instructor123"
STUDENT_EMAIL="student1@example.com"
STUDENT_PASS="student123"

# ================ TOKEN MANAGEMENT ================

# Get Admin Token
alias admin_login='curl -s -X POST $BASE_URL/api/public/auth/login -H "Content-Type: application/json" -d "{\"email\":\"$ADMIN_EMAIL\",\"password\":\"$ADMIN_PASS\"}" | jq -r ".token"'

# Get Instructor Token
alias inst_login='curl -s -X POST $BASE_URL/api/public/auth/login -H "Content-Type: application/json" -d "{\"email\":\"$INSTRUCTOR_EMAIL\",\"password\":\"$INSTRUCTOR_PASS\"}" | jq -r ".token"'

# Get Student Token
alias student_login='curl -s -X POST $BASE_URL/api/public/auth/login -H "Content-Type: application/json" -d "{\"email\":\"$STUDENT_EMAIL\",\"password\":\"$STUDENT_PASS\"}" | jq -r ".token"'

# Store tokens in variables
get_all_tokens() {
    export ADMIN_TOKEN=$(curl -s -X POST $BASE_URL/api/public/auth/login \
        -H "Content-Type: application/json" \
        -d "{\"email\":\"$ADMIN_EMAIL\",\"password\":\"$ADMIN_PASS\"}" | jq -r '.token')
    
    export INSTRUCTOR_TOKEN=$(curl -s -X POST $BASE_URL/api/public/auth/login \
        -H "Content-Type: application/json" \
        -d "{\"email\":\"$INSTRUCTOR_EMAIL\",\"password\":\"$INSTRUCTOR_PASS\"}" | jq -r '.token')
    
    export STUDENT_TOKEN=$(curl -s -X POST $BASE_URL/api/public/auth/login \
        -H "Content-Type: application/json" \
        -d "{\"email\":\"$STUDENT_EMAIL\",\"password\":\"$STUDENT_PASS\"}" | jq -r '.token')
    
    echo "✓ Tokens loaded into ADMIN_TOKEN, INSTRUCTOR_TOKEN, STUDENT_TOKEN"
}

# ================ SERVER MANAGEMENT ================

# Check server status
server_status() {
    if lsof -i :8080 > /dev/null; then
        echo "✓ Server is running on port 8080"
        lsof -i :8080 | head -2
    else
        echo "✗ Server is not running"
    fi
}

# Start server
start_server() {
    cd /Users/abdelkaderbouzomita/Sites/lms-go
    pkill -f lms-server 2>/dev/null || true
    CGO_ENABLED=0 go build -o lms-server . && ./lms-server > server.log 2>&1 &
    sleep 2
    server_status
}

# Stop server
stop_server() {
    pkill -f lms-server
    echo "Server stopped"
}

# Restart server
restart_server() {
    stop_server
    sleep 1
    start_server
}

# View logs
view_logs() {
    tail -f /Users/abdelkaderbouzomita/Sites/lms-go/server.log
}

# ================ DATABASE MANAGEMENT ================

# Reseed database
reseed_database() {
    cd /Users/abdelkaderbouzomita/Sites/lms-go
    make seed
}

# Connect to database
db_connect() {
    psql -U user -d lms_db -h localhost
}

# Check database status
db_status() {
    if psql -U user -d lms_db -h localhost -c "SELECT 1" > /dev/null 2>&1; then
        echo "✓ Database is connected"
    else
        echo "✗ Database is not accessible"
    fi
}

# ================ USER QUERIES ================

# Get current user profile
get_profile() {
    local token=${1:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/users/profile \
        -H "Authorization: Bearer $token" | jq '.'
}

# List instructors
list_instructors() {
    local token=${1:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/users/instructors \
        -H "Authorization: Bearer $token" | jq '.[] | {id, email, first_name, last_name}'
}

# List students
list_students() {
    local token=${1:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/users/students \
        -H "Authorization: Bearer $token" | jq '.[] | {id, email, first_name, last_name}'
}

# List all users (admin)
list_users() {
    local token=${1:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/admin/users \
        -H "Authorization: Bearer $token" | jq '.[] | {id, email, first_name, role}'
}

# ================ COURSE QUERIES ================

# Get all courses
get_courses() {
    local token=${1:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/courses \
        -H "Authorization: Bearer $token" | jq '.[] | {id, title, code, level}'
}

# Get first course ID
get_first_course() {
    local token=${1:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/courses \
        -H "Authorization: Bearer $token" | jq -r '.[0].id'
}

# Get course details
get_course() {
    local course_id=$1
    local token=${2:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/courses/$course_id \
        -H "Authorization: Bearer $token" | jq '.'
}

# Get course modules
get_modules() {
    local course_id=$1
    local token=${2:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/courses/$course_id/modules \
        -H "Authorization: Bearer $token" | jq '.[] | {id, title, order}'
}

# Get course assignments
get_assignments() {
    local course_id=$1
    local token=${2:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/courses/$course_id/assignments \
        -H "Authorization: Bearer $token" | jq '.[] | {id, title, due_date}'
}

# Get course quizzes
get_quizzes() {
    local course_id=$1
    local token=${2:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/courses/$course_id/quizzes \
        -H "Authorization: Bearer $token" | jq '.[] | {id, title, max_score}'
}

# ================ ENROLLMENT QUERIES ================

# Check enrollment status
check_enrollment() {
    local course_id=$1
    local token=${2:-$STUDENT_TOKEN}
    curl -s -X GET $BASE_URL/api/courses/$course_id/enrollment-status \
        -H "Authorization: Bearer $token" | jq '.'
}

# Enroll in course
enroll_course() {
    local course_id=$1
    local token=${2:-$STUDENT_TOKEN}
    curl -s -X POST $BASE_URL/api/courses/$course_id/enroll \
        -H "Authorization: Bearer $token" \
        -H "Content-Type: application/json" \
        -d '{}' | jq '.'
}

# Get course enrollments (instructor)
get_enrollments() {
    local course_id=$1
    local token=${2:-$INSTRUCTOR_TOKEN}
    curl -s -X GET $BASE_URL/api/courses/$course_id/enrollments \
        -H "Authorization: Bearer $token" | jq '.[] | {id, user_email, status}'
}

# ================ ASSIGNMENT QUERIES ================

# Get assignment details
get_assignment() {
    local assignment_id=$1
    local token=${2:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/assignments/$assignment_id \
        -H "Authorization: Bearer $token" | jq '.'
}

# Submit assignment
submit_assignment() {
    local assignment_id=$1
    local content="${2:-My submission}"
    local token=${3:-$STUDENT_TOKEN}
    curl -s -X POST $BASE_URL/api/assignments/$assignment_id/submit \
        -H "Authorization: Bearer $token" \
        -H "Content-Type: application/json" \
        -d "{\"content\":\"$content\",\"file_url\":\"\"}" | jq '.'
}

# Get submissions
get_submissions() {
    local assignment_id=$1
    local token=${2:-$INSTRUCTOR_TOKEN}
    curl -s -X GET $BASE_URL/api/assignments/$assignment_id/submissions \
        -H "Authorization: Bearer $token" | jq '.[] | {id, student_email, submitted_at}'
}

# ================ QUIZ QUERIES ================

# Get quiz details
get_quiz() {
    local quiz_id=$1
    local token=${2:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/quizzes/$quiz_id \
        -H "Authorization: Bearer $token" | jq '.'
}

# Get quiz questions
get_questions() {
    local quiz_id=$1
    local token=${2:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/quizzes/$quiz_id/questions \
        -H "Authorization: Bearer $token" | jq '.[] | {id, question, type}'
}

# Start quiz
start_quiz() {
    local quiz_id=$1
    local token=${2:-$STUDENT_TOKEN}
    curl -s -X POST $BASE_URL/api/quizzes/$quiz_id/start \
        -H "Authorization: Bearer $token" \
        -H "Content-Type: application/json" \
        -d '{}' | jq '.'
}

# ================ ADMIN QUERIES ================

# Get dashboard
get_dashboard() {
    local token=${1:-$ADMIN_TOKEN}
    curl -s -X GET $BASE_URL/api/admin/dashboard \
        -H "Authorization: Bearer $token" | jq '.'
}

# ================ QUICK TESTS ================

# Test all endpoints
test_all() {
    echo "🧪 Running comprehensive tests..."
    echo ""
    
    get_all_tokens
    echo ""
    
    echo "1️⃣  Testing Users..."
    echo "   Admin Profile:"
    get_profile $ADMIN_TOKEN | jq '.email'
    echo "   Instructors: $(list_instructors $ADMIN_TOKEN | jq 'length')"
    echo "   Students: $(list_students $ADMIN_TOKEN | jq 'length')"
    echo ""
    
    echo "2️⃣  Testing Courses..."
    COURSE_ID=$(get_first_course $ADMIN_TOKEN)
    echo "   Courses: $(get_courses $ADMIN_TOKEN | jq 'length')"
    echo "   First course: $COURSE_ID"
    echo ""
    
    echo "3️⃣  Testing Enrollments..."
    enroll_course $COURSE_ID $STUDENT_TOKEN | jq '.message'
    check_enrollment $COURSE_ID $STUDENT_TOKEN | jq '.status'
    echo ""
    
    echo "4️⃣  Testing Assignments..."
    ASSIGNMENT_ID=$(get_assignments $COURSE_ID $ADMIN_TOKEN | jq -r '.[0].id')
    echo "   First assignment: $ASSIGNMENT_ID"
    submit_assignment $ASSIGNMENT_ID "Test submission" $STUDENT_TOKEN | jq '.message'
    echo ""
    
    echo "5️⃣  Testing Quizzes..."
    QUIZ_ID=$(get_quizzes $COURSE_ID $ADMIN_TOKEN | jq -r '.[0].id')
    echo "   First quiz: $QUIZ_ID"
    start_quiz $QUIZ_ID $STUDENT_TOKEN | jq '.message // .attempt_id'
    echo ""
    
    echo "✅ All tests completed!"
}

# ================ HELP ================

# Display available functions
show_help() {
    cat << 'EOF'
LMS Backend - Quick Reference Commands

SERVER MANAGEMENT:
  server_status      - Check if server is running
  start_server       - Start the server
  stop_server        - Stop the server
  restart_server     - Restart the server
  view_logs          - View server logs

DATABASE:
  db_status          - Check database connection
  db_connect         - Connect to database
  reseed_database    - Reseed with test data

AUTHENTICATION:
  get_all_tokens     - Get all three tokens
  admin_login        - Get admin token
  inst_login         - Get instructor token
  student_login      - Get student token

USERS:
  get_profile        - Get user profile
  list_instructors   - List all instructors
  list_students      - List all students
  list_users         - List all users (admin)

COURSES:
  get_courses        - List all courses
  get_first_course   - Get first course ID
  get_course ID      - Get course details
  get_modules ID     - Get course modules
  get_assignments ID - Get course assignments
  get_quizzes ID     - Get course quizzes

ENROLLMENTS:
  check_enrollment ID         - Check enrollment status
  enroll_course ID [TOKEN]    - Enroll in course
  get_enrollments ID [TOKEN]  - Get course enrollments

ASSIGNMENTS:
  get_assignment ID           - Get assignment details
  submit_assignment ID [CONTENT] [TOKEN]  - Submit assignment
  get_submissions ID [TOKEN]  - Get submissions

QUIZZES:
  get_quiz ID                 - Get quiz details
  get_questions ID [TOKEN]    - Get quiz questions
  start_quiz ID [TOKEN]       - Start quiz attempt

ADMIN:
  get_dashboard [TOKEN]       - Get admin dashboard

TESTS:
  test_all           - Run comprehensive tests
  show_help          - Show this help message

USAGE:
  # Source this file first
  source quick-reference.sh
  
  # Get tokens
  get_all_tokens
  
  # Use commands
  get_courses
  get_course <course-id>
  enroll_course <course-id>
  submit_assignment <assignment-id>
EOF
}

# Display help on source
show_help
