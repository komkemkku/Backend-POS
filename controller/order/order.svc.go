package order

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
	"fmt"
	"time"
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

	for i := range resp {
		var orderItems []model.OrderItems
		err = db.NewSelect().Model(&orderItems).Where("order_id = ?", resp[i].ID).Scan(ctx)
		if err != nil {
			continue
		}

		var responseItems []response.OrderItemResponses
		for _, item := range orderItems {
			var menuItem model.MenuItems
			db.NewSelect().Model(&menuItem).Where("id = ?", item.MenuItemID).Scan(ctx)

			responseItems = append(responseItems, response.OrderItemResponses{
				ID:           item.ID,
				OrderID:      item.OrderID,
				MenuItemID:   item.MenuItemID,
				MenuName:     menuItem.Name,
				Quantity:     item.Quantity,
				PricePerItem: item.PricePerItem,
				SubTotal:     item.PricePerItem * float64(item.Quantity),
				Notes:        item.Notes,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			})
		}

		resp[i].OrderItems = responseItems
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

	var orderItems []model.OrderItems
	err = db.NewSelect().Model(&orderItems).Where("order_id = ?", order.ID).Scan(ctx)
	if err != nil {
		return order, nil
	}

	var responseItems []response.OrderItemResponses
	for _, item := range orderItems {
		var menuItem model.MenuItems
		db.NewSelect().Model(&menuItem).Where("id = ?", item.MenuItemID).Scan(ctx)

		responseItems = append(responseItems, response.OrderItemResponses{
			ID:           item.ID,
			OrderID:      item.OrderID,
			MenuItemID:   item.MenuItemID,
			MenuName:     menuItem.Name,
			Quantity:     item.Quantity,
			PricePerItem: item.PricePerItem,
			SubTotal:     item.PricePerItem * float64(item.Quantity),
			Notes:        item.Notes,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}

	order.OrderItems = responseItems
	return order, nil
}

func CreateOrderService(ctx context.Context, staffID int, req requests.OrderCreateRequest) (*response.OrderResponse, error) {
	order := &model.Orders{
		TableID:     req.TableID,
		StaffID:     staffID,
		Status:      req.Status,
		TotalAmount: req.TotalAmount,
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	_, err := db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetOrderByIdService(ctx, order.ID)
}

func PublicCreateOrderService(ctx context.Context, req requests.PublicOrderCreateRequest) (*response.PublicOrderResponse, error) {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", req.QrCodeIdentifier).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	var totalAmount float64
	for _, item := range req.Items {
		var menuItem model.MenuItems
		err := db.NewSelect().Model(&menuItem).Where("id = ? AND is_available = ?", item.MenuItemID, true).Scan(ctx)
		if err != nil {
			return nil, fmt.Errorf("ไม่พบเมนู ID: %d หรือเมนูไม่พร้อมใช้งาน", item.MenuItemID)
		}
		totalAmount += menuItem.Price * float64(item.Quantity)
	}

	order := &model.Orders{
		TableID:     table.ID,
		StaffID:     0,
		Status:      "pending",
		TotalAmount: totalAmount,
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	_, err = db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, err
	}

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
			Notes:        "",
		}
		orderItem.SetCreatedNow()
		orderItem.SetUpdateNow()

		_, err = db.NewInsert().Model(orderItem).Exec(ctx)
		if err != nil {
			return nil, err
		}

		responseItems = append(responseItems, response.OrderItemResponses{
			ID:           orderItem.ID,
			OrderID:      orderItem.OrderID,
			MenuItemID:   orderItem.MenuItemID,
			Quantity:     orderItem.Quantity,
			PricePerItem: orderItem.PricePerItem,
			SubTotal:     orderItem.PricePerItem * float64(orderItem.Quantity),
			Notes:        orderItem.Notes,
			CreatedAt:    orderItem.CreatedAt,
			UpdatedAt:    orderItem.UpdatedAt,
		})
	}

	return &response.PublicOrderResponse{
		ID:          order.ID,
		TableID:     order.TableID,
		TableNumber: table.TableNumber,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		Items:       responseItems,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		Message:     "ออเดอร์ของคุณได้รับการยืนยันแล้ว กรุณารอสักครู่",
	}, nil
}

func UpdateOrderService(ctx context.Context, id int, staffID int, req requests.OrderUpdateRequest) (*response.OrderResponse, error) {
	order := &model.Orders{
		ID:          id,
		TableID:     req.TableID,
		StaffID:     staffID,
		Status:      req.Status,
		TotalAmount: req.TotalAmount,
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

func PublicGetOrdersByTableService(ctx context.Context, qrCodeIdentifier string) ([]response.PublicOrderHistoryResponse, error) {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", qrCodeIdentifier).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	var orders []model.Orders
	err = db.NewSelect().Model(&orders).
		Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled') AND DATE(to_timestamp(created_at)) = CURRENT_DATE", table.ID).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var responses []response.PublicOrderHistoryResponse
	for _, order := range orders {
		var orderItems []model.OrderItems
		err = db.NewSelect().Model(&orderItems).Where("order_id = ?", order.ID).Scan(ctx)
		if err != nil {
			continue
		}

		var responseItems []response.OrderItemResponses
		for _, item := range orderItems {
			responseItems = append(responseItems, response.OrderItemResponses{
				ID:           item.ID,
				OrderID:      item.OrderID,
				MenuItemID:   item.MenuItemID,
				Quantity:     item.Quantity,
				PricePerItem: item.PricePerItem,
				SubTotal:     item.PricePerItem * float64(item.Quantity),
				Notes:        item.Notes,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			})
		}

		var statusText string
		switch order.Status {
		case "pending":
			statusText = "รอดำเนินการ"
		case "preparing":
			statusText = "กำลังเตรียม"
		case "ready":
			statusText = "พร้อมเสิร์ฟ"
		case "served":
			statusText = "เสิร์ฟแล้ว"
		case "completed":
			statusText = "เสร็จสิ้น"
		case "cancelled":
			statusText = "ยกเลิก"
		default:
			statusText = order.Status
		}

		responses = append(responses, response.PublicOrderHistoryResponse{
			ID:          order.ID,
			TableID:     order.TableID,
			TableNumber: table.TableNumber,
			Status:      order.Status,
			StatusText:  statusText,
			TotalAmount: order.TotalAmount,
			Items:       responseItems,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		})
	}

	return responses, nil
}

func PublicGetOrderStatusService(ctx context.Context, orderID int, qrCodeIdentifier string) (*response.PublicOrderStatusResponse, error) {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", qrCodeIdentifier).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	var order model.Orders
	err = db.NewSelect().Model(&order).
		Where("id = ? AND table_id = ?", orderID, table.ID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบออเดอร์ที่ระบุ")
	}

	var orderItems []model.OrderItems
	err = db.NewSelect().Model(&orderItems).Where("order_id = ?", order.ID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var responseItems []response.OrderItemResponses
	for _, item := range orderItems {
		responseItems = append(responseItems, response.OrderItemResponses{
			ID:           item.ID,
			OrderID:      item.OrderID,
			MenuItemID:   item.MenuItemID,
			Quantity:     item.Quantity,
			PricePerItem: item.PricePerItem,
			SubTotal:     item.PricePerItem * float64(item.Quantity),
			Notes:        item.Notes,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}

	var statusText string
	var statusColor string
	var estimatedTime string

	switch order.Status {
	case "pending":
		statusText = "รอดำเนินการ"
		statusColor = "orange"
		estimatedTime = "5-10 นาที"
	case "preparing":
		statusText = "กำลังเตรียม"
		statusColor = "blue"
		estimatedTime = "10-15 นาที"
	case "ready":
		statusText = "พร้อมเสิร์ฟ"
		statusColor = "green"
		estimatedTime = "พร้อมแล้ว"
	case "served":
		statusText = "เสิร์ฟแล้ว"
		statusColor = "purple"
		estimatedTime = "เสร็จสิ้น"
	case "completed":
		statusText = "เสร็จสิ้น"
		statusColor = "gray"
		estimatedTime = "เสร็จสิ้น"
	case "cancelled":
		statusText = "ยกเลิก"
		statusColor = "red"
		estimatedTime = "ยกเลิกแล้ว"
	default:
		statusText = order.Status
		statusColor = "gray"
		estimatedTime = "ไม่ทราบ"
	}

	return &response.PublicOrderStatusResponse{
		ID:            order.ID,
		TableID:       order.TableID,
		TableNumber:   table.TableNumber,
		Status:        order.Status,
		StatusText:    statusText,
		StatusColor:   statusColor,
		EstimatedTime: estimatedTime,
		TotalAmount:   order.TotalAmount,
		Items:         responseItems,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}, nil
}

func PublicGetAllOrderHistoryService(ctx context.Context, qrCodeIdentifier string) (*response.PublicOrderSummaryResponse, error) {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", qrCodeIdentifier).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	var currentOrders []model.Orders
	err = db.NewSelect().Model(&currentOrders).
		Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled') AND DATE(to_timestamp(created_at)) = CURRENT_DATE", table.ID).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var paidOrders []model.Orders
	err = db.NewSelect().Model(&paidOrders).
		Where("table_id = ? AND status IN ('paid', 'completed') AND DATE(to_timestamp(created_at)) = CURRENT_DATE", table.ID).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var totalSpent float64
	var totalOrders int = len(currentOrders) + len(paidOrders)

	for _, order := range currentOrders {
		totalSpent += order.TotalAmount
	}
	for _, order := range paidOrders {
		totalSpent += order.TotalAmount
	}
	var currentOrdersResp []response.PublicOrderHistoryResponse
	for _, order := range currentOrders {
		var orderItems []model.OrderItems
		db.NewSelect().Model(&orderItems).Where("order_id = ?", order.ID).Scan(ctx)

		var responseItems []response.OrderItemResponses
		for _, item := range orderItems {
			responseItems = append(responseItems, response.OrderItemResponses{
				ID:           item.ID,
				OrderID:      item.OrderID,
				MenuItemID:   item.MenuItemID,
				Quantity:     item.Quantity,
				PricePerItem: item.PricePerItem,
				SubTotal:     item.PricePerItem * float64(item.Quantity),
				Notes:        item.Notes,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			})
		}

		statusText := getStatusText(order.Status)
		currentOrdersResp = append(currentOrdersResp, response.PublicOrderHistoryResponse{
			ID:          order.ID,
			TableID:     order.TableID,
			TableNumber: table.TableNumber,
			Status:      order.Status,
			StatusText:  statusText,
			TotalAmount: order.TotalAmount,
			Items:       responseItems,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		})
	}

	var paidOrdersResp []response.PublicOrderHistoryResponse
	for _, order := range paidOrders {
		var orderItems []model.OrderItems
		db.NewSelect().Model(&orderItems).Where("order_id = ?", order.ID).Scan(ctx)

		var responseItems []response.OrderItemResponses
		for _, item := range orderItems {
			responseItems = append(responseItems, response.OrderItemResponses{
				ID:           item.ID,
				OrderID:      item.OrderID,
				MenuItemID:   item.MenuItemID,
				Quantity:     item.Quantity,
				PricePerItem: item.PricePerItem,
				SubTotal:     item.PricePerItem * float64(item.Quantity),
				Notes:        item.Notes,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			})
		}

		statusText := getStatusText(order.Status)
		paidOrdersResp = append(paidOrdersResp, response.PublicOrderHistoryResponse{
			ID:          order.ID,
			TableID:     order.TableID,
			TableNumber: table.TableNumber,
			Status:      order.Status,
			StatusText:  statusText,
			TotalAmount: order.TotalAmount,
			Items:       responseItems,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		})
	}

	return &response.PublicOrderSummaryResponse{
		TableInfo: response.TableInfo{
			ID:               table.ID,
			TableNumber:      table.TableNumber,
			QrCodeIdentifier: table.QrCodeIdentifier,
			Status:           table.Status,
		},
		CurrentOrders: currentOrdersResp,
		PaidOrders:    paidOrdersResp,
		Summary: response.OrderSummary{
			TotalOrders:    totalOrders,
			TotalSpent:     totalSpent,
			CurrentPending: len(currentOrdersResp),
			CompletedToday: len(paidOrdersResp),
		},
		Timestamp: time.Now().Unix(),
	}, nil
}

func PublicClearTableHistoryService(ctx context.Context, qrCodeIdentifier string, staffID int) error {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", qrCodeIdentifier).Scan(ctx)
	if err != nil {
		return fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	_, err = db.NewUpdate().
		Table("orders").
		Set("status = ?", "paid").
		Set("updated_at = ?", time.Now().Unix()).
		Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("ไม่สามารถอัปเดตสถานะออเดอร์ได้: %v", err)
	}

	_, err = db.NewUpdate().
		Table("tables").
		Set("status = ?", "available").
		Where("id = ?", table.ID).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("ไม่สามารถอัปเดตสถานะโต๊ะได้: %v", err)
	}

	return nil
}

func AdvancedClearTableHistoryService(ctx context.Context, qrCodeIdentifier string, staffID int, clearType string) (*response.AdvancedClearResponse, error) {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", qrCodeIdentifier).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	var clearedOrders []model.Orders

	switch clearType {
	case "payment":
		err = db.NewSelect().Model(&clearedOrders).
			Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
			Scan(ctx)
		if err != nil {
			return nil, err
		}

		_, err = db.NewUpdate().
			Table("orders").
			Set("status = ?", "paid").
			Set("updated_at = ?", time.Now().Unix()).
			Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
			Exec(ctx)

	case "cancel_all":
		err = db.NewSelect().Model(&clearedOrders).
			Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
			Scan(ctx)
		if err != nil {
			return nil, err
		}

		_, err = db.NewUpdate().
			Table("orders").
			Set("status = ?", "cancelled").
			Set("updated_at = ?", time.Now().Unix()).
			Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
			Exec(ctx)

	case "complete_all":
		err = db.NewSelect().Model(&clearedOrders).
			Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
			Scan(ctx)
		if err != nil {
			return nil, err
		}

		_, err = db.NewUpdate().
			Table("orders").
			Set("status = ?", "completed").
			Set("updated_at = ?", time.Now().Unix()).
			Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
			Exec(ctx)

	default:
		return nil, fmt.Errorf("ประเภทการล้างไม่ถูกต้อง: %s", clearType)
	}

	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถอัปเดตสถานะออเดอร์ได้: %v", err)
	}

	if clearType == "payment" || clearType == "complete_all" {
		_, err = db.NewUpdate().
			Table("tables").
			Set("status = ?", "available").
			Where("id = ?", table.ID).
			Exec(ctx)

		if err != nil {
			return nil, fmt.Errorf("ไม่สามารถอัปเดตสถานะโต๊ะได้: %v", err)
		}
	}

	var totalAmount float64
	var clearedOrdersResp []response.OrderResponse
	for _, order := range clearedOrders {
		totalAmount += order.TotalAmount
		clearedOrdersResp = append(clearedOrdersResp, response.OrderResponse{
			ID:          order.ID,
			TableID:     order.TableID,
			StaffID:     order.StaffID,
			Status:      order.Status,
			TotalAmount: order.TotalAmount,
			CreatedAt:   order.CreatedAt,
			CompletedAt: order.UpdatedAt,
		})
	}

	return &response.AdvancedClearResponse{
		Success:        true,
		ClearType:      clearType,
		OrdersAffected: len(clearedOrders),
		TotalAmount:    totalAmount,
		TableStatus:    "available",
		ClearedOrders:  clearedOrdersResp,
		Timestamp:      time.Now().Unix(),
		Message:        getClearMessage(clearType, len(clearedOrders)),
	}, nil
}

func CancelSpecificOrderService(ctx context.Context, orderID int, qrCodeIdentifier string, staffID int, reason string) error {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", qrCodeIdentifier).Scan(ctx)
	if err != nil {
		return fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	var order model.Orders
	err = db.NewSelect().Model(&order).
		Where("id = ? AND table_id = ?", orderID, table.ID).
		Scan(ctx)
	if err != nil {
		return fmt.Errorf("ไม่พบออเดอร์ที่ระบุในโต๊ะนี้")
	}

	if order.Status == "paid" || order.Status == "completed" || order.Status == "cancelled" {
		return fmt.Errorf("ไม่สามารถยกเลิกออเดอร์ที่มีสถานะ: %s", order.Status)
	}

	_, err = db.NewUpdate().
		Table("orders").
		Set("status = ?", "cancelled").
		Set("updated_at = ?", time.Now().Unix()).
		Where("id = ?", orderID).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("ไม่สามารถยกเลิกออเดอร์ได้: %v", err)
	}

	return nil
}

func UpdateOrderStatusService(ctx context.Context, orderID int, newStatus string, staffID int) (*response.OrderResponse, error) {
	allowedStatuses := []string{"pending", "preparing", "ready", "served", "completed", "cancelled"}
	isValidStatus := false
	for _, status := range allowedStatuses {
		if status == newStatus {
			isValidStatus = true
			break
		}
	}

	if !isValidStatus {
		return nil, fmt.Errorf("สถานะไม่ถูกต้อง: %s", newStatus)
	}

	_, err := db.NewUpdate().
		Table("orders").
		Set("status = ?", newStatus).
		Set("updated_at = ?", time.Now().Unix()).
		Where("id = ?", orderID).
		Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถอัปเดตสถานะได้: %v", err)
	}

	return GetOrderByIdService(ctx, orderID)
}

func PublicGetTableSummaryService(ctx context.Context, qrCodeIdentifier string) (*response.PublicTableSummaryResponse, error) {
	var table model.Tables
	err := db.NewSelect().Model(&table).Where("qr_code_identifier = ?", qrCodeIdentifier).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบโต๊ะที่ระบุ")
	}

	pendingCount, _ := db.NewSelect().Model((*model.Orders)(nil)).
		Where("table_id = ? AND status = 'pending'", table.ID).Count(ctx)

	preparingCount, _ := db.NewSelect().Model((*model.Orders)(nil)).
		Where("table_id = ? AND status = 'preparing'", table.ID).Count(ctx)

	readyCount, _ := db.NewSelect().Model((*model.Orders)(nil)).
		Where("table_id = ? AND status = 'ready'", table.ID).Count(ctx)

	var totalPending float64
	db.NewSelect().Model((*model.Orders)(nil)).
		Column("COALESCE(SUM(total_amount), 0)").
		Where("table_id = ? AND status NOT IN ('paid', 'completed', 'cancelled')", table.ID).
		Scan(ctx, &totalPending)

	return &response.PublicTableSummaryResponse{
		TableInfo: response.TableInfo{
			ID:               table.ID,
			TableNumber:      table.TableNumber,
			QrCodeIdentifier: table.QrCodeIdentifier,
			Status:           table.Status,
		},
		OrderCounts: response.OrderCounts{
			Pending:   int(pendingCount),
			Preparing: int(preparingCount),
			Ready:     int(readyCount),
			Total:     int(pendingCount + preparingCount + readyCount),
		},
		TotalPending: totalPending,
		LastUpdated:  time.Now().Unix(),
	}, nil
}

func getClearMessage(clearType string, ordersAffected int) string {
	switch clearType {
	case "payment":
		return fmt.Sprintf("ชำระเงินเรียบร้อย ได้ล้างประวัติ %d ออเดอร์", ordersAffected)
	case "cancel_all":
		return fmt.Sprintf("ยกเลิกออเดอร์ทั้งหมด %d รายการ", ordersAffected)
	case "complete_all":
		return fmt.Sprintf("เสร็จสิ้นออเดอร์ทั้งหมด %d รายการ", ordersAffected)
	default:
		return fmt.Sprintf("ดำเนินการเรียบร้อย %d ออเดอร์", ordersAffected)
	}
}

func getStatusText(status string) string {
	switch status {
	case "pending":
		return "รอดำเนินการ"
	case "preparing":
		return "กำลังเตรียม"
	case "ready":
		return "พร้อมเสิร์ฟ"
	case "served":
		return "เสิร์ฟแล้ว"
	case "paid":
		return "ชำระเงินแล้ว"
	case "completed":
		return "เสร็จสิ้น"
	case "cancelled":
		return "ยกเลิก"
	default:
		return status
	}
}
