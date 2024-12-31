package main

import (
	db "go-sqlite-crud-project/config"
	"go-sqlite-crud-project/controller"
	"go-sqlite-crud-project/middleware"
	"go-sqlite-crud-project/repository"
	"go-sqlite-crud-project/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	productRepo := repository.NewProductRepository(db.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Initialize Gin router
	r := gin.Default()

	// Middleware setup
	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.RateLimitingMiddleware(middleware.NewRateLimiter(10, time.Second)))
	// r.Use(middleware.InputValidationMiddleware([]string{"name", "email", "password"}))

	// Routes
	// r.POST("/register", userController.RegisterUser)
	// r.POST("/login", userController.AuthenticateUser)
	r.POST("/register", middleware.InputValidationMiddleware([]string{"name", "email", "password"}), userController.RegisterUser)
	r.POST("/login", middleware.InputValidationMiddleware([]string{"email", "password"}), userController.AuthenticateUser)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/", func(c *gin.Context) {
			userID := c.GetString("userID")
			c.JSON(http.StatusOK, gin.H{"App": "Welcome to The E-Commerce App!!", "message": "You have accessed a protected route!", "userID": userID})
		})

		// Routes
		authorized.POST("/products", middleware.InputValidationMiddleware([]string{"name", "description", "price", "stock", "category_id"}), productController.CreateProduct)
		authorized.GET("/products/:id", productController.GetProduct)
		authorized.GET("/products", productController.GetAllProducts)
		authorized.PUT("/products/:id", middleware.InputValidationMiddleware([]string{"name", "description", "price", "stock", "category_id"}), productController.UpdateProduct)
		authorized.DELETE("/products/:id", productController.DeleteProduct)

	}

	// Start the server
	r.Run(":8080")
}
