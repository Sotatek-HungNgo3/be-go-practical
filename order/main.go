package main

import (
	"fmt"
	"github.com/Sotatek-HungNgo3/be-practical-order/controller"
	_ "github.com/Sotatek-HungNgo3/be-practical-order/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description Swagger for order service.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := gin.Default()
	// Swagger UI route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	orderController := controller.NewOrderController()

	orderControllerRoute := router.Group("/orders")
	{
		orderControllerRoute.GET("", orderController.GetAllOrder)
		orderControllerRoute.GET(":id", orderController.GetOrderById)
		orderControllerRoute.POST("", orderController.CreateOrder)
		orderControllerRoute.PATCH(":id/cancel", orderController.CancelOrder)
	}

	fmt.Printf("Docs available in: http://localhost:%v/swagger/index.html", port)
	router.Run(":" + port)
}
