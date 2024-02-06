package rest

// ------------------------------INTERFACES---------------------------------- //

type IUserRestHandler interface {
	// Create(c *fiber.Ctx) error
}

// ------------------------------STRUCTS---------------------------------- //

type UserHandler struct {
	// Service ports.IUserServicePort
}

// ------------------------------METHODS---------------------------------- //

// func (h *UserHandler) Create(c *fiber.Ctx) error {
// 	createUserDto := new(dtos.CreateUserDto)
// 	c.BodyParser(&createUserDto)

// 	user, err := h.Service.Create(createUserDto.Name)
// 	if err != nil {
// 		return err
// 	}

// 	c.Status(http.StatusOK).JSON(fiber.Map{
// 		"message": "User created successfully",
// 		"user":    &user,
// 	})

// 	return nil
// }
