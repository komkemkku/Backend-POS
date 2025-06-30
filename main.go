package main

import (
	"Backend-POS/cmd"
	config "Backend-POS/configs"
	"Backend-POS/controller/auth"
	"Backend-POS/controller/categories"
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
)

func main() {
	if len(os.Args) > 1 {
		cmd.Execute()
		os.Exit(0)
	}

	// ==== ถ้าไม่มี arg ให้รันเป็น API server ปกติ ====
	config.Database()
	log.Println("Database connected successfully")
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"*"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             false,
	}))

	md := middlewares.AuthMiddleware()

	//auth
	r.POST("/auth/login", auth.LoginStaff)

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

	r.Run(":8080")
}
