# Deployment Guide - Railway Platform

## ปัญหาที่แก้ไขแล้ว

### 1. Main Redeclared Error ✅
**ปัญหา:** มีไฟล์ `create_staff_password.go` ใน root directory ที่มี `main` function ซ้ำกับ `main.go`

**การแก้ไข:**
- ลบไฟล์ `create_staff_password.go` ออกจาก root directory
- เหลือเพียงไฟล์ `tools/create_staff_password.go` สำหรับ utility

### 2. Docker Configuration ✅
**เพิ่มไฟล์:**
- `Dockerfile` - Multi-stage build สำหรับ Go application
- `.dockerignore` - ลดขนาด Docker context
- `railway.toml` - ตั้งค่า Railway platform

## ไฟล์ที่เพิ่ม/แก้ไข

### Dockerfile
```dockerfile
# Multi-stage build สำหรับ Go 1.23
FROM golang:1.23-alpine AS builder
# ... (build stage)
FROM alpine:latest
# ... (runtime stage)
```

### railway.toml
```toml
[build]
builder = "dockerfile"

[deploy]
healthcheckPath = "/health"
healthcheckTimeout = 100
restartPolicyType = "on_failure"
restartPolicyMaxRetries = 10
```

## Environment Variables สำหรับ Railway

ตั้งค่าใน Railway Dashboard:

```bash
# Database (PostgreSQL)
DATABASE_URL=postgresql://user:password@host:port/database

# JWT Secret
JWT_SECRET=your-super-secret-key-here

# Gin Mode
GIN_MODE=release

# PORT (Railway จะตั้งให้อัตโนมัติ)
# PORT=8080
```

## Health Check Endpoint

ระบบมี health check endpoint ที่ `/health`:

```bash
GET /health
Response: {
  "status": "ok",
  "message": "Server is running",
  "timestamp": 1704067200
}
```

## การ Deploy

1. **Push โค้ดไปยัง Git Repository**
2. **Connect กับ Railway:**
   - ไปที่ Railway Dashboard
   - เลือก "New Project"
   - Connect GitHub Repository
3. **ตั้งค่า Environment Variables**
4. **Deploy อัตโนมัติ**

## การทดสอบหลัง Deploy

### 1. Health Check
```bash
curl https://your-app-name.railway.app/health
```

### 2. Public API - ดูเมนู
```bash
curl https://your-app-name.railway.app/public/menu
```

### 3. Staff Login
```bash
curl -X POST https://your-app-name.railway.app/staff/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password"}'
```

## Troubleshooting

### 1. Build Failures
- ตรวจสอบว่าไม่มีไฟล์ที่มี `main` function ซ้ำ
- ตรวจสอบ `go.mod` และ dependencies

### 2. Runtime Errors
- ตรวจสอบ Environment Variables
- ดู logs ใน Railway Dashboard

### 3. Database Connection
- ตรวจสอบ `DATABASE_URL`
- ตรวจสอบว่า database accessible จากภายนอก

## Sample Data

หลังจาก deploy สำเร็จ ให้รัน migration:

```bash
# ถ้ามี CLI tools หรือใช้ manual SQL
# INSERT ข้อมูลจาก sample_data.sql
```

## Performance Optimization

1. **GOOS=linux CGO_ENABLED=0** - สำหรับ Alpine Linux
2. **Multi-stage build** - ลดขนาด Docker image
3. **Health check** - Railway ตรวจสอบสถานะ app
4. **Restart policy** - Auto restart เมื่อ app crash

---
**หมายเหตุ:** หลังจากแก้ไขปัญหา "main redeclared" แล้ว Railway ควรจะ build ผ่าน หากยังมีปัญหาให้ตรวจสอบ Environment Variables และ Database connection
