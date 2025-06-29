package requests

import "time"

type ExpenseRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ExpenseIdRequest struct {
	ID int `uri:"id"`
}

type ExpenseCreateRequest struct {
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Category    string    `json:"category"`
	ExpenseDate time.Time `json:"expense_date"`
	StaffID     int       `json:"staff_id"`
}

type ExpenseUpdateRequest struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Category    string    `json:"category"`
	ExpenseDate time.Time `json:"expense_date"`
	StaffID     int       `json:"staff_id"`
}

type ExpenseDeleteRequest struct {
	ID int `json:"id"`
}
