package reservation

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

func ListReservationService(ctx context.Context, req requests.ReservationRequest) ([]response.ReservationResponse, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}
	var items []model.Reservations
	query := db.NewSelect().Model(&items)
	if req.Search != "" {
		query = query.Where("customer_name ILIKE ?", "%"+req.Search+"%")
	}
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	err = query.OrderExpr("id DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]response.ReservationResponse, len(items))
	for i, r := range items {
		resp[i] = response.ReservationResponse{
			ID:              r.ID,
			TableID:         r.TableID,
			CustomerName:    r.CustomerName,
			CustomerPhone:   r.CustomerPhone,
			ReservationTime: strconv.FormatInt(r.ReservationTime, 10),
			NumberOfGuests:  r.NumberOfGuests,
			Status:          r.Status,
			Notes:           r.Notes,
			CreatedAt:       r.CreatedAt,
		}
	}
	return resp, total, nil
}

func GetReservationByIdService(ctx context.Context, id int) (*response.ReservationResponse, error) {
	var item model.Reservations
	err := db.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &response.ReservationResponse{
		ID:              item.ID,
		TableID:         item.TableID,
		CustomerName:    item.CustomerName,
		CustomerPhone:   item.CustomerPhone,
		ReservationTime: strconv.FormatInt(item.ReservationTime, 10),
		NumberOfGuests:  item.NumberOfGuests,
		Status:          item.Status,
		Notes:           item.Notes,
		CreatedAt:       item.CreatedAt,
	}, nil
}

func CreateReservationService(ctx context.Context, req requests.ReservationCreateRequest) (*response.ReservationResponse, error) {
	ts, err := strconv.ParseInt(req.ReservationTime, 10, 64)
	if err != nil {
		return nil, errors.New("invalid reservation_time")
	}
	item := &model.Reservations{
		TableID:         req.TableID,
		CustomerName:    req.CustomerName,
		CustomerPhone:   req.CustomerPhone,
		ReservationTime: ts,
		NumberOfGuests:  req.NumberOfGuests,
		Status:          req.Status,
		Notes:           req.Notes,
	}
	item.SetCreatedNow()

	_, err = db.NewInsert().Model(item).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetReservationByIdService(ctx, item.ID)
}

func UpdateReservationService(ctx context.Context, id int, req requests.ReservationUpdateRequest) (*response.ReservationResponse, error) {
	ts, err := strconv.ParseInt(req.ReservationTime, 10, 64)
	if err != nil {
		return nil, errors.New("invalid reservation_time")
	}
	item := &model.Reservations{
		ID:              id,
		TableID:         req.TableID,
		CustomerName:    req.CustomerName,
		CustomerPhone:   req.CustomerPhone,
		ReservationTime: ts,
		NumberOfGuests:  req.NumberOfGuests,
		Status:          req.Status,
		Notes:           req.Notes,
	}
	_, err = db.NewUpdate().Model(item).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetReservationByIdService(ctx, id)
}

func DeleteReservationService(ctx context.Context, id int) error {
	_, err := db.NewDelete().Model((*model.Reservations)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
