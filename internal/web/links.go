package web

import (
	"codeberg.org/lauralani/go-urlsh/internal/db"
	"codeberg.org/lauralani/go-urlsh/internal/misc"
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func HandleAdminLinkNewGet(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) {
		c.Location("/admin/")
		c.Status(fiber.StatusSeeOther)
		return nil
	}

	return c.Render("add_link", nil)
}

func HandleAdminLinkEditGet(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) {
		c.Location("/admin/")
		c.Status(fiber.StatusSeeOther)
		return nil
	}

	var link = new(models.Link)
	var id = c.Params("id", "")

	if id == "" {
		log.Println("[HandleAdminLinkEditGet] no id param given")
		return fiber.NewError(fiber.StatusBadRequest, "400 Bad Request")
	}

	err := models.DB.NewSelect().Model(link).Where("id = ?", id).Scan(context.Background())

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[HandleAdminLinkEditGet] Shortlink %v not found\n", id)
			return fiber.NewError(fiber.StatusNotFound, "404 Not Found")
		} else {
			log.Printf("[HandleAdminLinkEditGet] Error querying Shortlink %v from database: %v\n", id, err)
			return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
		}
	}

	return c.Render("edit_link", link)
}

func HandleAdminLinkIndexGet(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) {
		c.Location("/admin/login")
		c.Status(fiber.StatusSeeOther)
		return nil
	} else {
		var links []models.Link

		err := models.DB.NewSelect().Model(&links).Scan(context.Background())
		if err != nil {
			return fmt.Errorf("error querying links: %v", err.Error())
		}
		return c.Render("links", links)
	}
}
