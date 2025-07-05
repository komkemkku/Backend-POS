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

	// Get today's revenue from payments
	var todayRevenue float64
	err = db.NewSelect().
		Table("payments").
		Column("COALESCE(SUM(amount), 0)").
		Where("created_at >= ? AND created_at <= ?", startOfDay, endOfDay).
		Scan(ctx, &todayRevenue)
	if err != nil {
		return nil, err
	}
	summary.TodayRevenue = todayRevenue

	// Get today's customers count (unique table numbers from today's orders)
	var todayCustomers int
	err = db.NewSelect().
		Table("orders").
		Column("COUNT(DISTINCT table_number)").
		Where("created_at >= ? AND created_at <= ?", startOfDay, endOfDay).
		Scan(ctx, &todayCustomers)
	if err != nil {
		return nil, err
	}
	summary.TodayCustomers = todayCustomers

	// Get popular menu items (top 5 from today's order_items)
	var popularItems []response.PopularItemResponse
	err = db.NewSelect().
		Table("order_items").
		Column("mi.id", "mi.name", "c.name as category", "COUNT(*) as sold", "SUM(order_items.price * order_items.quantity) as revenue").
		Join("JOIN menu_items mi ON order_items.menu_item_id = mi.id").
		Join("JOIN categories c ON mi.category_id = c.id").
		Join("JOIN orders o ON order_items.order_id = o.id").
		Where("o.created_at >= ? AND o.created_at <= ?", startOfDay, endOfDay).
		Group("mi.id", "mi.name", "c.name").
		Order("sold DESC").
		Limit(5).
		Scan(ctx, &popularItems)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	summary.RecentOrders = recentOrders

	// Set average order time (mock data for now)
	summary.AvgOrderTime = 25

	return summary, nil
}
