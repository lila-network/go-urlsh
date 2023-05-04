package web

import (
	"codeberg.org/lauralani/go-urlsh/internal/db"
	"codeberg.org/lauralani/go-urlsh/internal/misc"
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

func HandleAdminLoginPost(c *fiber.Ctx) error {
	var login models.LoginRequest
	var user models.User

	if err := c.BodyParser(&login); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "400 Bad Request")
	}

	err := models.DB.NewSelect().Model(&user).Where("username = ?", login.Username).Scan(context.Background())
	if err != nil {
		log.Printf("Error authenticating User %v: User doesn't exist by IP %v\n", login.Username, c.IP())
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	}

	passwordsumbytes := sha256.Sum256([]byte(user.PasswordSalt + login.Password))
	passwordsum := hex.EncodeToString(passwordsumbytes[:])

	if passwordsum == user.PasswordHash {
		// Passwords match
		var expires time.Time

		if login.Remember == "on" {
			expires = time.Now().Add(30 * 24 * time.Hour)
		} else {
			expires = time.Now().Add(24 * time.Hour)
		}

		key := uuid.New().String()
		dblogin := new(models.Login)

		cookie := new(fiber.Cookie)
		cookie.Name = misc.CookieName
		cookie.Value = key
		cookie.Expires = expires

		dblogin.Expires = expires
		dblogin.Cookie = key
		dblogin.UserName = user.UserName

		_, err = models.DB.NewInsert().Model(dblogin).Exec(context.Background())
		if err != nil {
			log.Printf("DB Error inserting login cookie information for user %v: %v\n", dblogin.UserName, err.Error())
			return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
		}

		user.LastLogin = time.Now()
		_, err = models.DB.NewUpdate().Model(&user).WherePK().Exec(context.Background())
		if err != nil {
			log.Printf("DB Error updating last login information for user %v: %v\n", dblogin.UserName, err.Error())
			return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
		}
		c.Cookie(cookie)
		c.Status(fiber.StatusSeeOther)
		c.Location("/admin/")
		return nil
	} else {
		log.Printf("Error authenticating User %v: password mismatch by IP %v\n", login.Username, c.IP())
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	}
}

func HandleAdminLoginGet(c *fiber.Ctx) error {
	if db.IsCookieValid(c.Cookies(misc.CookieName, "")) {
		c.Location("/admin/")
		c.Status(fiber.StatusSeeOther)
		return nil
	} else {
		return c.Render("login", nil)
	}
}
