package main

import (
	"Architecture_Laba_2/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Static("/static", "./views/static")

	authStorage := &handlers.AuthStorage{Users: map[string]handlers.User{}}
	authHandler := &handlers.AuthHandler{Storage: authStorage}
	userHandler := &handlers.UserHandler{Storage: authStorage}

	handlers.SetupHandlers(app, authHandler, userHandler)

	logrus.Fatal(app.Listen(":8090"))
}
