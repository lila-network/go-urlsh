package api

import (
	"codeberg.org/lauralani/go-urlsh/internal/db"
	"codeberg.org/lauralani/go-urlsh/internal/misc"
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func HandleLinkGetAll(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) && !db.IsApiKeyValid(c.GetRespHeader("x-api-key", "")) {
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	}

	var links []models.Link

	err := models.DB.NewSelect().Model(&links).Scan(context.Background())
	if err != nil {
		return fmt.Errorf("error querying links: %v", err.Error())
	}

	for _, link := range links {
		log.Printf("id: %v, url: %v", link.ID, link.URL)
	}
	err = c.JSON(links)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func HandleLinkGet(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "501 Not Implemented")
}

func HandleLinkPost(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) && !db.IsApiKeyValid(c.GetRespHeader("x-api-key", "")) {
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	}

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

	err = c.JSON(newlink)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func HandleLinkPut(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) && !db.IsApiKeyValid(c.GetRespHeader("x-api-key", "")) {
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	}

	var dblink models.Link
	var editlink models.Link

	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "400 Bad Request")
	}

	err := models.DB.NewSelect().Model(&dblink).Where("id = ?", id).Scan(context.Background())

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[HandleAdminLinkEditGet] Shortlink %v not found\n", id)
			return fiber.NewError(fiber.StatusNotFound, "404 Not Found")
		} else {
			log.Printf("[HandleAdminLinkEditGet] Error querying Shortlink %v from database: %v\n", id, err)
			return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
		}
	}

	err = json.Unmarshal(c.Body(), &editlink)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "400 Bad Request: "+err.Error())
	}

	dblink.Description = editlink.Description
	dblink.URL = editlink.URL
	dblink.Modified = time.Now()

	_, err = models.DB.NewUpdate().Model(&dblink).WherePK().Exec(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error: "+err.Error())
	}

	c.Append("Location", c.BaseURL()+"/api/v1/links/"+editlink.ID)
	c.Status(fiber.StatusNoContent)

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

	c.Status(fiber.StatusNoContent)
	err = c.SendString("204 No Content")
	if err != nil {
		log.Println(err)
	}

	return nil
}
