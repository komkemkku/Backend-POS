# 🔐 Authentication System Update

## การเปลี่ยนแปลงระบบการเข้าสู่ระบบ

### ✅ สิ่งที่เปลี่ยนแปลง

#### 1. **Login Request Format**
```json
// ก่อน (ใช้ email)
{
  "email": "admin@restaurant.com",
  "password_hash": "hashed_password"
}

// หลัง (ใช้ username + plain password)
{
  "username": "admin",
  "password": "password"
}
```

#### 2. **Password Security**
- **หน้าบ้าน**: ส่ง password แบบ plain text
- **หลังบ้าน**: รับ plain text → hash ด้วย bcrypt → เปรียบเทียบกับ database
- **ความปลอดภัย**: ใช้ bcrypt cost 14 (สูงกว่ามาตรฐาน)

#### 3. **Response Format**
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

### 🔧 การใช้งานสำหรับ Frontend

#### JavaScript Example
```javascript
const login = async (username, password) => {
  try {
    const response = await fetch('http://localhost:8080/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: username,
        password: password  // ส่ง plain text
      })
    });
    
    const result = await response.json();
    
    if (result.success) {
      // เก็บ token สำหรับใช้ใน API อื่นๆ
      localStorage.setItem('token', result.data.token);
      localStorage.setItem('staff', JSON.stringify(result.data.staff));
      
      return result.data;
    } else {
      throw new Error(result.message);
    }
  } catch (error) {
    console.error('Login error:', error);
    throw error;
  }
};

// การใช้งาน
login('admin', 'password')
  .then(data => {
    console.log('Login successful:', data);
    // redirect หรือ update UI
  })
  .catch(error => {
    alert('เข้าสู่ระบบไม่สำเร็จ: ' + error.message);
  });
```

#### React Hook Example
```jsx
import { useState } from 'react';

const useAuth = () => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(false);
  
  const login = async (username, password) => {
    setLoading(true);
    try {
      const response = await fetch('/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
      });
      
      const result = await response.json();
      
      if (result.success) {
        setUser(result.data);
        localStorage.setItem('token', result.data.token);
        return result.data;
      } else {
        throw new Error(result.message);
      }
    } finally {
      setLoading(false);
    }
  };
  
  const logout = () => {
    setUser(null);
    localStorage.removeItem('token');
    localStorage.removeItem('staff');
  };
  
  return { user, login, logout, loading };
};
```

### 👥 ข้อมูล Staff สำหรับทดสอบ

```
Username: admin
Password: password
Role: admin

Username: staff1  
Password: password
Role: staff

Username: staff2
Password: password
Role: staff
```

### 🔒 Security Features

#### 1. **Password Hashing**
- ใช้ bcrypt algorithm
- Cost factor: 14 (2^14 iterations)
- Salt อัตโนมัติ

#### 2. **Error Messages**
- ไม่เปิดเผยว่า username หรือ password ผิด
- ใช้ข้อความเดียวกัน: "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง"

#### 3. **JWT Token**
- Contains staff information
- Used for API authorization
- Include in Authorization header: `Bearer <token>`

### 🛠️ Tools

#### Password Hash Generator
```bash
# รัน tool สำหรับสร้าง password hash
go run tools/create_staff_password.go
```

#### การเพิ่ม Staff ใหม่
```sql
-- แทนที่ <username>, <password_hash>, <full_name>, <role>
INSERT INTO staff (username, password_hash, full_name, role, created_at, updated_at) 
VALUES ('<username>', '<password_hash>', '<full_name>', '<role>', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));
```

### 🚨 Important Notes

1. **ไม่เก็บ password แบบ plain text** ในฐานข้อมูล
2. **ใช้ HTTPS** ในการส่ง password จริง
3. **Validate input** ก่อนส่งไป API
4. **Handle errors** อย่างเหมาะสม
5. **เก็บ token** อย่างปลอดภัย (หลีกเลี่ยง localStorage ใน production)

### 📝 Migration Guide

หากมีระบบเก่าที่ใช้ email:

1. **Update frontend forms** ให้ใช้ username field
2. **Update API calls** ให้ส่ง username แทน email  
3. **Update validation** ให้ตรงกับ format ใหม่
4. **Test authentication flow** ทั้งหมด

---

**✅ ระบบ Authentication พร้อมใช้งานแล้ว! ปลอดภัยและใช้งานง่าย**
