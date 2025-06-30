package orderitem

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
)

var db = config.Database()

func ListOrderItemService(ctx context.Context, req requests.OrderItemRequest) ([]response.OrderItemResponses, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}
	resp := []response.OrderItemResponses{}
	query := db.NewSelect().
		Model((*model.OrderItems)(nil)).
		Column("id", "order_id", "menu_item_id", "quantity", "price_per_item", "notes")

	// Filter
	if req.Search != "" {
		query = query.Where("notes ILIKE ?", "%"+req.Search+"%")
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

func GetOrderItemByIdService(ctx context.Context, id int) (*response.OrderItemResponses, error) {
	data := &response.OrderItemResponses{}
	err := db.NewSelect().
		Model((*model.OrderItems)(nil)).
		Column("id", "order_id", "menu_item_id", "quantity", "price_per_item", "notes").
		Where("id = ?", id).Scan(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreateOrderItemService(ctx context.Context, req requests.OrderItemCreateRequest) (*response.OrderItemResponses, error) {
	item := &model.OrderItems{
		OrderID:      req.OrderID,
		MenuItemID:   req.MenuItemID,
		Quantity:     req.Quantity,
		PricePerItem: req.PricePerItem,
		Notes:        req.Notes,
	}
	_, err := db.NewInsert().Model(item).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetOrderItemByIdService(ctx, item.ID)
}

func UpdateOrderItemService(ctx context.Context, id int, req requests.OrderItemUpdateRequest) (*response.OrderItemResponses, error) {
	item := &model.OrderItems{
		ID:           id,
		OrderID:      req.OrderID,
		MenuItemID:   req.MenuItemID,
		Quantity:     req.Quantity,
		PricePerItem: req.PricePerItem,
		Notes:        req.Notes,
	}
	_, err := db.NewUpdate().Model(item).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetOrderItemByIdService(ctx, id)
}

func DeleteOrderItemService(ctx context.Context, id int) error {
	_, err := db.NewDelete().Model((*model.OrderItems)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
