package initial

import (
	"chat-bot/src/core/ws"
	"chat-bot/src/initial/database"
)

func Main() {
	database.NewMariaDB()

	ws.NewHub()
	go ws.GetHub().Run()
}
