package main

import (
	"estore-gin/products"

	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()                       // returns a pointer to an Engine instance with embeded RouterGroup
	router.GET("/products", products.GetProducts) // associate the GET HTTP method and /products path with a handler function
	router.GET("/products/:id", products.GetProduct)
	router.POST("/products", products.PostProduct)
	router.DELETE("/products/:id", products.DeleteProduct)
	router.Run("localhost:8080") // attach the router to an http.Server and start the server
}
