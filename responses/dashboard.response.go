package response

type DashboardSummaryResponse struct {
	TotalTables   int     `json:"total_tables"`
	TodayRevenue  float64 `json:"today_revenue"`
	TodayOrders   int     `json:"today_orders"`
	PendingOrders int     `json:"pending_orders"`
}
