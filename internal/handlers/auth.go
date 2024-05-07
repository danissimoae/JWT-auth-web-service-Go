package handlers

import (
	"Architecture_Laba_2/internal/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type AuthHandler struct {
	Storage *AuthStorage
}

type AuthStorage struct {
	Users map[string]User
}

type User struct {
	Email    string
	Name     string
	Password string
}

type UserHandler struct {
	Storage *AuthStorage
}

var JwtSecretKey = []byte("very-secret-key")

var errBadCredentials = errors.New("email or password is incorrect")

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	regReq := models.RegisterRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %d", err)
	}

	if _, exists := h.Storage.Users[regReq.Email]; exists {
		return errors.New("user already exists")
	}

	h.Storage.Users[regReq.Email] = User{
		Email:    regReq.Email,
		Name:     regReq.Name,
		Password: regReq.Password,
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	regReq := models.LoginRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return errBadCredentials
	}

	user, exists := h.Storage.Users[regReq.Email]
	if !exists {
		return errBadCredentials
	}

	if user.Password != regReq.Password {
		return errBadCredentials
	}

	// Генерация полезных данных для токена
	payload := jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	// Генерация JWT-токена для пользователя
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(JwtSecretKey)
	if err != nil {
		logrus.WithError(err).Error("JWT token signing failed")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(models.LoginResponse{AccessToken: t})
}
