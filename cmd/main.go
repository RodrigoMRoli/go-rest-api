package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Repositories
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Usecases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Everything is fine!",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.PATCH("/product/:productId", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8080")
}
