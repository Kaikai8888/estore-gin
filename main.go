package main

import (
	"estore-gin/products"

	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // *Engine with embeded RouterGroup
	router.GET("/products", products.GetProducts)
	router.POST("/product", products.PostProduct)
	router.Run("localhost:8080")
}
