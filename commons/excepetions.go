package commons

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type GeneralError struct {
	Message    string
	StatusCode int `default:"400"`
}

func (e GeneralError) Error() string { return e.Message }

func ErrorHandler(c *gin.Context) {
	c.Next() // executes the pending handlers in the chain

	lastErr := c.Errors.Last()
	if lastErr == nil {
		return
	}

	statusCode := http.StatusBadRequest
	if e, ok := lastErr.Err.(GeneralError); ok {
		statusCode = e.StatusCode
	}

	c.JSON(statusCode, map[string]string{"message": lastErr.Err.Error()})

}
