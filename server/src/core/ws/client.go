// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	roomMessage "chat-bot/src/app/room/interactor/message"

	roomDomain "chat-bot/src/app/room/domain"
	"chat-bot/src/app/room/service"
	core_values "chat-bot/src/core/values"
	"chat-bot/src/core/ws/domain"
	"errors"

	"chat-bot/src/core/ws/values"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

	ctx *gin.Context

	svc service.RoomService
}

func NewClient(hub *Hub, conn *websocket.Conn, send chan []byte, ctx *gin.Context, svc service.RoomService) *Client {
	return &Client{
		hub,
		conn,
		send,
		ctx,
		svc,
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

mainRoop:
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		reqMsg := domain.Message{}
		if err := json.Unmarshal(message, &reqMsg); err != nil {
			log.Printf("error: %v", err)
			break
		}

		dataBytes, err := json.Marshal(reqMsg.Data)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		switch reqMsg.Type {
		case values.MessageTypeAuth:
			reqAuth := domain.AuthMessage{}
			json.Unmarshal(dataBytes, &reqAuth)

			if err := c.authorizeWsJwt(reqAuth.Token); err != nil {
				log.Printf("error: %v", err)
				break mainRoop
			}
		case values.MessageTypeChat:
			reqChat := domain.ChatMessage{}
			json.Unmarshal(dataBytes, &reqChat)

			domainChat := roomDomain.Chat{}
			domainChat.Content = reqChat.Content
			domainChat.RoomID = reqChat.RoomID
			workerID, ok := c.ctx.Get(core_values.WorkerIDKey)
			if !ok {
				err := errors.New("worker id does not exist.")
				log.Printf("error: %v", err)
				break mainRoop
			}
			domainChat.CreatorID = workerID.(uint)

			if _, err := c.svc.SaveChat(domainChat); err != nil {
				log.Printf("error: %v", err)
				break mainRoop
			}

			responseChat := roomMessage.ResponseChat{}
			responseChat.Build(&domainChat)

			chatBytes, err := json.Marshal(responseChat)
			if err != nil {
				log.Printf("error: %v", err)
				break mainRoop
			}

			c.Hub.broadcast <- chatBytes
		default:
			break mainRoop
		}

	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
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
			c.ctx.Set(core_values.WorkerIDKey, uint(workerID))
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
