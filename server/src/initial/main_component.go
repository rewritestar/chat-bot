package initial

import (
	"log"

	"chat-bot/src/core/ws"
	"chat-bot/src/initial/database"
	"chat-bot/src/initial/default_data"

	"github.com/joho/godotenv"
)

func Main() {
	database.NewMariaDB()

	ws.NewHub()
	go ws.GetHub().Run()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	default_data.SetDefaultData(database.GetMariaDB())
}
