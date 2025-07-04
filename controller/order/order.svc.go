package order

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
	"fmt"
)

var db = config.Database()

func ListOrderService(ctx context.Context, req requests.OrderRequest) ([]response.OrderResponse, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}
	resp := []response.OrderResponse{}

	query := db.NewSelect().
		Model((*model.Orders)(nil)).
		Column("id", "table_id", "staff_id", "status", "total_amount", "created_at", "completed_at")
	if req.Search != "" {
		query = query.Where("status ILIKE ?", "%"+req.Search+"%")
	}
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	err = query.OrderExpr("id DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}

func GetOrderByIdService(ctx context.Context, id int) (*response.OrderResponse, error) {
	order := &response.OrderResponse{}
	err := db.NewSelect().
		Model((*model.Orders)(nil)).
		Column("id", "table_id", "staff_id", "status", "total_amount", "created_at", "completed_at").
		Where("id = ?", id).
		Scan(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func CreateOrderService(ctx context.Context, staffID int, req requests.OrderCreateRequest) (*response.OrderResponse, error) {
	order := &model.Orders{
		TableID:     req.TableID,
		StaffID:     staffID,
		Status:      req.Status,
		TotalAmount: req.TotalAmount, // ✅ แก้ไขชื่อฟิลด์
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	_, err := db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetOrderByIdService(ctx, order.ID)
}

// PublicCreateOrderService สำหรับลูกค้าสร้างออเดอร์ (ไม่ต้องมี staff_id)
func PublicCreateOrderService(ctx context.Context, req requests.PublicOrderCreateRequest) (*response.PublicOrderResponse, error) {
	// ตรวจสอบว่าโต๊ะมีอยู่จริง
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", req.QrCodeIdentifier).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	// คำนวณราคารวม
	var totalAmount float64
	for _, item := range req.Items {
		var menuItem model.MenuItems
		err := db.NewSelect().Model(&menuItem).Where("id = ? AND is_available = ?", item.MenuItemID, true).Scan(ctx)
		if err != nil {
			return nil, fmt.Errorf("ไม่พบเมนู ID: %d หรือเมนูไม่พร้อมใช้งาน", item.MenuItemID)
		}
		totalAmount += menuItem.Price * float64(item.Quantity)
	}

	// สร้างออเดอร์
	order := &model.Orders{
		TableID:     table.ID,
		StaffID:     0,         // ลูกค้าสั่งเอง ไม่มี staff
		Status:      "pending", // สถานะรอดำเนินการ
		TotalAmount: totalAmount,
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	_, err = db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// สร้าง order items และเก็บ response items
	var responseItems []response.OrderItemResponses

	for _, item := range req.Items {
		var menuItem model.MenuItems
		err = db.NewSelect().Model(&menuItem).Where("id = ?", item.MenuItemID).Scan(ctx)
		if err != nil {
			return nil, fmt.Errorf("ไม่พบเมนู ID: %d", item.MenuItemID)
		}

		orderItem := &model.OrderItems{
			OrderID:      order.ID,
			MenuItemID:   item.MenuItemID,
			Quantity:     item.Quantity,
			PricePerItem: menuItem.Price,
			Notes:        "", // รับจาก request ถ้ามี
		}
		orderItem.SetCreatedNow()
		orderItem.SetUpdateNow()

		_, err = db.NewInsert().Model(orderItem).Exec(ctx)
		if err != nil {
			return nil, err
		}

		// เพิ่มใน response items
		responseItems = append(responseItems, response.OrderItemResponses{
			ID:           orderItem.ID,
			OrderID:      orderItem.OrderID,
			MenuItemID:   orderItem.MenuItemID,
			Quantity:     orderItem.Quantity,
			PricePerItem: orderItem.PricePerItem,
			SubTotal:     orderItem.PricePerItem * float64(orderItem.Quantity), // คำนวณ SubTotal
			Notes:        orderItem.Notes,
			CreatedAt:    orderItem.CreatedAt,
			UpdatedAt:    orderItem.UpdatedAt,
		})
	}

	return &response.PublicOrderResponse{
		ID:          order.ID,
		TableID:     order.TableID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		Items:       responseItems, // เพิ่ม items
		CreatedAt:   order.CreatedAt,
		Message:     "ออเดอร์ของคุณได้รับการยืนยันแล้ว กรุณารอสักครู่",
	}, nil
}

func UpdateOrderService(ctx context.Context, id int, staffID int, req requests.OrderUpdateRequest) (*response.OrderResponse, error) {
	order := &model.Orders{
		ID:          id,
		TableID:     req.TableID,
		StaffID:     staffID,
		Status:      req.Status,
		TotalAmount: req.TotalAmount, // ✅ แก้ไขชื่อฟิลด์
	}
	order.SetUpdateNow()

	_, err := db.NewUpdate().Model(order).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetOrderByIdService(ctx, id)
}

func DeleteOrderService(ctx context.Context, id int) error {
	_, err := db.NewDelete().Model((*model.Orders)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
