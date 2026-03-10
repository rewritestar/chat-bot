package main

import (
	"log"

	"chat-bot/src/app/auth"
	"chat-bot/src/app/room"
	"chat-bot/src/init/database"
	"chat-bot/src/init/database/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.NewMariaDB()
	r := gin.Default()

	router := r.Group("/api")
	{
		auth.Main(router)
	}

	router.Use(middleware.JwtMiddleware())
	{
		room.Main(router)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
