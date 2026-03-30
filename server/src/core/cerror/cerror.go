package cerror

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonError struct {
	Comment string
	Message string
}

func (e CommonError) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, statusCode int, err error) {
	if cErr, ok := err.(CommonError); ok {
		c.JSON(statusCode, gin.H{
			"error":   cErr.Comment,
			"message": cErr.Message,
		})
		return
	}

	var message string
	switch statusCode {
	case http.StatusBadRequest:
		message = BadRequestMsg
	case http.StatusNotFound:
		message = NotFoundMsg
	default:
		message = InternalServerErrorMsg
	}

	c.JSON(statusCode, gin.H{
		"error":   err.Error(),
		"message": message,
	})
}
