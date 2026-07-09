package web

import (
	"github.com/lila-network/go-urlsh/internal/db"
	"github.com/lila-network/go-urlsh/internal/misc"
	"github.com/lila-network/go-urlsh/models"
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func HandleRootGet(c *fiber.Ctx) error {
	if c.Params("id", "") == "" {
		if db.IsCookieValid(c.Cookies(misc.CookieName, "")) {
			c.Location("/admin/")
			c.Status(fiber.StatusSeeOther)
			return nil
		} else {
			err := c.SendString("This is the index page of go-urlsh.\n " +
				"See https://github.com/lila-network/go-urlsh for more info.")
			if err != nil {
				log.Println(err)
				return err
			}
			c.Status(fiber.StatusOK)
			return nil
		}
	} else {
		link := new(models.Link)

		err := models.DB.NewSelect().Model(link).Where("id = ?", c.Params("id")).Scan(context.Background())
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("Shortlink %v not found\n", c.Params("id"))
				return fiber.NewError(fiber.StatusNotFound, "404 Not Found")
			} else {
				log.Printf("Error querying Link %v from database: %v\n", c.Params("id"), err)
				return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
			}
		}

		c.Location(link.URL)
		c.Status(fiber.StatusSeeOther)
		return nil
	}
}
