#!/bin/bash

# pgAdmin 4 Connection Guide for LMS Database

# This script explains how to connect pgAdmin to your PostgreSQL database

cat << 'EOF'

╔═══════════════════════════════════════════════════════════════════════╗
║ ║
║ 🔧 CONNECTING pgADMIN 4 TO YOUR DATABASE ║
║ ║
╚═══════════════════════════════════════════════════════════════════════╝

📍 DATABASE CONNECTION DETAILS

Host: localhost
Port: 5432
Database: lms_db
Username: user
Password: password

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🚀 STEP-BY-STEP GUIDE TO CONNECT pgAdmin

Step 1: Open pgAdmin
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. Launch pgAdmin from Applications (or Applications folder)
2. It will open in your web browser (usually http://localhost:5050)
3. You may need to set a master password on first launch

Step 2: Add a New Server
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. In pgAdmin, look for "Servers" on the left sidebar
2. Right-click on "Servers" → "Create" → "Server"
3. Or click the "+" button next to "Servers"

Step 3: General Tab (Name your connection)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Name: LMS Database
(or any name you prefer)

Comments: (optional) My LMS Go Backend Database

Step 4: Connection Tab (Enter connection details)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Host name/address: localhost
Port: 5432
Maintenance database: postgres
Username: user
Password: password
Save password? ✓ Check this box (optional but recommended)

Step 5: Advanced Tab (Optional)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Leave defaults unless you have specific needs

Step 6: Click "Save"
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

The server should now appear in your Servers list
It might show a red X initially (still connecting)
Once connected, it will turn green

Step 7: Explore Your Database
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. Click on "LMS Database" to expand it
2. Navigate to: Databases → lms_db → Schemas → public → Tables
3. You'll see all 24 tables:
   - users
   - courses
   - enrollments
   - assignments
   - quizzes
   - submissions
   - grades
   - and more...

Step 8: View Table Data
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. Right-click on any table (e.g., "users")
2. Select "View/Edit Data" → "All Rows"
3. See all data in that table with full GUI editing capabilities

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎯 COMMON TASKS IN pgAdmin

View Users Table:
• Servers → LMS Database → Databases → lms_db → Schemas → public → Tables
• Right-click "users" → "View/Edit Data" → "All Rows"

View Courses:
• Same path, select "courses" table

Run SQL Query:
• Click on "LMS Database"
• Click "Tools" → "Query Tool"
• Write SQL: SELECT \* FROM users;
• Press F5 or click Execute

Create New Tables:
• Right-click "Tables" → "Create" → "Table"

Edit Data:
• View table → double-click cells to edit → Save

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

❓ TROUBLESHOOTING

Connection Failed?
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

❌ "Connection refused"
→ Make sure PostgreSQL is running:
$ docker ps | grep postgres
→ If not running:
$ docker start lms-postgres

❌ "Invalid password"
→ Password is: password
→ Check spelling carefully

❌ "Host not found"
→ Host should be: localhost
→ Port should be: 5432

❌ "Cannot create maintenance database"
→ In Connection tab, set "Maintenance database" to "postgres"

Server Shows Red X (Not Connected)?
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. Right-click the server → "Properties"
2. Check all connection details again
3. Click "Save"
4. Right-click server → "Connect Server"

Cannot See Tables?
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. Expand: Servers → LMS Database → Databases
2. Look for "lms_db"
3. Expand lms_db → Schemas → public → Tables
4. You should see 24 tables listed

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

✨ USEFUL pgAdmin FEATURES

Query Tool (Best for SQL):
• Tools → Query Tool
• Write and execute SQL queries
• See results in real-time

Data Viewer:
• Right-click table → View/Edit Data
• See all rows with full editing capabilities

Table Structure:
• Right-click table → Properties
• See columns, types, constraints

Statistics:
• Right-click database → Statistics
• See database size and performance info

Backup/Restore:
• Right-click database → Backup
• Export database to file

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📚 DATABASE STRUCTURE

Your database has these 24 tables:

Core Tables:
✓ users - All users (admin, instructors, students)
✓ courses - Course information
✓ modules - Course modules
✓ lessons - Lessons within modules
✓ resources - Course resources

Learning Content:
✓ assignments - Course assignments
✓ submissions - Student assignment submissions
✓ grades - Assignment grades
✓ quizzes - Quiz definitions
✓ questions - Quiz questions
✓ options - Quiz answer options
✓ quiz_answers - Student quiz answers
✓ quiz_attempts - Student quiz attempts

Enrollment & Progress:
✓ enrollments - Student course enrollments
✓ lesson_progresses - Student lesson progress
✓ course_enrollments - Course enrollment tracking
✓ course_instructors - Course instructor assignments
✓ certificates - Student certificates

Discussion & Communication:
✓ discussions - Course discussions
✓ forum_posts - Forum posts
✓ forum_replies - Replies to forum posts
✓ comments - General comments
✓ announcements - Course announcements
✓ notifications - User notifications

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎉 YOU'RE ALL SET!

Now you can:
✅ View all database tables in pgAdmin
✅ Edit data visually
✅ Run SQL queries
✅ Manage your database easily

Enjoy! 🚀

EOF
