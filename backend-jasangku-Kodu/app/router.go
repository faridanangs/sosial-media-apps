package app

import (
	"time"

	"github.com/faridanang/jasangku-kodu/internals/handler"
	"github.com/faridanang/jasangku-kodu/internals/handler/ws"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitializeRouter(app *fiber.App, db *gorm.DB, validate *validator.Validate, hub *ws.Hub) {
	InitializeRouterUser(app, db, validate)
	InitializeRouterPost(app, db, validate)
	InitializeRouterComment(app, db, validate)
	InitializeRouterLike(app, db, validate)
	InitializeRouterFriend(app, db, validate)
	InitializeRouterNotifikasi(app, db, validate)
	InitializeRouterWebsocker(app, validate, hub)
}
func InitializeRouterWebsocker(app *fiber.App, vlidate *validator.Validate, hub *ws.Hub) {
	ws := ws.NewHandler(hub, vlidate)

	api := app.Group("/api/ws")
	api.Get("/join-room/:roomId", websocket.New(func(c *websocket.Conn) {
		ws.JoinRoom(c)
	}, websocket.Config{
		Filter: func(c *fiber.Ctx) bool {
			return true
		},
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: time.Hour * 24,
	}))
	api.Post("/create-room", func(c *fiber.Ctx) error {
		return ws.CreateRoom(c)
	})
	api.Get("/room", func(c *fiber.Ctx) error {
		return ws.GetRooms(c)
	})
	api.Get("/client/:roomId", func(c *fiber.Ctx) error {
		return ws.GetClients(c)
	})
}

func InitializeRouterUser(app *fiber.App, db *gorm.DB, vlidate *validator.Validate) {
	api := app.Group("/api/user")
	handler := handler.NewUserHandler(db, vlidate)

	api.Post("/signup", func(c *fiber.Ctx) error {
		return handler.Create(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handler.Update(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handler.Delete(c)
	})
	api.Get("/:username", func(c *fiber.Ctx) error {
		return handler.GetByUserName(c)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return handler.GetAll(c)
	})
	api.Post("/signin", func(c *fiber.Ctx) error {
		return handler.CreateToken(c)
	})
	api.Get("/signin/:email", func(c *fiber.Ctx) error {
		return handler.GetByEmail(c)
	})
}
func InitializeRouterPost(app *fiber.App, db *gorm.DB, vlidate *validator.Validate) {
	api := app.Group("/api/post")
	handler := handler.NewPostHandler(db, vlidate)

	api.Post("/create", func(c *fiber.Ctx) error {
		return handler.Create(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handler.Update(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handler.Delete(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handler.GetByID(c)
	})
	api.Get("/comment/:id", func(c *fiber.Ctx) error {
		return handler.GetByIDPost(c)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return handler.GetAll(c)
	})
}
func InitializeRouterComment(app *fiber.App, db *gorm.DB, vlidate *validator.Validate) {
	api := app.Group("/api/comment")
	handler := handler.NewCommentHandler(db, vlidate)

	api.Post("/create", func(c *fiber.Ctx) error {
		return handler.Create(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handler.Update(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handler.Delete(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handler.GetByID(c)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return handler.GetAll(c)
	})
}
func InitializeRouterLike(app *fiber.App, db *gorm.DB, vlidate *validator.Validate) {
	api := app.Group("/api/like")
	handler := handler.NewLikeHandler(db, vlidate)

	api.Post("/create", func(c *fiber.Ctx) error {
		return handler.Create(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handler.Delete(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handler.GetByID(c)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return handler.GetAll(c)
	})
}
func InitializeRouterFriend(app *fiber.App, db *gorm.DB, vlidate *validator.Validate) {
	api := app.Group("/api/friend")
	handler := handler.NewFriendHandler(db, vlidate)

	api.Post("/create", func(c *fiber.Ctx) error {
		return handler.Create(c)
	})
	api.Delete("/delete", func(c *fiber.Ctx) error {
		return handler.Delete(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handler.GetByID(c)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return handler.GetAll(c)
	})
}
func InitializeRouterNotifikasi(app *fiber.App, db *gorm.DB, vlidate *validator.Validate) {
	api := app.Group("/api/notification")
	handler := handler.NewNotificationHandler(db, vlidate)

	api.Post("/create", func(c *fiber.Ctx) error {
		return handler.Create(c)
	})
	api.Put("/update", func(c *fiber.Ctx) error {
		return handler.Update(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handler.Delete(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handler.GetByID(c)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return handler.GetAll(c)
	})
}
