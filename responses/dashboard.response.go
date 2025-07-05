package response

type DashboardSummaryResponse struct {
	TotalTables       int                   `json:"total_tables"`
	TodayRevenue      float64               `json:"today_revenue"`
	TodayOrders       int                   `json:"today_orders"`
	PendingOrders     int                   `json:"pending_orders"`
	PopularItems      []PopularItemResponse `json:"popular_items"`
	RecentOrders      []RecentOrderResponse `json:"recent_orders"`
	TodayCustomers    int                   `json:"today_customers"`
	AvgOrderTime      int                   `json:"avg_order_time_minutes"`
	YesterdayRevenue  float64               `json:"yesterday_revenue"`
	YesterdayOrders   int                   `json:"yesterday_orders"`
	YesterdayCustomers int                  `json:"yesterday_customers"`
	LastWeekAvgTime   int                   `json:"last_week_avg_time_minutes"`
	SalesChart        SalesChartResponse    `json:"sales_chart"`
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

type SalesChartResponse struct {
	SevenDays   ChartDataResponse `json:"seven_days"`
	ThirtyDays  ChartDataResponse `json:"thirty_days"`
	NinetyDays  ChartDataResponse `json:"ninety_days"`
}

type ChartDataResponse struct {
	Labels []string  `json:"labels"`
	Data   []float64 `json:"data"`
}
