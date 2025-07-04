#!/bin/bash

# สคริปต์สำหรับเริ่มต้นระบบ Backend POS

echo "=== การตั้งค่าระบบ Backend POS ==="

# ตรวจสอบว่ามี Go หรือไม่
if ! command -v go &> /dev/null; then
    echo "❌ ไม่พบ Go บนระบบ กรุณาติดตั้ง Go ก่อน"
    exit 1
fi

echo "✅ พบ Go เวอร์ชัน: $(go version)"

# ตรวจสอบไฟล์ .env
if [ ! -f ".env" ]; then
    echo "📝 สร้างไฟล์ .env จาก .env.example"
    cp .env.example .env
    echo "⚠️  กรุณาแก้ไขไฟล์ .env ให้ตรงกับฐานข้อมูลของคุณ"
    echo "   เปิดไฟล์ .env และตั้งค่า:"
    echo "   - DB_HOST (เช่น localhost)"
    echo "   - DB_PORT (เช่น 5432)"
    echo "   - DB_DATABASE (เช่น pos_db)"
    echo "   - DB_USER (เช่น postgres)"
    echo "   - DB_PASSWORD (รหัสผ่านฐานข้อมูล)"
    echo ""
    echo "แล้วรันสคริปต์นี้อีกครั้ง"
    exit 1
fi

echo "✅ พบไฟล์ .env"

# โหลด environment variables
export $(cat .env | grep -v '#' | awk '/=/ {print $1}')

echo "📦 ติดตั้ง dependencies..."
go mod tidy

echo "🔨 สร้าง executable..."
go build -o pos-server

if [ $? -eq 0 ]; then
    echo "✅ สร้างเซิร์ฟเวอร์สำเร็จ"
    echo ""
    echo "🚀 วิธีการรัน:"
    echo "1. รัน migration (สร้างตาราง): ./pos-server migrate"
    echo "2. รันเซิร์ฟเวอร์: ./pos-server"
    echo ""
    echo "หรือรันตรงๆ: go run main.go"
    echo ""
    echo "📡 เซิร์ฟเวอร์จะรันที่: http://localhost:${PORT:-8080}"
    echo "🩺 Health check: http://localhost:${PORT:-8080}/health"
    echo "📖 API Documentation อยู่ในไฟล์ SETUP_GUIDE.md"
else
    echo "❌ สร้างเซิร์ฟเวอร์ไม่สำเร็จ กรุณาตรวจสอบ error ข้างต้น"
    exit 1
fi
