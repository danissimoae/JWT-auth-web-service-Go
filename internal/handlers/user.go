package handlers

import (
	"Architecture_Laba_2/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func jwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
	jwtToken, ok := c.Context().Value(models.ContextKeyUser).(*jwt.Token)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_context_value": c.Context().Value(models.ContextKeyUser),
		}).Error("wrong type of JWT token in context")
		return nil, false
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_claims": jwtToken.Claims,
		}).Error("wrong type of JWT token claims")
		return nil, false
	}

	return payload, true
}

// Profile Обработчик запросов на получение информации о пользователе
func (h *UserHandler) Profile(c *fiber.Ctx) error {
	jwtPayload, ok := jwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userInfo, ok := h.Storage.Users[jwtPayload["sub"].(string)]
	if !ok {
		return errors.New("user not found")
	}

	return c.JSON(models.ProfileResponse{
		Email: userInfo.Email,
		Name:  userInfo.Name,
	})
}
