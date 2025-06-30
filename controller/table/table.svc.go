package table

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
)

var db = config.Database()

func ListTableService(ctx context.Context, req requests.TableRequest) ([]response.TableResponses, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}
	var tables []model.Tables
	query := db.NewSelect().Model(&tables)
	if req.Search != "" {
		query = query.Where("CAST(table_number AS TEXT) ILIKE ?", "%"+req.Search+"%")
	}
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	err = query.OrderExpr("id DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]response.TableResponses, len(tables))
	for i, t := range tables {
		resp[i] = response.TableResponses{
			ID:               t.ID,
			TableNumber:      t.TableNumber,
			Capacity:         t.Capacity,
			Status:           t.Status,
			QrCodeIdentifier: t.QrCodeIdentifier,
		}
	}
	return resp, total, nil
}

func GetTableByIdService(ctx context.Context, id int) (*response.TableResponses, error) {
	var t model.Tables
	err := db.NewSelect().Model(&t).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &response.TableResponses{
		ID:               t.ID,
		TableNumber:      t.TableNumber,
		Capacity:         t.Capacity,
		Status:           t.Status,
		QrCodeIdentifier: t.QrCodeIdentifier,
	}, nil
}

func CreateTableService(ctx context.Context, req requests.TableCreateRequest) (*response.TableResponses, error) {
	table := &model.Tables{
		TableNumber:      req.TableNumber,
		Capacity:         req.Capacity,
		Status:           req.Status,
		QrCodeIdentifier: req.QrCodeIdentifier,
	}
	_, err := db.NewInsert().Model(table).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetTableByIdService(ctx, table.ID)
}

func UpdateTableService(ctx context.Context, id int, req requests.TableUpdateRequest) (*response.TableResponses, error) {
	table := &model.Tables{
		ID:               id,
		TableNumber:      req.TableNumber,
		Capacity:         req.Capacity,
		Status:           req.Status,
		QrCodeIdentifier: req.QrCodeIdentifier,
	}
	_, err := db.NewUpdate().Model(table).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetTableByIdService(ctx, id)
}

func DeleteTableService(ctx context.Context, id int) error {
	_, err := db.NewDelete().Model((*model.Tables)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
