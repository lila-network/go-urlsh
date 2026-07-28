package api

import "github.com/gofiber/fiber/v3"

// WriteJSON sends a JSON response with the specified HTTP status and value.
func WriteJSON(c fiber.Ctx, status int, value any) error {
	return c.Status(status).JSON(value)
}

// WriteError sends a JSON error response with the specified HTTP status and message.
func WriteError(c fiber.Ctx, status int, message string) error {
	return WriteJSON(c, status, fiber.Map{
		"status":  "error",
		"message": message,
	})
}

// WriteSuccess sends a JSON response with the status "success," the given HTTP status code, and the provided value.
func WriteSuccess(c fiber.Ctx, status int, value any) error {
	return WriteJSON(c, status, fiber.Map{
		"status": "success",
		"data":   value,
	})
}

// WriteNoContent sends a JSON response with the status "204 No Content" and empty body.
func WriteNoContent(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
