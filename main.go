package main

import (
	"estore-gin/commons"
	"estore-gin/products"

	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // returns a pointer to an Engine instance with embeded RouterGroup
	router.Use(commons.ErrorHandler)

	productRoute := router.Group("/products")
	{
		productRoute.GET("", products.GetProducts) // associate the GET HTTP method and /products path with a handler function
		productRoute.GET("/:id", products.GetProduct)
		productRoute.POST("", products.PostProduct)
		productRoute.DELETE("/:id", products.DeleteProduct)
	}

	router.Run("localhost:8080") // attach the router to an http.Server and start the server
}
