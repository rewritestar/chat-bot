package auth

import (
	"fmt"
	"net/http"

	"chat-bot/src/auth/message"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.RouterGroup) {
    r.POST("/signin", func(c *gin.Context) {
	req := message.RequestSignin{}
    if err := c.ShouldBindJSON(&req); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

	fmt.Println(req)

    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
}