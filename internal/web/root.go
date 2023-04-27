package web

import (
	"codeberg.org/lauralani/go-urlsh/internal/db"
	"github.com/gofiber/fiber/v2"
)

func HandleAdminRootGet(c *fiber.Ctx) error {
	cookie := c.Cookies("gourlsh_auth")
	if cookie == "" {
		c.Location("/admin/login")
		c.Status(fiber.StatusSeeOther)
		return nil
	}

	if db.IsCookieValid(cookie) {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "It works",
			"Plat":  "almost",
		})
	} else {
		c.ClearCookie("gourlsh_auth")
		c.Location("/admin/login")
		c.Status(fiber.StatusSeeOther)
		return nil
	}
}
