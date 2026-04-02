package push

import (
	"net/http"

	"chat-bot/src/app/push/interactor"
	"chat-bot/src/app/push/repository"
	"chat-bot/src/app/push/service"
	"chat-bot/src/initial/database"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.RouterGroup) {
	repo := repository.NewRepository(database.GetMariaDB())
	service := service.NewService(repo)

	r.POST("/push", func(ctx *gin.Context) {
		reqData, err := interactor.PushSaveController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		_, err = service.Save(*reqData)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.CreatedPresenter(ctx)
	})
}
