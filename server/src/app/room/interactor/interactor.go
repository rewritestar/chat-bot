package interactor

import (
	"errors"
	"net/http"

	"chat-bot/src/app/room/domain"
	"chat-bot/src/app/room/interactor/message"
	core_values "chat-bot/src/core"

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

func ErrorPresenter(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
