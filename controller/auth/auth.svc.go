package auth

import (
	config "Backend-POS/configs"
	"Backend-POS/model"
	"Backend-POS/requests"
	"Backend-POS/utils"
	"context"
	"errors"
)

var db = config.Database()

func LoginUserService(ctx context.Context, req requests.LoginRequest) (*model.Staff, error) {
	// ตรวจสอบว่ามี username นี้อยู่จริง
	ex, err := db.NewSelect().TableExpr("staff").Where("username = ?", req.UserName).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง")
	}

	staff := &model.Staff{}

	// ดึงข้อมูล staff ตาม username
	err = db.NewSelect().Model(staff).Where("username = ?", req.UserName).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// เปรียบเทียบรหัสผ่าน (plain text vs hashed)
	isPasswordValid := utils.CheckPasswordHash(req.Password, staff.PasswordHash)

	if !isPasswordValid {
		return nil, errors.New("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง")
	}

	return staff, nil
}
