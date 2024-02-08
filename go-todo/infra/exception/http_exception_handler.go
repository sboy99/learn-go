package exception

import "github.com/gofiber/fiber/v2"

func HandleHttpException(c *fiber.Ctx, err *HttpException) {
	c.Status(err.StatusCode).JSON(fiber.Map{
		"status":  err.StatusCode,
		"message": err.Message,
		"error":   err,
	})
}
