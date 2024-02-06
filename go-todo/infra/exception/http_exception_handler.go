package exception

import "github.com/gofiber/fiber/v2"

func HandleHttpException(c *fiber.Ctx, err *HttpException) {
	c.Status(err.StatusCode).JSON(fiber.Map{
		"message": err.Message,
		"error":   err,
	})
}
