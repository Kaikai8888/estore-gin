package products

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetProduct(c *gin.Context) {
	var productRepository ProductRepository
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	product, err := productRepository.FindById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, product)
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

func DeleteProduct(c *gin.Context) {
	var productRepository ProductRepository
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	if err := productRepository.Delete(id); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	response := map[string]string{"message": "success"}
	c.JSON(http.StatusOK, response)

}
