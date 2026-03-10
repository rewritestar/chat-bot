package interactor

import (
	"chat-bot/src/app/auth/domain"
	"chat-bot/src/app/auth/interactor/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SigninController(ctx *gin.Context) (*domain.Worker, error) {
	reqData := message.RequestSignin{}
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, err
	}
	return reqData.ToWorker(), nil
}

func SigninPresenter(ctx *gin.Context, worker *domain.Worker) {
	response := message.ResponseSignin{}
	response.Build(*worker)
	ctx.JSON(http.StatusOK, response)
}

func LoginController(ctx *gin.Context) (*domain.Worker, error) {
	reqData := message.RequestLogin{}
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, err
	}
	return reqData.ToWorker(), nil
}

func LoginPresenter(ctx *gin.Context, token *domain.Token) {
	response := message.ResponseLogin{}
	response.Build(token)
	ctx.JSON(http.StatusOK, response)
}

func ErrorPresenter(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
