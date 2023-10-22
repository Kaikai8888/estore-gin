package products

import (
	"estore-gin/commons"
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var productRepository ProductRepository
	products, err := productRepository.Find()
	if err != nil {
		commons.HandleError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	var productRepository ProductRepository
	product, err := productRepository.FindById(c.Param("id"))
	if err != nil {
		commons.HandleError(c, err)
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
		commons.HandleError(c, err)
		return
	}

	newProduct, err = productRepository.Create(newProduct)
	if err != nil {
		commons.HandleError(c, err)
		return
	}
	fmt.Printf("error: %v", newProduct)
	c.IndentedJSON(http.StatusOK, newProduct)
}

func DeleteProduct(c *gin.Context) {
	var productRepository ProductRepository

	if err := productRepository.Delete(c.Param("id")); err != nil {
		commons.HandleError(c, err)
		return
	}

	response := map[string]string{"message": "success"}
	c.JSON(http.StatusOK, response)

}
