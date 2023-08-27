package products

import (
	"net/http"
	gin "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var productRepository ProductRepository
	products, err := productRepository.Find()
	if err != nil {
		// TODO: http error response
	}
	c.IndentedJSON(http.StatusOK, products)
}