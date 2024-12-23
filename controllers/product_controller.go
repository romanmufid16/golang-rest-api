package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romanmufid16/golang-rest-api/models"
)

var product = []models.Product{
	{ID: "1", Name: "Laptop", Price: 1000, Quantity: 10},
	{ID: "2", Name: "Mouse", Price: 10, Quantity: 50},
}

func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, product)
}

func AddProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	product = append(product, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	for _, p := range product {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}
