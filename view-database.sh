#!/bin/bash

# LMS Database Viewer - Simple and Working
# Usage: ./view-database.sh

export PGPASSWORD="password"
export PATH="/opt/homebrew/opt/postgresql@15/bin:$PATH"

echo ""
echo "╔═══════════════════════════════════════════════════════════════════╗"
echo "║                  📊 LMS DATABASE VIEWER                          ║"
echo "║              View all tables and their contents                   ║"
echo "╚═══════════════════════════════════════════════════════════════════╝"
echo ""

# Check connection
if ! psql -U user -d lms_db -h localhost -c "SELECT 1" > /dev/null 2>&1; then
    echo "❌ ERROR: Cannot connect to database"
    echo ""
    echo "   Make sure PostgreSQL is running:"
    echo "   $ docker ps | grep postgres"
    echo ""
    exit 1
fi

echo "✅ Connected to lms_db database"
echo ""

# Show all tables
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📋 ALL 24 TABLES:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
psql -U user -d lms_db -h localhost -c "\dt" | tail -n +3 | head -n -1 | awk '{print "  ✓", $3}'
echo ""

# Show record counts
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📊 RECORD COUNTS:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

tables=(
    "users" "courses" "modules" "lessons" "assignments" "quizzes" 
    "questions" "submissions" "grades" "discussions" "forum_posts" 
    "enrollments" "announcements" "notifications" "resources"
)

for table in "${tables[@]}"; do
    result=$(psql -U user -d lms_db -h localhost -t -c "SELECT COUNT(*) FROM $table" 2>/dev/null)
    printf "  %-25s: %5s records\n" "$table" "$result"
done

echo ""

# Users
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "👤 USERS:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
psql -U user -d lms_db -h localhost -c "
SELECT email, first_name, last_name, role 
FROM users 
ORDER BY created_at DESC;"
echo ""

# Courses
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📚 COURSES:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
psql -U user -d lms_db -h localhost -c "
SELECT title, code, level, status 
FROM courses 
ORDER BY created_at DESC;"
echo ""

# Enrollments
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ ENROLLMENTS:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
psql -U user -d lms_db -h localhost -c "
SELECT u.email, c.title, e.status 
FROM enrollments e
JOIN users u ON e.user_id = u.id
JOIN courses c ON e.course_id = c.id
LIMIT 15;"
echo ""

# Assignments
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📝 ASSIGNMENTS:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
psql -U user -d lms_db -h localhost -c "
SELECT a.title, c.title as course, a.due_date 
FROM assignments a
JOIN courses c ON a.course_id = c.id
LIMIT 10;"
echo ""

# Quizzes
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🎯 QUIZZES:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
psql -U user -d lms_db -h localhost -c "
SELECT q.title, c.title as course, q.max_score 
FROM quizzes q
JOIN courses c ON q.course_id = c.id
LIMIT 10;"
echo ""

# Submissions
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "� SUBMISSIONS:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
psql -U user -d lms_db -h localhost -c "
SELECT u.email, a.title, s.status, s.submitted_at 
FROM submissions s
JOIN users u ON s.user_id = u.id
JOIN assignments a ON s.assignment_id = a.id
LIMIT 10;"
echo ""

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ Database viewer complete!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "💡 QUICK TIPS:"
echo ""
echo "  View all data interactively:"
echo "  $ PGPASSWORD=password /opt/homebrew/opt/postgresql@15/bin/psql -U user -d lms_db -h localhost"
echo ""
echo "  Then use these commands:"
echo "  \\dt                     - List all tables"
echo "  SELECT * FROM users;    - View users table"
echo "  \\q                      - Exit"

