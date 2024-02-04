package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/learn-go/go-todo/adapters/handlers/rest/dtos"
	"github.com/sboy99/learn-go/go-todo/internal/ports"
)

// ------------------------------INTERFACES---------------------------------- //

type IUserRestHandler interface {
	Create(c *fiber.Ctx) error
}

// ------------------------------STRUCTS---------------------------------- //

type UserHandler struct {
	Service ports.IUserServicePort
}

// ------------------------------METHODS---------------------------------- //

func (h *UserHandler) Create(c *fiber.Ctx) error {
	createUserDto := new(dtos.CreateUserDto)
	c.BodyParser(&createUserDto)

	user, err := h.Service.Create(createUserDto.Name)
	if err != nil {
		return err
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    &user,
	})

	return nil
}
