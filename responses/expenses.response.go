package response

type ExpenseResponses struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	ExpenseDate string `json:"expense_date"`
	StaffID     int    `json:"staff_id"`
	CreatedAt   int64  `json:"created_at"`
}
