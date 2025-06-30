package expense

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
	"errors"
	"strconv"
)

var db = config.Database()

type ExpenseWithStaffFlat struct {
	ID          int
	Description string
	Amount      float64
	Category    string
	ExpenseDate int
	CreatedAt   int64

	StaffID       int
	StaffUsername string
	StaffFullName string
	StaffRole     string
}

func ListExpenseService(ctx context.Context, req requests.ExpenseRequest) ([]response.ExpenseResponses, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	var temp []ExpenseWithStaffFlat

	// สร้าง query
	query := db.NewSelect().
		TableExpr("expenses AS e").
		Column(
			"e.id", "e.description", "e.amount", "e.category", "e.expense_date", "e.created_at",
		).
		ColumnExpr("s.id AS staff_id, s.username AS staff_username, s.full_name AS staff_full_name, s.role AS staff_role").
		Join("LEFT JOIN staff s ON e.staff_id = s.id")

	if req.Search != "" {
		query = query.Where("e.category ILIKE ?", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	err = query.OrderExpr("e.category DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &temp)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]response.ExpenseResponses, len(temp))
	for i, e := range temp {
		resp[i] = response.ExpenseResponses{
			ID:          e.ID,
			Description: e.Description,
			Amount:      e.Amount,
			Category:    e.Category,
			ExpenseDate: strconv.Itoa(e.ExpenseDate),
			CreatedAt:   e.CreatedAt,
			StaffID: response.StaffExpenseResponses{
				ID:       e.StaffID,
				Username: e.StaffUsername,
				FullName: e.StaffFullName,
				Role:     e.StaffRole,
			},
		}
	}

	return resp, total, nil
}

func GetByIdExpenseService(ctx context.Context, ID int) (*response.ExpenseResponses, error) {
	ex, err := db.NewSelect().Table("expenses").Where("id = ?", ID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("expenses not found")
	}

	// Struct สำหรับรับค่าจาก DB (flat)
	type ExpenseWithStaffFlat struct {
		ID            int
		Description   string
		Amount        float64
		Category      string
		ExpenseDate   int
		CreatedAt     int64
		StaffID       int
		StaffUsername string
		StaffFullName string
		StaffRole     string
	}

	var temp ExpenseWithStaffFlat

	// สร้าง query
	err = db.NewSelect().
		TableExpr("expenses AS e").
		Column(
			"e.id", "e.description", "e.amount", "e.category", "e.expense_date", "e.created_at",
		).
		ColumnExpr("s.id AS staff_id, s.username AS staff_username, s.full_name AS staff_full_name, s.role AS staff_role").
		Join("LEFT JOIN staff s ON e.staff_id = s.id").
		Where("e.id = ?", ID). // ต้องระบุ table alias "e.id"
		Scan(ctx, &temp)
	if err != nil {
		return nil, err
	}

	// Map ไป struct response ที่ nested
	expense := &response.ExpenseResponses{
		ID:          temp.ID,
		Description: temp.Description,
		Amount:      temp.Amount,
		Category:    temp.Category,
		ExpenseDate: strconv.Itoa(temp.ExpenseDate),
		CreatedAt:   temp.CreatedAt,
		StaffID: response.StaffExpenseResponses{
			ID:       temp.StaffID,
			Username: temp.StaffUsername,
			FullName: temp.StaffFullName,
			Role:     temp.StaffRole,
		},
	}

	return expense, nil
}

func CreateExpenseService(ctx context.Context, req requests.ExpenseCreateRequest) (*model.Expenses, error) {
	if req.Category == "" {
		return nil, errors.New("name categories are required")
	}

	expense := &model.Expenses{
		Description: req.Description,
		Amount:      float64(req.Amount),
		Category:    req.Category,
		ExpenseDate: req.ExpenseDate,
		StaffID:     req.StaffID,
	}

	_, err := db.NewInsert().Model(expense).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func UpdateExpenseService(ctx context.Context, ID int, req requests.ExpenseUpdateRequest) (*response.ExpenseResponses, error) {
	// 1. Check ว่า expense มีจริง
	ex, err := db.NewSelect().Table("expenses").Where("id = ?", ID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("expense not found")
	}

	// 2. Prepare struct สำหรับ update
	expense := &model.Expenses{
		ID:          ID,
		Description: req.Description,
		Amount:      float64(req.Amount),
		Category:    req.Category,
		ExpenseDate: req.ExpenseDate,
		StaffID:     req.StaffID,
	}
	expense.SetCreatedNow()

	// 3. Update ข้อมูล
	_, err = db.NewUpdate().Model(expense).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}

	// 4. Return รายการล่าสุด (join ข้อมูล staff ด้วย)
	return GetByIdExpenseService(ctx, ID)
}

func DeleteExpenseService(ctx context.Context, ID int) error {
	// 1. ตรวจสอบว่า expense มีจริงไหม
	ex, err := db.NewSelect().Model((*model.Expenses)(nil)).Where("id = ?", ID).Exists(ctx)
	if err != nil {
		return err
	}
	if !ex {
		return errors.New("expense not found")
	}

	// 2. ลบข้อมูล
	_, err = db.NewDelete().Model((*model.Expenses)(nil)).Where("id = ?", ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
