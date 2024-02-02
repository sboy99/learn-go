package exception

import "github.com/gofiber/fiber/v2"

func HandleHttpException(c *fiber.Ctx, err *HttpException) {
	c.JSON(err.StatusCode, *err.Message)
}
