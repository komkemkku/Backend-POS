# 📋 Backend POS API Documentation - สำหรับ Frontend Integration

## � Base URL
```
Development: http://localhost:8080
Production: https://your-api-domain.com
```

## � Authentication
- **Public API**: ไม่ต้อง Authentication (สำหรับลูกค้า)
- **Staff API**: ต้องใส่ `Authorization: Bearer <token>` ใน Header

---

## 📱 PUBLIC API (สำหรับลูกค้า)

### 1️⃣ ดูเมนูตาม QR Code โต๊ะ
```http
GET /public/menu/{qrCodeIdentifier}
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/public/menu/table_001')
```

**Response:**
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

### 2️⃣ ดูเมนูทั้งหมด
```http
GET /public/menu-items
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/public/menu-items')
```

**Response:**
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

### 3️⃣ สร้างออเดอร์ (ลูกค้าสั่งอาหาร)
```http
POST /public/orders/create
Content-Type: application/json
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
      "menu_item_id": 3,
      "quantity": 1
    }
  ]
}
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/public/orders/create', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    qr_code_identifier: 'table_001',
    items: [
      { menu_item_id: 1, quantity: 2 },
      { menu_item_id: 3, quantity: 1 }
    ]
  })
})
```

**Response:**
```json
{
  "success": true,
  "message": "สร้างออเดอร์สำเร็จ",
  "data": {
    "id": 123,
    "table_id": 1,
    "table_number": 1,
    "status": "pending",
    "total_amount": 270.00,
    "items": [
      {
        "id": 1,
        "menu_item_id": 1,
        "quantity": 2,
        "price_per_item": 60.00,
        "sub_total": 120.00,
        "notes": ""
      },
      {
        "id": 2,
        "menu_item_id": 3,
        "quantity": 1,
        "price_per_item": 150.00,
        "sub_total": 150.00,
        "notes": ""
      }
    ],
    "created_at": 1704067200,
    "updated_at": 1704067200,
    "message": "ออเดอร์ของคุณได้รับการยืนยันแล้ว กรุณารอสักครู่"
  }
}
```

### 4️⃣ ดูประวัติออเดอร์ปัจจุบัน (ยังไม่ชำระเงิน)
```http
GET /public/orders/table/{qrCodeIdentifier}
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/public/orders/table/table_001')
```

**Response:**
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
      "total_amount": 270.00,
      "items": [
        {
          "id": 1,
          "order_id": 123,
          "menu_item_id": 1,
          "quantity": 2,
          "price_per_item": 60.00,
          "sub_total": 120.00,
          "notes": "",
          "created_at": 1704067200,
          "updated_at": 1704067200
        }
      ],
      "created_at": 1704067200,
      "updated_at": 1704067300
    }
  ]
}
```

### 5️⃣ ดูสถานะออเดอร์เฉพาะ
```http
GET /public/orders/{orderID}/table/{qrCodeIdentifier}
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/public/orders/123/table/table_001')
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 123,
    "table_id": 1,
    "table_number": 1,
    "status": "preparing",
    "status_text": "กำลังเตรียม",
    "status_color": "blue",
    "estimated_time": "10-15 นาที",
    "total_amount": 270.00,
    "items": [
      {
        "id": 1,
        "order_id": 123,
        "menu_item_id": 1,
        "quantity": 2,
        "price_per_item": 60.00,
        "sub_total": 120.00,
        "notes": "",
        "created_at": 1704067200,
        "updated_at": 1704067200
      }
    ],
    "created_at": 1704067200,
    "updated_at": 1704067300
  }
}
```

### 6️⃣ ดูประวัติทั้งหมด (แยกปัจจุบัน/ชำระแล้ว)
```http
GET /public/orders/history/{qrCodeIdentifier}
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/public/orders/history/table_001')
```

**Response:**
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
        "table_id": 1,
        "table_number": 1,
        "status": "preparing",
        "status_text": "กำลังเตรียม",
        "total_amount": 270.00,
        "items": [...],
        "created_at": 1704067200,
        "updated_at": 1704067300
      }
    ],
    "paid_orders": [
      {
        "id": 120,
        "table_id": 1,
        "table_number": 1,
        "status": "paid",
        "status_text": "ชำระเงินแล้ว",
        "total_amount": 180.00,
        "items": [...],
        "created_at": 1704060000,
        "updated_at": 1704063600
      }
    ],
    "summary": {
      "total_orders": 5,
      "total_spent": 850.00,
      "current_pending": 2,
      "completed_today": 3
    },
    "timestamp": 1704067500
  }
}
```

### 7️⃣ ดูสรุปโต๊ะ
```http
GET /public/table/summary/{qrCodeIdentifier}
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/public/table/summary/table_001')
```

**Response:**
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

## 👨‍💼 STAFF API (ต้อง Authentication)

## 👨‍💼 STAFF API (ต้อง Authentication)

### 🔑 Authentication - เข้าสู่ระบบ
```http
POST /staff/login
Content-Type: application/json
```

**Request Body:**
```json
{
  "username": "admin",
  "password": "password"
}
```

**ตัวอย่าง Request:**
```javascript
const authResponse = await fetch('http://localhost:8080/staff/login', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    username: 'admin',
    password: 'password'
  })
});
const authData = await authResponse.json();
```

**Response:**
```json
{
  "success": true,
  "message": "เข้าสู่ระบบสำเร็จ",
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

### 1️⃣ ล้างประวัติโต๊ะหลังชำระเงิน
```http
POST /staff/orders/clear-table/{qrCodeIdentifier}
Authorization: Bearer {token}
```

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/staff/orders/clear-table/table_001', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
```

**Response:**
```json
{
  "success": true,
  "message": "ล้างประวัติโต๊ะเรียบร้อยแล้ว"
}
```

### 2️⃣ ล้างประวัติแบบละเอียด
```http
POST /staff/orders/advanced-clear/{qrCodeIdentifier}?type={clearType}
Authorization: Bearer {token}
```

**Query Parameters:**
- `type`: ประเภทการล้าง
  - `payment` (default) - ชำระเงินแล้ว
  - `cancel_all` - ยกเลิกทั้งหมด  
  - `complete_all` - เสร็จสิ้นทั้งหมด

**ตัวอย่าง Request:**
```javascript
// ชำระเงินแล้ว
fetch('http://localhost:8080/staff/orders/advanced-clear/table_001?type=payment', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`
  }
})

// ยกเลิกทั้งหมด
fetch('http://localhost:8080/staff/orders/advanced-clear/table_001?type=cancel_all', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
```

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
    "cleared_orders": [
      {
        "id": 123,
        "table_id": 1,
        "staff_id": 0,
        "status": "paid",
        "total_amount": 270.00,
        "created_at": 1704067200,
        "completed_at": 1704067500
      }
    ],
    "timestamp": 1704067500,
    "message": "ชำระเงินเรียบร้อย ได้ล้างประวัติ 3 ออเดอร์"
  }
}
```

### 3️⃣ ยกเลิกออเดอร์เฉพาะ
```http
POST /staff/orders/cancel/{orderID}/table/{qrCodeIdentifier}?reason={reason}
Authorization: Bearer {token}
```

**Query Parameters:**
- `reason` (optional): เหตุผลการยกเลิก

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/staff/orders/cancel/123/table/table_001?reason=ลูกค้าไม่ต้องการ', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
```

**Response:**
```json
{
  "success": true,
  "message": "ยกเลิกออเดอร์เรียบร้อยแล้ว"
}
```

### 4️⃣ อัปเดตสถานะออเดอร์
```http
PATCH /staff/orders/{orderID}/status
Authorization: Bearer {token}
Content-Type: application/json
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

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/staff/orders/123/status', {
  method: 'PATCH',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    status: 'preparing'
  })
})
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 123,
    "table_id": 1,
    "staff_id": 1,
    "status": "preparing",
    "total_amount": 270.00,
    "created_at": 1704067200,
    "completed_at": 0
  }
}
```

### 5️⃣ ดูรายการออเดอร์ทั้งหมด (Staff)
```http
GET /orders?page=1&size=10&search=
Authorization: Bearer {token}
```

**Query Parameters:**
- `page` (optional): หน้าที่ต้องการ (default: 1)
- `size` (optional): จำนวนรายการต่อหน้า (default: 10)
- `search` (optional): คำค้นหา

**ตัวอย่าง Request:**
```javascript
fetch('http://localhost:8080/orders?page=1&size=10', {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 123,
      "table_id": 1,
      "staff_id": 1,
      "status": "preparing",
      "total_amount": 270.00,
      "created_at": 1704067200,
      "completed_at": 0
    }
  ],
  "total": 25,
  "page": 1,
  "size": 10
}
```

---

## 🎨 ระบบสถานะและสี

### สถานะออเดอร์
| Status | ภาษาไทย | สี | เวลาประมาณการ | สำหรับ Frontend |
|--------|---------|----|----|----------------|
| `pending` | รอดำเนินการ | `#FFA500` (orange) | 5-10 นาที | 🟠 |
| `preparing` | กำลังเตรียม | `#0066CC` (blue) | 10-15 นาที | 🔵 |
| `ready` | พร้อมเสิร์ฟ | `#00CC00` (green) | พร้อมแล้ว | 🟢 |
| `served` | เสิร์ฟแล้ว | `#9900CC` (purple) | เสร็จสิ้น | 🟣 |
| `paid` | ชำระเงินแล้ว | `#999999` (gray) | เสร็จสิ้น | ⚪ |
| `completed` | เสร็จสิ้น | `#999999` (gray) | เสร็จสิ้น | ⚪ |
| `cancelled` | ยกเลิก | `#CC0000` (red) | ยกเลิกแล้ว | 🔴 |

### สถานะโต๊ะ
| Status | ภาษาไทย | สำหรับ Frontend |
|--------|---------|----------------|
| `available` | ว่าง | สีเขียว |
| `occupied` | มีลูกค้า | สีส้ม |
| `reserved` | จอง | สีน้ำเงิน |
| `maintenance` | ปิดปรับปรุง | สีแดง |

---

## 🔧 การจัดการ Error

### Error Response Format
```json
{
  "success": false,
  "message": "ข้อความ error ภาษาไทย",
  "error": "technical_error_code"
}
```

### HTTP Status Codes
- `200` - สำเร็จ
- `400` - Bad Request (ข้อมูลไม่ถูกต้อง)
- `401` - Unauthorized (ไม่มีสิทธิ์)
- `404` - Not Found (ไม่พบข้อมูล)
- `500` - Internal Server Error (ข้อผิดพลาดของเซิร์ฟเวอร์)

### ตัวอย่าง Error Responses
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

// ไม่มีสิทธิ์
{
  "success": false,
  "message": "ไม่พบข้อมูล staff"
}
```

---

## 💻 ตัวอย่างการใช้งานใน Frontend

### React/Next.js Example
```jsx
import { useState, useEffect } from 'react';

// Hook สำหรับจัดการ API
const useAPI = () => {
  const baseURL = 'http://localhost:8080';
  
  const api = async (endpoint, options = {}) => {
    const response = await fetch(`${baseURL}${endpoint}`, {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    });
    return response.json();
  };
  
  return { api };
};

// Component สำหรับดูเมนู
const MenuPage = ({ qrCode }) => {
  const [menu, setMenu] = useState(null);
  const [cart, setCart] = useState([]);
  const { api } = useAPI();
  
  useEffect(() => {
    loadMenu();
  }, [qrCode]);
  
  const loadMenu = async () => {
    const result = await api(`/public/menu/${qrCode}`);
    if (result.success) {
      setMenu(result.data);
    }
  };
  
  const addToCart = (item, quantity) => {
    setCart(prev => [...prev, { menu_item_id: item.id, quantity }]);
  };
  
  const createOrder = async () => {
    const result = await api('/public/orders/create', {
      method: 'POST',
      body: JSON.stringify({
        qr_code_identifier: qrCode,
        items: cart
      })
    });
    
    if (result.success) {
      alert(result.data.message);
      setCart([]);
      // Redirect to order tracking
    }
  };
  
  return (
    <div>
      <h1>โต๊ะ {menu?.table_info?.table_number}</h1>
      {menu?.menu_items?.map(item => (
        <div key={item.id}>
          <h3>{item.name} - ฿{item.price}</h3>
          <p>{item.description}</p>
          <button onClick={() => addToCart(item, 1)}>
            เพิ่มลงตะกร้า
          </button>
        </div>
      ))}
      
      {cart.length > 0 && (
        <button onClick={createOrder}>สั่งอาหาร</button>
      )}
    </div>
  );
};

// Component สำหรับติดตามออเดอร์
const OrderTracking = ({ qrCode, orderId }) => {
  const [order, setOrder] = useState(null);
  const { api } = useAPI();
  
  useEffect(() => {
    const interval = setInterval(loadOrderStatus, 5000); // Refresh ทุก 5 วินาที
    return () => clearInterval(interval);
  }, [orderId]);
  
  const loadOrderStatus = async () => {
    const result = await api(`/public/orders/${orderId}/table/${qrCode}`);
    if (result.success) {
      setOrder(result.data);
    }
  };
  
  const getStatusColor = (status) => {
    const colors = {
      pending: '#FFA500',
      preparing: '#0066CC',
      ready: '#00CC00',
      served: '#9900CC',
      paid: '#999999',
      completed: '#999999',
      cancelled: '#CC0000'
    };
    return colors[status] || '#999999';
  };
  
  return (
    <div>
      <h1>ติดตามออเดอร์ #{order?.id}</h1>
      <div style={{ color: getStatusColor(order?.status) }}>
        <h2>{order?.status_text}</h2>
        <p>เวลาประมาณการ: {order?.estimated_time}</p>
      </div>
      
      <h3>รายการอาหาร:</h3>
      {order?.items?.map(item => (
        <div key={item.id}>
          <p>{item.quantity}x - ฿{item.sub_total}</p>
        </div>
      ))}
      
      <h3>ยอดรวม: ฿{order?.total_amount}</h3>
    </div>
  );
};
```

### Vue.js Example
```vue
<template>
  <div>
    <h1>โต๊ะ {{ tableInfo?.table_number }}</h1>
    
    <!-- Menu List -->
    <div v-for="item in menuItems" :key="item.id">
      <h3>{{ item.name }} - ฿{{ item.price }}</h3>
      <p>{{ item.description }}</p>
      <button @click="addToCart(item)">เพิ่มลงตะกร้า</button>
    </div>
    
    <!-- Cart -->
    <div v-if="cart.length > 0">
      <h3>ตะกร้าสินค้า</h3>
      <button @click="createOrder">สั่งอาหาร</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableInfo: null,
      menuItems: [],
      cart: []
    };
  },
  
  async mounted() {
    await this.loadMenu();
  },
  
  methods: {
    async loadMenu() {
      const response = await fetch(`/public/menu/${this.$route.params.qrCode}`);
      const result = await response.json();
      
      if (result.success) {
        this.tableInfo = result.data.table_info;
        this.menuItems = result.data.menu_items;
      }
    },
    
    addToCart(item) {
      this.cart.push({
        menu_item_id: item.id,
        quantity: 1
      });
    },
    
    async createOrder() {
      const response = await fetch('/public/orders/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          qr_code_identifier: this.$route.params.qrCode,
          items: this.cart
        })
      });
      
      const result = await response.json();
      
      if (result.success) {
        alert(result.data.message);
        this.$router.push(`/order/${result.data.id}`);
      }
    }
  }
};
</script>
```

---

## 🚀 การ Deploy และ Environment

### Environment Variables
```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=pos_db
DB_USER=postgres
DB_PASSWORD=your_password

# Server
PORT=8080
GIN_MODE=production

# JWT
JWT_SECRET=your_jwt_secret_key_here
```

### CORS Configuration
Backend รองรับ domains ต่อไปนี้:
- `http://localhost:3000` (React dev)
- `http://localhost:5173` (Vite dev)
- `https://*.vercel.app` (Vercel deployment)
- `https://your-frontend-domain.com`

### Health Check
```http
GET /health
GET /ping
```

---

## 📋 Checklist สำหรับ Frontend Developer

### ✅ การเตรียมการ
- [ ] ตั้งค่า Base URL ของ API
- [ ] เตรียม function สำหรับเรียก API
- [ ] ตั้งค่า Error handling
- [ ] เตรียม Loading states

### ✅ ฟีเจอร์หลัก
- [ ] หน้าดูเมนูจาก QR Code
- [ ] ระบบตะกร้าสินค้า
- [ ] หน้าสั่งอาหาร
- [ ] หน้าติดตามออเดอร์ (real-time)
- [ ] หน้าประวัติออเดอร์

### ✅ ฟีเจอร์ Staff (ถ้ามี)
- [ ] หน้า Login
- [ ] หน้าจัดการออเดอร์
- [ ] ระบบอัปเดตสถานะ
- [ ] ระบบล้างประวัติโต๊ะ

### ✅ UX/UI
- [ ] แสดงสถานะด้วยสีและ icon
- [ ] Loading spinner/skeleton
- [ ] Error messages ภาษาไทย
- [ ] Responsive design
- [ ] Offline handling (optional)

---

**🎉 เอกสารนี้ครอบคลุมทุก API ที่จำเป็นสำหรับพัฒนา Frontend ที่ทำงานร่วมกับ Backend POS System ได้อย่างสมบูรณ์!**
