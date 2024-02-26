package app

import (
	"fmt"
	"log"
	"os"

	"github.com/faridanang/jasangku-kodu/internals/handler/ws"
	"github.com/faridanang/jasangku-kodu/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func getEnv(key, defVal string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	fmt.Print(env)

	return defVal
}

func RunServer() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",             // Perbaiki penulisan OPTIONS
		AllowHeaders: "Content-Type, Authorization, Origin, Accept", // Perbaiki penulisan Content-Type
	}))

	app.Use(middleware.Authorization()) // ini middlewarenya

	validate := validator.New()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading.env file")
	}

	bbConfig := DbConfig{
		Host:     getEnv(os.Getenv("DB_HOST"), "localhost"),
		Port:     getEnv(os.Getenv("DB_PORT"), "5432"),
		User:     getEnv(os.Getenv("DB_USER"), "jasaku_user"),
		Name:     getEnv(os.Getenv("DB_NAME"), "jasangku_kodu"),
		Password: getEnv(os.Getenv("DB_PASSWORD"), "jasangku_kodu_user_role"),
		Schema:   getEnv(os.Getenv("DB_SCHEMA"), "jasangku_kodu"),
	}
	db := ConnectToDatabase(bbConfig)
	hub := ws.NewHub()
	go hub.Run()

	InitializeRouter(app, db, validate, hub)

	app.Listen(":8000")
}
