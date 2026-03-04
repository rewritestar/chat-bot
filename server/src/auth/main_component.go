package auth

import (
	"net/http"

	"chat-bot/src/auth/message"
	"chat-bot/src/init/database"
	"chat-bot/src/model"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.RouterGroup) {
	db := database.GetMariaDB()

	r.POST("/signin", func(c *gin.Context) {
		req := message.RequestSignin{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		worker := model.Worker{}
		worker.Email = req.Email
		worker.Password = req.Password

		err := db.Create(&worker).
			Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := message.ResponseSignin{
			WorkerID: worker.ID,
			Email:    worker.Email,
		}

		// Return JSON response
		c.JSON(http.StatusOK, res)
	})
}
