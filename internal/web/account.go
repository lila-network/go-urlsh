package web

import (
	"context"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/lila-network/go-urlsh/internal/misc"
	"github.com/lila-network/go-urlsh/models"
)

func HandleAdminAccountGet(c fiber.Ctx) error {

	cookie := c.Cookies(misc.CookieName, "")
	var session models.Session
	var user models.User

	err := models.DB.NewSelect().Model(&session).Where("cookie = ?", cookie).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[HandleUserMe] Session %v not found\n", cookie)
		} else {
			log.Printf("[HandleUserMe] Error querying session %v from database: %v\n", cookie, err)
		}
		return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
	}

	err = models.DB.NewSelect().Model(&user).Where("username = ?", session.UserName).Scan(context.Background())
	if err != nil {
		log.Printf("[HandleUserMe] Error querying user %v from database: %v\n", session.UserName, err)
		return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
	}

	return c.Render("account", user)
}
