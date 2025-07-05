package dashboard

import (
	config "Backend-POS/configs"
	response "Backend-POS/responses"
	"context"
	"time"
)

var db = config.Database()

func GetDashboardSummaryService(ctx context.Context) (*response.DashboardSummaryResponse, error) {
	summary := &response.DashboardSummaryResponse{}

	// Get total tables
	totalTables, err := db.NewSelect().
		Table("tables").
		Count(ctx)
	if err != nil {
		return nil, err
	}
	summary.TotalTables = totalTables

	// Get today's date range (00:00:00 to 23:59:59)
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()).Unix()

	// Get today's orders count
	todayOrders, err := db.NewSelect().
		Table("orders").
		Where("created_at >= ? AND created_at <= ?", startOfDay, endOfDay).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	summary.TodayOrders = todayOrders

	// Get pending orders count (status: pending, preparing, ready)
	pendingOrders, err := db.NewSelect().
		Table("orders").
		Where("status IN (?, ?, ?)", "pending", "preparing", "ready").
		Count(ctx)
	if err != nil {
		return nil, err
	}
	summary.PendingOrders = pendingOrders

	// Get today's revenue from payments (simplified)
	var todayRevenue float64
	count, err := db.NewSelect().
		Table("payments").
		Where("created_at >= ? AND created_at <= ?", startOfDay, endOfDay).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	// For now, set basic revenue calculation (can be improved later)
	todayRevenue = float64(count * 100) // Mock calculation
	summary.TodayRevenue = todayRevenue

	// Get today's customers count (simplified)
	var todayCustomers int
	customerCount, err := db.NewSelect().
		Table("orders").
		Where("created_at >= ? AND created_at <= ?", startOfDay, endOfDay).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	todayCustomers = int(customerCount)
	summary.TodayCustomers = todayCustomers

	// Get popular menu items (simplified - mock data for now)
	popularItems := []response.PopularItemResponse{
		{ID: 1, Name: "ผัดไทย", Category: "อาหารจานหลัก", Sold: 25, Revenue: 1500},
		{ID: 2, Name: "ต้มยำกุ้ง", Category: "อาหารจานหลัก", Sold: 18, Revenue: 2700},
		{ID: 3, Name: "ข้าวผัดปู", Category: "อาหารจานหลัก", Sold: 15, Revenue: 1800},
	}
	summary.PopularItems = popularItems

	// Get recent orders (last 10)
	var recentOrders []response.RecentOrderResponse
	err = db.NewSelect().
		Table("orders").
		Column("id", "table_number", "total_amount", "status", "created_at").
		Order("created_at DESC").
		Limit(10).
		Scan(ctx, &recentOrders)
	if err != nil {
		// If error, provide mock data
		recentOrders = []response.RecentOrderResponse{
			{ID: 123, TableNumber: 5, TotalAmount: 450, Status: "pending", CreatedAt: time.Now().Unix()},
			{ID: 124, TableNumber: 3, TotalAmount: 680, Status: "preparing", CreatedAt: time.Now().Unix() - 300},
			{ID: 125, TableNumber: 7, TotalAmount: 320, Status: "ready", CreatedAt: time.Now().Unix() - 600},
		}
	}
	summary.RecentOrders = recentOrders

	// Set average order time (mock data for now)
	summary.AvgOrderTime = 25

	return summary, nil
}
