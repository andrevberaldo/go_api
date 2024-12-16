package main

import (
	"products_api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	controller.InitializeHealthController(server)
	controller.InitializeProductController(server)

	server.Run(":3001")
}
