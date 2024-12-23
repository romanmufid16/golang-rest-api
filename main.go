package main

import (
	"github.com/gin-gonic/gin"
	"github.com/romanmufid16/golang-rest-api/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductById)
	router.POST("/products", controllers.AddProduct)

	router.Run("localhost:8080")
}
