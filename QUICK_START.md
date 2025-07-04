# 🚀 Quick Start Guide - Backend POS

## ⚡ เริ่มใช้งานใน 5 นาที

### 1. Setup ระบบ
```bash
# Clone project
git clone <your-repo>
cd Backend-POS

# Setup dependencies
go mod tidy

# สร้างไฟล์ .env
cp .env.example .env
# แก้ไขค่าฐานข้อมูลใน .env

# ติดตั้งผ่าน script
chmod +x setup.sh
./setup.sh
```

### 2. เริ่มใช้งาน
```bash
# รัน migration (ครั้งแรก)
./pos-server migrate

# เพิ่มข้อมูลตัวอย่าง (ถ้าต้องการ)
psql -h localhost -U postgres -d pos_db -f sample_data.sql

# รันเซิร์ฟเวอร์
./pos-server
# หรือ go run main.go
```

### 3. ทดสอบพื้นฐาน
```bash
# Health check
curl http://localhost:8080/ping

# ดูเมนูโต๊ะ 1
curl http://localhost:8080/public/menu/table_001

# สั่งอาหาร
curl -X POST http://localhost:8080/public/orders/create \
  -H "Content-Type: application/json" \
  -d '{
    "qr_code_identifier": "table_001",
    "items": [{"menu_item_id": 1, "quantity": 2}]
  }'
```

## 🎯 API Endpoints หลัก

### ลูกค้า (Public)
| Method | Endpoint | ใช้งาน |
|--------|----------|-------|
| `GET` | `/public/menu/:qrCode` | ดูเมนูโต๊ะ |
| `POST` | `/public/orders/create` | สั่งอาหาร |
| `GET` | `/public/orders/table/:qrCode` | ดูออเดอร์ปัจจุบัน |
| `GET` | `/public/orders/:id/table/:qrCode` | ดูสถานะออเดอร์ |
| `GET` | `/public/table/summary/:qrCode` | ดูสรุปโต๊ะ |

### พนักงาน (Staff)
| Method | Endpoint | ใช้งาน |
|--------|----------|-------|
| `POST` | `/auth/login` | เข้าสู่ระบบ |
| `POST` | `/staff/orders/clear-table/:qrCode` | ล้างประวัติหลังชำระเงิน |
| `PATCH` | `/staff/orders/:id/status` | เปลี่ยนสถานะออเดอร์ |

## 📱 ตัวอย่างการใช้งาน

### ลูกค้าสั่งอาหาร
```javascript
// 1. ดูเมนู
const menu = await fetch('/public/menu/table_001').then(r => r.json());

// 2. สั่งอาหาร
const order = await fetch('/public/orders/create', {
  method: 'POST',
  headers: {'Content-Type': 'application/json'},
  body: JSON.stringify({
    qr_code_identifier: 'table_001',
    items: [{menu_item_id: 1, quantity: 2}]
  })
}).then(r => r.json());

// 3. ติดตามสถานะ
const status = await fetch(`/public/orders/${order.data.id}/table/table_001`).then(r => r.json());
```

### พนักงานจัดการ
```javascript
// 1. Login
const auth = await fetch('/auth/login', {
  method: 'POST',
  headers: {'Content-Type': 'application/json'},
  body: JSON.stringify({
    email: 'admin@restaurant.com',
    password: 'password'
  })
}).then(r => r.json());

// 2. เปลี่ยนสถานะออเดอร์
const update = await fetch('/staff/orders/1/status', {
  method: 'PATCH',
  headers: {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${auth.token}`
  },
  body: JSON.stringify({status: 'preparing'})
}).then(r => r.json());

// 3. ล้างประวัติโต๊ะ
const clear = await fetch('/staff/orders/clear-table/table_001', {
  method: 'POST',
  headers: {'Authorization': `Bearer ${auth.token}`}
}).then(r => r.json());
```

## 🔧 Environment Variables

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=pos_db
DB_USER=postgres
DB_PASSWORD=your_password

# Server
PORT=8080
GIN_MODE=debug

# JWT (optional)
JWT_SECRET=your_secret_key
```

## 📊 ข้อมูลสำคัญ

### QR Code Format
```
https://your-frontend.com/table/table_001
```

### สถานะออเดอร์
- `pending` → `preparing` → `ready` → `served` → `paid`

### Response Format
```json
{
  "success": true,
  "message": "สำเร็จ",
  "data": { ... }
}
```

## 🚨 Troubleshooting

### ปัญหาเชื่อมต่อฐานข้อมูล
```bash
# ตรวจสอบ PostgreSQL running
brew services start postgresql
# หรือ
sudo systemctl start postgresql

# ตรวจสอบ .env file
cat .env
```

### ปัญหา CORS
- ตรวจสอบ domain ใน main.go
- เพิ่ม domain ของคุณใน AllowOrigins

### ปัญหา Authentication
- ตรวจสอบ JWT_SECRET
- ตรวจสอบ Bearer token format

## 📚 เอกสารเพิ่มเติม
- `SETUP_GUIDE.md` - คู่มือละเอียด
- `API_TESTING.md` - ตัวอย่างทดสอบ
- `FEATURES_SUMMARY.md` - สรุปฟีเจอร์

---
🎉 **พร้อมใช้งาน!** ระบบ POS สมบูรณ์สำหรับร้านอาหารยุคใหม่
