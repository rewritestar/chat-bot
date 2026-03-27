package room

import (
	"net/http"

	"chat-bot/src/app/room/interactor"
	"chat-bot/src/app/room/repository"
	"chat-bot/src/app/room/service"
	"chat-bot/src/initial/database"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.RouterGroup) {
	repo := repository.NewRoomRepository(database.GetMariaDB())
	service := service.NewRoomService(repo)

	roomRouter := r.Group("rooms")

	roomRouter.GET("", func(ctx *gin.Context) {
		workerID, err := interactor.RoomIndexController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		roomList, err := service.IndexRoom(workerID)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.RoomIndexPresenter(ctx, roomList)
	})

	roomRouter.POST("", func(ctx *gin.Context) {
		reqData, err := interactor.RoomSaveController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		room, err := service.SaveRoom(*reqData)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.RoomPresenter(ctx, room)
	})

	roomRouter.PUT("/:id", func(ctx *gin.Context) {
		reqData, err := interactor.RoomUpdateController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		room, err := service.UpdateRoom(*reqData)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.RoomPresenter(ctx, room)
	})

	roomRouter.GET("/:id", func(ctx *gin.Context) {
		roomID, err := interactor.RoomShowController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		room, err := service.ShowRoom(roomID)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.RoomPresenter(ctx, room)
	})
}
