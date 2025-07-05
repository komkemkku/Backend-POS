# 📋 Backend POS API - สรุปสำหรับทีม Frontend

## 🎯 **สรุปการทำงาน: Backend ส่งอะไรให้ Frontend / Frontend ส่งอะไรมาให้ Backend**

---

## 📱 **PUBLIC API (สำหรับลูกค้า - ไม่ต้อง Authentication)**

### 1️⃣ **ดูเมนูตาม QR Code โต๊ะ**
```
GET /public/menu/{qrCodeIdentifier}
```

**Frontend ส่งมา:**
- `qrCodeIdentifier` ใน URL (เช่น "table_001")

**Backend ส่งกลับ:**
```json
{
  "success": true,
  "message": "เมนูโต๊ะ 1",
  "data": {
    "table_info": {
      "id": 1,
      "table_number": 1,
      "qr_code_identifier": "table_001",
      "status": "available"
    },
    "menu_items": [
      {
        "id": 1,
        "name": "ผัดไทย",
        "description": "ผัดไทยแสนอร่อยเส้นหมี่ใหญ่",
        "price": 60.00,
        "category_id": 1,
        "image_url": "https://example.com/padthai.jpg",
        "is_available": true,
        "category_name": "อาหารจานหลัก"
      }
    ]
  }
}
```

---

### 2️⃣ **ดูเมนูทั้งหมด**
```
GET /public/menu
```

**Frontend ส่งมา:**
- ไม่ต้องส่งอะไร

**Backend ส่งกลับ:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "ผัดไทย",
      "description": "ผัดไทยแสนอร่อยเส้นหมี่ใหญ่",
      "price": 60.00,
      "category_id": 1,
      "image_url": "https://example.com/padthai.jpg",
      "is_available": true
    }
  ]
}
```

---

### 3️⃣ **สร้างออเดอร์ (ลูกค้าสั่งอาหาร)**
```
POST /public/orders/create
```

**Frontend ส่งมา:**
```json
{
  "qr_code_identifier": "table_001",
  "items": [
    {
      "menu_item_id": 1,
      "quantity": 2
    },
    {
      "menu_item_id": 3,
      "quantity": 1
    }
  ]
}
```

**Backend ส่งกลับ:**
```json
{
  "success": true,
  "message": "สร้างออเดอร์สำเร็จ",
  "data": {
    "id": 123,
    "table_id": 1,
    "table_number": 1,
    "status": "pending",
    "status_text": "รอดำเนินการ",
    "total_amount": 270.00,
    "items": [
      {
        "id": 1,
        "menu_item_id": 1,
        "menu_item_name": "ผัดไทย",
        "quantity": 2,
        "price_per_item": 60.00,
        "sub_total": 120.00,
        "notes": ""
      }
    ],
    "created_at": 1704067200,
    "message": "ออเดอร์ของคุณได้รับการยืนยันแล้ว กรุณารอสักครู่"
  }
}
```

---

### 4️⃣ **ดูประวัติออเดอร์ปัจจุบัน (ยังไม่ชำระเงิน)**
```
GET /public/orders/table/{qrCodeIdentifier}
```

**Frontend ส่งมา:**
- `qrCodeIdentifier` ใน URL

**Backend ส่งกลับ:**
```json
{
  "success": true,
  "data": [
    {
      "id": 123,
      "table_id": 1,
      "table_number": 1,
      "status": "preparing",
      "status_text": "กำลังเตรียม",
      "status_color": "#0066CC",
      "estimated_time": "10-15 นาที",
      "total_amount": 270.00,
      "items": [
        {
          "id": 1,
          "order_id": 123,
          "menu_item_id": 1,
          "menu_item_name": "ผัดไทย",
          "quantity": 2,
          "price_per_item": 60.00,
          "sub_total": 120.00,
          "notes": ""
        }
      ],
      "created_at": 1704067200,
      "updated_at": 1704067300
    }
  ]
}
```

---

### 5️⃣ **ดูสถานะออเดอร์เฉพาะ**
```
GET /public/orders/{orderID}/table/{qrCodeIdentifier}
```

**Frontend ส่งมา:**
- `orderID` ใน URL (เช่น 123)
- `qrCodeIdentifier` ใน URL (เช่น "table_001")

**Backend ส่งกลับ:**
```json
{
  "success": true,
  "data": {
    "id": 123,
    "table_id": 1,
    "table_number": 1,
    "status": "preparing",
    "status_text": "กำลังเตรียม",
    "status_color": "#0066CC",
    "estimated_time": "10-15 นาที",
    "total_amount": 270.00,
    "items": [...],
    "created_at": 1704067200,
    "updated_at": 1704067300
  }
}
```

---

### 6️⃣ **ดูประวัติทั้งหมด (แยกปัจจุบัน/ชำระแล้ว)**
```
GET /public/orders/history/{qrCodeIdentifier}
```

**Frontend ส่งมา:**
- `qrCodeIdentifier` ใน URL

**Backend ส่งกลับ:**
```json
{
  "success": true,
  "data": {
    "table_info": {
      "id": 1,
      "table_number": 1,
      "qr_code_identifier": "table_001",
      "status": "occupied"
    },
    "current_orders": [
      {
        "id": 123,
        "status": "preparing",
        "status_text": "กำลังเตรียม",
        "total_amount": 270.00,
        "items": [...],
        "created_at": 1704067200
      }
    ],
    "paid_orders": [
      {
        "id": 120,
        "status": "paid",
        "status_text": "ชำระเงินแล้ว",
        "total_amount": 180.00,
        "items": [...],
        "created_at": 1704060000
      }
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

---

### 7️⃣ **ดูสรุปโต๊ะ**
```
GET /public/table/summary/{qrCodeIdentifier}
```

**Frontend ส่งมา:**
- `qrCodeIdentifier` ใน URL

**Backend ส่งกลับ:**
```json
{
  "success": true,
  "data": {
    "table_info": {
      "id": 1,
      "table_number": 1,
      "qr_code_identifier": "table_001",
      "status": "occupied"
    },
    "order_counts": {
      "pending": 1,
      "preparing": 2,
      "ready": 0,
      "total": 3
    },
    "total_pending": 450.00,
    "last_updated": 1704067500
  }
}
```

---

## 👨‍💼 **STAFF API (ต้อง Authentication)**

### 🔐 **Authentication - เข้าสู่ระบบ**
```
POST /staff/login
```

**Frontend ส่งมา:**
```json
{
  "username": "admin",
  "password": "password"
}
```

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "staff": {
      "id": 1,
      "username": "admin",
      "full_name": "ผู้ดูแลระบบ",
      "role": "admin"
    }
  }
}
```

---

### 👤 **ข้อมูลพนักงานที่ล็อกอิน**
```
GET /staff/info
Authorization: Bearer <token>
```

**Frontend ส่งมา:**
- `Authorization: Bearer <token>` ใน Header

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "username": "admin",
    "full_name": "ผู้ดูแลระบบ",
    "role": "admin",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### 📊 **สรุปข้อมูล Dashboard** ✨ **NEW**
```
GET /summary
Authorization: Bearer <token>
```

**Frontend ส่งมา:**
- `Authorization: Bearer <token>` ใน Header

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "data": {
    "total_tables": 5,        // จำนวนโต๊ะทั้งหมดในระบบ
    "today_revenue": 2350.75, // รายได้วันนี้ (จากการชำระเงิน)
    "today_orders": 15,       // จำนวนออเดอร์วันนี้
    "pending_orders": 3       // ออเดอร์ที่รอดำเนินการ (pending, preparing, ready)
  }
}
```

---

### 🔄 **อัปเดตสถานะออเดอร์**
```
PATCH /staff/orders/{orderID}/status
Authorization: Bearer <token>
```

**Frontend ส่งมา:**
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

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "data": {
    "id": 123,
    "table_id": 1,
    "staff_id": 1,
    "status": "preparing",
    "status_text": "กำลังเตรียม",
    "total_amount": 270.00,
    "created_at": 1704067200,
    "updated_at": 1704067300
  }
}
```

---

### 🗑️ **ล้างประวัติโต๊ะหลังชำระเงิน**
```
POST /staff/orders/clear-table/{qrCodeIdentifier}
Authorization: Bearer <token>
```

**Frontend ส่งมา:**
- `qrCodeIdentifier` ใน URL
- `Authorization: Bearer <token>` ใน Header

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "message": "ล้างประวัติโต๊ะเรียบร้อยแล้ว",
  "data": {
    "orders_cleared": 3,
    "total_amount": 450.00
  }
}
```

---

### 🗑️ **ล้างประวัติแบบละเอียด**
```
POST /staff/orders/advanced-clear/{qrCodeIdentifier}?type={clearType}
Authorization: Bearer <token>
```

**Frontend ส่งมา:**
- `qrCodeIdentifier` ใน URL
- `type` ใน Query Parameters:
  - `payment` (default) - ชำระเงินแล้ว
  - `cancel_all` - ยกเลิกทั้งหมด  
  - `complete_all` - เสร็จสิ้นทั้งหมด
- `Authorization: Bearer <token>` ใน Header

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "data": {
    "success": true,
    "clear_type": "payment",
    "orders_affected": 3,
    "total_amount": 450.00,
    "table_status": "available",
    "cleared_orders": [
      {
        "id": 123,
        "table_id": 1,
        "status": "paid",
        "total_amount": 270.00,
        "created_at": 1704067200,
        "completed_at": 1704067500
      }
    ],
    "message": "ชำระเงินเรียบร้อย ได้ล้างประวัติ 3 ออเดอร์"
  }
}
```

---

### ❌ **ยกเลิกออเดอร์เฉพาะ**
```
POST /staff/orders/cancel/{orderID}/table/{qrCodeIdentifier}?reason={reason}
Authorization: Bearer <token>
```

**Frontend ส่งมา:**
- `orderID` ใน URL
- `qrCodeIdentifier` ใน URL
- `reason` ใน Query Parameters (optional)
- `Authorization: Bearer <token>` ใน Header

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "message": "ยกเลิกออเดอร์เรียบร้อยแล้ว",
  "data": {
    "order_id": 123,
    "reason": "ลูกค้าไม่ต้องการ",
    "cancelled_at": 1704067500
  }
}
```

---

### 📋 **ดูรายการออเดอร์ทั้งหมด (Staff)**
```
GET /orders?page=1&size=10&search=
Authorization: Bearer <token>
```

**Frontend ส่งมา:**
- `page` ใน Query (optional, default: 1)
- `size` ใน Query (optional, default: 10)
- `search` ใน Query (optional)
- `Authorization: Bearer <token>` ใน Header

**Backend ส่งกลับ:**
```json
{
  "status": "success",
  "data": [
    {
      "id": 123,
      "table_id": 1,
      "table_number": 1,
      "staff_id": 1,
      "status": "preparing",
      "status_text": "กำลังเตรียม",
      "total_amount": 270.00,
      "items_count": 2,
      "created_at": 1704067200,
      "updated_at": 1704067300
    }
  ],
  "pagination": {
    "total": 25,
    "page": 1,
    "size": 10,
    "total_pages": 3
  }
}
```

---

## 🎨 **ระบบสถานะและสี (สำหรับ Frontend แสดงผล)**

### 📋 **สถานะออเดอร์**
| Status | Text (ภาษาไทย) | Color | Estimated Time | สำหรับ UI |
|--------|----------------|-------|----------------|-----------|
| `pending` | รอดำเนินการ | `#FFA500` | 5-10 นาที | 🟠 Orange |
| `preparing` | กำลังเตรียม | `#0066CC` | 10-15 นาที | 🔵 Blue |
| `ready` | พร้อมเสิร์ฟ | `#00CC00` | พร้อมแล้ว | 🟢 Green |
| `served` | เสิร์ฟแล้ว | `#9900CC` | เสร็จสิ้น | 🟣 Purple |
| `paid` | ชำระเงินแล้ว | `#999999` | เสร็จสิ้น | ⚪ Gray |
| `completed` | เสร็จสิ้น | `#999999` | เสร็จสิ้น | ⚪ Gray |
| `cancelled` | ยกเลิก | `#CC0000` | ยกเลิกแล้ว | 🔴 Red |

### 🪑 **สถานะโต๊ะ**
| Status | Text (ภาษาไทย) | Color | สำหรับ UI |
|--------|----------------|-------|-----------|
| `available` | ว่าง | `#00CC00` | 🟢 Green |
| `occupied` | มีลูกค้า | `#FFA500` | 🟠 Orange |
| `reserved` | จอง | `#0066CC` | 🔵 Blue |
| `maintenance` | ปิดปรับปรุง | `#CC0000` | 🔴 Red |

---

## 🔧 **Error Handling**

### ❌ **Error Response Format**
```json
{
  "success": false,
  "message": "ข้อความ error ภาษาไทย",
  "error": "technical_error_code"
}
```

### 📊 **HTTP Status Codes**
- `200` - สำเร็จ
- `400` - Bad Request (ข้อมูลไม่ถูกต้อง)
- `401` - Unauthorized (ไม่มีสิทธิ์/token หมดอายุ)
- `404` - Not Found (ไม่พบข้อมูล)
- `500` - Internal Server Error (ข้อผิดพลาดของเซิร์ฟเวอร์)

### 🚨 **ตัวอย่าง Error Responses**
```json
// โต๊ะไม่พบ
{
  "success": false,
  "message": "ไม่พบโต๊ะที่ระบุ"
}

// เมนูไม่พร้อมใช้งาน
{
  "success": false,
  "message": "ไม่พบเมนู ID: 999 หรือเมนูไม่พร้อมใช้งาน"
}

// Token หมดอายุ
{
  "success": false,
  "message": "Token หมดอายุ กรุณาเข้าสู่ระบบใหม่"
}

// ข้อมูลไม่ครบ
{
  "success": false,
  "message": "กรุณาระบุ qr_code_identifier และ items"
}
```

---

## 🌐 **CORS และ Environment**

### 🔗 **Base URLs**
```javascript
// Development
const API_BASE_URL = 'http://localhost:8080';

// Production
const API_BASE_URL = 'https://backend-pos-production.up.railway.app';
```

### 🌍 **CORS Support**
Backend รองรับ Frontend domains:
- `http://localhost:3000` (React dev)
- `http://localhost:5173` (Vite dev)
- `https://*.vercel.app` (Vercel deployment)
- `https://komkemkty-frontend-pos.vercel.app`
- `https://frontend-pos-jade.vercel.app`

### ⚡ **Health Check**
```
GET /health
GET /ping
```

---

## 💡 **ข้อแนะนำสำหรับทีม Frontend**

### 🔄 **Real-time Updates**
- ใช้ `setInterval` refresh สถานะออเดอร์ทุก 5-10 วินาที
- เรียก `/public/orders/table/{qrCode}` เพื่อดูสถานะล่าสุด

### 🎨 **UI/UX Recommendations**
- แสดงสถานะด้วยสีตามตารางข้างต้น
- ใช้ Loading skeleton ระหว่างรอข้อมูล
- แสดง Error message ภาษาไทยที่เข้าใจง่าย
- ใช้ Toast notification เมื่อสำเร็จ

### 🔐 **Token Management**
- เก็บ token ใน localStorage หรือ sessionStorage
- ตรวจสอบ token หมดอายุและ redirect ไป login
- ส่ง token ใน Header: `Authorization: Bearer <token>`

### 📱 **Mobile Responsive**
- QR Code scanning บนมือถือ
- Touch-friendly UI สำหรับโต๊ะ tablet
- Offline mode สำหรับ network issues

---

## 🧪 **ตัวอย่างการใช้งาน JavaScript**

### 🚀 **API Helper Function**
```javascript
const API_BASE = 'http://localhost:8080';

const api = async (endpoint, options = {}) => {
  try {
    const response = await fetch(`${API_BASE}${endpoint}`, {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    });
    
    const result = await response.json();
    
    if (!result.success) {
      throw new Error(result.message);
    }
    
    return result;
  } catch (error) {
    console.error('API Error:', error);
    throw error;
  }
};

// ใช้งาน
const menu = await api('/public/menu/table_001');
const order = await api('/public/orders/create', {
  method: 'POST',
  body: JSON.stringify({
    qr_code_identifier: 'table_001',
    items: [{ menu_item_id: 1, quantity: 2 }]
  })
});
```

### 🔐 **Authentication Helper**
```javascript
class AuthManager {
  static setToken(token) {
    localStorage.setItem('pos_token', token);
  }
  
  static getToken() {
    return localStorage.getItem('pos_token');
  }
  
  static clearToken() {
    localStorage.removeItem('pos_token');
  }
  
  static async login(username, password) {
    const result = await api('/staff/login', {
      method: 'POST',
      body: JSON.stringify({ username, password })
    });
    
    this.setToken(result.data.token);
    return result.data.staff;
  }
  
  static async getStaffInfo() {
    return await api('/staff/info', {
      headers: {
        'Authorization': `Bearer ${this.getToken()}`
      }
    });
  }
}
```

---

## 📋 **Checklist สำหรับ Frontend Developer**

### ✅ **พื้นฐาน**
- [ ] ตั้งค่า Base URL
- [ ] สร้าง API helper functions
- [ ] จัดการ Error handling
- [ ] ทำ Loading states

### ✅ **ฟีเจอร์ลูกค้า (Public)**
- [ ] หน้าดูเมนูจาก QR Code
- [ ] ระบบตะกร้าสินค้า
- [ ] หน้าสั่งอาหาร
- [ ] หน้าติดตามออเดอร์
- [ ] หน้าประวัติออเดอร์

### ✅ **ฟีเจอร์พนักงาน (Staff)**
- [ ] หน้า Login
- [ ] หน้า Dashboard (ใช้ /summary)
- [ ] หน้าจัดการออเดอร์
- [ ] ระบบอัปเดตสถานะ
- [ ] ระบบล้างประวัติโต๊ะ

### ✅ **UX/UI**
- [ ] แสดงสถานะด้วยสี
- [ ] Real-time updates
- [ ] Mobile responsive
- [ ] Error messages ภาษาไทย

---

## 🎉 **สรุป: พร้อมใช้งาน 100%!**

✅ **API ครบทั้งหมด** - ทั้ง Public และ Staff  
✅ **Response format สม่ำเสมอ** - มี success, message, data  
✅ **Error handling ครบ** - มี HTTP status และ message ภาษาไทย  
✅ **Authentication ใช้งานได้** - JWT token system  
✅ **CORS รองรับ Vercel** - Deploy production ได้เลย  
✅ **Documentation ละเอียด** - มีตัวอย่างทุก endpoint  

**🚀 ทีม Frontend สามารถเริ่มพัฒนาและ integrate ได้ทันที!**
