package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)

	productUsecase := usecase.NewProductUsecase(productRepository)

	productController := controller.NewProductController(productUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pongue",
		})
	})

	server.GET(
		"/products",
		productController.GetProducts,
	)

	server.GET(
		"/product/:product-id",
		productController.GetProductById,
	)

	server.POST(
		"/product/create",
		productController.CreateProduct,
	)

	server.Run(":8000")
}
