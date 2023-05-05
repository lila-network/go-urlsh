package api

import (
	"codeberg.org/lauralani/go-urlsh/internal/misc"
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"crypto/sha256"
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
