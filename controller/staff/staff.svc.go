package staff

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"Backend-POS/utils"
	"context"
	"errors"
)

var db = config.Database()

func ListStaffService(ctx context.Context, req requests.StaffRequest) ([]response.StaffResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.StaffResponses{}

	query := db.NewSelect().
		TableExpr("staff AS s ").
		Column("s.id", "s.username", "s.full_name", "s.role", "s.created_at", "s.updated_at")

	if req.Search != "" {
		query.Where("s.username ILIKE ? OR s.full_name ILIKE ?",
			"%"+req.Search+"%", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	err = query.OrderExpr("s.created_at DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func GetByIdStaffService(ctx context.Context, ID int) (*response.StaffResponses, error) {
	ex, err := db.NewSelect().Table("staff").Where("id = ?", ID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("staff not found")
	}

	staff := &response.StaffResponses{}

	err = db.NewSelect().
		TableExpr("staff AS s ").
		Column("s.id", "s.username", "s.full_name", "s.role", "s.created_at", "s.updated_at").
		Where("s.id = ?", ID).
		Scan(ctx, staff)

	if err != nil {
		return nil, err
	}

	return staff, nil
}

func CreateStaffService(ctx context.Context, req requests.StaffCreateRequest) (*model.Staff, error) {
	if req.Username == "" || req.FullName == "" || req.Role == "" {
		return nil, errors.New("username, full_name and role are required")
	}

	exists, err := db.NewSelect().Table("staff").Where("username = ?", req.Username).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	hashpassword, _ := utils.HashPassword(req.PasswordHash)
	staff := &model.Staff{
		UserName:     req.Username,
		PasswordHash: hashpassword,
		FullName:     req.FullName,
		Role:         req.Role,
	}
	staff.SetCreatedNow()
	staff.SetUpdateNow()

	_, err = db.NewInsert().Model(staff).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func UpdateStaffService(ctx context.Context, id int, req requests.StaffUpdateRequest) (*model.Staff, error) {
	if req.Username == "" || req.FullName == "" || req.Role == "" {
		return nil, errors.New("username, full_name and role are required")
	}

	staff := new(model.Staff)
	err := db.NewSelect().Model(staff).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, errors.New("staff not found")
	}

	exists, err := db.NewSelect().
		Model((*model.Staff)(nil)).
		Where("username = ? AND id != ?", req.Username, id).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	if req.PasswordHash != "" {
		hashpassword, _ := utils.HashPassword(req.PasswordHash)
		staff.PasswordHash = hashpassword
	}
	staff.UserName = req.Username
	staff.FullName = req.FullName
	staff.Role = req.Role
	staff.SetUpdateNow()

	_, err = db.NewUpdate().Model(staff).
		Where("id = ?", id).
		OmitZero().
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func DeleteStaffService(ctx context.Context, ID int) error {
	ex, err := db.NewSelect().TableExpr("staff").Where("id=?", ID).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("staff not found")
	}

	_, err = db.NewDelete().TableExpr("staff").Where("id =?", ID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
