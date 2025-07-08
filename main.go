package main

import (
	"Backend-POS/cmd"
	config "Backend-POS/configs"
	"Backend-POS/controller/auth"
	"Backend-POS/controller/categories"
	"Backend-POS/controller/dashboard"
	"Backend-POS/controller/expense"
	menuitem "Backend-POS/controller/menu_item"
	"Backend-POS/controller/order"
	orderitem "Backend-POS/controller/order_item"
	"Backend-POS/controller/payment"
	"Backend-POS/controller/reservation"
	"Backend-POS/controller/staff"
	"Backend-POS/controller/table"
	"Backend-POS/middlewares"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// โหลดไฟล์ .env สำหรับการพัฒนาในเครื่อง
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}

	if len(os.Args) > 1 {
		cmd.Execute()
		os.Exit(0)
	}

	// ==== ถ้าไม่มี arg ให้รันเป็น API server ปกติ ====
	config.Database()
	log.Println("Database connected successfully")
	r := gin.Default()

	// ปรับ CORS สำหรับ production และรองรับ localhost frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", // Create React App/Next.js
			"http://localhost:3001", // Vite alternative port
			"http://localhost:3002", // Vite port 3002
			"http://localhost:5173", // Vite default port
			"http://localhost:8080", // Local development
			"https://*.vercel.app",
			"https://komkemkty-frontend-pos.vercel.app",
			"https://frontend-pos-jade.vercel.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))

	md := middlewares.AuthMiddleware()

	// Auth endpoints
	r.POST("/staff/login", auth.LoginStaff)

	// Dashboard endpoints
	r.GET("/summary", md, dashboard.GetDashboardSummary)

	// Staff endpoints
	r.GET("/staff/info", md, staff.GetInfoStaff)
	r.GET("/staff", md, staff.StaffList)
	r.GET("/staff/:id", md, staff.GetStaffByID)
	r.POST("/staff/create", staff.CreateStaff)
	r.PATCH("/staff/:id", md, staff.UpdateStaff)
	r.DELETE("/staff/:id", md, staff.DeleteStaff)

	// Categories endpoints
	r.GET("/categories", md, categories.CategoryList)
	r.GET("/categories/:id", md, categories.GetCategoryByID)
	r.POST("/categories/create", md, categories.CreateCategory)
	r.PATCH("/categories/:id", md, categories.UpdateCategory)
	r.DELETE("/categories/:id", md, categories.DeleteCategory)

	// Expense endpoints
	r.GET("/expenses", md, expense.ExpenseList)
	r.GET("/expenses/:id", md, expense.GetExpenseByID)
	r.POST("/expenses/create", md, expense.CreateExpense)
	r.PATCH("/expenses/:id", md, expense.UpdateExpense)
	r.DELETE("/expenses/:id", md, expense.DeleteExpense)

	// Menu Item endpoints
	r.GET("/menu-items", md, menuitem.ListMenuItems)
	r.GET("/menu-items/:id", md, menuitem.GetMenuItemByID)
	r.POST("/menu-items/create", md, menuitem.CreateMenuItem)
	r.PATCH("/menu-items/:id", md, menuitem.UpdateMenuItem)
	r.DELETE("/menu-items/:id", md, menuitem.DeleteMenuItem)

	// Order endpoints
	r.GET("/orders", md, order.ListOrders)
	r.GET("/orders/:id", md, order.GetOrderById)
	r.POST("/orders/create", md, order.CreateOrder)
	r.PATCH("/orders/:id", md, order.UpdateOrder)
	r.DELETE("/orders/:id", md, order.DeleteOrder)

	// Order Item endpoints
	r.GET("/order-items", md, orderitem.ListOrderItems)
	r.GET("/order-items/:id", md, orderitem.GetOrderItemById)
	r.POST("/order-items/create", md, orderitem.CreateOrderItem)
	r.PATCH("/order-items/:id", md, orderitem.UpdateOrderItem)
	r.DELETE("/order-items/:id", md, orderitem.DeleteOrderItem)

	// Payment endpoints
	r.GET("/payments", md, payment.ListPayments)
	r.GET("/payments/:id", md, payment.GetPaymentById)
	r.POST("/payments/create", md, payment.CreatePayment)
	r.PATCH("/payments/:id", md, payment.UpdatePayment)
	r.DELETE("/payments/:id", md, payment.DeletePayment)

	// Reservation endpoints
	r.GET("/reservations", md, reservation.ListReservations)
	r.GET("/reservations/:id", md, reservation.GetReservationById)
	r.POST("/reservations/create", md, reservation.CreateReservation)
	r.PATCH("/reservations/:id", md, reservation.UpdateReservation)
	r.DELETE("/reservations/:id", md, reservation.DeleteReservation)

	// Table endpoints
	r.GET("/tables", md, table.ListTables)
	r.GET("/tables/:id", md, table.GetTableById)
	r.POST("/tables/create", md, table.CreateTable)
	r.PATCH("/tables/:id", md, table.UpdateTable)
	r.DELETE("/tables/:id", md, table.DeleteTable)
	// ===== PUBLIC ENDPOINTS สำหรับลูกค้า (ไม่ต้อง auth) =====

	// Public - สแกน QR Code โต๊ะเพื่อดูเมนู
	r.GET("/public/menu/:qrCodeIdentifier", table.PublicMenuByQrCode)

	// Public - ดูเมนูทั้งหมด (ไม่จำกัดจำนวน)
	r.GET("/public/menu", menuitem.PublicListMenuItems)

	// Public - สร้างออเดอร์ (ลูกค้าสั่งอาหาร)
	r.POST("/public/orders/create", order.PublicCreateOrder)

	// Public - ดูประวัติออเดอร์ตามโต๊ะ (เฉพาะยังไม่ชำระเงิน)
	r.GET("/public/orders/table/:qrCodeIdentifier", order.PublicGetOrdersByTable)

	// Public - ดูสถานะออเดอร์เฉพาะ
	r.GET("/public/orders/:orderID/table/:qrCodeIdentifier", order.PublicGetOrderStatus)

	// Public - ดูประวัติออเดอร์ทั้งหมด (รวมที่ชำระแล้ว)
	r.GET("/public/orders/history/:qrCodeIdentifier", order.PublicGetAllOrderHistory)

	// Public - ดูสรุปโต๊ะ
	r.GET("/public/table/summary/:qrCodeIdentifier", order.PublicGetTableSummary)

	// Staff - ล้างประวัติโต๊ะหลังชำระเงิน (ต้อง auth)
	r.POST("/staff/orders/clear-table/:qrCodeIdentifier", md, order.PublicClearTableHistory)

	// Staff - ล้างประวัติแบบละเอียด (เพิ่มฟีเจอร์)
	r.POST("/staff/orders/advanced-clear/:qrCodeIdentifier", md, order.AdvancedClearTableHistory)

	// Staff - ยกเลิกออเดอร์เฉพาะ
	r.POST("/staff/orders/cancel/:orderID/table/:qrCodeIdentifier", md, order.CancelSpecificOrder)

	// Staff - อัปเดตสถานะออเดอร์
	r.PATCH("/staff/orders/:orderID/status", md, order.UpdateOrderStatus)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"message":   "Server is running",
			"timestamp": func() int64 { return 1704067200 }(), // Unix timestamp
		})
	})

	// Emergency fallback endpoint for testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Environment: %s", os.Getenv("GIN_MODE"))
	r.Run(":" + port)
}
