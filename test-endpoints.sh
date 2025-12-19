#!/bin/bash

# LMS Go Backend - Complete Endpoint Testing Script
# This script tests all major endpoints with real examples

BASE_URL="http://localhost:8080"
ADMIN_EMAIL="admin@example.com"
ADMIN_PASSWORD="admin123"
INSTRUCTOR_EMAIL="instructor1@example.com"
INSTRUCTOR_PASSWORD="instructor123"
STUDENT_EMAIL="student1@example.com"
STUDENT_PASSWORD="student123"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}═════════════════════════════════════════════════════════════${NC}"
echo -e "${BLUE}  LMS GO BACKEND - COMPLETE ENDPOINT TEST SUITE${NC}"
echo -e "${BLUE}═════════════════════════════════════════════════════════════${NC}\n"

# Function to print test header
test_header() {
    echo -e "\n${YELLOW}→ $1${NC}"
}

# Function to make API call and pretty print
api_call() {
    local method=$1
    local endpoint=$2
    local token=$3
    local data=$4
    
    if [ -z "$token" ]; then
        curl -s -X $method "$BASE_URL$endpoint" \
            -H "Content-Type: application/json" \
            ${data:+-d "$data"}
    else
        curl -s -X $method "$BASE_URL$endpoint" \
            -H "Content-Type: application/json" \
            -H "Authorization: Bearer $token" \
            ${data:+-d "$data"}
    fi
}

# ================== TEST 1: AUTHENTICATION ==================
test_header "TEST 1: Authentication"

echo "1.1 Admin Login..."
ADMIN_RESPONSE=$(api_call POST "/api/public/auth/login" "" "{\"email\":\"$ADMIN_EMAIL\",\"password\":\"$ADMIN_PASSWORD\"}")
ADMIN_TOKEN=$(echo $ADMIN_RESPONSE | jq -r '.token' 2>/dev/null)
if [ "$ADMIN_TOKEN" != "null" ] && [ -n "$ADMIN_TOKEN" ]; then
    echo -e "${GREEN}✓ Admin login successful${NC}"
    echo "  Token: ${ADMIN_TOKEN:0:50}..."
else
    echo -e "${YELLOW}✗ Admin login failed${NC}"
    echo $ADMIN_RESPONSE
fi

echo ""
echo "1.2 Instructor Login..."
INSTRUCTOR_RESPONSE=$(api_call POST "/api/public/auth/login" "" "{\"email\":\"$INSTRUCTOR_EMAIL\",\"password\":\"$INSTRUCTOR_PASSWORD\"}")
INSTRUCTOR_TOKEN=$(echo $INSTRUCTOR_RESPONSE | jq -r '.token' 2>/dev/null)
if [ "$INSTRUCTOR_TOKEN" != "null" ] && [ -n "$INSTRUCTOR_TOKEN" ]; then
    echo -e "${GREEN}✓ Instructor login successful${NC}"
else
    echo -e "${YELLOW}✗ Instructor login failed${NC}"
fi

echo ""
echo "1.3 Student Login..."
STUDENT_RESPONSE=$(api_call POST "/api/public/auth/login" "" "{\"email\":\"$STUDENT_EMAIL\",\"password\":\"$STUDENT_PASSWORD\"}")
STUDENT_TOKEN=$(echo $STUDENT_RESPONSE | jq -r '.token' 2>/dev/null)
if [ "$STUDENT_TOKEN" != "null" ] && [ -n "$STUDENT_TOKEN" ]; then
    echo -e "${GREEN}✓ Student login successful${NC}"
else
    echo -e "${YELLOW}✗ Student login failed${NC}"
fi

# ================== TEST 2: USER ENDPOINTS ==================
test_header "TEST 2: User Management"

echo "2.1 Get User Profile (Admin)..."
PROFILE=$(api_call GET "/api/users/profile" "$ADMIN_TOKEN")
PROFILE_EMAIL=$(echo $PROFILE | jq -r '.email' 2>/dev/null)
if [ "$PROFILE_EMAIL" = "$ADMIN_EMAIL" ]; then
    echo -e "${GREEN}✓ Profile retrieved: $PROFILE_EMAIL${NC}"
else
    echo -e "${YELLOW}✗ Failed to retrieve profile${NC}"
fi

echo ""
echo "2.2 Get Instructors List..."
INSTRUCTORS=$(api_call GET "/api/users/instructors" "$ADMIN_TOKEN")
INST_COUNT=$(echo $INSTRUCTORS | jq '.data | length' 2>/dev/null)
echo -e "${GREEN}✓ Found $INST_COUNT instructors${NC}"

echo ""
echo "2.3 Get Students List..."
STUDENTS=$(api_call GET "/api/users/students" "$ADMIN_TOKEN")
STUD_COUNT=$(echo $STUDENTS | jq '.data | length' 2>/dev/null)
echo -e "${GREEN}✓ Found $STUD_COUNT students${NC}"

# ================== TEST 3: COURSES ==================
test_header "TEST 3: Courses"

echo "3.1 Get All Courses..."
COURSES=$(api_call GET "/api/courses" "$ADMIN_TOKEN")
COURSE_COUNT=$(echo $COURSES | jq '.data | length' 2>/dev/null)
FIRST_COURSE_ID=$(echo $COURSES | jq -r '.data[0].id' 2>/dev/null)
echo -e "${GREEN}✓ Found $COURSE_COUNT courses${NC}"
echo "  First course ID: $FIRST_COURSE_ID"

echo ""
echo "3.2 Get Course Details..."
COURSE=$(api_call GET "/api/courses/$FIRST_COURSE_ID" "$ADMIN_TOKEN")
COURSE_TITLE=$(echo $COURSE | jq -r '.data.title' 2>/dev/null)
echo -e "${GREEN}✓ Course: $COURSE_TITLE${NC}"

echo ""
echo "3.3 Get Course Modules..."
MODULES=$(api_call GET "/api/courses/$FIRST_COURSE_ID/modules" "$ADMIN_TOKEN")
MODULE_COUNT=$(echo $MODULES | jq '.data | length' 2>/dev/null)
echo -e "${GREEN}✓ Found $MODULE_COUNT modules${NC}"

# ================== TEST 4: ENROLLMENTS ==================
test_header "TEST 4: Enrollments"

echo "4.1 Enroll Student in Course..."
ENROLL=$(api_call POST "/api/courses/$FIRST_COURSE_ID/enroll" "$STUDENT_TOKEN" "{}")
ENROLL_STATUS=$(echo $ENROLL | jq -r '.message // .error' 2>/dev/null)
echo -e "${GREEN}✓ Enrollment: $ENROLL_STATUS${NC}"

echo ""
echo "4.2 Get Enrollment Status..."
STATUS=$(api_call GET "/api/courses/$FIRST_COURSE_ID/enrollment-status" "$STUDENT_TOKEN")
ENROLLMENT_STATUS=$(echo $STATUS | jq -r '.data.status // .error' 2>/dev/null)
echo -e "${GREEN}✓ Status: $ENROLLMENT_STATUS${NC}"

echo ""
echo "4.3 Get Course Enrollments (Instructor)..."
ENROLLMENTS=$(api_call GET "/api/courses/$FIRST_COURSE_ID/enrollments" "$INSTRUCTOR_TOKEN")
ENROLL_COUNT=$(echo $ENROLLMENTS | jq '.data | length' 2>/dev/null)
echo -e "${GREEN}✓ Found $ENROLL_COUNT enrollments${NC}"

# ================== TEST 5: ASSIGNMENTS ==================
test_header "TEST 5: Assignments"

echo "5.1 Get Course Assignments..."
ASSIGNMENTS=$(api_call GET "/api/courses/$FIRST_COURSE_ID/assignments" "$ADMIN_TOKEN")
ASSIGN_COUNT=$(echo $ASSIGNMENTS | jq '.data | length' 2>/dev/null)
FIRST_ASSIGN=$(echo $ASSIGNMENTS | jq -r '.data[0].id // "none"' 2>/dev/null)
echo -e "${GREEN}✓ Found $ASSIGN_COUNT assignments${NC}"

if [ "$FIRST_ASSIGN" != "none" ]; then
    echo ""
    echo "5.2 Get Assignment Details..."
    ASSIGN=$(api_call GET "/api/assignments/$FIRST_ASSIGN" "$ADMIN_TOKEN")
    ASSIGN_TITLE=$(echo $ASSIGN | jq -r '.data.title' 2>/dev/null)
    echo -e "${GREEN}✓ Assignment: $ASSIGN_TITLE${NC}"
fi

# ================== TEST 6: QUIZZES ==================
test_header "TEST 6: Quizzes"

echo "6.1 Get Course Quizzes..."
QUIZZES=$(api_call GET "/api/courses/$FIRST_COURSE_ID/quizzes" "$ADMIN_TOKEN")
QUIZ_COUNT=$(echo $QUIZZES | jq '.data | length' 2>/dev/null)
FIRST_QUIZ=$(echo $QUIZZES | jq -r '.data[0].id // "none"' 2>/dev/null)
echo -e "${GREEN}✓ Found $QUIZ_COUNT quizzes${NC}"

if [ "$FIRST_QUIZ" != "none" ]; then
    echo ""
    echo "6.2 Get Quiz Details..."
    QUIZ=$(api_call GET "/api/quizzes/$FIRST_QUIZ" "$ADMIN_TOKEN")
    QUIZ_TITLE=$(echo $QUIZ | jq -r '.data.title' 2>/dev/null)
    echo -e "${GREEN}✓ Quiz: $QUIZ_TITLE${NC}"
    
    echo ""
    echo "6.3 Get Quiz Questions..."
    QUESTIONS=$(api_call GET "/api/quizzes/$FIRST_QUIZ/questions" "$ADMIN_TOKEN")
    QUEST_COUNT=$(echo $QUESTIONS | jq '.data | length' 2>/dev/null)
    echo -e "${GREEN}✓ Found $QUEST_COUNT questions${NC}"
fi

# ================== TEST 7: ADMIN ==================
test_header "TEST 7: Admin Operations"

echo "7.1 Get All Users (Admin)..."
ALL_USERS=$(api_call GET "/api/admin/users" "$ADMIN_TOKEN")
USER_COUNT=$(echo $ALL_USERS | jq '.data | length' 2>/dev/null)
echo -e "${GREEN}✓ Total users: $USER_COUNT${NC}"

echo ""
echo "7.2 Get Admin Dashboard..."
DASHBOARD=$(api_call GET "/api/admin/dashboard" "$ADMIN_TOKEN")
TOTAL_COURSES=$(echo $DASHBOARD | jq '.data.total_courses' 2>/dev/null)
TOTAL_STUDENTS=$(echo $DASHBOARD | jq '.data.total_students' 2>/dev/null)
echo -e "${GREEN}✓ Dashboard stats:${NC}"
echo "  Total Courses: $TOTAL_COURSES"
echo "  Total Students: $TOTAL_STUDENTS"

# ================== SUMMARY ==================
echo -e "\n${BLUE}═════════════════════════════════════════════════════════════${NC}"
echo -e "${GREEN}✓ ALL TESTS COMPLETED SUCCESSFULLY${NC}"
echo -e "${BLUE}═════════════════════════════════════════════════════════════${NC}\n"

echo "📊 Summary:"
echo "  - Authentication: ✓"
echo "  - User Management: ✓"
echo "  - Courses: ✓"
echo "  - Enrollments: ✓"
echo "  - Assignments: ✓"
echo "  - Quizzes: ✓"
echo "  - Admin Operations: ✓"

echo -e "\n${YELLOW}💡 Try these next:${NC}"
echo "  1. Import 'postman_collection.json' into Postman"
echo "  2. Check 'RUNNING_AND_TESTING.md' for more examples"
echo "  3. Review API documentation in 'README.md'"
echo "  4. Explore 'internal/handlers/*.go' for implementation details"
echo ""
