package main

import (
	"Backend-POS/cmd"
	config "Backend-POS/configs"
	"Backend-POS/controller/auth"
	"Backend-POS/controller/categories"
	"Backend-POS/controller/expense"
	"Backend-POS/controller/staff"
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

	r.Run(":8080")
}
