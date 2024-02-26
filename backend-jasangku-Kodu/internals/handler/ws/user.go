package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Message  chan *Message
	RoomID   string `json:"room_id"`
	Conn     *websocket.Conn
}

type Message struct {
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	RoomID    string    `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id"`
}

func (c *User) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

type msg struct {
	Content string `json:"content"`
	ID      string `json:"id"`
}

func (c *User) readMessage(hub *Hub) {
	defer func() {
		hub.UnRegister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		reciveMsg := msg{}
		json.Unmarshal(m, &reciveMsg)

		msg := &Message{
			Content:  reciveMsg.Content,
			RoomID:   c.RoomID,
			Username: c.Username,
			UserID:   reciveMsg.ID,
		}
		hub.Broadcast <- msg
	}
}
