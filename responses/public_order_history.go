package response

type PublicOrderHistoryResponse struct {
	ID          int                  `json:"id"`
	TableID     int                  `json:"table_id"`
	TableNumber int                  `json:"table_number"`
	Status      string               `json:"status"`
	StatusText  string               `json:"status_text"`
	TotalAmount float64              `json:"total_amount"`
	Items       []OrderItemResponses `json:"items"`
	CreatedAt   int64                `json:"created_at"`
	UpdatedAt   int64                `json:"updated_at"`
}

type PublicOrderStatusResponse struct {
	ID            int                  `json:"id"`
	TableID       int                  `json:"table_id"`
	TableNumber   int                  `json:"table_number"`
	Status        string               `json:"status"`
	StatusText    string               `json:"status_text"`
	StatusColor   string               `json:"status_color"`
	EstimatedTime string               `json:"estimated_time"`
	TotalAmount   float64              `json:"total_amount"`
	Items         []OrderItemResponses `json:"items"`
	CreatedAt     int64                `json:"created_at"`
	UpdatedAt     int64                `json:"updated_at"`
}

type PublicOrderSummaryResponse struct {
	TableInfo     TableInfo                    `json:"table_info"`
	CurrentOrders []PublicOrderHistoryResponse `json:"current_orders"`
	PaidOrders    []PublicOrderHistoryResponse `json:"paid_orders"`
	Summary       OrderSummary                 `json:"summary"`
	Timestamp     int64                        `json:"timestamp"`
}

type PublicTableSummaryResponse struct {
	TableInfo    TableInfo   `json:"table_info"`
	OrderCounts  OrderCounts `json:"order_counts"`
	TotalPending float64     `json:"total_pending"`
	LastUpdated  int64       `json:"last_updated"`
}

type OrderSummary struct {
	TotalOrders    int     `json:"total_orders"`
	TotalSpent     float64 `json:"total_spent"`
	CurrentPending int     `json:"current_pending"`
	CompletedToday int     `json:"completed_today"`
}

type OrderCounts struct {
	Pending   int `json:"pending"`
	Preparing int `json:"preparing"`
	Ready     int `json:"ready"`
	Total     int `json:"total"`
}

type AdvancedClearResponse struct {
	Success        bool            `json:"success"`
	ClearType      string          `json:"clear_type"`
	OrdersAffected int             `json:"orders_affected"`
	TotalAmount    float64         `json:"total_amount"`
	TableStatus    string          `json:"table_status"`
	ClearedOrders  []OrderResponse `json:"cleared_orders"`
	Timestamp      int64           `json:"timestamp"`
	Message        string          `json:"message"`
}
