package ws

import (
	"chat-bot/src/common-service/chat"
	"os"

	"chat-bot/src/core/ollama"
	core_values "chat-bot/src/core/values"
	webpush "chat-bot/src/core/web_push"

	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	svc chat.ChatService

	ollamaSvc ollama.OllamaService

	pushSvc webpush.WebPushService

	userID *uint

	CurrentRoomID *uint
}

func NewClient(hub *Hub, conn *websocket.Conn, send chan []byte, svc chat.ChatService, ollamaSvc ollama.OllamaService, pushSvc webpush.WebPushService) *Client {
	return &Client{
		hub,
		conn,
		send,
		svc,
		ollamaSvc,
		pushSvc,
		nil,
		nil,
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage() //프론트에서만 trigger 된다.
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.readHandler(message)
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send: // chatHandler 를 거치고 프론트로 보낸다.
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) authorizeWsJwt(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return err
	}
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		workerID, ok := claims[core_values.WorkerIDKey].(float64)
		if ok {
			userID := uint(workerID)
			c.userID = &userID
		}
	}

	duration := time.Until(exp.Time)
	time.AfterFunc(duration, func() {
		c.conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(4001, "token expired"))
		c.conn.Close()
	})
	return nil
}
