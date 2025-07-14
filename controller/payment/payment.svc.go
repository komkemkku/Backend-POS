package payment

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"context"
	"strconv"
	"time"
)

var db = config.Database()

func ListPaymentService(ctx context.Context, req requests.PaymentRequest) ([]response.PaymentResponse, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}
	dbQuery := db.NewSelect().Model((*model.Payments)(nil))
	if req.Search != "" {
		dbQuery = dbQuery.Where("payment_method ILIKE ?", "%"+req.Search+"%")
	}

	total, err := dbQuery.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	var pays []model.Payments
	err = dbQuery.OrderExpr("id DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &pays)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]response.PaymentResponse, len(pays))
	for i, p := range pays {
		resp[i] = response.PaymentResponse{
			ID:              p.ID,
			OrderID:         p.OrderID,
			PaymentMethod:   p.PaymentMethod,
			AmountPaid:      p.AmountPaid,
			TransactionTime: strconv.FormatInt(p.TransactionTime, 10),
		}
	}
	return resp, total, nil
}

func GetPaymentByIdService(ctx context.Context, id int) (*response.PaymentResponse, error) {
	var pay model.Payments
	err := db.NewSelect().Model(&pay).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &response.PaymentResponse{
		ID:              pay.ID,
		OrderID:         pay.OrderID,
		PaymentMethod:   pay.PaymentMethod,
		AmountPaid:      pay.AmountPaid,
		TransactionTime: strconv.FormatInt(pay.TransactionTime, 10),
	}, nil
}

func CreatePaymentService(ctx context.Context, req requests.PaymentCreateRequest) (*response.PaymentResponse, error) {
	var tt int64
	var err error

	tt, err = strconv.ParseInt(req.TransactionTime, 10, 64)
	if err != nil {
		if t, parseErr := time.Parse(time.RFC3339, req.TransactionTime); parseErr == nil {
			tt = t.Unix()
		} else {
			tt = time.Now().Unix()
		}
	}

	pay := &model.Payments{
		OrderID:         req.OrderID,
		PaymentMethod:   req.PaymentMethod,
		AmountPaid:      req.AmountPaid,
		TransactionTime: tt,
	}
	_, err = db.NewInsert().Model(pay).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetPaymentByIdService(ctx, pay.ID)
}

func UpdatePaymentService(ctx context.Context, id int, req requests.PaymentUpdateRequest) (*response.PaymentResponse, error) {
	var tt int64
	var err error

	tt, err = strconv.ParseInt(req.TransactionTime, 10, 64)
	if err != nil {
		if t, parseErr := time.Parse(time.RFC3339, req.TransactionTime); parseErr == nil {
			tt = t.Unix()
		} else {
			tt = time.Now().Unix()
		}
	}

	pay := &model.Payments{
		ID:              id,
		OrderID:         req.OrderID,
		PaymentMethod:   req.PaymentMethod,
		AmountPaid:      req.AmountPaid,
		TransactionTime: tt,
	}
	_, err = db.NewUpdate().Model(pay).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return GetPaymentByIdService(ctx, id)
}

func DeletePaymentService(ctx context.Context, id int) error {
	_, err := db.NewDelete().Model((*model.Payments)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
