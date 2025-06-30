package menuitem

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
	"errors"
)

var db = config.Database()

func ListMenuItemService(ctx context.Context, req requests.MenuItemRequest) ([]response.MenuItemResponses, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.MenuItemResponses{}

	query := db.NewSelect().
		TableExpr("menu_items AS m").
		Column("m.id", "m.category_id", "m.name", "m.description", "m.price", "m.image_url", "m.is_available", "m.created_at", "m.updated_at")

	if req.Search != "" {
		query = query.Where("m.name ILIKE ?", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	err = query.OrderExpr("m.created_at DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func CreateMenuItemService(ctx context.Context, req requests.MenuItemCreateRequest) (*response.MenuItemResponses, error) {
	menu := &model.MenuItems{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		IsAvailable: req.IsAvailable,
	}
	menu.SetCreatedNow()
	menu.SetUpdateNow()

	_, err := db.NewInsert().Model(menu).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &response.MenuItemResponses{
		ID:          menu.ID,
		CategoryID:  menu.CategoryID,
		Name:        menu.Name,
		Description: menu.Description,
		Price:       menu.Price,
		ImageUrl:    menu.ImageURL,
		IsAvailable: menu.IsAvailable,
		CreatedAt:   menu.CreatedAt,
		UpdatedAt:   menu.UpdatedAt,
	}, nil
}

func UpdateMenuItemService(ctx context.Context, id int, req requests.MenuItemUpdateRequest) (*response.MenuItemResponses, error) {
	ex, err := db.NewSelect().Table("menu_items").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("menu item not found")
	}
	menu := &model.MenuItems{
		ID:          id,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		IsAvailable: req.IsAvailable,
	}
	menu.SetUpdateNow()

	_, err = db.NewUpdate().Model(menu).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	// ดึงข้อมูลล่าสุด
	return GetMenuItemByIDService(ctx, id)
}

func DeleteMenuItemService(ctx context.Context, id int) error {
	_, err := db.NewDelete().Model((*model.MenuItems)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}

func GetMenuItemByIDService(ctx context.Context, id int) (*response.MenuItemResponses, error) {
	item := &response.MenuItemResponses{}
	err := db.NewSelect().Model((*model.MenuItems)(nil)).Where("id = ?", id).Scan(ctx, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
