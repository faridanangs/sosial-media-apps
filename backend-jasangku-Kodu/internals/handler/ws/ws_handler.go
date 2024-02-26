package ws

import (
	"time"

	"github.com/faridanang/jasangku-kodu/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	hub      *Hub
	validate *validator.Validate
}

func NewHandler(hub *Hub, validate *validator.Validate) *Handler {
	return &Handler{hub: hub, validate: validate}
}

func (h *Handler) CreateRoom(ctx *fiber.Ctx) error {
	req := model.CreateRoom{}
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.validate.Struct(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:    req.ID,
		Name:  req.Name,
		Users: make(map[string]*User),
	}

	return ctx.Status(fiber.StatusOK).JSON(model.CreateRoom{
		ID:   req.ID,
		Name: req.Name,
	})
}

func (h *Handler) GetRooms(ctx *fiber.Ctx) error {
	rooms := make([]Room, 0)
	for _, v := range h.hub.Rooms {
		rooms = append(rooms, Room{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(rooms)
}

func (h *Handler) JoinRoom(conn *websocket.Conn) {
	roomID := conn.Params("roomId")
	userID := conn.Query("userId")
	username := conn.Query("username")
	usr := &User{
		ID:       userID,
		Username: username,
		Message:  make(chan *Message, 10),
		RoomID:   roomID,
		Conn:     conn,
	}
	m := &Message{
		Username:  username,
		Content:   username + " baru saja memasuki ruang obrolan",
		RoomID:    roomID,
		CreatedAt: time.Now().Local(),
	}

	h.hub.Register <- usr
	h.hub.Broadcast <- m

	go usr.writeMessage()
	usr.readMessage(h.hub)
}

type UserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *fiber.Ctx) error {
	var clients []UserRes
	roomId := c.Params("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]UserRes, 0)
		c.JSON(clients)
	}

	for _, r := range h.hub.Rooms[roomId].Users {
		clients = append(clients, UserRes{
			ID:       r.ID,
			Username: r.Username,
		})
	}

	return c.JSON(clients)

}
