package auth

import (
	"net/http"

	"chat-bot/src/app/auth/interactor"
	"chat-bot/src/app/auth/repository"
	"chat-bot/src/app/auth/service"
	"chat-bot/src/initial/database"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.RouterGroup) {
	repo := repository.NewAuthRepository(database.GetMariaDB())
	service := service.NewAuthService(repo)

	r.POST("/signin", func(ctx *gin.Context) {
		reqData, err := interactor.SigninController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		worker, err := service.Signin(*reqData)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.SigninPresenter(ctx, worker)
	})

	r.POST("/login", func(ctx *gin.Context) {
		reqData, err := interactor.LoginController(ctx)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusBadRequest, err)
			return
		}

		token, err := service.Login(*reqData)
		if err != nil {
			interactor.ErrorPresenter(ctx, http.StatusInternalServerError, err)
			return
		}

		interactor.LoginPresenter(ctx, token)
	})
}
