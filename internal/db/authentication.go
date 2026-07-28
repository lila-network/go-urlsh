package db

import (
	"context"
	"fmt"

	"github.com/lila-network/go-urlsh/models"
	"github.com/uptrace/bun"
)

//func IsAuthenticated(db *bun.DB, ctx context.Context, c fiber.Ctx) bool {
//	if isCookieValid(db, ctx, c.Cookies(misc.CookieName, "")) {
//		return true
//	}
//
//	if isApiKeyValid(db, ctx, c.GetRespHeader("x-api-key", "")) {
//		return true
//	}
//
//	return false
//}

func GetUserFromApiKey(db *bun.DB, ctx context.Context, key string) (user models.User, err error) {
	if key == "" {
		err = fmt.Errorf("api key is empty")
		return
	}

	var dbApiKey models.ApiKey
	err = db.NewSelect().Model(&dbApiKey).Where("key = ?", key).Scan(ctx)
	if err != nil {
		return
	}

	err = db.NewSelect().Model(&user).Where("username = ?", dbApiKey.UserName).Scan(ctx)
	if err != nil {
		return
	}

	return
}

func GetUserFromCookie(db *bun.DB, ctx context.Context, cookie string) (user models.User, err error) {
	if cookie == "" {
		err = fmt.Errorf("cookie is empty")
		return
	}

	var session models.Session
	err = db.NewSelect().Model(&session).Where("cookie = ?", cookie).Scan(ctx)
	if err != nil {
		return
	}

	err = db.NewSelect().Model(&user).Where("username = ?", session.UserName).Scan(ctx)
	if err != nil {
		return
	}

	return
}
