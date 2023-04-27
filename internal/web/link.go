package web

import "github.com/gofiber/fiber/v2"

func HandleAdminLinkAddGet(c *fiber.Ctx) error {
	return c.Render("add_link", nil)
}
