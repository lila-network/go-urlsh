package web

import (
	"codeberg.org/lauralani/go-urlsh/internal/db"
	"codeberg.org/lauralani/go-urlsh/internal/misc"
	"github.com/gofiber/fiber/v2"
)

func HandleAdminRootGet(c *fiber.Ctx) error {
	cookie := c.Cookies(misc.CookieName)
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
		c.ClearCookie(misc.CookieName)
		c.Location("/admin/login")
		c.Status(fiber.StatusSeeOther)
		return nil
	}
}
