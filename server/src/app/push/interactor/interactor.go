package interactor

import (
	"errors"
	"net/http"

	"chat-bot/src/app/push/domain"
	"chat-bot/src/app/push/interactor/message"
	"chat-bot/src/core/cerror"
	core_values "chat-bot/src/core/values"

	"github.com/gin-gonic/gin"
)

func PushSaveController(ctx *gin.Context) (*domain.PushSubscription, error) {
	reqData := message.RequestPush{}
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, err
	}
	workerID, ok := ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		return nil, err
	}
	return reqData.ToPushSub(workerID.(uint)), nil
}

func CreatedPresenter(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, nil)
}

func ErrorPresenter(ctx *gin.Context, statusCode int, err error) {
	cerror.HandleError(ctx, statusCode, err)
}
