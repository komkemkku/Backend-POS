# Backend POS System - API Documentation

## ระบบ Backend POS สำหรับร้านอาหาร
ระบบจัดการหลังบ้านและ Public API สำหรับลูกค้าของร้านอาหาร รองรับการสั่งอาหารผ่าน QR Code

### ขั้นตอนการติดตั้งและรัน

#### 1. ติดตั้ง Dependencies
```bash
go mod tidy
```

#### 2. ตั้งค่าฐานข้อมูล
สร้างไฟล์ `.env` จาก `.env.example`:
```bash
cp .env.example .env
```

แก้ไขไฟล์ `.env` ให้ตรงกับฐานข้อมูล PostgreSQL ของคุณ:
```env
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=pos_db
DB_USER=postgres
DB_PASSWORD=your_password
```

#### 3. รันการ Migration (สร้างตาราง)
```bash
go run main.go migrate
```

#### 4. รันเซิร์ฟเวอร์
```bash
# Development
go run main.go

# หรือ build และรัน
go build -o pos-server
./pos-server
```

### Public API Endpoints สำหรับลูกค้า

#### 1. ดูเมนูตาม QR Code โต๊ะ
```
GET /public/menu/:qrCodeIdentifier
```
**Response:**
```json
{
  "success": true,
  "message": "เมนูโต๊ะ 1",
  "data": {
    "table_info": {
      "id": 1,
      "table_number": "1",
      "qr_code_identifier": "table_001",
      "status": "available"
    },
    "menu_items": [
      {
        "id": 1,
        "name": "ผัดไทย",
        "description": "ผัดไทยแสนอร่อย",
        "price": 60.00,
        "category_id": 1,
        "image_url": "https://example.com/padthai.jpg",
        "is_available": true
      }
    ]
  }
}
```

#### 2. ดูเมนูทั้งหมด
```
GET /public/menu-items
```

#### 3. สร้างออเดอร์ (ลูกค้าสั่งอาหาร)
```
POST /public/orders/create
```
**Request Body:**
```json
{
  "qr_code_identifier": "table_001",
  "items": [
    {
      "menu_item_id": 1,
      "quantity": 2
    },
    {
      "menu_item_id": 2,
      "quantity": 1
    }
  ]
}
```

**Response:**
```json
{
  "success": true,
  "message": "สร้างออเดอร์สำเร็จ",
  "data": {
    "id": 1,
    "table_id": 1,
    "table_number": "1",
    "status": "pending",
    "total_amount": 180.00,
    "items": [
      {
        "id": 1,
        "menu_item_id": 1,
        "quantity": 2,
        "price_per_item": 60.00,
        "sub_total": 120.00
      }
    ],
    "created_at": 1704067200,
    "message": "ออเดอร์ของคุณได้รับการยืนยันแล้ว กรุณารอสักครู่"
  }
}
```

#### 4. ดูประวัติออเดอร์ตามโต๊ะ (เฉพาะยังไม่ชำระเงิน)
```
GET /public/orders/table/:qrCodeIdentifier
```

#### 5. ดูสถานะออเดอร์เฉพาะ
```
GET /public/orders/:orderID/table/:qrCodeIdentifier
```
**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "table_number": "1",
    "status": "preparing",
    "status_text": "กำลังเตรียม",
    "status_color": "blue",
    "estimated_time": "10-15 นาที",
    "total_amount": 180.00,
    "items": [...],
    "created_at": 1704067200
  }
}
```

#### 6. ดูประวัติออเดอร์ทั้งหมด (รวมที่ชำระแล้ว)
```
GET /public/orders/history/:qrCodeIdentifier
```
**Response:**
```json
{
  "success": true,
  "data": {
    "table_info": {
      "table_number": "1",
      "qr_code_identifier": "table_001"
    },
    "current_orders": [
      // ออเดอร์ที่ยังไม่ชำระเงิน
    ],
    "paid_orders": [
      // ออเดอร์ที่ชำระแล้ว
    ],
    "summary": {
      "total_orders": 5,
      "total_spent": 850.00,
      "current_pending": 2,
      "completed_today": 3
    }
  }
}
```

#### 7. ดูสรุปโต๊ะ
```
GET /public/table/summary/:qrCodeIdentifier
```
**Response:**
```json
{
  "success": true,
  "data": {
    "table_info": {
      "table_number": "1",
      "status": "occupied"
    },
    "order_counts": {
      "pending": 1,
      "preparing": 2,
      "ready": 0,
      "total": 3
    },
    "total_pending": 280.00,
    "last_updated": 1704067200
  }
}
```

### Staff API Endpoints

#### ล้างประวัติโต๊ะหลังชำระเงิน (ต้อง Authentication)
```
POST /staff/orders/clear-table/:qrCodeIdentifier
Authorization: Bearer <token>
```

#### ล้างประวัติแบบละเอียด (เพิ่มฟีเจอร์)
```
POST /staff/orders/advanced-clear/:qrCodeIdentifier?type=payment
Authorization: Bearer <token>
```
**Query Parameters:**
- `type`: ประเภทการล้าง
  - `payment` (default) - ชำระเงินแล้ว
  - `cancel_all` - ยกเลิกทั้งหมด
  - `complete_all` - เสร็จสิ้นทั้งหมด

**Response:**
```json
{
  "success": true,
  "data": {
    "success": true,
    "clear_type": "payment",
    "orders_affected": 3,
    "total_amount": 450.00,
    "table_status": "available",
    "cleared_orders": [...],
    "timestamp": 1704067200,
    "message": "ชำระเงินเรียบร้อย ได้ล้างประวัติ 3 ออเดอร์"
  }
}
```

#### ยกเลิกออเดอร์เฉพาะ
```
POST /staff/orders/cancel/:orderID/table/:qrCodeIdentifier?reason=เหตุผล
Authorization: Bearer <token>
```

#### อัปเดตสถานะออเดอร์
```
PATCH /staff/orders/:orderID/status
Authorization: Bearer <token>
```
**Request Body:**
```json
{
  "status": "preparing"
}
```

**สถานะที่อนุญาต:**
- `pending` - รอดำเนินการ
- `preparing` - กำลังเตรียม
- `ready` - พร้อมเสิร์ฟ
- `served` - เสิร์ฟแล้ว
- `completed` - เสร็จสิ้น
- `cancelled` - ยกเลิก

### การทำงานของระบบ

1. **ลูกค้าสแกน QR Code** → ดูเมนูของโต๊ะนั้น
2. **ลูกค้าสั่งอาหาร** → ระบบสร้างออเดอร์สถานะ "pending"
3. **ลูกค้าติดตามสถานะ** → pending → preparing → ready → served
4. **ลูกค้าดูประวัติ** → เห็นเฉพาะออเดอร์ที่ยังไม่ชำระเงิน
5. **Staff ชำระเงิน** → เรียก API clear-table → เปลี่ยนสถานะเป็น "paid"
6. **ลูกค้าดูประวัติทั้งหมด** → เห็นได้ทั้งออเดอร์ปัจจุบันและที่ชำระแล้ว

### ฟีเจอร์ขั้นสูง

#### การล้างประวัติแบบละเอียด
- **payment**: ชำระเงินแล้ว → เปลี่ยนสถานะเป็น "paid" + โต๊ะว่าง
- **cancel_all**: ยกเลิกทั้งหมด → เปลี่ยนสถานะเป็น "cancelled"
- **complete_all**: เสร็จสิ้นทั้งหมด → เปลี่ยนสถานะเป็น "completed" + โต๊ะว่าง

#### การยกเลิกออเดอร์เฉพาะ
- Staff สามารถยกเลิกออเดอร์ใดออเดอร์หนึ่งได้
- ตรวจสอบว่าออเดอร์ยังสามารถยกเลิกได้ (ไม่ใช่ paid/completed/cancelled)
- บันทึกเหตุผลการยกเลิก

#### การอัปเดตสถานะออเดอร์
- Staff สามารถเปลี่ยนสถานะออเดอร์ได้ตามขั้นตอนการทำงาน
- ตรวจสอบสถานะที่อนุญาตเท่านั้น
- อัปเดต timestamp อัตโนมัติ

### สถานะออเดอร์ (Order Status)
- `pending` - รอดำเนินการ (สีส้ม)
- `preparing` - กำลังเตรียม (สีน้ำเงิน)
- `ready` - พร้อมเสิร์ฟ (สีเขียว)
- `served` - เสิร์ฟแล้ว (สีม่วง)
- `paid` - ชำระเงินแล้ว (สีเทา)
- `completed` - เสร็จสิ้น (สีเทา)
- `cancelled` - ยกเลิก (สีแดง)

### CORS Configuration
ระบบรองรับการเรียกใช้จาก:
- localhost:3000 (React)
- localhost:5173 (Vite)
- *.vercel.app (Vercel deployment)
- frontend-pos-jade.vercel.app (โดเมนเฉพาะ)

### การใช้งานกับ Frontend

#### การสแกน QR Code
QR Code ควรมี URL format:
```
https://your-frontend.vercel.app/table/table_001
```
โดย `table_001` คือ `qr_code_identifier` ในฐานข้อมูล

#### Flow การใช้งาน
1. Customer สแกน QR → Frontend รับ qr_code_identifier
2. Frontend เรียก `/public/menu/:qrCodeIdentifier` → แสดงเมนู
3. Customer สั่งอาหาร → Frontend เรียก `/public/orders/create`
4. Frontend แสดงหน้าติดตามออเดอร์ → เรียก `/public/orders/:orderID/table/:qrCodeIdentifier`
5. Customer ดูประวัติ → Frontend เรียก `/public/orders/table/:qrCodeIdentifier`
6. Staff ชำระเงิน → Frontend เรียก `/staff/orders/clear-table/:qrCodeIdentifier`

### ข้อมูลเพิ่มเติม

- ระบบใช้ PostgreSQL เป็นฐานข้อมูล
- ใช้ Go Gin Framework
- มี JWT Authentication สำหรับ Staff
- Public API ไม่ต้อง Authentication
- รองรับ CORS สำหรับ Frontend deployment
- มี Health Check endpoint: `/health` และ `/ping`

### หมายเหตุสำคัญ
- ระบบแยกออเดอร์ที่ "ยังไม่ชำระเงิน" กับ "ชำระแล้ว" อย่างชัดเจน
- ลูกค้าจะเห็นเฉพาะออเดอร์ปัจจุบัน ไม่เห็นของคนอื่นที่เคยใช้โต๊ะนั้น
- เมื่อ Staff ล้างโต๊ะ = เปลี่ยนออเดอร์เป็น "paid" + เปลี่ยนโต๊ะเป็น "available"
- ระบบจะนับเฉพาะออเดอร์วันเดียวกัน เพื่อไม่ให้ข้อมูลซ้อนกับวันอื่น
