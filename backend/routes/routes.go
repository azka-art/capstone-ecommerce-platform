// backend/routes/routes.go
package routes

import (
	"ecommerce-platform/backend/controllers"
	"ecommerce-platform/backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Protected Routes
	auth := r.Group("/").Use(middleware.AuthMiddleware())
	{
		auth.GET("/products", controllers.GetProducts)
		auth.POST("/products", controllers.CreateProduct)
		auth.GET("/cart", controllers.GetCartItems)
		auth.POST("/cart", controllers.AddToCart)
		auth.DELETE("/cart/:id", controllers.RemoveFromCart)
	}

	return r
}
