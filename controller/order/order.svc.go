package order

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
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
		TableID:    req.TableID,
		StaffID:    staffID,
		Status:     req.Status,
		TotalAmout: req.TotalAmount,
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	_, err := db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetOrderByIdService(ctx, order.ID)
}

func UpdateOrderService(ctx context.Context, id int, staffID int, req requests.OrderUpdateRequest) (*response.OrderResponse, error) {
	order := &model.Orders{
		ID:         id,
		TableID:    req.TableID,
		StaffID:    staffID,
		Status:     req.Status,
		TotalAmout: req.TotalAmount,
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
