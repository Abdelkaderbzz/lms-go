# 🚀 Quick Setup Guide - LMS Go Backend

**Status**: ✅ Project Ready, Awaiting Database Setup

---

## 📋 Setup Options

### Option 1: Using Docker (Recommended - No Installation Needed)

#### Step 1: Start PostgreSQL in Docker

```bash
# Start PostgreSQL container
docker run -d \
  --name lms-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=lms_db \
  -p 5432:5432 \
  postgres:15
```

#### Step 2: Wait for Database to be Ready

```bash
# Wait 5-10 seconds for database to start
sleep 10

# Test connection
docker exec lms-postgres psql -U user -d lms_db -c "SELECT 1"
```

#### Step 3: Seed Database

```bash
cd /Users/abdelkaderbouzomita/Sites/lms-go
make seed
```

#### Step 4: Start Application

```bash
make run
```

#### Step 5: Test in New Terminal

```bash
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}'
```

---

### Option 2: Using docker-compose (Easiest)

#### Step 1: Check docker-compose.yml

```bash
cat /Users/abdelkaderbouzomita/Sites/lms-go/docker-compose.yml
```

#### Step 2: Start Services

```bash
cd /Users/abdelkaderbouzomita/Sites/lms-go
make docker-up
```

#### Step 3: Wait for Database

```bash
sleep 10
```

#### Step 4: Seed & Run

```bash
make seed
make run
```

---

### Option 3: Install PostgreSQL Locally (macOS)

#### Using Homebrew

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

#### Then Run Project

```bash
cd /Users/abdelkaderbouzomita/Sites/lms-go
make seed
make run
```

---

## 🎯 Quick Start (Choose One Option Above First)

### After Database is Running:

**Terminal 1 - Start Application:**

```bash
cd /Users/abdelkaderbouzomita/Sites/lms-go
make seed    # Seed test data (one-time)
make run     # Start server
```

**Terminal 2 - Test Endpoints:**

```bash
# Login
curl -X POST http://localhost:8080/api/public/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}'

# Copy the token from response
```

---

## 🧪 Quick Test Commands

After getting token, save it:

```bash
export TOKEN="your_token_here"
```

### Test 1: Get Courses

```bash
curl -X GET http://localhost:8080/api/courses \
  -H "Authorization: Bearer $TOKEN" | jq
```

### Test 2: Get User Profile

```bash
curl -X GET http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer $TOKEN" | jq
```

### Test 3: Get All Users (Admin Only)

```bash
curl -X GET http://localhost:8080/api/admin/users \
  -H "Authorization: Bearer $TOKEN" | jq
```

---

## 📊 Test Users Available

After running `make seed`:

```
ADMIN:
  Email: admin@example.com
  Password: admin123

INSTRUCTOR:
  Email: instructor1@example.com
  Password: instructor123

STUDENT:
  Email: student1@example.com
  Password: student123
```

---

## 🐳 Docker Option (Fastest)

If you have Docker installed:

```bash
# 1. Start PostgreSQL
docker run -d \
  --name lms-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=lms_db \
  -p 5432:5432 \
  postgres:15

# 2. Wait for it to start
sleep 10

# 3. Seed database
cd /Users/abdelkaderbouzomita/Sites/lms-go
make seed

# 4. Run application
make run
```

**Stop Docker when done:**

```bash
docker stop lms-postgres
docker rm lms-postgres
```

---

## ✅ Verification Checklist

Before running, ensure:

- [ ] Database option chosen (Docker, docker-compose, or local PostgreSQL)
- [ ] `make seed` completes successfully
- [ ] Server starts with `make run`
- [ ] Can curl login endpoint

---

## 📝 First-Time Setup Script

Create this as `setup.sh`:

```bash
#!/bin/bash

echo "🚀 LMS Go Backend Setup"
echo "======================="

# Check if Docker is available
if command -v docker &> /dev/null; then
    echo "✅ Docker found"
    echo ""
    echo "Starting PostgreSQL container..."

    docker run -d \
      --name lms-postgres \
      -e POSTGRES_USER=user \
      -e POSTGRES_PASSWORD=password \
      -e POSTGRES_DB=lms_db \
      -p 5432:5432 \
      postgres:15

    echo "⏳ Waiting for PostgreSQL to start..."
    sleep 10

    echo "✅ PostgreSQL started"
else
    echo "⚠️  Docker not found. Please install Docker or PostgreSQL manually."
    exit 1
fi

echo ""
echo "📦 Running tests..."
make test

echo ""
echo "🌱 Seeding database..."
make seed

echo ""
echo "✅ Setup complete!"
echo ""
echo "🚀 To start the server, run: make run"
echo "📝 To test endpoints, see: RUNNING_AND_TESTING.md"
```

Run with:

```bash
chmod +x setup.sh
./setup.sh
```

---

## 🔍 Troubleshooting Setup

### Can't connect to database

```bash
# Verify DATABASE_URL is set correctly
echo $DATABASE_URL

# Or set it manually
export DATABASE_URL="postgres://user:password@localhost:5432/lms_db"
```

### Docker issues

```bash
# Check if container is running
docker ps | grep lms-postgres

# View logs
docker logs lms-postgres

# Stop and remove
docker stop lms-postgres
docker rm lms-postgres
```

### Port 5432 already in use

```bash
# Use different port
docker run -d \
  --name lms-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=lms_db \
  -p 5433:5432 \
  postgres:15

# Update DATABASE_URL
export DATABASE_URL="postgres://user:password@localhost:5433/lms_db"
```

---

## 📚 Next Steps After Setup

1. **Run tests** to verify setup

   ```bash
   make test
   ```

2. **Seed database**

   ```bash
   make seed
   ```

3. **Start server**

   ```bash
   make run
   ```

4. **Test endpoints** (new terminal)
   - See RUNNING_AND_TESTING.md for curl commands
   - Or use Postman with postman_collection.json

---

**Ready to go! 🎉**
