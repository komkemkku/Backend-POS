package response

type ExpenseResponses struct {
	ID          int                   `json:"id"`
	Description string                `json:"description"`
	Amount      float64               `json:"amount"`
	Category    string                `json:"category"`
	ExpenseDate string                `json:"expense_date"`
	CreatedAt   int64                 `json:"created_at"`
	StaffID     StaffExpenseResponses `json:"staff"`
}

type StaffExpenseResponses struct {
	ID       int    `json:"staff_id"`
	Username string `json:"staff_username"`
	FullName string `json:"staff_full_name"`
	Role     string `json:"staff_role"`
}
