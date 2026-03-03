package main

import (
	"chat-bot/src/auth"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
  // Create a Gin router with default middleware (logger and recovery)
  router := gin.Default()

  apiRouter := router.Group("/api")
  // apiRouter.GET("/test", func(c *gin.Context) {
  //   // Return JSON response
  //   c.JSON(http.StatusOK, gin.H{
  //     "message": "pong",
  //   })
  // })
  auth.Main(apiRouter)


  if err := router.Run(); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}