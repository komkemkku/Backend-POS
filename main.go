package main

import (
	config "Backend-POS/configs"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
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

	// md := middlewares.AuthMiddleware()

	r.Run(":8080")
}
