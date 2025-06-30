package response

type ExpenseResponses struct {
	ID          int                   `json:"id"`
	Description string                `json:"description"`
	Amount      float64               `json:"amount"`
	Category    string                `json:"category"`
	ExpenseDate string                `json:"expense_date"`
	StaffID     StaffExpenseResponses `json:"staff"`
	CreatedAt   int64                 `json:"created_at"`
}

type StaffExpenseResponses struct {
	StaffID       int    `json:"staff_id"`
	StaffUsername string `json:"staff_username"`
	StaffFullName string `json:"staff_full_name"`
	StaffRole     string `json:"staff_role"`
}
