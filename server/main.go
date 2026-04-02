package main

import (
	"log"

	"chat-bot/src/app/auth"
	"chat-bot/src/app/push"
	"chat-bot/src/app/room"
	"chat-bot/src/app/ws"
	"chat-bot/src/initial"
	"chat-bot/src/initial/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	initial.Main()
	r := gin.Default()

	router := r.Group("/api")
	{
		auth.Main(router)
	}

	router.Use(middleware.JwtMiddleware())
	{
		room.Main(router)
		push.Main(router)
	}

	wsRouter := r.Group("/ws")
	{
		ws.Main(wsRouter)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
