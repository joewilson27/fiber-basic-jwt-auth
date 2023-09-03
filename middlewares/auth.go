package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Middleware JWT Function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ErrorHandler: jwtError,
	})
}

type BaseResponseFailed struct {
	Message string `json:"message"`
}

// using func error handler
func jwtError(c *fiber.Ctx, err error) error {
	fmt.Println("jwt error", err)
	var resp BaseResponseFailed
	if err.Error() == "Missing or malformed JWT" {
		resp.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}
	resp.Message = err.Error()
	return c.Status(fiber.StatusUnauthorized).JSON(resp)
}
