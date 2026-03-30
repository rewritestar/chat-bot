package interactor

import (
	"chat-bot/src/core/cerror"

	"github.com/gin-gonic/gin"
)

func ErrorPresenter(ctx *gin.Context, statusCode int, err error) {
	cerror.HandleError(ctx, statusCode, err)
}
