package api

import (
	"github.com/lila-network/go-urlsh/internal/db"
	"github.com/lila-network/go-urlsh/internal/misc"
	"github.com/lila-network/go-urlsh/models"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func HandleUserPost(c *fiber.Ctx) error {
	var newuser models.LoginRequest

	err := json.Unmarshal(c.Body(), &newuser)
	if err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "400 Bad Request")
	}

	usercount, err := models.DB.NewSelect().Model((*models.User)(nil)).Count(context.Background())
	if err != nil {
		log.Printf("[POST /api/v1/users] Error querying database for users: %v\n", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
	}
	if usercount != 0 {
		log.Printf("[POST /api/v1/users] someone trying to create user but user already exists\n")
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	} else {
		salt := misc.RandomString(15)
		created := time.Now()
		hashbytes := sha256.Sum256([]byte(salt + newuser.Password))

		hash := hex.EncodeToString(hashbytes[:])

		user := new(models.User)
		user.UserName = newuser.Username
		user.PasswordSalt = salt
		user.PasswordHash = hash
		user.Created = created

		_, err = models.DB.NewInsert().Model(user).Exec(context.Background())
		if err != nil {
			log.Printf("[POST /api/v1/users] Error adding user %v to database : %v\n", newuser.Username, err.Error())
			return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
		}

		userresponse := models.UserResponse{UserName: newuser.Username, Created: created}

		c.Status(fiber.StatusCreated)
		err = c.JSON(userresponse)
		if err != nil {
			log.Println(err)
		}
		return nil
	}
}

func HandleUserMe(c *fiber.Ctx) error {
	if !db.IsCookieValid(c.Cookies(misc.CookieName, "")) && !db.IsApiKeyValid(c.GetRespHeader("x-api-key", "")) {
		return fiber.NewError(fiber.StatusUnauthorized, "401 Unauthorized")
	}

	cookie := c.Cookies(misc.CookieName, "")
	apikey := c.GetRespHeader("x-api-key", "")

	var authmethod string
	var user models.User

	if cookie != "" {
		authmethod = "cookie"
	} else {
		authmethod = "apikey"
	}

	switch authmethod {
	case "cookie":
		var session models.Session
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

	case "apikey":
		var key models.ApiKey
		err := models.DB.NewSelect().Model(&key).Where("key = ?", apikey).Scan(context.Background())
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("[HandleUserMe] ApiKey %v not found\n", apikey)
			} else {
				log.Printf("[HandleUserMe] Error querying apikey %v from database: %v\n", apikey, err)
			}
			return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
		}

		err = models.DB.NewSelect().Model(&user).Where("username = ?", key.UserName).Scan(context.Background())
		if err != nil {
			log.Printf("[HandleUserMe] Error querying user %v from database: %v\n", key.UserName, err)
			return fiber.NewError(fiber.StatusInternalServerError, "500 Internal Server Error")
		}
	}

	var userresponse = new(models.UserResponse)
	userresponse.UserName = user.UserName
	userresponse.Created = user.Created

	err := c.JSON(userresponse)
	if err != nil {
		return err
	}
	return nil
}
