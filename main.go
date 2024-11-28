package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"product-management-project/controllers"
	"product-management-project/models"
	"product-management-project/repositories"
	"product-management-project/services"
)

func main() {
	dsn := os.Getenv("DB_VALUE") //root:root@tcp(127.0.0.1:3306)/product_management?parseTime=true

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
		log.Fatal("Failed to connect to the database")
	}

	db.AutoMigrate(&models.UserModel{}, &models.ProductModel{})

	if err != nil {
		log.Fatal(err)
	}

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	ProductController := controllers.NewProductController(productService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	v1 := r.Group("/api")

	{
		v1.GET("/product/get-all", ProductController.FindAllProducts)
		v1.GET("/product/:id", ProductController.FindProductById)
		v1.POST("/product/:userId", ProductController.CreateProduct)
		v1.PUT("/product/:id", ProductController.UpdateProduct)
		v1.DELETE("/product/:id", ProductController.DeleteProduct)

		// User routes
		v1.GET("/user", userController.FindAllUsers)
		v1.GET("/user/:id", userController.FindUserById)
		v1.POST("/user", userController.CreateUser)
		v1.PUT("/user", userController.UpdateUser)
		v1.DELETE("/user/:id", userController.DeleteUser)
	}

	r.Run()
}
