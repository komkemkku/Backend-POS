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
		// If database error, use mock data
		totalTables = 20
	}
	summary.TotalTables = totalTables

	// Mock data for now (since database schema might not match)
	summary.TodayRevenue = 15240.50
	summary.TodayOrders = 48
	summary.PendingOrders = 5
	summary.TodayCustomers = 124
	summary.AvgOrderTime = 25

	// เพิ่มข้อมูลสำหรับคำนวณเปอร์เซ็นต์การเปลี่ยนแปลง
	summary.YesterdayRevenue = 13567.80
	summary.YesterdayOrders = 44
	summary.YesterdayCustomers = 128
	summary.LastWeekAvgTime = 27

	// ข้อมูลกราฟยอดขาย (mock data)
	summary.SalesChart = response.SalesChartResponse{
		SevenDays: response.ChartDataResponse{
			Labels: []string{"จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์", "อาทิตย์"},
			Data:   []float64{12000, 15000, 18000, 22000, 25000, 28000, 20000},
		},
		ThirtyDays: response.ChartDataResponse{
			Labels: []string{"สัปดาห์ 1", "สัปดาห์ 2", "สัปดาห์ 3", "สัปดาห์ 4"},
			Data:   []float64{85000, 92000, 105000, 98000},
		},
		NinetyDays: response.ChartDataResponse{
			Labels: []string{"เดือน 1", "เดือน 2", "เดือน 3"},
			Data:   []float64{350000, 380000, 420000},
		},
	}

	// Popular items (mock data)
	summary.PopularItems = []response.PopularItemResponse{
		{ID: 1, Name: "ผัดไทย", Category: "อาหารจานหลัก", Sold: 25, Revenue: 1500},
		{ID: 2, Name: "ต้มยำกุ้ง", Category: "อาหารจานหลัก", Sold: 18, Revenue: 2700},
		{ID: 3, Name: "ข้าวผัดปู", Category: "อาหารจานหลัก", Sold: 15, Revenue: 1800},
		{ID: 4, Name: "ส้มตำ", Category: "อาหารเรียกน้ำย่อย", Sold: 12, Revenue: 600},
	}

	// Recent orders (mock data)
	now := time.Now().Unix()
	summary.RecentOrders = []response.RecentOrderResponse{
		{ID: 123, TableNumber: 5, TotalAmount: 450, Status: "pending", CreatedAt: now - 300},
		{ID: 124, TableNumber: 3, TotalAmount: 680, Status: "preparing", CreatedAt: now - 600},
		{ID: 125, TableNumber: 7, TotalAmount: 320, Status: "ready", CreatedAt: now - 900},
		{ID: 126, TableNumber: 2, TotalAmount: 550, Status: "completed", CreatedAt: now - 1200},
		{ID: 127, TableNumber: 8, TotalAmount: 750, Status: "pending", CreatedAt: now - 1500},
	}

	return summary, nil
}
