package ws

type Hub struct {
	Register chan *Client

	unregister chan *Client

	broadcast chan *BroadCast

	RoomList map[uint]*RoomHub
}

type BroadCast struct {
	Client  *Client
	RoomID  *uint
	Content []byte
}

type RoomHub struct {
	clients map[*Client]bool
}

var hub *Hub

func NewHub() {
	newHub := &Hub{
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *BroadCast),
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
			for _, roomID := range client.roomIDs {
				_, ok := h.RoomList[roomID]
				if !ok {
					h.RoomList[roomID] = &RoomHub{
						clients: make(map[*Client]bool),
					}
				}
				h.RoomList[roomID].clients[client] = true
			}
		case client := <-h.unregister:
			close(client.send)
			for _, roomID := range client.roomIDs {
				roomHub, ok := h.RoomList[roomID]
				if ok {
					delete(roomHub.clients, client)
				}
			}
		case broadCast := <-h.broadcast:
			roomHub, ok := h.RoomList[*broadCast.RoomID]
			if !ok {
				h.RoomList[*broadCast.RoomID] = &RoomHub{
					clients: make(map[*Client]bool),
				}
				h.RoomList[*broadCast.RoomID].clients[broadCast.Client] = true
				roomHub = h.RoomList[*broadCast.RoomID]
			}
			for client := range roomHub.clients {
				select {
				case client.send <- broadCast.Content:
				default:
					close(client.send)
					delete(roomHub.clients, client)
				}
			}
		}
	}
}
