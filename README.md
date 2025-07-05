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

# ติดตั้ง dependencies
go mod tidy

# ตั้งค่า environment variables
cp .env.example .env
# แก้ไขค่าในไฟล์ .env

# รัน migration
go run cmd/migrateCmd.go

# รันเซิร์ฟเวอร์
go run main.go
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
