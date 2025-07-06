# POS Backend API

Backend API สำหรับระบบ Point of Sale (POS) ร้านอาหาร

## 🔗 Related Repository
- **Frontend**: [Frontend-POS](https://github.com/komkemkku/Frontend-POS.git)
- **Backend**: [Backend-POS](https://github.com/komkemkku/Backend-POS.git) (this repository)

## เทคโนโลยี
- Go 1.23.4
- Gin Framework
- PostgreSQL + Bun ORM
- JWT Authentication

## การติดตั้ง

```bash
# Clone โปรเจค Backend
git clone https://github.com/komkemkku/Backend-POS.git
cd Backend-POS

# Clone โปรเจค Frontend (ในอีก terminal)
git clone https://github.com/komkemkku/Frontend-POS.git

# Setup development environment
make setup

# ติดตั้ง dependencies
make deps

# ตั้งค่า environment variables
# แก้ไขค่าในไฟล์ .env ที่สร้างจาก make setup

# รัน migration
make migrate

# รันเซิร์ฟเวอร์
make dev
```

## API Endpoints

### Public (สำหรับลูกค้า)
- `GET /public/menu/:qrCode` - ดูเมนูโต๊ะ
- `POST /public/orders/create` - สั่งอาหาร
- `GET /public/orders/table/:qrCode` - ดูออเดอร์ปัจจุบัน

### Staff (ต้อง Authentication)
- `POST /staff/login` - เข้าสู่ระบบ
- `GET /orders` - ดูออเดอร์ทั้งหมด
- `PATCH /orders/:id/status` - เปลี่ยนสถานะออเดอร์

## โครงสร้างโปรเจค
```
Backend-POS/
├── cmd/          # Commands & migrations
├── configs/      # Database configuration
├── controller/   # API controllers
├── middlewares/  # Auth middleware
├── model/        # Database models
├── requests/     # Request structs
├── responses/    # Response structs
├── utils/        # Utility functions
└── main.go       # Entry point
```

## 🌐 CORS Configuration
API รองรับ CORS สำหรับ Frontend domains:
- `http://localhost:3000` (Next.js development)
- `http://localhost:5173` (Vite development)
- `https://*.vercel.app` (Production deployment)

## 📱 Development Workflow
1. เริ่ม Backend server: `go run main.go` (port 8080)
2. เริ่ม Frontend development server ในโฟลเดอร์ Frontend-POS
3. Frontend จะเชื่อมต่อกับ Backend API ผ่าน localhost:8080

## 🚀 Auto-Deploy (Backend)
Backend มีระบบ auto-deploy ที่จะ commit และ push ทุกครั้งที่มีการแก้ไข:

```bash
# วิธีใช้ auto-deploy
make deploy
# หรือ
./auto-deploy.sh

# คำสั่งอื่นๆ
make dev        # รัน development server
make migrate    # รัน database migration
make test       # รัน tests
make help       # ดูคำสั่งทั้งหมด
```

## 🚀 Auto-Deploy

โปรเจคนี้มีระบบ auto-deploy สำหรับ Railway โดยอัตโนมัติ

### วิธีใช้ Auto-Deploy

**สำหรับ Windows (PowerShell):**
```powershell
# Auto commit และ push การเปลี่ยนแปลงทั้งหมด
make deploy-win
```

**สำหรับ Linux/Mac:**
```bash
# Auto commit และ push การเปลี่ยนแปลงทั้งหมด
make deploy
```

**การใช้งานโดยตรง:**
```powershell
# Windows
powershell -ExecutionPolicy Bypass -File ./auto-deploy.ps1

# Linux/Mac
./auto-deploy.sh
```

### คุณสมบัติ Auto-Deploy

- ✅ ตรวจสอบการเปลี่ยนแปลงอัตโนมัติ
- ✅ สร้าง commit message ที่มีรายละเอียด
- ✅ แสดงไฟล์ที่เปลี่ยนแปลง
- ✅ Push ไปยัง Git repository
- ✅ Railway จะทำการ deploy อัตโนมัติ
- ✅ แสดงสถานะการ deploy

### หลังจาก Deploy

หลังจากรัน auto-deploy แล้ว:
1. Railway จะตรวจจับการเปลี่ยนแปลงใน Git
2. ทำการ build และ deploy backend อัตโนมัติ
3. ตรวจสอบสถานะได้ที่: https://railway.app
