package db

import (
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"log"
)

// IsCookieValid checks the database if Cookie val is valid.
//
// Returns true if it's valid, false if not.
func IsCookieValid(val string) bool {
	count, err := models.DB.NewSelect().Model((*models.Login)(nil)).Where("cookie = ?", val).Count(context.Background())
	if err != nil {
		log.Printf("Error checking cookie validity for cookie %v\n", val)
		return false
	}
	if count < 1 {
		return false
	} else {
		return true
	}
}
