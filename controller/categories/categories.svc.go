package categories

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
	"errors"
)

var db = config.Database()

func ListCategoryService(ctx context.Context, req requests.CategoryRequest) ([]response.CategoryResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.CategoryResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("categories AS c ").
		Column("c.id", "c.name", "c.description", "c.display_order")

	if req.Search != "" {
		query.Where("c.name ILIKE ?",
			"%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Execute query
	err = query.OrderExpr("c.display_order DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func GetByIdCategoryService(ctx context.Context, ID int) (*response.CategoryResponses, error) {
	ex, err := db.NewSelect().Table("categories").Where("id = ?", ID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("category not found")
	}

	category := &response.CategoryResponses{}

	// สร้าง query
	err = db.NewSelect().
		Table("categories").
		Where("id = ?", ID).
		Scan(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func CreateCategoryService(ctx context.Context, req requests.CategoryCreateRequest) (*model.Categories, error) {
	if req.Name == "" {
		return nil, errors.New("name categories are required")
	}

	// ตรวจสอบว่า username มีอยู่แล้วหรือไม่
	exists, err := db.NewSelect().Table("categories").Where("name = ?", req.Name).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("name categories already exists")
	}

	category := &model.Categories{
		Name:         req.Name,
		Description:  req.Description,
		DisplayOrder: req.DisplayOrder,
	}

	_, err = db.NewInsert().Model(category).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func UpdateCategoryService(ctx context.Context, ID int, req requests.CategoryUpdateRequest) (*response.CategoryResponses, error) {
    // 1. Check ว่าหมวดหมู่มีจริง
    ex, err := db.NewSelect().Table("categories").Where("id = ?", ID).Exists(ctx)
    if err != nil {
        return nil, err
    }
    if !ex {
        return nil, errors.New("category not found")
    }

    // 2. ดึงข้อมูลเก่า
    category := new(model.Categories)
    err = db.NewSelect().Model(category).Where("id = ?", ID).Scan(ctx)
    if err != nil {
        return nil, errors.New("category not found")
    }

    // 3. ตรวจสอบชื่อซ้ำ
    exists, err := db.NewSelect().
        Model((*model.Categories)(nil)).
        Where("name = ? AND id != ?", req.Name, ID).
        Exists(ctx)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("category name already exists")
    }

    // 4. Update field
    category.Name = req.Name
    category.Description = req.Description
    category.DisplayOrder = req.DisplayOrder

    // 5. อัปเดตลง DB
    _, err = db.NewUpdate().Model(category).
        Where("id = ?", ID).
        Exec(ctx)
    if err != nil {
        return nil, err
    }

    // 6. Mapping กลับเป็น response
    resp := &response.CategoryResponses{
        ID:           category.ID,
        Name:         category.Name,
        Description:  category.Description,
        DisplayOrder: category.DisplayOrder,
    }

    return resp, nil
}


func DeleteCategoryService(ctx context.Context, ID int) error {
	ex, err := db.NewSelect().TableExpr("categories").Where("id=?", ID).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("categories not found")
	}

	_, err = db.NewDelete().TableExpr("categories").Where("id =?", ID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
