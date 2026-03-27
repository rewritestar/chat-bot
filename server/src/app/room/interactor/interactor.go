package interactor

import (
	"errors"
	"net/http"
	"strconv"

	"chat-bot/src/app/room/domain"
	"chat-bot/src/app/room/interactor/message"
	core_values "chat-bot/src/core/values"

	"github.com/gin-gonic/gin"
)

func RoomIndexController(ctx *gin.Context) (uint, error) {
	workerID, ok := ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		return 0, err
	}
	return workerID.(uint), nil
}

func RoomShowController(ctx *gin.Context) (uint, error) {
	roomID, err := getID(ctx)
	if err != nil {
		return 0, err
	}
	return roomID, nil
}

func RoomSaveController(ctx *gin.Context) (*domain.Room, error) {
	workerID, ok := ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		return nil, err
	}
	reqData := message.RequestRoom{}
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, err
	}
	return reqData.ToRoom(workerID.(uint)), nil
}

func RoomUpdateController(ctx *gin.Context) (*domain.Room, error) {
	roomID, err := getID(ctx)
	if err != nil {
		return nil, err
	}
	workerID, ok := ctx.Get(core_values.WorkerIDKey)
	if !ok {
		err := errors.New("worker id does not exist.")
		return nil, err
	}
	reqData := message.UpdateRoom{}
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, err
	}
	return reqData.ToRoom(uint(roomID), workerID.(uint)), nil
}

func RoomIndexPresenter(ctx *gin.Context, roomList *domain.RoomList) {
	response := message.ResponseRoomList{}
	response.Build(*roomList)
	ctx.JSON(http.StatusOK, response)
}

func RoomPresenter(ctx *gin.Context, room *domain.Room) {
	response := message.ResponseRoom{}
	response.Build(*room)
	ctx.JSON(http.StatusOK, response)
}

func CreatedPresenter(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, nil)
}

func ErrorPresenter(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}

func getID(ctx *gin.Context) (uint, error) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		err := errors.New("id parsing error.")
		return 0, err
	}
	return uint(id), nil
}
