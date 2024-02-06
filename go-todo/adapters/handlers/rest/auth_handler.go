package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/learn-go/go-todo/adapters/handlers/rest/dtos"
	"github.com/sboy99/learn-go/go-todo/infra/exception"
	"github.com/sboy99/learn-go/go-todo/internal/ports"
)

// ------------------------------INTERFACES---------------------------------- //

type IAuthRestHandler interface {
	Register(c *fiber.Ctx) error
}

// ------------------------------STRUCTS---------------------------------- //

type AuthHandler struct {
	Service ports.IAuthServicePort
}

// ------------------------------METHODS---------------------------------- //

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	dto := new(dtos.CreateUserDto)
	c.BodyParser(&dto)

	user, err := h.Service.Register(dto.Email, dto.Pass, dto.Name)
	if err != nil {
		exception.HandleHttpException(c, err)
		return nil
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"user": *user,
	})

	return nil
}
