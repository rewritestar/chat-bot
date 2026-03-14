package interactor

import (
	"errors"
	"net/http"

	"chat-bot/src/app/room/domain"
	"chat-bot/src/app/room/interactor/message"
	core_values "chat-bot/src/core/values"

	"github.com/gin-gonic/gin"
)

func RoomController(ctx *gin.Context) (*domain.Room, error) {
	workerID, ok := ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		return nil, err
	}
	reqData := message.RequestRoom{}
	return reqData.ToRoom(workerID.(uint)), nil
}

func RoomPresenter(ctx *gin.Context, room *domain.Room) {
	response := message.ResponseRoom{}
	response.Build(*room)
	ctx.JSON(http.StatusOK, response)
}

func ChatController(ctx *gin.Context) (*domain.Chat, error) {
	reqData := message.RequestChat{}
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, err
	}
	workerID, ok := ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		return nil, err
	}
	return reqData.ToChat(workerID.(uint)), nil
}

func CreatedPresenter(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, nil)
}

func ErrorPresenter(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
