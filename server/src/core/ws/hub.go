package ws

import (
	"encoding/json"
	"log"
)

type Hub struct {
	Register chan *Client

	unregister chan *Client

	leaveRoom chan *Client

	JoinRoom chan *JoinRoom

	Broadcast chan *BroadCast

	//온라인 client 리스트
	Clients map[uint]*Client

	//websocket 메시지 받을 room 별 client. 로그인이 되면 유저는 소속된 모든 room 에 join 된다.
	RoomList map[uint]*RoomHub
}

type JoinRoom struct {
	RoomID uint
	Client *Client
}

type BroadCast struct {
	Type    string `json:"type"`
	RoomID  uint   `json:"roomId"`
	Content any    `json:"content"`
}

type RoomHub struct {
	clients map[*Client]bool
}

var hub *Hub

func NewHub() {
	newHub := &Hub{
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		leaveRoom:  make(chan *Client),
		JoinRoom:   make(chan *JoinRoom),
		Broadcast:  make(chan *BroadCast),
		Clients:    make(map[uint]*Client),
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
			if client.userID != nil {
				h.Clients[*client.userID] = client
			}
		case client := <-h.unregister:
			close(client.send)
			delete(h.Clients, *client.userID)
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
			joinRoom.Client.CurrentRoomID = &joinRoom.RoomID
		case client := <-h.leaveRoom:
			client.CurrentRoomID = nil
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

func (h *Hub) IsPushTarget(userID, roomID uint) bool {
	client, ok := h.Clients[userID]
	if ok && client.CurrentRoomID != nil && *client.CurrentRoomID == roomID {
		// User is online && join target room
		return false
	}
	return true
}
