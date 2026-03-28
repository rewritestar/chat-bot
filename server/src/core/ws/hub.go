package ws

import (
	"encoding/json"
	"log"
)

type Hub struct {
	Register chan *Client

	unregister chan *Client

	JoinRoom chan *JoinRoom

	Broadcast chan *BroadCast

	clients map[*Client]bool

	RoomList map[uint]*RoomHub
}

type JoinRoom struct {
	RoomID uint
	Client *Client
}

type BroadCast struct {
	RoomID  uint `json:"roomId"`
	Content any  `json:"content"`
}

type RoomHub struct {
	clients map[*Client]bool
}

var hub *Hub

func NewHub() {
	newHub := &Hub{
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		JoinRoom:   make(chan *JoinRoom),
		Broadcast:  make(chan *BroadCast),
		clients:    make(map[*Client]bool),
		RoomList:   make(map[uint]*RoomHub),
	}
	hub = newHub
}

func GetHub() *Hub {
	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.unregister:
			close(client.send)
			delete(h.clients, client)
			for _, roomhub := range h.RoomList {
				ok := roomhub.clients[client]
				if ok {
					delete(roomhub.clients, client)
				}
			}
		case joinRoom := <-h.JoinRoom:
			room, ok := h.RoomList[joinRoom.RoomID]
			if !ok {
				room = &RoomHub{
					clients: make(map[*Client]bool),
				}
				h.RoomList[joinRoom.RoomID] = room
			}
			room.clients[joinRoom.Client] = true
		case broadCast := <-h.Broadcast:
			roomHub, ok := h.RoomList[broadCast.RoomID]
			if ok {
				send, err := json.Marshal(broadCast)
				if err != nil {
					log.Println(err.Error())
					continue
				}
				for client := range roomHub.clients {
					select {
					case client.send <- send:
					default:
						close(client.send)
						delete(roomHub.clients, client)
					}
				}
			}
		}
	}
}
