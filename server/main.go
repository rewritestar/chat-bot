package main

import (
	"chat-bot/src/auth"
	"chat-bot/src/init/database"
	"chat-bot/src/init/database/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.NewMariaDB()
	r := gin.Default()

	router := r.Group("/api")
	auth.Main(router)

	router.Use(middleware.JwtMiddleware())
	router.GET("test", func(c *gin.Context) {
		fmt.Println("excuted.")
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
