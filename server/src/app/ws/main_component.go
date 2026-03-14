package ws

import (
	"chat-bot/src/app/room/repository"
	"chat-bot/src/app/room/service"
	"chat-bot/src/app/ws/interactor"
	"chat-bot/src/app/ws/values"
	"chat-bot/src/core/ws"
	"chat-bot/src/initial/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.RouterGroup) {
	repo := repository.NewRoomRepository(database.GetMariaDB())
	service := service.NewRoomService(repo)
	hub := ws.GetHub()

	r.GET("", func(ctx *gin.Context) {
		conn, err := values.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}
		client := ws.NewClient(hub, conn, make(chan []byte, 256), ctx, service)
		client.Hub.Register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.WritePump()
		go client.ReadPump()
	})
}
