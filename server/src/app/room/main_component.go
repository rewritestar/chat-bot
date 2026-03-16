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

	roomRouter.POST("", func(ctx *gin.Context) {
		reqData, err := interactor.RoomController(ctx)
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

	roomRouter.GET("", func(ctx *gin.Context) {
		reqData, err := interactor.RoomController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		room, err := service.SingleRoom(*reqData)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.RoomPresenter(ctx, room)
	})
}
