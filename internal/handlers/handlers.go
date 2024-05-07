package handlers

import (
	"Architecture_Laba_2/internal/models"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// SetupHandlers Установка обработчиков
func SetupHandlers(app *fiber.App, authHandler *AuthHandler, userHandler *UserHandler) {
	// Обработчик login
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})

	// Обработчик register
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", nil)
	})

	// Обработчик index
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index",
			fiber.Map{
				"Title": "Здравствуй, пользователь! 👋",
				"Description": "Это веб-сервис с JWT-авторизацией ✨, " +
					"JWT (JSON Web Token) — это открытый стандарт для создания токенов доступа," +
					" основанный на формате JSON. Обычно он используется для" +
					" передачи данных для аутентификации пользователей в" +
					" клиент-серверных приложениях. Токены создаются сервером," +
					" подписываются секретным ключом и передаются юзеру, который в" +
					" дальнейшем использует их для подтверждения своей личности.",
			})
	})

	publicGroup := app.Group("")
	publicGroup.Post("/register", authHandler.Register)
	publicGroup.Post("/login", authHandler.Login)

	authorizedGroup := app.Group("")
	authorizedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: JwtSecretKey,
		},
		ContextKey: models.ContextKeyUser,
	}))
	authorizedGroup.Get("/profile", userHandler.Profile)
}
