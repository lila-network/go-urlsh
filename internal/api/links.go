package api

import (
	"codeberg.org/lauralani/go-urlsh/internal/db"
	"codeberg.org/lauralani/go-urlsh/internal/misc"
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func HandleLinkGetAll(c *fiber.Ctx) error {
	var links []models.Link

	err := models.DB.NewSelect().Model(&links).Scan(context.Background())
	if err != nil {
		return fmt.Errorf("error querying links: %v", err.Error())
	}

	for _, link := range links {
		log.Printf("id: %v, url: %v", link.ID, link.URL)
	}
	c.JSON(links)
	return nil
}

func HandleLinkGet(c *fiber.Ctx) error {
	return nil
}

func HandleLinkPost(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) {
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	}

	// TODO: Add API-Key Auth
	var newlink models.Link

	err := json.Unmarshal(c.Body(), &newlink)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "400 Bad Request: "+err.Error())
	}

	// TODO: check if ID is already taken when ID is given

	if newlink.ID == "" {
		newlink.ID = misc.RandomString(6)
	}

	now := time.Now()
	newlink.Created = now
	newlink.Modified = now

	_, err = models.DB.NewInsert().Model(&newlink).Exec(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error: "+err.Error())
	}

	c.Append("Location", c.BaseURL()+"/api/v1/links/"+newlink.ID)
	c.Status(fiber.StatusCreated)
	c.JSON(newlink)

	return nil
}

func HandleLinkPut(c *fiber.Ctx) error {
	return nil
}

func HandleLinkDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "400 Bad Request")
	}

	numrows, err := models.DB.NewSelect().Model((*models.Link)(nil)).Where("id = ?", id).Count(context.Background())
	if err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
	}

	if numrows < 1 {
		return fiber.NewError(fiber.StatusNotFound, "404 Not Found")
	}

	_, err = models.DB.NewDelete().Model((*models.Link)(nil)).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
	}
	return nil
}
