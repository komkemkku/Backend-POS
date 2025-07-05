package response

type DashboardSummaryResponse struct {
	TotalTables    int                 `json:"total_tables"`
	TodayRevenue   float64             `json:"today_revenue"`
	TodayOrders    int                 `json:"today_orders"`
	PendingOrders  int                 `json:"pending_orders"`
	PopularItems   []PopularItemResponse `json:"popular_items"`
	RecentOrders   []RecentOrderResponse `json:"recent_orders"`
	TodayCustomers int                 `json:"today_customers"`
	AvgOrderTime   int                 `json:"avg_order_time_minutes"`
}

type PopularItemResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Sold     int     `json:"sold"`
	Revenue  float64 `json:"revenue"`
}

type RecentOrderResponse struct {
	ID          int     `json:"id"`
	TableNumber int     `json:"table_number"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
	CreatedAt   int64   `json:"created_at"`
}
