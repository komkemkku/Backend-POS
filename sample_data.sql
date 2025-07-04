-- Sample Data สำหรับทดสอบระบบ Backend POS
-- รันคำสั่งนี้หลังจาก migration เสร็จแล้ว

-- Tables
INSERT INTO tables (table_number, qr_code_identifier, status, created_at, updated_at) VALUES
(1, 'table_001', 'available', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
(2, 'table_002', 'available', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
(3, 'table_003', 'occupied', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
(4, 'table_004', 'available', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
(5, 'table_005', 'available', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));

-- Categories
INSERT INTO categories (name, description, created_at, updated_at) VALUES
('อาหารจานหลัก', 'อาหารจานหลักต่างๆ', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('เครื่องดื่ม', 'เครื่องดื่มร้อนและเย็น', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ของหวาน', 'ขนมหวานและของหวาน', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('อาหารเรียกน้ำย่อย', 'อาหารว่างและเครื่องเคียง', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));

-- Menu Items
INSERT INTO menu_items (name, description, price, category_id, image_url, is_available, created_at, updated_at) VALUES
('ผัดไทย', 'ผัดไทยแสนอร่อยเส้นหมี่ใหญ่', 60.00, 1, 'https://example.com/padthai.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ข้าวผัดปู', 'ข้าวผัดปูสไตล์ไทยแท้', 120.00, 1, 'https://example.com/crabrice.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ต้มยำกุ้ง', 'ต้มยำกุ้งรสจัดจ้าน', 150.00, 1, 'https://example.com/tomyum.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ส้มตำไทย', 'ส้มตำไทยใส่ปลาร้า', 50.00, 4, 'https://example.com/somtam.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ชาไทย', 'ชาไทยเย็นหวานมัน', 30.00, 2, 'https://example.com/thaitea.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('กาแฟเย็น', 'กาแฟเย็นสูตรเฮ้าส์', 35.00, 2, 'https://example.com/icecoffee.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('น้ำส้มคั้นฝรั่ง', 'น้ำส้มคั้นสดใหม่', 40.00, 2, 'https://example.com/orange.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ข้าวเหนียวมะม่วง', 'ข้าวเหนียวมะม่วงหวานเหนียวนุ่ม', 80.00, 3, 'https://example.com/mango.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ทับทิมกรอบ', 'ทับทิมกรอบเย็นชื่นใจ', 45.00, 3, 'https://example.com/tubtim.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('แกงเขียวหวานไก่', 'แกงเขียวหวานไก่รสชาติเข้มข้น', 90.00, 1, 'https://example.com/greencurry.jpg', true, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));

-- Staff (password ต้อง hash ด้วย bcrypt จริงๆ)
INSERT INTO staff (first_name, last_name, email, password, phone, position, created_at, updated_at) VALUES
('Admin', 'System', 'admin@restaurant.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', '0812345678', 'admin', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('John', 'Doe', 'staff1@restaurant.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', '0823456789', 'staff', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('Jane', 'Smith', 'staff2@restaurant.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', '0834567890', 'staff', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));

-- Sample Orders (สำหรับโต๊ะ 3 ที่ occupied)
INSERT INTO orders (table_id, staff_id, status, total_amount, created_at, updated_at) VALUES
(3, 0, 'pending', 210.00, EXTRACT(EPOCH FROM NOW()) - 600, EXTRACT(EPOCH FROM NOW()) - 600),
(3, 0, 'preparing', 120.00, EXTRACT(EPOCH FROM NOW()) - 300, EXTRACT(EPOCH FROM NOW()) - 150),
(3, 0, 'paid', 180.00, EXTRACT(EPOCH FROM NOW()) - 7200, EXTRACT(EPOCH FROM NOW()) - 7200);

-- Sample Order Items
INSERT INTO order_items (order_id, menu_item_id, quantity, price_per_item, notes, created_at, updated_at) VALUES
-- Order 1 (pending)
(1, 1, 2, 60.00, '', EXTRACT(EPOCH FROM NOW()) - 600, EXTRACT(EPOCH FROM NOW()) - 600),
(1, 3, 1, 150.00, 'ไม่เอาผักชี', EXTRACT(EPOCH FROM NOW()) - 600, EXTRACT(EPOCH FROM NOW()) - 600),

-- Order 2 (preparing)  
(2, 2, 1, 120.00, '', EXTRACT(EPOCH FROM NOW()) - 300, EXTRACT(EPOCH FROM NOW()) - 150),

-- Order 3 (paid - ประวัติเก่า)
(3, 1, 1, 60.00, '', EXTRACT(EPOCH FROM NOW()) - 7200, EXTRACT(EPOCH FROM NOW()) - 7200),
(3, 2, 1, 120.00, '', EXTRACT(EPOCH FROM NOW()) - 7200, EXTRACT(EPOCH FROM NOW()) - 7200);

-- Update table 3 status to occupied
UPDATE tables SET status = 'occupied' WHERE id = 3;

-- Sample payments
INSERT INTO payments (order_id, amount, payment_method, status, created_at, updated_at) VALUES
(3, 180.00, 'cash', 'completed', EXTRACT(EPOCH FROM NOW()) - 7200, EXTRACT(EPOCH FROM NOW()) - 7200);

-- Sample expenses
INSERT INTO expenses (name, description, amount, category, created_at, updated_at) VALUES
('วัตถุดิบอาหาร', 'ซื้อวัตถุดิบทำอาหารประจำวัน', 2500.00, 'food', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ค่าไฟฟ้า', 'ค่าไฟฟ้าประจำเดือน', 1200.00, 'utilities', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('ค่าแก๊ส', 'ค่าแก๊สทำครัว', 800.00, 'utilities', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));

-- Sample reservations
INSERT INTO reservations (customer_name, customer_phone, table_id, reservation_date, reservation_time, guest_count, status, notes, created_at, updated_at) VALUES
('นาย สมชาย ใจดี', '0812345678', 1, '2024-12-28', '19:00:00', 4, 'confirmed', 'งานเลี้ยงวันเกิด', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())),
('นางสาว สมหญิง รักสะอาด', '0823456789', 2, '2024-12-29', '18:30:00', 2, 'pending', '', EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()));
