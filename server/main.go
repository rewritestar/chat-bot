package main

import (
	"chat-bot/src/auth"
	"chat-bot/src/init/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.NewMariaDB()
	router := gin.Default()

	apiRouter := router.Group("/api")
	auth.Main(apiRouter)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
