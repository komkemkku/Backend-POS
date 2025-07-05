# 🎉 Backend POS - พร้อมส่งมอบให้ Frontend

## ✅ **API ที่พร้อมใช้งาน:**

### 🔐 **Authentication APIs**
```bash
POST /staff/login               # เข้าสู่ระบบ
GET  /staff/info                # ข้อมูลพนักงานที่ล็อกอิน ✅ ส่ง full_name, role
```

### 📊 **Dashboard API**
```bash
GET  /summary                   # สรุปข้อมูล Dashboard ✅ ใหม่!
```

**Response `/summary`:**
```json
{
  "status": "success",
  "data": {
    "total_tables": 5,          # จำนวนโต๊ะทั้งหมด
    "today_revenue": 2350.75,   # รายได้วันนี้
    "today_orders": 15,         # ออเดอร์วันนี้
    "pending_orders": 3         # ออเดอร์รอดำเนินการ
  }
}
```

### 🍽️ **Public APIs (สำหรับลูกค้า)**
```bash
GET  /public/menu/:qrCode                              # ดูเมนูตาม QR Code
GET  /public/menu                                      # ดูเมนูทั้งหมด
POST /public/orders/create                             # สร้างออเดอร์
GET  /public/orders/table/:qrCode                      # ดูออเดอร์ปัจจุบัน
GET  /public/orders/:orderID/table/:qrCode             # ดูสถานะออเดอร์
GET  /public/orders/history/:qrCode                    # ดูประวัติออเดอร์
GET  /public/table/summary/:qrCode                     # สรุปโต๊ะ
```

### 👨‍💼 **Staff APIs (ต้อง Authentication)**
```bash
# Orders Management
GET    /orders                                         # รายการออเดอร์ทั้งหมด
PATCH  /staff/orders/:orderID/status                   # อัปเดตสถานะออเดอร์
POST   /staff/orders/clear-table/:qrCode               # ล้างประวัติโต๊ะ
POST   /staff/orders/cancel/:orderID/table/:qrCode     # ยกเลิกออเดอร์

# Menu Management
GET    /menu-items                                     # รายการเมนู
POST   /menu-items/create                              # เพิ่มเมนู
PATCH  /menu-items/:id                                 # แก้ไขเมนู
DELETE /menu-items/:id                                 # ลบเมนู

# Table Management
GET    /tables                                         # รายการโต๊ะ
POST   /tables/create                                  # เพิ่มโต๊ะ
PATCH  /tables/:id                                     # แก้ไขโต๊ะ

# และอื่นๆ (categories, staff, payments, reservations, expenses)
```

---

## 🚀 **สำหรับทีม Frontend:**

### ✅ **Authentication Flow**
```javascript
// 1. Login
const loginResponse = await fetch('/staff/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ username: 'admin', password: 'password' })
});

const { data } = await loginResponse.json();
const token = data.token;

// 2. Get Staff Info
const staffInfo = await fetch('/staff/info', {
  headers: { 'Authorization': `Bearer ${token}` }
});

// 3. Get Dashboard Summary
const summary = await fetch('/summary', {
  headers: { 'Authorization': `Bearer ${token}` }
});
```

### ✅ **CORS Configuration**
Frontend domains ที่รองรับ:
- `http://localhost:3000`
- `http://localhost:5173`
- `https://*.vercel.app`
- `https://komkemkty-frontend-pos.vercel.app`
- `https://frontend-pos-jade.vercel.app`

### ✅ **Error Handling**
Response format มาตรฐาน:
```json
// Success
{
  "status": "success",
  "data": { ... }
}

// Error
{
  "status": "error",
  "message": "รายละเอียดข้อผิดพลาด"
}
```

---

## 📋 **ข้อมูลตัวอย่างในระบบ:**

### 👤 **Staff Account**
```
Username: admin
Password: password
Role: admin
Full Name: ผู้ดูแลระบบ
```

### 🍽️ **Sample Menu Items**
- ผัดไทย (60 บาท)
- ต้มยำกุ้ง (80 บาท) 
- ข้าวผัดกุ้ง (70 บาท)

### 🪑 **Sample Tables**
- โต๊ะ 1: QR Code = "table_001"
- โต๊ะ 2: QR Code = "table_002"
- โต๊ะ 3: QR Code = "table_003"

---

## 🔗 **Production URLs:**

### Backend API
```
https://backend-pos-production.up.railway.app
```

### Frontend Demo
```
https://frontend-7t7jmbu4p-komkems-projects.vercel.app
```

---

## 🧪 **ทดสอบ API ด้วย curl:**

```bash
# Health Check
curl https://backend-pos-production.up.railway.app/health

# Login
curl -X POST https://backend-pos-production.up.railway.app/staff/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password"}'

# Dashboard Summary (ใส่ token ที่ได้จาก login)
curl https://backend-pos-production.up.railway.app/summary \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# Public Menu
curl https://backend-pos-production.up.railway.app/public/menu
```

---

## 🎯 **Status: 100% Ready!**

✅ **API `/staff/info`** - ส่ง full_name, role  
✅ **API `/summary`** - ส่ง total_tables, today_revenue, today_orders, pending_orders  
✅ **CORS** - รองรับ Vercel domains  
✅ **Authentication** - JWT tokens  
✅ **Error Handling** - Response format มาตรฐาน  
✅ **Production Deploy** - Railway  

**🚀 Frontend สามารถเชื่อมต่อและใช้งานได้เต็มรูปแบบแล้ว!**
