package interactor

import "github.com/gin-gonic/gin"

func ErrorPresenter(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
