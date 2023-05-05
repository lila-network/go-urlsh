package db

import (
	"codeberg.org/lauralani/go-urlsh/models"
	"context"
	"log"
)

// IsApiKeyValid checks the database if ApiKey val is valid.
//
// Returns true if it's valid, false if not.
func IsApiKeyValid(val string) bool {
	if val == "" {
		return false
	}

	count, err := models.DB.NewSelect().Model((*models.ApiKey)(nil)).Where("key = ?", val).Count(context.Background())
	if err != nil {
		log.Printf("Error checking apikey validity for key %v\n", val)
		return false
	}
	if count < 1 {
		return false
	} else {
		return true
	}
}
