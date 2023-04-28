package web

import (
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func HandleAdminLinkAddGet(c *fiber.Ctx) error {
	return c.Render("add_link", nil)
}

func HandleAdminLinkIndexGet(c *fiber.Ctx) error {
	var links []models.Link

	err := models.DB.NewSelect().Model(&links).Scan(context.Background())
	if err != nil {
		return fmt.Errorf("error querying links: %v", err.Error())
	}
	return c.Render("links", links)
}
