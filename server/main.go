package main

import (
	"log"
	"net/http"
	"strings"

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

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 파일 요청이면 404
		if strings.Contains(path, ".") {
			c.Status(http.StatusNotFound)
			return
		}

		// 페이지 요청이면 index.html
		c.File("./index.html")
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
