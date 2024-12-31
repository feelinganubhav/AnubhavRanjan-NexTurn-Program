package main

import (
	"fmt"
	db "go-blog-management-system/config"
	"go-blog-management-system/controller"
	"go-blog-management-system/middleware"
	"go-blog-management-system/repository"
	"go-blog-management-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Routes
	http.HandleFunc("/register", middleware.loggingMiddleware(middleware.InputValidationMiddleware([]string{"name", "email", "password"}, blogController.RegisterUser)))
	http.HandleFunc("/login", middleware.loggingMiddleware(middleware.InputValidationMiddleware([]string{"email", "password"}, blogController.AuthenticateUser)))

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

	loggedRoutes := middleware.LoggingMiddleware(http.DefaultServeMux)

	// Start the server
	fmt.Println("Server is running on port 8081")
	if err := http.ListenAndServe(":8081", loggedRoutes); err != nil {
		fmt.Println("Error Starting Server:", err)
	}
}
