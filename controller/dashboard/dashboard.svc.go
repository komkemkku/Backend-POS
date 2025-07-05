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
		Where("deleted_at IS NULL").
		Count(ctx)
	if err != nil {
		return nil, err
	}
	summary.TotalTables = totalTables

	// Get today's date range (00:00:00 to 23:59:59)
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Get today's orders count
	todayOrders, err := db.NewSelect().
		Table("orders").
		Where("created_at >= ? AND created_at <= ?", startOfDay, endOfDay).
		Where("deleted_at IS NULL").
		Count(ctx)
	if err != nil {
		return nil, err
	}
	summary.TodayOrders = todayOrders

	// Get pending orders count (status: pending, preparing, ready)
	pendingOrders, err := db.NewSelect().
		Table("orders").
		Where("status IN (?, ?, ?)", "pending", "preparing", "ready").
		Where("deleted_at IS NULL").
		Count(ctx)
	if err != nil {
		return nil, err
	}
	summary.PendingOrders = pendingOrders

	// Get today's revenue from payments
	var todayRevenue float64
	err = db.NewSelect().
		Table("payments").
		Column("COALESCE(SUM(amount), 0) as revenue").
		Where("created_at >= ? AND created_at <= ?", startOfDay, endOfDay).
		Where("deleted_at IS NULL").
		Scan(ctx, &todayRevenue)
	if err != nil {
		return nil, err
	}
	summary.TodayRevenue = todayRevenue

	return summary, nil
}
