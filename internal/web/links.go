package web

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/lila-network/go-urlsh/internal/links"
	"github.com/lila-network/go-urlsh/models"
)

func HandleAdminLinkNewGet(c fiber.Ctx) error {

	return c.Render("add_link", nil)
}

func HandleAdminLinkEditGet(c fiber.Ctx) error {

	var link = new(links.Link)
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

func HandleAdminLinkIndexGet(c fiber.Ctx) error {
	var links []links.Link

	err := models.DB.NewSelect().Model(&links).Scan(context.Background())
	if err != nil {
		return fmt.Errorf("error querying links: %v", err.Error())
	}
	return c.Render("links", links)
}
