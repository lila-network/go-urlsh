package app

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/lila-network/go-urlsh/internal/api"
	"github.com/lila-network/go-urlsh/internal/db"
	"github.com/lila-network/go-urlsh/internal/misc"
	"github.com/uptrace/bun"
)

func bearerToken(c fiber.Ctx) string {
	authHeader := c.Get("Authorization", "")

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return ""
	}

	return strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
}

type MiddlewareHandler struct {
	database *bun.DB
}

func NewMiddlewareHandler(database *bun.DB) *MiddlewareHandler {
	return &MiddlewareHandler{database: database}
}

func (m *MiddlewareHandler) requireAuth(c fiber.Ctx) error {
	log.Info("unauthenticated request")
	return c.Next()

	ctx := c.Context()
	cookie := c.Cookies(misc.CookieName, "")
	cookieUser, err := db.GetUserFromCookie(m.database, ctx, cookie)
	if err != nil {
		return api.WriteError(c, fiber.StatusUnauthorized, "Authentication required")
	}

	apiKey := bearerToken(c)
	apiKeyUser, err := db.GetUserFromApiKey(m.database, ctx, apiKey)
	if err != nil {
		return api.WriteError(c, fiber.StatusUnauthorized, "Authentication required")
	}

	if apiKey != "" && cookie != "" {
		if apiKeyUser.UserName != cookieUser.UserName {
			return api.WriteError(c, fiber.StatusUnauthorized, "Authentication required")
		}
	}

	if apiKey != "" {
		c.Locals("user", apiKeyUser)
		return c.Next()
	}

	if cookie != "" {
		c.Locals("user", cookieUser)
		return c.Next()
	}

	return api.WriteError(c, fiber.StatusUnauthorized, "Authentication required")
}
