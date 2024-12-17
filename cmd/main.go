package main

import (
	"fmt"
	"os"
	"products_api/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env file")
	}
}

func main() {
	server := gin.Default()

	controller.InitializeHealthController(server)
	controller.InitializeProductController(server)

	server.Run(os.Getenv("PORT"))
}
