# 🍽️ POS Backend API

> **Backend API สำหรับระบบ Point of Sale (POS) ร้านอาหาร**

![Go](https://img.shields.io/badge/Go-1.23.4-00ADD8?style=flat-square&logo=go)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat-square)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?style=flat-square&logo=postgresql)
![JWT](https://img.shields.io/badge/JWT-Authentication-000000?style=flat-square&logo=jsonwebtokens)

## 🔗 Related Repository
- **Frontend**: [Frontend-POS](https://github.com/komkemkku/Frontend-POS.git)
- **Backend**: [Backend-POS](https://github.com/komkemkku/Backend-POS.git) *(this repository)*

---

## 🛠️ เทคโนโลยี

| Technology | Version | Purpose |
|------------|---------|---------|
| **Go** | 1.23.4 | Backend Language |
| **Gin Framework** | Latest | HTTP Web Framework |
| **PostgreSQL** | Latest | Primary Database |
| **Bun ORM** | Latest | Database ORM |
| **JWT** | Latest | Authentication |

---

## 🚀 การติดตั้งและใช้งาน

### 1. Clone โปรเจกต์

```bash
# Clone Backend
git clone https://github.com/komkemkku/Backend-POS.git
cd Backend-POS

# Clone Frontend (optional)
git clone https://github.com/komkemkku/Frontend-POS.git
```

### 2. ติดตั้ง Dependencies

```bash
# ติดตั้ง Go modules
go mod download
go mod tidy
```

### 3. ตั้งค่า Environment Variables

สร้างไฟล์ `.env` และกำหนดค่าตัวแปร:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=pos_database

# JWT Configuration
TOKEN_SECRET=your_jwt_secret_key
TOKEN_DURATION=24h

# Server Configuration
PORT=8080
GIN_MODE=debug
```

### 4. รัน Database Migration

```bash
go run main.go migrate
```

### 5. เริ่มเซิร์ฟเวอร์

```bash
# Development mode
go run main.go

# Production build
go build -o pos-backend .
./pos-backend
```

เซิร์ฟเวอร์จะทำงานที่ `http://localhost:8080`

---

## 📚 API Documentation

### 🌐 Public Endpoints (สำหรับลูกค้า)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/public/menu/:qrCodeIdentifier` | ดูเมนูจาก QR Code โต๊ะ |
| `GET` | `/public/menu` | ดูเมนูทั้งหมด |
| `POST` | `/public/orders/create` | สั่งอาหาร |
| `GET` | `/public/orders/table/:qrCodeIdentifier` | ดูออเดอร์ปัจจุบันของโต๊ะ |
| `GET` | `/public/orders/:orderID/table/:qrCodeIdentifier` | ดูสถานะออเดอร์เฉพาะ |
| `GET` | `/public/orders/history/:qrCodeIdentifier` | ดูประวัติออเดอร์ทั้งหมด |
| `GET` | `/public/table/summary/:qrCodeIdentifier` | ดูสรุปโต๊ะ |

### 🔐 Staff Endpoints (ต้อง Authentication)

#### Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/staff/login` | เข้าสู่ระบบพนักงาน |

#### Dashboard
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/summary` | ดูสรุปข้อมูลแดชบอร์ด |

#### Staff Management
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/staff/info` | ดูข้อมูลพนักงานปัจจุบัน |
| `GET` | `/staff` | ดูรายการพนักงานทั้งหมด |
| `GET` | `/staff/:id` | ดูข้อมูลพนักงานตาม ID |
| `POST` | `/staff/create` | สร้างพนักงานใหม่ |
| `PATCH` | `/staff/:id` | แก้ไขข้อมูลพนักงาน |
| `DELETE` | `/staff/:id` | ลบพนักงาน |

#### Order Management
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/orders` | ดูออเดอร์ทั้งหมด |
| `GET` | `/orders/:id` | ดูออเดอร์ตาม ID |
| `POST` | `/orders/create` | สร้างออเดอร์ใหม่ |
| `PATCH` | `/orders/:id` | แก้ไขออเดอร์ |
| `DELETE` | `/orders/:id` | ลบออเดอร์ |
| `PATCH` | `/staff/orders/:orderID/status` | เปลี่ยนสถานะออเดอร์ |
| `POST` | `/staff/orders/clear-table/:qrCodeIdentifier` | ล้างประวัติโต๊ะ |
| `POST` | `/staff/orders/cancel/:orderID/table/:qrCodeIdentifier` | ยกเลิกออเดอร์ |

#### Menu Management
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/menu-items` | ดูเมนูทั้งหมด |
| `GET` | `/menu-items/:id` | ดูเมนูตาม ID |
| `POST` | `/menu-items/create` | สร้างเมนูใหม่ |
| `PATCH` | `/menu-items/:id` | แก้ไขเมนู |
| `DELETE` | `/menu-items/:id` | ลบเมนู |

#### Other Management
- **Categories**: `/categories/*` - จัดการหมวดหมู่
- **Tables**: `/tables/*` - จัดการโต๊ะ
- **Payments**: `/payments/*` - จัดการการชำระเงิน
- **Expenses**: `/expenses/*` - จัดการค่าใช้จ่าย
- **Reservations**: `/reservations/*` - จัดการการจอง

### 🔧 System Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/ping` | Simple ping test |

---

## 📁 โครงสร้างโปรเจกต์

```
Backend-POS/
├── 📁 cmd/              # Commands & CLI tools
│   ├── cmd.go
│   ├── migrateCmd.go
│   └── ...
├── 📁 configs/          # Database configuration
│   └── database.go
├── 📁 controller/       # API controllers
│   ├── auth/           # Authentication
│   ├── categories/     # Category management
│   ├── dashboard/      # Dashboard data
│   ├── expense/        # Expense management
│   ├── menu_item/      # Menu management
│   ├── order/          # Order management
│   ├── payment/        # Payment processing
│   ├── reservation/    # Table reservations
│   ├── staff/          # Staff management
│   └── table/          # Table management
├── 📁 middlewares/      # HTTP middlewares
│   └── auth.middleware.go
├── 📁 model/           # Database models
│   ├── base.model.go
│   ├── categories.model.go
│   ├── expenses.model.go
│   └── ...
├── 📁 requests/        # Request DTOs
├── 📁 responses/       # Response DTOs
├── 📁 utils/           # Utility functions
│   ├── hasing.go
│   └── jwt/
└── 📄 main.go          # Application entry point
```

---

## 🌐 CORS Configuration

API รองรับ CORS สำหรับ Frontend domains:

- `http://localhost:3000` *(Next.js development)*
- `http://localhost:3001` *(Alternative port)*
- `http://localhost:3002` *(Alternative port)*
- `http://localhost:5173` *(Vite development)*
- `http://localhost:8080` *(Backend port)*
- `https://*.vercel.app` *(Production deployment)*
- `https://komkemkty-frontend-pos.vercel.app`
- `https://frontend-pos-jade.vercel.app`

---

## � Authentication

API ใช้ **JWT (JSON Web Token)** สำหรับการยืนยันตัวตน:

1. **Login**: `POST /staff/login` พร้อม username และ password
2. **Receive Token**: ระบบจะส่ง JWT token กลับมา
3. **Use Token**: ส่ง token ใน Header: `Authorization: Bearer <token>`

**Token Format:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "staff": {
    "id": 1,
    "username": "admin",
    "full_name": "Administrator",
    "role": "admin"
  }
}
```

---

## �📱 Development Workflow

### สำหรับการพัฒนา Full-Stack:

1. **เริ่ม Backend**:
   ```bash
   cd Backend-POS
   go run main.go
   ```
   *(Backend จะทำงานที่ port 8080)*

2. **เริ่ม Frontend** (ใน terminal อื่น):
   ```bash
   cd Frontend-POS
   npm run dev
   ```
   *(Frontend จะทำงานที่ port 3000)*

3. **Test API**: ใช้ Postman หรือ curl ทดสอบ endpoints
4. **Database**: ตรวจสอบข้อมูลใน PostgreSQL

---

## 🧪 Testing

### สร้างข้อมูลทดสอบ:

```bash
# สร้างพนักงานแรก
go run tools/create_staff_password.go

# รัน migration ใหม่
go run main.go migrate
```

### ทดสอบ API:

```bash
# Health check
curl http://localhost:8080/health

# Ping test
curl http://localhost:8080/ping

# Login test
curl -X POST http://localhost:8080/staff/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

---

## 📈 Features

### ✅ **Completed Features**

- 🔐 **Authentication System** - JWT-based login
- 👥 **Staff Management** - CRUD operations
- 🍽️ **Menu Management** - Food items and categories
- 📋 **Order System** - Create, track, and manage orders
- 🏷️ **QR Code Integration** - Table-based ordering
- 💰 **Payment Processing** - Multiple payment methods
- 📊 **Dashboard Analytics** - Sales and order statistics
- 🗄️ **Database Migration** - Automated schema updates
- 🌐 **CORS Support** - Cross-origin resource sharing

### 🚧 **Future Enhancements**

- 📱 **Mobile App API** - Extended mobile support
- 📈 **Advanced Analytics** - Detailed reporting
- 🔔 **Real-time Notifications** - WebSocket integration
- 💳 **Payment Gateway** - External payment integration
- 📦 **Inventory Management** - Stock tracking

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit changes: `git commit -m 'Add amazing feature'`
4. Push to branch: `git push origin feature/amazing-feature`
5. Submit a Pull Request

---

## 📄 License

This project is licensed under the MIT License.

---

## 📞 Contact

- **GitHub**: [komkemkku](https://github.com/komkemkku)
- **Repository**: [Backend-POS](https://github.com/komkemkku/Backend-POS)

---

<div align="center">

**Made with ❤️ for Restaurant Management**

</div>
