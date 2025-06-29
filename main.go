package main

import (
	"Backend-POS/cmd"
	config "Backend-POS/configs"
	"Backend-POS/controller/auth"
	"Backend-POS/controller/staff"
	"Backend-POS/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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

	r.Run(":8080")
}
