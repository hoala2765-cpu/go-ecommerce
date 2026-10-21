package main

import (
	"log"
	"os"

	"go-ecommerce/internal/handler"
	"go-ecommerce/internal/model"
	"go-ecommerce/internal/repository"
	"go-ecommerce/internal/service"
	"go-ecommerce/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Không tìm thấy .env, dùng biến môi trường hệ thống")
	}

	database.Connect()

	if err := database.DB.AutoMigrate(&model.Product{}, &model.User{}); err != nil {
		log.Fatal("Migrate thất bại:", err)
	}
	log.Println("✅ Migrate thành công")

	productRepo := repository.NewProductRepository(database.DB)
	productSvc := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetAll)
			products.GET("/:id", productHandler.GetByID)
			products.POST("", productHandler.Create)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server chạy tại http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
