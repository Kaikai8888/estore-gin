package main

import (
	gin "github.com/gin-gonic/gin"
	"estore-gin/products"
)

func main() {
	router := gin.Default()
	router.GET("/products", products.GetProducts)

	router.Run("localhost:8080")
}