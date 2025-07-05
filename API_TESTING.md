# API Testing Examples

## ทดสอบ Public API ด้วย curl

### 1. Health Check
```bash
curl -X GET http://localhost:8080/health
curl -X GET http://localhost:8080/ping
```

### 2. ดูเมนูตาม QR Code (ต้องมีข้อมูลโต๊ะในฐานข้อมูลก่อน)
```bash
curl -X GET http://localhost:8080/public/menu/table_001
```

### 3. ดูเมนูทั้งหมด
```bash
curl -X GET http://localhost:8080/public/menu
```

### 4. สร้างออเดอร์
```bash
curl -X POST http://localhost:8080/public/orders/create \
  -H "Content-Type: application/json" \
  -d '{
    "qr_code_identifier": "table_001",
    "items": [
      {
        "menu_item_id": 1,
        "quantity": 2
      }
    ]
  }'
```

### 5. ดูประวัติออเดอร์ของโต๊ะ
```bash
curl -X GET http://localhost:8080/public/orders/table/table_001
```

### 6. ดูสถานะออเดอร์เฉพาะ
```bash
curl -X GET http://localhost:8080/public/orders/1/table/table_001
```

### 7. ดูประวัติออเดอร์ทั้งหมด
```bash
curl -X GET http://localhost:8080/public/orders/history/table_001
```

### 8. ดูสรุปโต๊ะ
```bash
curl -X GET http://localhost:8080/public/table/summary/table_001
```

## ทดสอบ Staff API (ต้อง Login ก่อน)

### 1. Login
```bash
curl -X POST http://localhost:8080/staff/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password"
  }'
```

**Response ที่ได้:**
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

**คัดลอก token ไปใช้ใน API อื่นๆ**

### 2. ล้างประวัติโต๊ะ (ต้องใส่ token ที่ได้จาก login)
```bash
curl -X POST http://localhost:8080/staff/orders/clear-table/table_001 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 3. ล้างประวัติแบบละเอียด
```bash
# ชำระเงินแล้ว (default)
curl -X POST http://localhost:8080/staff/orders/advanced-clear/table_001?type=payment \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# ยกเลิกทั้งหมด
curl -X POST http://localhost:8080/staff/orders/advanced-clear/table_001?type=cancel_all \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# เสร็จสิ้นทั้งหมด
curl -X POST http://localhost:8080/staff/orders/advanced-clear/table_001?type=complete_all \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 4. ยกเลิกออเดอร์เฉพาะ
```bash
curl -X POST http://localhost:8080/staff/orders/cancel/1/table/table_001?reason=ลูกค้าไม่ต้องการ \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 5. อัปเดตสถานะออเดอร์
```bash
curl -X PATCH http://localhost:8080/staff/orders/1/status \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "preparing"
  }'
```

### 6. Dashboard Summary
```bash
# ต้องใส่ token ที่ได้จาก login
curl -X GET http://localhost:8080/summary \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

**Response ที่ได้:**
```json
{
  "status": "success",
  "data": {
    "total_tables": 5,
    "today_revenue": 2350.75,
    "today_orders": 15,
    "pending_orders": 3
  }
}
```

## ข้อมูลตัวอย่างที่ต้องมีในฐานข้อมูล

### Tables
```sql
INSERT INTO tables (table_number, qr_code_identifier, status, created_at, updated_at) 
VALUES (1, 'table_001', 'available', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));
```

### Categories
```sql
INSERT INTO categories (name, description, created_at, updated_at) 
VALUES ('อาหารจานหลัก', 'อาหารจานหลักต่างๆ', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));
```

### Menu Items
```sql
INSERT INTO menu_items (name, description, price, category_id, image_url, is_available, created_at, updated_at) 
VALUES ('ผัดไทย', 'ผัดไทยแสนอร่อย', 60.00, 1, '', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));
```

### Staff (สำหรับ login)
```sql
INSERT INTO staff (first_name, last_name, email, password, phone, position, created_at, updated_at) 
VALUES ('Admin', 'User', 'admin@example.com', 'hashed_password', '0812345678', 'admin', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));
```

## Frontend Integration

### JavaScript Example
```javascript
const API_BASE = 'http://localhost:8080';

// ดูเมนูตาม QR Code
async function getMenuByQR(qrCode) {
  const response = await fetch(`${API_BASE}/public/menu/${qrCode}`);
  return await response.json();
}

// สร้างออเดอร์
async function createOrder(qrCode, items) {
  const response = await fetch(`${API_BASE}/public/orders/create`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      qr_code_identifier: qrCode,
      items: items
    })
  });
  return await response.json();
}

// ดูสถานะออเดอร์
async function getOrderStatus(orderId, qrCode) {
  const response = await fetch(`${API_BASE}/public/orders/${orderId}/table/${qrCode}`);
  return await response.json();
}

// ดูประวัติออเดอร์
async function getOrderHistory(qrCode) {
  const response = await fetch(`${API_BASE}/public/orders/table/${qrCode}`);
  return await response.json();
}
```

### React Example
```jsx
import { useState, useEffect } from 'react';

function MenuComponent({ qrCode }) {
  const [menu, setMenu] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(`http://localhost:8080/public/menu/${qrCode}`)
      .then(res => res.json())
      .then(data => {
        setMenu(data.data);
        setLoading(false);
      })
      .catch(err => {
        console.error('Error:', err);
        setLoading(false);
      });
  }, [qrCode]);

  if (loading) return <div>กำลังโหลด...</div>;
  if (!menu) return <div>ไม่พบข้อมูลเมนู</div>;

  return (
    <div>
      <h2>โต๊ะ {menu.table_info.table_number}</h2>
      <div className="menu-items">
        {menu.menu_items.map(item => (
          <div key={item.id} className="menu-item">
            <h3>{item.name}</h3>
            <p>{item.description}</p>
            <p>ราคา: {item.price} บาท</p>
          </div>
        ))}
      </div>
    </div>
  );
}
```

## สำคัญ!

1. **ต้องมีฐานข้อมูล PostgreSQL พร้อมข้อมูลตัวอย่าง**
2. **ต้องรัน migration ก่อน**: `./pos-server migrate`
3. **ตรวจสอบไฟล์ .env ให้ถูกต้อง**
4. **CORS ตั้งค่าไว้รองรับ localhost และ Vercel แล้ว**

## คำแนะนำ Debug

```bash
# ดู log ของเซิร์ฟเวอร์
go run main.go

# ดูข้อมูลฐานข้อมูล
psql -h localhost -U postgres -d pos_db -c "SELECT * FROM tables;"

# ทดสอบ API แบบง่าย
curl -X GET http://localhost:8080/ping
```
