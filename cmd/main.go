package main

import (
	"go-api/controller"
	"go-api/usecase"
	"go-api/repository"
	"go-api/db"

	"github.com/gin-gonic/gin"
)
import (
	"net/http"
)
func main() {
	server := gin.Default()

	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	//Camada de repositórios
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada de usecases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetProductByID)
	server.PUT("/product/:id", ProductController.UpdateProduct)

	server.Run(":8080")
}