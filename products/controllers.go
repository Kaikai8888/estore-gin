package products

import (
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var productRepository ProductRepository
	products, err := productRepository.Find()
	if err != nil {
		// TODO: http error response
		fmt.Printf("error: %v\n", err)
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func PostProduct(c *gin.Context) {
	var productRepository ProductRepository
	var newProduct Product
	var err error

	// parsed the request body as JSON, decodes the json payload into the struct specified as a pointer
	if err := c.BindJSON(&newProduct); err != nil {
		// TODO: http error response
		fmt.Printf("error: %v\n", err)
		fmt.Println(err)

		return
	}

	newProduct, err = productRepository.Create(newProduct)
	if err != nil {
		// TODO: http error response
		fmt.Printf("error: %v\n", err)
		fmt.Println(err)

		return
	}
	fmt.Printf("error: %v", newProduct)
	c.IndentedJSON(http.StatusOK, newProduct)
}
