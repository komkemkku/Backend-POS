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

	var orderItems []model.OrderItems
	query := db.NewSelect().Model(&orderItems)

	if req.Search != "" {
		query = query.Where("notes ILIKE ?", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	err = query.OrderExpr("id DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]response.OrderItemResponses, len(orderItems))
	for i, item := range orderItems {
		resp[i] = response.OrderItemResponses{
			ID:           item.ID,
			OrderID:      item.OrderID,
			MenuItemID:   item.MenuItemID,
			Quantity:     item.Quantity,
			PricePerItem: item.PricePerItem,
			SubTotal:     item.PricePerItem * float64(item.Quantity),
			Notes:        item.Notes,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	return resp, total, nil
}

func GetOrderItemByIdService(ctx context.Context, id int) (*response.OrderItemResponses, error) {
	var item model.OrderItems
	err := db.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &response.OrderItemResponses{
		ID:           item.ID,
		OrderID:      item.OrderID,
		MenuItemID:   item.MenuItemID,
		Quantity:     item.Quantity,
		PricePerItem: item.PricePerItem,
		SubTotal:     item.PricePerItem * float64(item.Quantity),
		Notes:        item.Notes,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}, nil
}

func CreateOrderItemService(ctx context.Context, req requests.OrderItemCreateRequest) (*response.OrderItemResponses, error) {
	orderItem := &model.OrderItems{
		OrderID:      req.OrderID,
		MenuItemID:   req.MenuItemID,
		Quantity:     req.Quantity,
		PricePerItem: req.PricePerItem,
		Notes:        req.Notes,
	}
	orderItem.SetCreatedNow()
	orderItem.SetUpdateNow()

	_, err := db.NewInsert().Model(orderItem).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return GetOrderItemByIdService(ctx, orderItem.ID)
}

func UpdateOrderItemService(ctx context.Context, id int, req requests.OrderItemUpdateRequest) (*response.OrderItemResponses, error) {
	orderItem := &model.OrderItems{
		ID:           id,
		OrderID:      req.OrderID,
		MenuItemID:   req.MenuItemID,
		Quantity:     req.Quantity,
		PricePerItem: req.PricePerItem,
		Notes:        req.Notes,
	}
	orderItem.SetUpdateNow()

	_, err := db.NewUpdate().Model(orderItem).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return GetOrderItemByIdService(ctx, id)
}

func DeleteOrderItemService(ctx context.Context, id int) error {
	_, err := db.NewDelete().Model((*model.OrderItems)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
